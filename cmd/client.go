package cmd

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_client"
)

func init() {
	getCmd.Flags().String("id", "", "identifier")
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {

		id, iErr := cmd.Flags().GetString("id")

		if iErr != nil {
			log.Fatal("no id given")
		}

		clientConfig := grpc_client.Config{
			Username: config.Api.Username,
			Password: config.Api.Password,
			Host:     config.Api.Host,
			Port:     config.Api.Port,
			Insecure: config.ApiClient.Insecure,
			CaCert:   config.ApiClient.CaCert,
		}

		err := grpc_client.Open(clientConfig, func(c v1.PeopleClient) error {

			req := v1.GetPersonRequest{Id: id}
			res, err := c.GetPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			person := res.GetPerson()

			bytes, err := json.Marshal(person)

			if err != nil {
				return err
			}

			os.Stdout.Write(bytes)

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}
