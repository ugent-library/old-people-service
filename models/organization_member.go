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

func (om OrganizationMember) Dup() OrganizationMember {
	return OrganizationMember{
		ID:          om.ID,
		DateCreated: copyTime(om.DateCreated),
		DateUpdated: copyTime(om.DateUpdated),
		From:        copyTime(om.From),
		Until:       copyTime(om.Until),
	}
}

type ByOrganizationMember []OrganizationMember

func (orgMembers ByOrganizationMember) Len() int {
	return len(orgMembers)
}
func (orgMembers ByOrganizationMember) Swap(i, j int) {
	orgMembers[i], orgMembers[j] = orgMembers[j], orgMembers[i]
}
func (orgMembers ByOrganizationMember) Less(i, j int) bool {
	if !orgMembers[i].From.Equal(*orgMembers[j].From) {
		return orgMembers[i].From.Before(*orgMembers[j].From)
	}
	return orgMembers[i].ID < orgMembers[j].ID
}
