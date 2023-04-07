package models

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

type PersonService interface {
	Create(context.Context, *Person) (*Person, error)
	Update(context.Context, *Person) (*Person, error)
	Get(context.Context, string) (*Person, error)
	Delete(context.Context, string) error
	Each(context.Context, func(*Person) bool) error
}
