package models

import "time"

type Organization struct {
	Id          string     `json:"id,omitempty"`
	GismoId     string     `json:"gismo_id,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Type        string     `json:"type,omitempty"`
	NameDut     string     `json:"name_dut,omitempty"`
	NameEng     string     `json:"name_eng,omitempty"`
	ParentId    string     `json:"parent_id,omitempty"`
	OtherId     IdRefs     `json:"other_id,omitempty"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func NewOrganization() *Organization {
	org := &Organization{}
	org.OtherId = IdRefs{}
	return org
}
