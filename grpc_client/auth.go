package grpc_client

import (
	"context"
	"encoding/base64"
)

type BasicAuth struct {
	Username string
	Password string
}

func (b BasicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	enc := base64.StdEncoding.EncodeToString([]byte(b.Username + ":" + b.Password))
	return map[string]string{
		"Authorization": "Basic " + enc,
	}, nil
}

func (b BasicAuth) RequireTransportSecurity() bool {
	return false
}
