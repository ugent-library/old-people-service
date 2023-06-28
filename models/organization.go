package models

import (
	v1 "github.com/ugent-library/person-service/api/v1"
)

type Organization struct {
	*v1.Organization
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func NewOrganization() *Organization {
	return &Organization{
		Organization: &v1.Organization{},
	}
}
