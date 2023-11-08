package models

import (
	"time"
)

type OrganizationMember struct {
	ID          string     `json:"id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	Until       *time.Time `json:"until,omitempty"`
}

func (om *OrganizationMember) Dup() *OrganizationMember {
	return &OrganizationMember{
		ID:          om.ID,
		DateCreated: copyTime(om.DateCreated),
		DateUpdated: copyTime(om.DateUpdated),
		From:        copyTime(om.From),
		Until:       copyTime(om.Until),
	}
}
