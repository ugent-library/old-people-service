package grpc_server

import (
	"context"
	"fmt"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerConfig struct {
	Logger        *zap.SugaredLogger
	TlsEnabled    bool
	TlsServerCert string
	TlsServerKey  string
	Username      string
	Password      string
	PersonService models.PersonService
}

type server struct {
	v1.UnimplementedPeopleServer
	personService models.PersonService
	logger        *zap.SugaredLogger
}

func NewServer(serverConfig *ServerConfig) *grpc.Server {
	zap_opt := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel

			case codes.Internal:
				l = zapcore.ErrorLevel

			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)
	// Defaults to an insecure connection
	tlsOption := grpc.Creds(nil)

	// If set, enabled server-side TLS secure connection
	if serverConfig.TlsEnabled {
		creds, err := LoadCredentials(serverConfig.TlsServerCert, serverConfig.TlsServerKey)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		tlsOption = grpc.Creds(creds)
	}

	authCb := newAuth(AuthConfig{
		Username: serverConfig.Username,
		Password: serverConfig.Password,
	})

	gsrv := grpc.NewServer(
		tlsOption,
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(serverConfig.Logger.Desugar(), zap_opt),
			grpc_auth.StreamServerInterceptor(authCb),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(serverConfig.Logger.Desugar(), zap_opt),
			grpc_auth.UnaryServerInterceptor(authCb),
		),
	)

	// Enable the gRPC reflection API
	// e.g. grpcurl host:port list -> list all available services & methods
	reflection.Register(gsrv)

	srv := &server{
		personService: serverConfig.PersonService,
		logger:        serverConfig.Logger,
	}

	v1.RegisterPeopleServer(gsrv, srv)

	return gsrv
}

func (srv *server) GetPerson(ctx context.Context, getPersonRequest *v1.GetPersonRequest) (*v1.GetPersonResponse, error) {
	person, err := srv.personService.Get(ctx, getPersonRequest.Id)

	if err != nil && err == models.ErrNotFound {
		grpcErr := status.New(codes.InvalidArgument, fmt.Errorf("could not find person with id %s", getPersonRequest.Id).Error())
		return &v1.GetPersonResponse{
			Response: &v1.GetPersonResponse_Error{
				Error: grpcErr.Proto(),
			},
		}, nil
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve person with id '%s': %s", getPersonRequest.Id, err.Error())
	}

	return &v1.GetPersonResponse{
		Response: &v1.GetPersonResponse_Person{
			Person: mapPerson(person),
		},
	}, nil
}

func mapPerson(p *models.Person) *v1.Person {
	v := &v1.Person{
		Id:                 p.ID,
		Active:             p.Active,
		DateCreated:        timestamppb.New(*p.DateCreated),
		DateUpdated:        timestamppb.New(*p.DateUpdated),
		FullName:           p.FullName,
		FirstName:          p.FirstName,
		LastName:           p.LastName,
		Email:              p.Email,
		Orcid:              p.Orcid,
		OrcidToken:         p.OrcidToken,
		PreferredFirstName: p.PreferredFirstName,
		PreferredLastName:  p.PreferedLastName,
		BirthDate:          p.BirthDate,
		JobTitle:           p.JobTitle,
	}

	v.OrganizationId = append([]string{}, p.OrganizationID...)
	v.OtherId = make([]*v1.IdRef, 0, len(p.OtherID))
	for _, otherId := range p.OtherID {
		v.OtherId = append(v.OtherId, &v1.IdRef{
			Id:   otherId.ID,
			Type: otherId.Type,
		})
	}

	return v
}
