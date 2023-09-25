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

func (s *Service) GetPersonByOtherId(ctx context.Context, req *GetPersonByOtherIdRequest) (*Person, error) {
	person, err := s.repository.GetPersonByAnyOtherId(ctx, string(req.Type), req.ID)
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

func (s *Service) GetOrganizationByOtherId(ctx context.Context, req *GetOrganizationByOtherIdRequest) (*Organization, error) {
	org, err := s.repository.GetOrganizationByAnyOtherId(ctx, string(req.Type), req.ID)
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

func (s *Service) AddPerson(ctx context.Context, p *Person) (*Person, error) {
	var person *models.Person

	if p.ID.Set {
		oldPerson, err := s.repository.GetPerson(ctx, p.ID.Value)
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("cannot find person record %s to update", p.ID.Value)
		} else if err != nil {
			return nil, err
		}
		person = oldPerson
	} else {
		person = models.NewPerson()
	}

	person.Active = p.Active.Value
	person.BirthDate = p.BirthDate.Value
	person.Email = p.Email.Value
	person.ExpirationDate = p.ExpirationDate.Value
	person.FirstName = p.FirstName.Value
	person.LastName = p.LastName.Value
	person.FullName = p.FullName.Value
	person.GismoId = p.GismoID.Value
	person.JobCategory = p.JobCategory
	person.ObjectClass = p.ObjectClass
	person.Orcid = p.Orcid.Value
	person.OrcidToken = p.OrcidToken.Value
	person.PreferredFirstName = p.PreferredFirstName.Value
	person.PreferredLastName = p.PreferredLastName.Value
	person.OtherId = models.IdRefs(p.OtherID.Value)
	person.Role = p.Role
	person.Settings = p.Settings.Value
	person.Title = p.Title.Value

	for _, orgRef := range p.Organization {
		newOrgRef := models.NewOrganizationRef(orgRef.ID)
		newOrgRef.From = &orgRef.From
		newOrgRef.Until = &orgRef.Until
		person.Organization = append(person.Organization, newOrgRef)
	}

	if newPerson, err := s.repository.SavePerson(ctx, person); err != nil {
		return nil, err
	} else {
		person = newPerson
	}

	return mapToExternalPerson(person), nil
}

func (s *Service) AddOrganization(ctx context.Context, o *Organization) (*Organization, error) {
	var org *models.Organization

	if o.ID.Set {
		oldOrg, err := s.repository.GetOrganization(ctx, o.ID.Value)
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("cannot find organization record \"%s\" to update", o.ID.Value)
		} else if err != nil {
			return nil, err
		}
		org = oldOrg
	} else {
		org = models.NewOrganization()
	}

	org.GismoId = o.GismoID.Value
	org.NameDut = o.NameDut.Value
	org.NameEng = o.NameEng.Value
	org.OtherId = models.IdRefs(o.OtherID.Value)
	org.ParentId = o.ParentID.Value
	org.Type = o.Type.Value

	if newOrg, err := s.repository.SaveOrganization(ctx, org); err != nil {
		return nil, err
	} else {
		org = newOrg
	}

	return mapToExternalOrganization(org), nil
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
	if errors.Is(err, models.ErrInvalidReference) {
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
	p.ID = NewOptString(person.Id)
	p.Active = NewOptBool(person.Active)
	if person.BirthDate != "" {
		p.BirthDate = NewOptString(person.BirthDate)
	}
	p.DateCreated = NewOptDateTime(*person.DateCreated)
	p.DateUpdated = NewOptDateTime(*person.DateUpdated)
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
			DateCreated: NewOptDateTime(*orgRef.DateCreated),
			DateUpdated: NewOptDateTime(*orgRef.DateUpdated),
			From:        *orgRef.From,
		}
		if orgRef.Until != nil {
			oRef.Until = *orgRef.Until
		}
		p.Organization = append(p.Organization, oRef)
	}
	if len(person.OtherId) > 0 {
		p.OtherID = NewOptIdRefs(IdRefs(person.OtherId))
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
	o.ID = NewOptString(org.Id)
	if org.GismoId != "" {
		o.GismoID = NewOptString(org.GismoId)
	}
	o.DateCreated = NewOptDateTime(*org.DateCreated)
	o.DateUpdated = NewOptDateTime(*org.DateUpdated)
	if org.NameDut != "" {
		o.NameDut = NewOptString(org.NameDut)
	}
	if org.NameEng != "" {
		o.NameEng = NewOptString(org.NameEng)
	}
	if len(org.OtherId) > 0 {
		o.OtherID = NewOptIdRefs(IdRefs(org.OtherId))
	}
	if org.ParentId != "" {
		o.ParentID = NewOptString(org.ParentId)
	}
	o.Type = NewOptString(org.Type)

	return o
}
