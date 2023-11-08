package models

import "time"

type OrganizationParent struct {
	ID          string     `json:"id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	Until       *time.Time `json:"until,omitempty"`
}

func (op *OrganizationParent) Dup() *OrganizationParent {
	return &OrganizationParent{
		ID:          op.ID,
		DateCreated: copyTime(op.DateCreated),
		DateUpdated: copyTime(op.DateUpdated),
		From:        copyTime(op.From),
		Until:       copyTime(op.Until),
	}
}
