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
)

type ServerConfig struct {
	Logger                    *zap.SugaredLogger
	TlsEnabled                bool
	TlsServerCert             string
	TlsServerKey              string
	Username                  string
	Password                  string
	PersonService             models.PersonService
	PersonSearchService       models.PersonSearchService
	OrganizationService       models.OrganizationService
	OrganizationSearchService models.OrganizationSearchService
}

type server struct {
	v1.UnimplementedPeopleServer
	personService             models.PersonService
	personSearchService       models.PersonSearchService
	organizationService       models.OrganizationService
	organizationSearchService models.OrganizationSearchService
	logger                    *zap.SugaredLogger
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
		personService:             serverConfig.PersonService,
		personSearchService:       serverConfig.PersonSearchService,
		organizationService:       serverConfig.OrganizationService,
		organizationSearchService: serverConfig.OrganizationSearchService,
		logger:                    serverConfig.Logger,
	}

	v1.RegisterPeopleServer(gsrv, srv)

	return gsrv
}

func (srv *server) GetPerson(ctx context.Context, req *v1.GetPersonRequest) (*v1.GetPersonResponse, error) {
	person, err := srv.personService.GetPerson(ctx, req.Id)

	if err != nil && err == models.ErrNotFound {
		grpcErr := status.New(codes.InvalidArgument, fmt.Errorf("could not find person with id %s", req.Id).Error())
		return &v1.GetPersonResponse{
			Response: &v1.GetPersonResponse_Error{
				Error: grpcErr.Proto(),
			},
		}, nil
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve person with id '%s': %s", req.Id, err.Error())
	}

	return &v1.GetPersonResponse{
		Response: &v1.GetPersonResponse_Person{
			Person: &person.Person,
		},
	}, nil
}

func (srv *server) GetAllPerson(req *v1.GetAllPersonRequest, stream v1.People_GetAllPersonServer) error {

	return srv.personService.EachPerson(context.Background(), func(p *models.Person) bool {
		streamErr := stream.Send(&v1.GetAllPersonResponse{
			Person: &p.Person,
		})
		if streamErr != nil {
			srv.logger.Errorf("unable to send message to stream: %w", streamErr)
			return false
		}
		return true
	})

}

