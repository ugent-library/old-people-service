package models

import (
	"fmt"

	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/validation"
)

type Person struct {
	v1.Person
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func (person *Person) Validate() validation.Errors {
	var errs validation.Errors

	for i, otherId := range person.GetOtherId() {
		if otherId.Id == "" {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/id", i),
				Code:    "other_id.id.required",
			})
		}
		if otherId.Type == "" {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/type", i),
				Code:    "other_id.type.required",
			})
		} else if !validation.InArray(schema.IdRefTypes, otherId.Type) {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/type", i),
				Code:    "other_id.type.invalid",
			})
		}
	}

	return errs
}

func (p *Person) Dup() *Person {
	// *p copies mutex values too..
	np := &Person{
		Person: v1.Person{
			Id:                 p.Id,
			Active:             p.Active,
			FullName:           p.FullName,
			FirstName:          p.FirstName,
			LastName:           p.LastName,
			Email:              p.Email,
			Orcid:              p.Orcid,
			OrcidToken:         p.OrcidToken,
			JobCategory:        append([]string{}, p.JobCategory...),
			OrganizationId:     append([]string{}, p.OrganizationId...),
			PreferredFirstName: p.PreferredFirstName,
			PreferredLastName:  p.PreferredLastName,
			BirthDate:          p.BirthDate,
			Title:              p.Title,
		},
	}

	if p.DateCreated != nil {
		t := *p.DateCreated
		np.DateCreated = &t
	}
	if p.DateUpdated != nil {
		t := *p.DateUpdated
		np.DateUpdated = &t
	}
	np.OtherId = make([]*v1.IdRef, 0, len(p.OtherId))
	for _, oldIdRef := range p.GetOtherId() {
		np.OtherId = append(np.OtherId, &v1.IdRef{
			Id:   oldIdRef.Id,
			Type: oldIdRef.Type,
		})
	}

	return np
}
