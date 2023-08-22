package models

import (
	v1 "github.com/ugent-library/people-service/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func NewOrganizationRef(id string) *v1.OrganizationRef {
	return &v1.OrganizationRef{
		Id:    id,
		From:  timestamppb.New(BeginningOfTime),
		Until: timestamppb.New(EndOfTime),
	}
}
