package cmd

import (
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_client"
)

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
