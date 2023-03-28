package grpc_server

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc/credentials"
)

func LoadCredentials(cert string, key string) (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, fmt.Errorf("cannot load TLS credentials: %w", err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(tlsConfig), nil
}
