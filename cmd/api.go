package cmd

import (
	"net/http"
	"time"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ory/graceful"
	"github.com/spf13/cobra"
	"github.com/ugent-library/person-service/api/v1/v1connect"
	"github.com/ugent-library/person-service/grpcserver"
	"github.com/ugent-library/zaphttp"
	"github.com/ugent-library/zaphttp/zapchi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func init() {
	apiCmd.AddCommand(apiStartCmd)
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api [command]",
	Short: "api commands",
}

var apiStartCmd = &cobra.Command{
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

		services := Services()
		grpcPath, grpcHandler := grpcserver.NewHandler(&grpcserver.ServerConfig{
			Logger:                     logger,
			PersonService:              services.PersonService,
			PersonSuggestService:       services.PersonSuggestService,
			OrganizationService:        services.OrganizationService,
			OrganizationSuggestService: services.OrganizationSuggestService,
			Username:                   config.Api.Username,
			Password:                   config.Api.Password,
		})
		// important: use Mount instead of Handle. Otherwise reflection api does not work
		mux.Mount(grpcPath, grpcHandler)

		reflector := grpcreflect.NewStaticReflector(v1connect.PersonServiceName)
		mux.Mount(grpcreflect.NewHandlerV1(reflector))
		mux.Mount(grpcreflect.NewHandlerV1Alpha(reflector))

		checker := grpchealth.NewStaticChecker(v1connect.PersonServiceName)
		mux.Mount(grpchealth.NewHandler(checker))

		handler := h2c.NewHandler(mux, &http2.Server{})

		srv := graceful.WithDefaults(&http.Server{
			Addr:         config.Api.Addr(),
			Handler:      handler,
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
