package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/api/v1/v1connect"
	"github.com/ugent-library/person-service/models"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Logger                     *zap.SugaredLogger
	Username                   string
	Password                   string
	PersonService              models.PersonService
	PersonSuggestService       models.PersonSuggestService
	OrganizationService        models.OrganizationService
	OrganizationSuggestService models.OrganizationSuggestService
}

type server struct {
	v1connect.UnimplementedPersonServiceHandler
	personService              models.PersonService
	personSuggestService       models.PersonSuggestService
	organizationService        models.OrganizationService
	organizationSuggestService models.OrganizationSuggestService
	logger                     *zap.SugaredLogger
}

func NewHandler(serverConfig *ServerConfig) (string, http.Handler) {

	// streaming handlers not available for regular http (so only GetPerson and GetOrganization)
	// those handlers always return status 415 (unsupported media type)
	srv := &server{
		personService:              serverConfig.PersonService,
		personSuggestService:       serverConfig.PersonSuggestService,
		organizationService:        serverConfig.OrganizationService,
		organizationSuggestService: serverConfig.OrganizationSuggestService,
		logger:                     serverConfig.Logger,
	}

	interceptors := connect.WithInterceptors(
		newAuthInterceptor(AuthConfig{
			Username: serverConfig.Username,
			Password: serverConfig.Password,
		}),
	)

	return v1connect.NewPersonServiceHandler(
		srv,
		interceptors,
	)
}

func (srv *server) GetPerson(ctx context.Context, req *connect.Request[v1.GetPersonRequest]) (*connect.Response[v1.GetPersonResponse], error) {
	person, err := srv.personService.GetPerson(ctx, req.Msg.Id)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find person with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call GetPerson failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.GetPersonResponse](&v1.GetPersonResponse{
		Person: person.Person,
	}), nil
}

func (srv *server) GetAllPerson(ctx context.Context, req *connect.Request[v1.GetAllPersonRequest], stream *connect.ServerStream[v1.GetAllPersonResponse]) error {
	return srv.personService.EachPerson(ctx, func(p *models.Person) bool {
		streamErr := stream.Send(&v1.GetAllPersonResponse{
			Person: p.Person,
		})
		if streamErr != nil {
			srv.logger.Errorf("unable to send message to stream: %w", streamErr)
			return false
		}
		return true
	})

}

func (srv *server) SuggestPerson(ctx context.Context, req *connect.Request[v1.SuggestPersonRequest], stream *connect.ServerStream[v1.SuggestPersonResponse]) error {
	persons, _ := srv.personSuggestService.SuggestPerson(ctx, req.Msg.Query)

	for _, person := range persons {
		if err := stream.Send(&v1.SuggestPersonResponse{
			Person: person.Person,
		}); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (srv *server) SetPersonOrcidToken(ctx context.Context, req *connect.Request[v1.SetPersonOrcidTokenRequest]) (*connect.Response[v1.SetPersonOrcidTokenResponse], error) {
	err := srv.personService.SetOrcidToken(ctx, req.Msg.Id, req.Msg.OrcidToken)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find person with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call SetPersonOrcidToken failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.SetPersonOrcidTokenResponse](&v1.SetPersonOrcidTokenResponse{
		Message: "ok",
	}), nil
}

func (srv *server) GetOrganization(ctx context.Context, req *connect.Request[v1.GetOrganizationRequest]) (*connect.Response[v1.GetOrganizationResponse], error) {
	org, err := srv.organizationService.GetOrganization(ctx, req.Msg.Id)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find organization with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call GetOrganization failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.GetOrganizationResponse](&v1.GetOrganizationResponse{
		Organization: org.Organization,
	}), nil
}

func (srv *server) GetAllOrganization(ctx context.Context, req *connect.Request[v1.GetAllOrganizationRequest], stream *connect.ServerStream[v1.GetAllOrganizationResponse]) error {

	err := srv.organizationService.EachOrganization(ctx, func(o *models.Organization) bool {
		stream.Send(&v1.GetAllOrganizationResponse{
			Organization: o.Organization,
		})
		return true
	})

	if err != nil {
		return fmt.Errorf("unable to retrieve all organizations from the database: %w", err)
	}

	return nil
}

func (srv *server) SuggestOrganization(ctx context.Context, req *connect.Request[v1.SuggestOrganizationRequest], stream *connect.ServerStream[v1.SuggestOrganizationResponse]) error {
	organizations, _ := srv.organizationSuggestService.SuggestOrganization(ctx, req.Msg.Query)

	for _, org := range organizations {
		if err := stream.Send(&v1.SuggestOrganizationResponse{
			Organization: org.Organization,
		}); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (srv *server) SetPersonOrcid(ctx context.Context, req *connect.Request[v1.SetPersonOrcidRequest]) (*connect.Response[v1.SetPersonOrcidResponse], error) {
	err := srv.personService.SetOrcid(ctx, req.Msg.Id, req.Msg.Orcid)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find person with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call SetPersonOrcid failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.SetPersonOrcidResponse](&v1.SetPersonOrcidResponse{
		Message: "ok",
	}), nil
}

func (srv *server) SetPersonRole(ctx context.Context, req *connect.Request[v1.SetPersonRoleRequest]) (*connect.Response[v1.SetPersonRoleResponse], error) {
	err := srv.personService.SetRole(ctx, req.Msg.Id, req.Msg.Role)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find person with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call SetPersonRole failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.SetPersonRoleResponse](&v1.SetPersonRoleResponse{
		Message: "ok",
	}), nil
}

func (srv *server) SetPersonSettings(ctx context.Context, req *connect.Request[v1.SetPersonSettingsRequest]) (*connect.Response[v1.SetPersonSettingsResponse], error) {
	err := srv.personService.SetSettings(ctx, req.Msg.Id, req.Msg.Settings)

	if err != nil && err == models.ErrNotFound {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("could not find person with id %s", req.Msg.Id),
		)
	} else if err != nil {
		srv.logger.Errorf("grpc call SetPersonSettings failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	return connect.NewResponse[v1.SetPersonSettingsResponse](&v1.SetPersonSettingsResponse{
		Message: "ok",
	}), nil
}
