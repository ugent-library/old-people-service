package grpcserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/api/v1/v1connect"
	"github.com/ugent-library/person-service/models"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Logger     *zap.SugaredLogger
	Username   string
	Password   string
	Repository models.Repository
}

type server struct {
	v1connect.UnimplementedPersonServiceHandler
	repository models.Repository
	logger     *zap.SugaredLogger
}

func NewHandler(serverConfig *ServerConfig) (string, http.Handler) {
	// streaming handlers not available for regular http (so only GetPerson and GetOrganization)
	// those handlers always return status 415 (unsupported media type)
	srv := &server{
		repository: serverConfig.Repository,
		logger:     serverConfig.Logger,
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
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}

	person, err := srv.repository.GetPerson(ctx, req.Msg.Id)

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
	return srv.repository.EachPerson(ctx, func(p *models.Person) bool {
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

func (srv *server) SuggestPerson(ctx context.Context, req *connect.Request[v1.SuggestPersonRequest]) (*connect.Response[v1.SuggestPersonResponse], error) {
	if req.Msg.Query == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute query not given"),
		)
	}

	persons, err := srv.repository.SuggestPerson(ctx, req.Msg.Query)

	if err != nil {
		srv.logger.Errorf("grpc call SuggestPerson failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	v1People := make([]*v1.Person, 0, len(persons))
	for _, p := range persons {
		v1People = append(v1People, p.Person)
	}

	return connect.NewResponse[v1.SuggestPersonResponse](&v1.SuggestPersonResponse{
		Person: v1People,
	}), nil
}

func (srv *server) SetPersonOrcidToken(ctx context.Context, req *connect.Request[v1.SetPersonOrcidTokenRequest]) (*connect.Response[v1.SetPersonOrcidTokenResponse], error) {
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}
	if req.Msg.OrcidToken == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute orcid_token not given"),
		)
	}

	err := srv.repository.SetPersonOrcidToken(ctx, req.Msg.Id, req.Msg.OrcidToken)

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
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}

	org, err := srv.repository.GetOrganization(ctx, req.Msg.Id)

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
	err := srv.repository.EachOrganization(ctx, func(o *models.Organization) bool {
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

func (srv *server) SuggestOrganization(ctx context.Context, req *connect.Request[v1.SuggestOrganizationRequest]) (*connect.Response[v1.SuggestOrganizationResponse], error) {
	if req.Msg.Query == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute query not given"),
		)
	}

	organizations, err := srv.repository.SuggestOrganization(ctx, req.Msg.Query)

	if err != nil {
		srv.logger.Errorf("grpc call SuggestOrganization failed: %s", err)
		return nil, connect.NewError(
			connect.CodeInternal,
			fmt.Errorf("internal server error"),
		)
	}

	v1Orgs := make([]*v1.Organization, 0, len(organizations))
	for _, org := range organizations {
		v1Orgs = append(v1Orgs, org.Organization)
	}

	return connect.NewResponse[v1.SuggestOrganizationResponse](&v1.SuggestOrganizationResponse{
		Organization: v1Orgs,
	}), nil
}

func (srv *server) SetPersonOrcid(ctx context.Context, req *connect.Request[v1.SetPersonOrcidRequest]) (*connect.Response[v1.SetPersonOrcidResponse], error) {
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}
	if req.Msg.Orcid == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute orcid not given"),
		)
	}

	err := srv.repository.SetPersonOrcid(ctx, req.Msg.Id, req.Msg.Orcid)

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
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}
	if len(req.Msg.Role) == 0 {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute role not given"),
		)
	}

	err := srv.repository.SetPersonRole(ctx, req.Msg.Id, req.Msg.Role)

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
	if req.Msg.Id == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("attribute id not given"),
		)
	}

	err := srv.repository.SetPersonSettings(ctx, req.Msg.Id, req.Msg.Settings)

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
