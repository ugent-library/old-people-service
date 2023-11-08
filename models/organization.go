package models

import (
	"sort"
	"time"
)

type Organization struct {
	ID          string               `json:"id,omitempty"`
	DateCreated *time.Time           `json:"date_created,omitempty"`
	DateUpdated *time.Time           `json:"date_updated,omitempty"`
	Type        string               `json:"type,omitempty"`
	NameDut     string               `json:"name_dut,omitempty"`
	NameEng     string               `json:"name_eng,omitempty"`
	Parent      []OrganizationParent `json:"parent,omitempty"`
	Identifier  []Identifier         `json:"identifier,omitempty"`
	Acronym     string               `json:"acronym,omitempty"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func NewOrganization() *Organization {
	org := &Organization{}
	return org
}

func (org *Organization) AddIdentifier(propertyID string, value string) {
	org.Identifier = append(org.Identifier, NewIdentifier(propertyID, value))
	sort.Sort(ByIdentifier(org.Identifier))
}

func (org *Organization) SetIdentifier(ids ...Identifier) {
	sort.Sort(ByIdentifier(ids))
	org.Identifier = ids
}

func (org *Organization) ClearIdentifier() {
	org.Identifier = nil
}

func (org *Organization) GetIdentifierValues(propertyID string) []string {
	vals := make([]string, 0, len(org.Identifier))
	for _, id := range org.Identifier {
		if id.PropertyID == propertyID {
			vals = append(vals, id.Value)
		}
	}
	return vals
}

func (org *Organization) GetIdentifierValue(propertyID string) string {
	vals := org.GetIdentifierValues(propertyID)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (org *Organization) SetParent(parents ...OrganizationParent) {
	sort.Sort(ByOrganizationParent(parents))
	org.Parent = parents
}

func (org *Organization) AddParent(parents ...OrganizationParent) {
	org.Parent = append(org.Parent, parents...)
	sort.Sort(ByOrganizationParent(org.Parent))
}

func (org *Organization) Dup() *Organization {
	newOrg := &Organization{
		ID:          org.ID,
		Type:        org.Type,
		NameDut:     org.NameDut,
		NameEng:     org.NameEng,
		Acronym:     org.Acronym,
		DateCreated: copyTime(org.DateCreated),
		DateUpdated: copyTime(org.DateUpdated),
	}

	for _, id := range org.Identifier {
		newOrg.Identifier = append(newOrg.Identifier, *id.Dup())
	}
	for _, op := range org.Parent {
		newOrg.Parent = append(newOrg.Parent, *op.Dup())
	}

	return newOrg
}