func (srv *server) ReindexPerson(req *v1.ReindexPersonRequest, stream v1.People_ReindexPersonServer) error {

	ctx := stream.Context()
	idxSwitcher, err := srv.personSearchService.NewIndexSwitcher(models.BulkIndexerConfig{
		OnError: func(err error) {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error: %w").Error())
			if err := stream.Send(&v1.ReindexPersonResponse{
				Response: &v1.ReindexPersonResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
		},
		OnIndexError: func(id string, err error) {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error for %s: %w", id, err).Error())
			if err := stream.Send(&v1.ReindexPersonResponse{
				Response: &v1.ReindexPersonResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
		},
	})
	if err != nil {
		log.Fatal("unable to initialize index switcher: %s", err.Error())
	}

	numIndexed := 0
	srv.personService.EachPerson(ctx, func(person *models.Person) bool {
		if err := idxSwitcher.Index(ctx, person); err != nil {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error for record %s: %w", person.Id, err).Error())
			if err := stream.Send(&v1.ReindexPersonResponse{
				Response: &v1.ReindexPersonResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
			return false
		}
		numIndexed++
		return true
	})

	if err := idxSwitcher.Switch(ctx); err != nil {
		grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error: %s", err).Error())
		if err := stream.Send(&v1.ReindexPersonResponse{
			Response: &v1.ReindexPersonResponse_Error{
				Error: grpcErr.Proto(),
			},
		}); err != nil {
			log.Fatal(err)
		}
	}

	if err := stream.Send(&v1.ReindexPersonResponse{
		Response: &v1.ReindexPersonResponse_Message{
			Message: fmt.Sprintf("%d records reindexed", numIndexed),
		},
	}); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (srv *server) SuggestPerson(req *v1.SuggestPersonRequest, stream v1.People_SuggestPersonServer) error {
	persons, _ := srv.personSearchService.Suggest(req.Query)

	for _, person := range persons {
		if err := stream.Send(&v1.SuggestPersonResponse{
			Person: &person.Person,
		}); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (srv *server) GetOrganization(ctx context.Context, req *v1.GetOrganizationRequest) (*v1.GetOrganizationResponse, error) {
	org, err := srv.organizationService.GetOrganization(ctx, req.Id)

	if err != nil && err == models.ErrNotFound {
		grpcErr := status.New(codes.InvalidArgument, fmt.Errorf("could not find organization with id %s", req.Id).Error())
		return &v1.GetOrganizationResponse{
			Response: &v1.GetOrganizationResponse_Error{
				Error: grpcErr.Proto(),
			},
		}, nil
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to retrieve organization with id '%s': %s", req.Id, err.Error())
	}

	return &v1.GetOrganizationResponse{
		Response: &v1.GetOrganizationResponse_Organization{
			Organization: &org.Organization,
		},
	}, nil
}

func (srv *server) GetAllOrganization(req *v1.GetAllOrganizationRequest, stream v1.People_GetAllOrganizationServer) error {

	err := srv.organizationService.EachOrganization(stream.Context(), func(o *models.Organization) bool {
		stream.Send(&v1.GetAllOrganizationResponse{
			Organization: &o.Organization,
		})
		return true
	})

	if err != nil {
		return fmt.Errorf("unable to retrieve all organizations from the database: %w", err)
	}

	return nil
}

func (srv *server) ReindexOrganization(req *v1.ReindexOrganizationRequest, stream v1.People_ReindexOrganizationServer) error {

	ctx := stream.Context()
	idxSwitcher, err := srv.organizationSearchService.NewIndexSwitcher(models.BulkIndexerConfig{
		OnError: func(err error) {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error: %w").Error())
			if err := stream.Send(&v1.ReindexOrganizationResponse{
				Response: &v1.ReindexOrganizationResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
		},
		OnIndexError: func(id string, err error) {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error for %s: %w", id, err).Error())
			if err := stream.Send(&v1.ReindexOrganizationResponse{
				Response: &v1.ReindexOrganizationResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
		},
	})
	if err != nil {
		log.Fatal("unable to initialize index switcher: %s", err.Error())
	}

	numIndexed := 0
	srv.organizationService.EachOrganization(ctx, func(org *models.Organization) bool {
		if err := idxSwitcher.Index(ctx, org); err != nil {
			grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error for record %s: %w", org.Id, err).Error())
			if err := stream.Send(&v1.ReindexOrganizationResponse{
				Response: &v1.ReindexOrganizationResponse_Error{
					Error: grpcErr.Proto(),
				},
			}); err != nil {
				log.Fatal(err)
			}
			return false
		}
		numIndexed++
		return true
	})

	if err := idxSwitcher.Switch(ctx); err != nil {
		grpcErr := status.New(codes.Internal, fmt.Errorf("es6 index error: %s", err).Error())
		if err := stream.Send(&v1.ReindexOrganizationResponse{
			Response: &v1.ReindexOrganizationResponse_Error{
				Error: grpcErr.Proto(),
			},
		}); err != nil {
			log.Fatal(err)
		}
	}

	if err := stream.Send(&v1.ReindexOrganizationResponse{
		Response: &v1.ReindexOrganizationResponse_Message{
			Message: fmt.Sprintf("%d records reindexed", numIndexed),
		},
	}); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (srv *server) SuggestOrganization(req *v1.SuggestOrganizationRequest, stream v1.People_SuggestOrganizationServer) error {
	organizations, _ := srv.organizationSearchService.Suggest(req.Query)

	for _, org := range organizations {
		if err := stream.Send(&v1.SuggestOrganizationResponse{
			Organization: &org.Organization,
		}); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
