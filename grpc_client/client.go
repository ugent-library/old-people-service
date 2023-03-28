package grpc_client

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/ugent-library/people/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Open(config Config, cb func(c v1.PeopleClient) error) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Set TLS encryption
	var dialOptionSecureConn grpc.DialOption
	if config.Insecure {
		dialOptionSecureConn = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		creds, err := LoadCredentials(config.CaCert)
		if err != nil {
			return fmt.Errorf("unable to load cacert: %w", err)
		}
		dialOptionSecureConn = grpc.WithTransportCredentials(creds)
	}

	// Set up the connection and the API client with Basic Authentication
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := grpc.DialContext(ctx, addr,
		dialOptionSecureConn,
		grpc.WithPerRPCCredentials(BasicAuth{
			Username: config.Username,
			Password: config.Password,
		}),
		grpc.WithBlock(),
	)

	if err != nil {
		return fmt.Errorf("unable to open connection: %w", err)
	}

	defer conn.Close()

	client := v1.NewPeopleClient(conn)

	err = cb(client)
	if err != nil {
		return fmt.Errorf("client callback failed: %w", err)
	}

	return nil
}
