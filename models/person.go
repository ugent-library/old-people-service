package models

import (
	v1 "github.com/ugent-library/people/api/v1"
)

type Person struct {
	*v1.Person
	// unconfirmed organization identifiers
	OtherOrganizationId []string `json:"-"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}
