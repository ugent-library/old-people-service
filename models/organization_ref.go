package models

import (
	"time"
)

type OrganizationRef struct {
	Id          string     `json:"id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	From        *time.Time `json:"from,omitempty"`
	Until       *time.Time `json:"until,omitempty"`
}
