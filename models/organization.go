package models

import "time"

type Organization struct {
	ID          string       `json:"id,omitempty"`
	DateCreated *time.Time   `json:"date_created,omitempty"`
	DateUpdated *time.Time   `json:"date_updated,omitempty"`
	Type        string       `json:"type,omitempty"`
	NameDut     string       `json:"name_dut,omitempty"`
	NameEng     string       `json:"name_eng,omitempty"`
	ParentID    string       `json:"parent_id,omitempty"`
	Identifier  []Identifier `json:"identifier,omitempty"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func NewOrganization() *Organization {
	org := &Organization{}
	return org
}

func (org *Organization) AddIdentifier(typ string, val string) {
	org.Identifier = append(org.Identifier, NewIdentifier(typ, val))
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
