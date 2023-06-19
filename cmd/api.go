package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/ugent-library/people/grpc_server"
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

		services := Services()
		srv := grpc_server.NewServer(&grpc_server.ServerConfig{
			Logger:                     logger,
			PersonService:              services.PersonService,
			PersonSuggestService:       services.PersonSuggestService,
			OrganizationService:        services.OrganizationService,
			OrganizationSuggestService: services.OrganizationSuggestService,
			Username:                   config.Api.Username,
			Password:                   config.Api.Password,
			TlsEnabled:                 config.Api.TlsEnabled,
			TlsServerCert:              config.Api.TlsServerCert,
			TlsServerKey:               config.Api.TlsServerKey,
		})

		addr := fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port)
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
