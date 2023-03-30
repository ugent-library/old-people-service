package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_client"
	"google.golang.org/grpc/status"
)

func init() {
	getCmd.Flags().String("id", "", "identifier")
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(allCmd)
}

func openClient(cb func(client v1.PeopleClient) error) error {
	clientConfig := grpc_client.Config{
		Username: config.Api.Username,
		Password: config.Api.Password,
		Host:     config.Api.Host,
		Port:     config.Api.Port,
		Insecure: config.ApiClient.Insecure,
		CaCert:   config.ApiClient.CaCert,
	}
	return grpc_client.Open(clientConfig, cb)
}

var getCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {

		id, iErr := cmd.Flags().GetString("id")

		if iErr != nil {
			log.Fatal("no id given")
		}

		err := openClient(func(c v1.PeopleClient) error {

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

var allCmd = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetAllPersonRequest{}
			stream, err := c.GetAllPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}

				if err != nil {
					if st, ok := status.FromError(err); ok {
						return errors.New(st.Message())
					}

					return err
				}

				if p := res.GetPerson(); p != nil {
					bytes, _ := json.Marshal(p)
					fmt.Printf("%s\n", string(bytes))
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}
