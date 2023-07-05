package models

import (
	"context"
)

type PersonService interface {
	SavePerson(context.Context, *Person) (*Person, error)
	CreatePerson(context.Context, *Person) (*Person, error)
	UpdatePerson(context.Context, *Person) (*Person, error)
	GetPerson(context.Context, string) (*Person, error)
	GetPersonByOtherId(context.Context, string, string) (*Person, error)
	GetPersonByGismoId(context.Context, string) (*Person, error)
	DeletePerson(context.Context, string) error
	EachPerson(context.Context, func(*Person) bool) error
	SetPersonOrcidToken(context.Context, string, string) error
	SetPersonOrcid(context.Context, string, string) error
	SetPersonRole(context.Context, string, []string) error
	SetPersonSettings(context.Context, string, map[string]string) error
}
