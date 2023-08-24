package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/ugent-library/people-service/models"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Logger     *zap.SugaredLogger
	Repository models.Repository
}

type Service struct {
	repository models.Repository
	logger     *zap.SugaredLogger
}

func NewService(serverConfig *ServerConfig) *Service {
	return &Service{
		logger:     serverConfig.Logger,
		repository: serverConfig.Repository,
	}
}

func (svc *Service) GetOrganization(ctx context.Context, params GetOrganizationParams) (*Organization, error) {
	org, err := svc.repository.GetOrganization(ctx, params.OrganizationId)

	if err != nil && err == models.ErrNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return mapToExternalOrganization(org), nil
}

func (svc *Service) GetOrganizations(ctx context.Context, params GetOrganizationsParams) (*PagedOrganizationListResponse, error) {
	var organizations []*models.Organization
	var err error
	var cursor string

	if params.Cursor.Set {
		organizations, cursor, err = svc.repository.GetMoreOrganizations(ctx, params.Cursor.Value)
	} else {
		organizations, cursor, err = svc.repository.GetOrganizations(ctx)
	}
	if err != nil {
		return nil, err
	}

	res := &PagedOrganizationListResponse{
		Data: make([]Organization, 0, len(organizations)),
	}
	if cursor != "" {
		res.Cursor = NewOptString(cursor)
	}
	for _, org := range organizations {
		res.Data = append(res.Data, *mapToExternalOrganization(org))
	}

	return res, nil
}

func (svc *Service) SuggestOrganizations(ctx context.Context, params SuggestOrganizationsParams) (*PagedOrganizationListResponse, error) {
	orgs, err := svc.repository.SuggestOrganization(ctx, params.Query)
	if err != nil {
		return nil, err
	}

	res := &PagedOrganizationListResponse{
		Data: make([]Organization, 0, len(orgs)),
	}
	for _, org := range orgs {
		res.Data = append(res.Data, *mapToExternalOrganization(org))
	}

	return res, nil
}

func (svc *Service) GetPeople(ctx context.Context, params GetPeopleParams) (*PagedPersonListResponse, error) {
	var people []*models.Person
	var err error
	var cursor string

	if params.Cursor.Set {
		people, cursor, err = svc.repository.GetMorePeople(ctx, params.Cursor.Value)
	} else {
		people, cursor, err = svc.repository.GetPeople(ctx)
	}
	if err != nil {
		return nil, err
	}

	res := &PagedPersonListResponse{
		Data: make([]Person, 0, len(people)),
	}
	if cursor != "" {
		res.Cursor = NewOptString(cursor)
	}
	for _, person := range people {
		res.Data = append(res.Data, *mapToExternalPerson(person))
	}

	return res, nil
}

func (svc *Service) SuggestPeople(ctx context.Context, params SuggestPeopleParams) (*PagedPersonListResponse, error) {
	people, err := svc.repository.SuggestPerson(ctx, params.Query)
	if err != nil {
		return nil, err
	}

	res := &PagedPersonListResponse{
		Data: make([]Person, 0, len(people)),
	}
	for _, person := range people {
		res.Data = append(res.Data, *mapToExternalPerson(person))
	}

	return res, nil
}

