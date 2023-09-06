package models

import (
	"time"
)

type Person struct {
	Id                 string             `json:"id,omitempty"`
	GismoId            string             `json:"gismo_id,omitempty"`
	Active             bool               `json:"active,omitempty"`
	DateCreated        *time.Time         `json:"date_created,omitempty"`
	DateUpdated        *time.Time         `json:"date_updated,omitempty"`
	FullName           string             `json:"full_name,omitempty"`
	FirstName          string             `json:"first_name,omitempty"`
	LastName           string             `json:"last_name,omitempty"`
	Email              string             `json:"email,omitempty"`
	Orcid              string             `json:"orcid,omitempty"`
	OrcidToken         string             `json:"orcid_token,omitempty"`
	PreferredFirstName string             `json:"preferred_first_name,omitempty"`
	PreferredLastName  string             `json:"preferred_last_name,omitempty"`
	BirthDate          string             `json:"birth_date,omitempty"`
	Title              string             `json:"title,omitempty"`
	OtherId            IdRefs             `json:"other_id,omitempty"`
	Organization       []*OrganizationRef `json:"organization,omitempty"`
	JobCategory        []string           `json:"job_category,omitempty"`
	Role               []string           `json:"role,omitempty"`
	Settings           map[string]string  `json:"settings,omitempty"`
	ObjectClass        []string           `json:"object_class,omitempty"`
	ExpirationDate     string             `json:"expiration_date,omitempty"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func NewPerson() *Person {
	p := &Person{}
	p.OtherId = IdRefs{}
	return p
}

func NewOrganizationRef(id string) *OrganizationRef {
	return &OrganizationRef{
		Id:    id,
		From:  &BeginningOfTime,
		Until: &EndOfTime,
	}
}
