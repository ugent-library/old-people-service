// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddOrganization implements AddOrganization operation.
	//
	// Insert/update a single organization record.
	//
	// POST /add-organization
	AddOrganization(ctx context.Context, req *Organization) (*Organization, error)
	// AddPerson implements AddPerson operation.
	//
	// Insert/update a single person record.
	//
	// POST /add-person
	AddPerson(ctx context.Context, req *Person) (*Person, error)
	// GetOrganization implements GetOrganization operation.
	//
	// Get single organization record.
	//
	// POST /get-organization
	GetOrganization(ctx context.Context, req *GetOrganizationRequest) (*Organization, error)
	// GetOrganizationById implements GetOrganizationById operation.
	//
	// Get single organization record by one of its extra identifiers.
	//
	// POST /get-organization-by-id
	GetOrganizationById(ctx context.Context, req *GetOrganizationByIdRequest) (*Organization, error)
	// GetOrganizations implements GetOrganizations operation.
	//
	// Get all organization records.
	//
	// POST /get-organizations
	GetOrganizations(ctx context.Context, req *GetOrganizationsRequest) (*OrganizationListResponse, error)
	// GetPeople implements GetPeople operation.
	//
	// Get all person records.
	//
	// POST /get-people
	GetPeople(ctx context.Context, req *GetPeopleRequest) (*PersonListResponse, error)
	// GetPerson implements GetPerson operation.
	//
	// Retrieve a single person record.
	//
	// POST /get-person
	GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error)
	// GetPersonById implements GetPersonById operation.
	//
	// Retrieve a single person record by one of its extra identifiers.
	//
	// POST /get-person-by-id
	GetPersonById(ctx context.Context, req *GetPersonByIdRequest) (*Person, error)
	// SetPersonOrcid implements SetPersonOrcid operation.
	//
	// Update person ORCID.
	//
	// POST /set-person-orcid
	SetPersonOrcid(ctx context.Context, req *SetPersonOrcidRequest) (*Person, error)
	// SetPersonOrcidToken implements SetPersonOrcidToken operation.
	//
	// Update person ORCID token.
	//
	// POST /set-person-orcid-token
	SetPersonOrcidToken(ctx context.Context, req *SetPersonOrcidTokenRequest) (*Person, error)
	// SetPersonRole implements SetPersonRole operation.
	//
	// Update person role.
	//
	// POST /set-person-role
	SetPersonRole(ctx context.Context, req *SetPersonRoleRequest) (*Person, error)
	// SetPersonSettings implements SetPersonSettings operation.
	//
	// Update person settings.
	//
	// POST /set-person-settings
	SetPersonSettings(ctx context.Context, req *SetPersonSettingsRequest) (*Person, error)
	// SuggestOrganizations implements SuggestOrganizations operation.
	//
	// Search on organization records.
	//
	// POST /suggest-organizations
	SuggestOrganizations(ctx context.Context, req *SuggestOrganizationsRequest) (*OrganizationListResponse, error)
	// SuggestPeople implements SuggestPeople operation.
	//
	// Search on person records.
	//
	// POST /suggest-people
	SuggestPeople(ctx context.Context, req *SuggestPeopleRequest) (*PersonListResponse, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
