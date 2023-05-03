package models

import "context"

type OrganizationSuggestService interface {
	SuggestOrganization(context.Context, string) ([]*Organization, error)
}