func (svc *Service) GetPerson(ctx context.Context, params GetPersonParams) (*Person, error) {
	person, err := svc.repository.GetPerson(ctx, params.PersonId)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (svc *Service) SetPersonOrcid(ctx context.Context, req *SetPersonOrcidRequest, params SetPersonOrcidParams) error {
	return svc.repository.SetPersonOrcid(ctx, params.PersonId, req.Orcid)
}

func (svc *Service) SetPersonOrcidToken(ctx context.Context, req *SetPersonOrcidTokenRequest, params SetPersonOrcidTokenParams) error {
	return svc.repository.SetPersonOrcidToken(ctx, params.PersonId, req.OrcidToken)
}

func (svc *Service) SetPersonRole(ctx context.Context, req *SetPersonRoleRequest, params SetPersonRoleParams) error {
	return svc.repository.SetPersonRole(ctx, params.PersonId, req.Role)
}

func (svc *Service) SetPersonSettings(ctx context.Context, req *SetPersonSettingsRequest, params SetPersonSettingsParams) error {
	if req.Settings == nil {
		return fmt.Errorf("%w: attribute settings is missing in request body", models.ErrMissingArgument)
	}
	return svc.repository.SetPersonSettings(ctx, params.PersonId, req.Settings)
}

func (svc *Service) NewError(ctx context.Context, err error) *ErrorStatusCode {
	if errors.Is(err, models.ErrNotFound) {
		return &ErrorStatusCode{
			StatusCode: 404,
			Response: Error{
				Code:    404,
				Message: err.Error(),
			},
		}
	}
	if errors.Is(err, models.ErrMissingArgument) {
		return &ErrorStatusCode{
			StatusCode: 400,
			Response: Error{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &ErrorStatusCode{
		StatusCode: 500,
		Response: Error{
			Code:    500,
			Message: err.Error(),
		},
	}
}

func mapToExternalOrganization(org *models.Organization) *Organization {
	o := &Organization{}
	o.ID = org.Id
	if org.GismoId != "" {
		o.GismoID = NewOptString(org.GismoId)
	}
	o.DateCreated = org.DateCreated.AsTime()
	o.DateUpdated = org.DateUpdated.AsTime()
	if org.NameDut != "" {
		o.NameDut = NewOptString(org.NameDut)
	}
	if org.NameEng != "" {
		o.NameEng = NewOptString(org.NameEng)
	}
	for _, otherId := range org.OtherId {
		o.OtherID = append(o.OtherID, IdRef{
			ID:   otherId.Id,
			Type: otherId.Type,
		})
	}
	if org.ParentId != "" {
		o.ParentID = NewOptString(org.ParentId)
	}
	o.Type = org.Type

	return o
}

func mapToExternalPerson(person *models.Person) *Person {
	p := &Person{}
	p.ID = person.Id
	p.Active = person.Active
	if person.BirthDate != "" {
		p.BirthDate = NewOptString(person.BirthDate)
	}
	p.DateCreated = person.DateCreated.AsTime()
	p.DateUpdated = person.DateUpdated.AsTime()
	if person.Email != "" {
		p.Email = NewOptString(person.Email)
	}
	if person.ExpirationDate != "" {
		p.ExpirationDate = NewOptString(person.ExpirationDate)
	}
	if person.FirstName != "" {
		p.FirstName = NewOptString(person.FirstName)
	}
	if person.LastName != "" {
		p.LastName = NewOptString(person.LastName)
	}
	if person.FullName != "" {
		p.FullName = NewOptString(person.FullName)
	}
	if person.PreferredFirstName != "" {
		p.PreferredFirstName = NewOptString(person.PreferredFirstName)
	}
	if person.PreferredLastName != "" {
		p.PreferredLastName = NewOptString(person.PreferredLastName)
	}
	if person.GismoId != "" {
		p.GismoID = NewOptString(person.GismoId)
	}
	p.JobCategory = append(p.JobCategory, person.JobCategory...)
	p.ObjectClass = append(p.ObjectClass, person.ObjectClass...)
	if person.Orcid != "" {
		p.Orcid = NewOptString(person.Orcid)
	}
	if person.OrcidToken != "" {
		p.OrcidToken = NewOptString(person.OrcidToken)
	}
	for _, orgRef := range person.Organization {
		oRef := OrganizationRef{
			ID:          orgRef.Id,
			DateCreated: orgRef.DateCreated.AsTime(),
			DateUpdated: orgRef.DateUpdated.AsTime(),
			From:        orgRef.From.AsTime(),
		}
		if orgRef.Until != nil {
			oRef.Until = NewOptDateTime(orgRef.Until.AsTime())
		}
		p.Organization = append(p.Organization, oRef)
	}
	for _, otherId := range person.OtherId {
		p.OtherID = append(p.OtherID, IdRef{
			ID:   otherId.Id,
			Type: otherId.Type,
		})
	}
	p.Role = append(p.Role, person.Role...)
	if person.Settings != nil {
		pSettings := PersonSettings{}
		for k, v := range person.Settings {
			pSettings[k] = v
		}
		p.Settings = NewOptPersonSettings(pSettings)
	}
	if person.Title != "" {
		p.Title = NewOptString(person.Title)
	}

	return p
}
