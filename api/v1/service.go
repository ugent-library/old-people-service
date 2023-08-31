package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/ugent-library/people-service/models"
)

type Service struct {
	repository models.Repository
}

func NewService(repository models.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPerson(ctx context.Context, req *GetPersonRequest) (*Person, error) {
	person, err := s.repository.GetPerson(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (s *Service) GetPeople(ctx context.Context, req *GetPeopleRequest) (*PersonListResponse, error) {
	var people []*models.Person
	var err error
	var cursor string

	if req.Cursor != "" {
		people, cursor, err = s.repository.GetMorePeople(ctx, req.Cursor)
	} else {
		people, cursor, err = s.repository.GetPeople(ctx)
	}
	if err != nil {
		return nil, err
	}

	res := &PersonListResponse{
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

func (s *Service) SuggestPeople(ctx context.Context, req *SuggestPeopleRequest) (*PersonListResponse, error) {
	people, err := s.repository.SuggestPeople(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	res := &PersonListResponse{
		Data: make([]Person, 0, len(people)),
	}
	for _, person := range people {
		res.Data = append(res.Data, *mapToExternalPerson(person))
	}

	return res, nil
}

func (s *Service) SetPersonOrcid(ctx context.Context, req *SetPersonOrcidRequest) (*Person, error) {
	if err := s.repository.SetPersonOrcid(ctx, req.ID, req.Orcid); err != nil {
		return nil, err
	}
	person, err := s.repository.GetPerson(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (s *Service) SetPersonOrcidToken(ctx context.Context, req *SetPersonOrcidTokenRequest) (*Person, error) {
	if err := s.repository.SetPersonOrcidToken(ctx, req.ID, req.OrcidToken); err != nil {
		return nil, err
	}
	person, err := s.repository.GetPerson(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (s *Service) SetPersonRole(ctx context.Context, req *SetPersonRoleRequest) (*Person, error) {
	if err := s.repository.SetPersonRole(ctx, req.ID, req.Role); err != nil {
		return nil, err
	}
	person, err := s.repository.GetPerson(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (s *Service) SetPersonSettings(ctx context.Context, req *SetPersonSettingsRequest) (*Person, error) {
	if req.Settings == nil {
		return nil, fmt.Errorf("%w: attribute settings is missing in request body", models.ErrMissingArgument)
	}
	if err := s.repository.SetPersonSettings(ctx, req.ID, req.Settings); err != nil {
		return nil, err
	}
	person, err := s.repository.GetPerson(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
}

func (s *Service) GetOrganization(ctx context.Context, req *GetOrganizationRequest) (*Organization, error) {
	org, err := s.repository.GetOrganization(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return mapToExternalOrganization(org), nil
}

func (s *Service) GetOrganizations(ctx context.Context, req *GetOrganizationsRequest) (*OrganizationListResponse, error) {
	var organizations []*models.Organization
	var err error
	var cursor string

	if req.Cursor != "" {
		organizations, cursor, err = s.repository.GetMoreOrganizations(ctx, req.Cursor)
	} else {
		organizations, cursor, err = s.repository.GetOrganizations(ctx)
	}
	if err != nil {
		return nil, err
	}

	res := &OrganizationListResponse{
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

func (s *Service) SuggestOrganizations(ctx context.Context, req *SuggestOrganizationsRequest) (*OrganizationListResponse, error) {
	orgs, err := s.repository.SuggestOrganizations(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	res := &OrganizationListResponse{
		Data: make([]Organization, 0, len(orgs)),
	}
	for _, org := range orgs {
		res.Data = append(res.Data, *mapToExternalOrganization(org))
	}

	return res, nil
}

func (s *Service) NewError(ctx context.Context, err error) *ErrorStatusCode {
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

func mapToExternalPerson(person *models.Person) *Person {
	p := &Person{}
	p.ID = person.Id
	p.Active = person.Active
	if person.BirthDate != "" {
		p.BirthDate = NewOptString(person.BirthDate)
	}
	p.DateCreated = *person.DateCreated
	p.DateUpdated = *person.DateUpdated
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
			DateCreated: *orgRef.DateCreated,
			DateUpdated: *orgRef.DateUpdated,
			From:        *orgRef.From,
		}
		if orgRef.Until != nil {
			oRef.Until = NewOptDateTime(*orgRef.Until)
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

func mapToExternalOrganization(org *models.Organization) *Organization {
	o := &Organization{}
	o.ID = org.Id
	if org.GismoId != "" {
		o.GismoID = NewOptString(org.GismoId)
	}
	o.DateCreated = *org.DateCreated
	o.DateUpdated = *org.DateUpdated
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
