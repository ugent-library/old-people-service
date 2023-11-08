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
	sort.Slice(org.Identifier, func(i, j int) bool {
		if org.Identifier[i].PropertyID != org.Identifier[j].PropertyID {
			return org.Identifier[i].PropertyID < org.Identifier[j].PropertyID
		}
		return org.Identifier[i].Value < org.Identifier[j].Value
	})
}

func (org *Organization) SetIdentifier(ids ...Identifier) {
	sort.Slice(ids, func(i, j int) bool {
		if ids[i].PropertyID != ids[j].PropertyID {
			return ids[i].PropertyID < ids[j].PropertyID
		}
		return ids[i].Value < ids[j].Value
	})
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
	sort.Slice(parents, func(i, j int) bool {
		if !parents[i].From.Equal(*parents[j].From) {
			return parents[i].From.Before(*parents[j].From)
		}
		return parents[i].ID < parents[j].ID
	})
	org.Parent = parents
}

func (org *Organization) AddParent(parents ...OrganizationParent) {
	org.Parent = append(org.Parent, parents...)
	sort.Slice(org.Parent, func(i, j int) bool {
		if !org.Parent[i].From.Equal(*org.Parent[j].From) {
			return org.Parent[i].From.Before(*org.Parent[j].From)
		}
		return org.Parent[i].ID < org.Parent[j].ID
	})
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
