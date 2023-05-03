package models

import "context"

type PersonSuggestService interface {
	SuggestPerson(context.Context, string) ([]*Person, error)
}
