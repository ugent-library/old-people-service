package models

import (
	v1 "github.com/ugent-library/person-service/api/v1"
)

type Person struct {
	*v1.Person
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func NewPerson() *Person {
	return &Person{
		Person: &v1.Person{},
	}
}
