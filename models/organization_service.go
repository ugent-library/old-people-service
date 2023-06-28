package models

import "context"

type OrganizationService interface {
	CreateOrganization(context.Context, *Organization) (*Organization, error)
	UpdateOrganization(context.Context, *Organization) (*Organization, error)
	GetOrganization(context.Context, string) (*Organization, error)
	GetOrganizationByGismoId(context.Context, string) (*Organization, error)
	GetOrganizationsByGismoId(context.Context, ...string) ([]*Organization, error)
	DeleteOrganization(context.Context, string) error
	EachOrganization(context.Context, func(*Organization) bool) error
}
