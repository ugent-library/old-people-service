package models

import (
	"time"
)

type Person struct {
	ID                  string                `json:"id,omitempty"`
	Active              bool                  `json:"active,omitempty"`
	DateCreated         *time.Time            `json:"date_created,omitempty"`
	DateUpdated         *time.Time            `json:"date_updated,omitempty"`
	Name                string                `json:"name,omitempty"`
	GivenName           string                `json:"given_name,omitempty"`
	FamilyName          string                `json:"family_name,omitempty"`
	Email               string                `json:"email,omitempty"`
	Token               []Token               `json:"token"`
	PreferredGivenName  string                `json:"preferred_given_name,omitempty"`
	PreferredFamilyName string                `json:"preferred_family_name,omitempty"`
	BirthDate           string                `json:"birth_date,omitempty"`
	HonorificPrefix     string                `json:"honorific_prefix,omitempty"`
	Identifier          []Identifier          `json:"identifier,omitempty"`
	Organization        []*OrganizationMember `json:"organization,omitempty"`
	JobCategory         []string              `json:"job_category,omitempty"`
	Role                []string              `json:"role,omitempty"`
	Settings            map[string]string     `json:"settings,omitempty"`
	ObjectClass         []string              `json:"object_class,omitempty"`
	ExpirationDate      string                `json:"expiration_date,omitempty"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func NewPerson() *Person {
	p := &Person{}
	return p
}

func NewOrganizationMember(id string) *OrganizationMember {
	return &OrganizationMember{
		Id:    id,
		From:  &BeginningOfTime,
		Until: nil,
	}
}

func (p *Person) AddIdentifier(propertyID string, value string) {
	p.Identifier = append(p.Identifier, NewIdentifier(propertyID, value))
}

func (p *Person) ClearIdentifier() {
	p.Identifier = nil
}

func (p *Person) RemoveIdentifierByPropertyID(propertyID string) {
	var newIds []Identifier
	for _, id := range p.Identifier {
		if id.PropertyID != propertyID {
			newIds = append(newIds, id)
		}
	}
	p.Identifier = newIds
}

func (p *Person) GetIdentifierValue(propertyID string) string {
	for _, id := range p.Identifier {
		if id.PropertyID == propertyID {
			return id.PropertyID
		}
	}
	return ""
}

func (p *Person) GetIdentifierValues(propertyID string) []string {
	vals := make([]string, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		if id.PropertyID == propertyID {
			vals = append(vals, id.Value)
		}
	}
	return vals
}

func (p *Person) AddToken(propertyID string, value string) {
	p.Token = append(p.Token, NewToken(propertyID, value))
}

func (p *Person) ClearToken() {
	p.Token = nil
}

func (p *Person) GetTokenValues(propertyID string) []string {
	vals := make([]string, 0, len(p.Token))
	for _, token := range p.Token {
		if token.PropertyID == propertyID {
			vals = append(vals, token.Value)
		}
	}
	return vals
}
