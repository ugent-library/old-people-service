package models

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")
var ErrFatal = errors.New("fatal error")
var ErrNonFatal = errors.New("non fatal error")

type PersonService interface {
	CreatePerson(context.Context, *Person) (*Person, error)
	UpdatePerson(context.Context, *Person) (*Person, error)
	GetPerson(context.Context, string) (*Person, error)
	DeletePerson(context.Context, string) error
	EachPerson(context.Context, func(*Person) bool) error
	SetOrcidToken(context.Context, string, string) error
	SetOrcid(context.Context, string, string) error
	SetRole(context.Context, string, []string) error
	SetSettings(context.Context, string, map[string]string) error
}
