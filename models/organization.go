package models

import (
	v1 "github.com/ugent-library/people/api/v1"
)

type Organization struct {
	*v1.Organization
	OtherParentId string `json:"-"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}
