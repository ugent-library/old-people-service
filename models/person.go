package models

import (
	"time"

	"github.com/ugent-library/people/ent/schema"
)

type Person struct {
	Active      bool           `json:"active"`
	BirthDate   string         `json:"birth_date,omitempty"`
	DateCreated *time.Time     `json:"date_created"`
	DateUpdated *time.Time     `json:"date_updated"`
	Email       string         `json:"email,omitempty"`
	OtherID     []schema.IdRef `json:"other_id,omitempty"`
	FirstName   string         `json:"first_name,omitempty"`
	FullName    string         `json:"full_name,omitempty"`
	ID          string         `json:"id,omitempty"`
	LastName    string         `json:"last_name,omitempty"`
	Category    []string       `json:"category,omitempty"`
	Orcid       string         `json:"orcid,omitempty"`
	//TODO: to encrypt
	//TODO: mark as sensitive to prevent logging
	OrcidToken         string   `json:"orcid_token,omitempty"`
	OrganizationID     []string `json:"organization_id,omitempty"`
	PreferedLastName   string   `json:"preferred_last_name,omitempty"`
	PreferredFirstName string   `json:"preferred_first_name,omitempty"`
	JobTitle           string   `json:"job_title,omitempty"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}
