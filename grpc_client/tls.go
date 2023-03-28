package grpc_client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

func LoadCredentials(caCrt string) (credentials.TransportCredentials, error) {
	// Use system CA certificates
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("unable to load system CA certificates: %w", err)
	}

	// Override if a custom CA cert is provided
	if caCrt != "" {
		pemServerCA, err := os.ReadFile(caCrt)
		if err != nil {
			return nil, fmt.Errorf("unable to load ca certificate from file: %w", err)
		}
		certPool = x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			return nil, fmt.Errorf("failed to add server CA's certificate")
		}
	}

	// Create the credentials and return it
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(tlsConfig), nil
}
