package models

import (
	"context"
)

type PersonService interface {
	SavePerson(context.Context, *Person) (*Person, error)
	CreatePerson(context.Context, *Person) (*Person, error)
	UpdatePerson(context.Context, *Person) (*Person, error)
	GetPerson(context.Context, string) (*Person, error)
	GetPersonByIdentifier(context.Context, ...URN) (*Person, error)
	GetPeopleByIdentifier(context.Context, ...URN) ([]*Person, error)
	DeletePerson(context.Context, string) error
	EachPerson(context.Context, func(*Person) bool) error
	SetPersonOrcidToken(context.Context, string, string) error
	SetPersonOrcid(context.Context, string, string) error
	SetPersonRole(context.Context, string, []string) error
	SetPersonSettings(context.Context, string, map[string]string) error
	AutoExpirePeople(context.Context) (int64, error)
	GetPeople(context.Context) ([]*Person, string, error)
	GetMorePeople(context.Context, string) ([]*Person, string, error)
}
