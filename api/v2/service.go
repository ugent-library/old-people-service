package api

import (
	"context"
	"errors"

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

func (s *Service) GetPerson(ctx context.Context, params GetPersonParams) (*Person, error) {
	person, err := s.repository.GetPerson(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return mapToExternalPerson(person), nil
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
