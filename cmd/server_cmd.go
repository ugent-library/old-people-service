package cmd

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ory/graceful"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/api/v2"
	"github.com/ugent-library/zaphttp"
	"github.com/ugent-library/zaphttp/zapchi"
)

func init() {
	serverCmd.AddCommand(serverStartCmd)
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [command]",
	Short: "server commands",
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api server",
	Run: func(cmd *cobra.Command, args []string) {
		mux := chi.NewMux()

		mux.Use(middleware.RequestID)
		if config.Production {
			mux.Use(middleware.RealIP)
		}
		mux.Use(zaphttp.SetLogger(logger.Desugar(), zapchi.RequestID))
		mux.Use(middleware.RequestLogger(zapchi.LogFormatter()))
		mux.Use(middleware.Recoverer)

		apiServer, err := api.NewServer(
			api.NewService(&api.ServerConfig{
				Logger:     logger,
				Repository: Repository(),
			}),
			api.WithErrorHandler(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
				status := ogenerrors.ErrorCode(err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				data, _ := json.Marshal(ErrorMessage{
					Code:    status,
					Message: err.Error(),
				})
				w.Write(data)
			}),
		)
		if err != nil {
			logger.Fatal(err)
		}

		mux.Mount("/api/v1", http.StripPrefix("/api/v1", apiServer))

		srv := graceful.WithDefaults(&http.Server{
			Addr:         config.Api.Addr(),
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})

		logger.Infof("starting server at %s", config.Api.Addr())
		if err := graceful.Graceful(srv.ListenAndServe, srv.Shutdown); err != nil {
			logger.Fatal(err)
		}
		logger.Info("gracefully stopped server")
	},
}
