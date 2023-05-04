package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_client"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	getPersonCmd.Flags().String("id", "", "identifier")
	suggestPersonCmd.Flags().String("query", "", "query")
	suggestOrganizationCmd.Flags().String("query", "", "query")

	personCmd.AddCommand(getPersonCmd)
	personCmd.AddCommand(getAllPersonCmd)
	personCmd.AddCommand(suggestPersonCmd)

	organizationCmd.AddCommand(allOrganizationCmd)
	organizationCmd.AddCommand(suggestOrganizationCmd)

	rootCmd.AddCommand(personCmd)
	rootCmd.AddCommand(organizationCmd)
}

func openClient(cb func(client v1.PeopleClient) error) error {
	clientConfig := grpc_client.ClientConfig{
		Username: config.ApiClient.Username,
		Password: config.ApiClient.Password,
		Host:     config.ApiClient.Host,
		Port:     config.ApiClient.Port,
		Insecure: config.ApiClient.Insecure,
		CaCert:   config.ApiClient.CaCert,
	}
	return grpc_client.Open(clientConfig, cb)
}

var personCmd = &cobra.Command{
	Use: "person",
}

var getPersonCmd = &cobra.Command{
	Use: "get [id]",
	Run: func(cmd *cobra.Command, args []string) {

		id, iErr := cmd.Flags().GetString("id")

		if iErr != nil {
			log.Fatal("no id given")
		}
		if id == "" {
			log.Fatal("no id given")
		}

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetPersonRequest{Id: id}
			res, err := c.GetPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			person := res.GetPerson()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				UseProtoNames:   true,
			}

			bytes, err := marshaller.Marshal(person)

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

var getAllPersonCmd = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetAllPersonRequest{}
			stream, err := c.GetAllPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				// alternative to protobuf field option "json_name"
				UseProtoNames: true,
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
					bytes, _ := marshaller.Marshal(p)
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

var suggestPersonCmd = &cobra.Command{
	Use: "suggest",
	Run: func(cmd *cobra.Command, args []string) {

		query, err := cmd.Flags().GetString("query")
		if err != nil {
			log.Fatal("no query given")
		}
		if query == "" {
			log.Fatal("no query given")
		}

		cErr := openClient(func(c v1.PeopleClient) error {

			req := v1.SuggestPersonRequest{Query: query}
			stream, err := c.SuggestPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				// alternative to protobuf field option "json_name"
				UseProtoNames: true,
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
					bytes, _ := marshaller.Marshal(p)
					fmt.Printf("%s\n", string(bytes))
				}
			}

			return nil
		})

		if cErr != nil {
			log.Fatal(cErr)
		}

	},
}

var organizationCmd = &cobra.Command{
	Use: "organization",
}

var allOrganizationCmd = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetAllOrganizationRequest{}
			stream, err := c.GetAllOrganization(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				// alternative to protobuf field option "json_name"
				UseProtoNames: true,
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

				if p := res.GetOrganization(); p != nil {
					bytes, _ := marshaller.Marshal(p)
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

var suggestOrganizationCmd = &cobra.Command{
	Use: "suggest",
	Run: func(cmd *cobra.Command, args []string) {

		query, err := cmd.Flags().GetString("query")
		if err != nil {
			log.Fatal("no query given")
		}
		if query == "" {
			log.Fatal("no query given")
		}

		cErr := openClient(func(c v1.PeopleClient) error {

			req := v1.SuggestOrganizationRequest{Query: query}
			stream, err := c.SuggestOrganization(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				// alternative to protobuf field option "json_name"
				UseProtoNames: true,
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

				if p := res.GetOrganization(); p != nil {
					bytes, _ := marshaller.Marshal(p)
					fmt.Printf("%s\n", string(bytes))
				}
			}

			return nil
		})

		if cErr != nil {
			log.Fatal(cErr)
		}

	},
}
