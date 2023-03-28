package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ugent-library/people/grpc_server"
	"github.com/ugent-library/people/models"
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

		ps, err := models.NewPersonService(&models.PersonConfig{
			DB: config.Db.Url,
		})

		if err != nil {
			logger.Fatal(err)
		}

		srv := grpc_server.NewServer(&grpc_server.ServerConfig{
			Logger:        logger,
			PersonService: ps,
			Username:      viper.GetString("api.username"),
			Password:      viper.GetString("api.password"),
			TlsEnabled:    viper.GetBool("api.tls_enabled"),
			TlsServerCert: viper.GetString("api.tls_server_crt"),
			TlsServerKey:  viper.GetString("api.tls_server_key"),
		})

		addr := fmt.Sprintf("%s:%d", viper.GetString("api.host"), viper.GetInt("api.port"))
		logger.Infof("Listening at %s", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		if err := srv.Serve(listener); err != nil {
			// TODO not everything is a fatal error.
			log.Println(err)
		}
	},
}
