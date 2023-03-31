package models

import (
	"fmt"
	"time"

	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/validation"
)

type Person struct {
	Active             bool           `json:"active"`
	BirthDate          string         `json:"birth_date,omitempty"`
	DateCreated        *time.Time     `json:"date_created"`
	DateUpdated        *time.Time     `json:"date_updated"`
	Email              string         `json:"email,omitempty"`
	OtherID            []schema.IdRef `json:"other_id,omitempty"`
	FirstName          string         `json:"first_name,omitempty"`
	FullName           string         `json:"full_name,omitempty"`
	ID                 string         `json:"id,omitempty"`
	LastName           string         `json:"last_name,omitempty"`
	JobCategory        []string       `json:"job_category,omitempty"`
	Orcid              string         `json:"orcid,omitempty"`
	OrcidToken         string         `json:"orcid_token,omitempty"`
	OrganizationID     []string       `json:"organization_id,omitempty"`
	PreferedLastName   string         `json:"preferred_last_name,omitempty"`
	PreferredFirstName string         `json:"preferred_first_name,omitempty"`
	Title              string         `json:"title,omitempty"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func (person *Person) Validate() validation.Errors {
	var errs validation.Errors

	for i, otherId := range person.OtherID {
		idErrs := otherId.Validate()
		for _, idErr := range idErrs {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d%s", i, idErr.Pointer),
				Code:    "other_id." + idErr.Code,
			})
		}
	}

	return errs
}

func (person *Person) Dup() *Person {
	np := *person

	if person.DateCreated != nil {
		t := *person.DateCreated
		np.DateCreated = &t
	}
	if person.DateUpdated != nil {
		t := *person.DateUpdated
		np.DateUpdated = &t
	}
	idRefs := make([]schema.IdRef, 0, len(person.OtherID))
	for _, idRef := range person.OtherID {
		idRefs = append(idRefs, idRef.Dup())
	}
	np.OtherID = idRefs
	np.JobCategory = append([]string{}, person.JobCategory...)
	np.OrganizationID = append([]string{}, person.OrganizationID...)

	return &np
}
