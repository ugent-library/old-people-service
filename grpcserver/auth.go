package grpcserver

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
)

type AuthConfig struct {
	Username string
	Password string
}

type AuthInterceptor struct {
	AuthConfig
}

func newAuthInterceptor(config AuthConfig) *AuthInterceptor {
	return &AuthInterceptor{
		AuthConfig: config,
	}
}

// cf. https://connect.build/docs/go/interceptors
// cf. https://connect.build/docs/go/streaming
func (ai *AuthInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		if err := ai.checkAuth(req.Header()); err != nil {
			return nil, err
		}
		return next(ctx, req)
	})

}

func (ai *AuthInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	})
}

func (ai *AuthInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {

		if err := ai.checkAuth(conn.RequestHeader()); err != nil {
			return err
		}
		return next(ctx, conn)
	})
}

func extractAuthorization(headers http.Header) (string, error) {
	h := headers.Get("authorization")
	h = strings.TrimSpace(h)
	h = strings.ReplaceAll(h, "  ", " ")
	if h == "" {
		return "", errors.New("not found")
	}
	parts := strings.Split(h, " ")
	if strings.ToLower(parts[0]) != "basic" {
		return "", errors.New("not found")
	}
	return parts[1], nil
}

func (ai *AuthInterceptor) checkAuth(headers http.Header) *connect.Error {
	token, err := extractAuthorization(headers)

	if err != nil {
		return connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("authentication failed"),
		)
	}

	c, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("authentication failed"),
		)
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("authentication failed"),
		)
	}
	user, password := cs[:s], cs[s+1:]

	if user != ai.Username || password != ai.Password {
		return connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("authentication failed"),
		)
	}

	return nil
}
