package models

import (
	"sort"
	"strings"
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
		ID:    id,
		From:  &BeginningOfTime,
		Until: nil,
	}
}

func (p *Person) SetEmail(email string) {
	p.Email = strings.ToLower(email)
}

func (p *Person) AddIdentifier(propertyID string, value string) {
	p.Identifier = append(p.Identifier, NewIdentifier(propertyID, value))
	sort.Slice(p.Identifier, func(i, j int) bool {
		return p.Identifier[i].PropertyID < p.Identifier[j].PropertyID ||
			p.Identifier[i].Value < p.Identifier[j].Value
	})
}

func (p *Person) SetIdentifier(ids ...Identifier) {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].PropertyID < ids[j].PropertyID ||
			ids[i].Value < ids[j].Value
	})
	p.Identifier = ids
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

func (p *Person) HasIdentifier(propertyID string, value string) bool {
	for _, id := range p.Identifier {
		if id.PropertyID == propertyID && id.Value == value {
			return true
		}
	}
	return false
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
	sort.Slice(p.Token, func(i, j int) bool {
		return p.Token[i].PropertyID < p.Token[j].PropertyID ||
			p.Token[i].Value < p.Token[j].Value
	})
}

func (p *Person) SetToken(tokens ...Token) {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].PropertyID < tokens[j].PropertyID ||
			tokens[i].Value < tokens[j].Value
	})
	p.Token = tokens
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

func (p *Person) SetRole(role ...string) {
	sort.Strings(role)
	p.Role = role
}

func (p *Person) AddRole(role ...string) {
	p.Role = append(p.Role, role...)
	sort.Strings(p.Role)
}

func (p *Person) SetObjectClass(objectClass ...string) {
	sort.Strings(objectClass)
	p.ObjectClass = objectClass
}

func (p *Person) AddObjectClass(objectClass ...string) {
	p.ObjectClass = append(p.ObjectClass, objectClass...)
	sort.Strings(p.ObjectClass)
}

func (p *Person) SetJobCategory(jobCategory ...string) {
	sort.Strings(jobCategory)
	p.JobCategory = jobCategory
}

func (p *Person) AddJobCategory(jobCategory ...string) {
	p.JobCategory = append(p.JobCategory, jobCategory...)
	sort.Strings(p.JobCategory)
}

func (p *Person) AddOrganizationMember(orgMembers ...*OrganizationMember) {
	p.Organization = append(p.Organization, orgMembers...)
	sort.Slice(p.Organization, func(i, j int) bool {
		return p.Organization[i].From.Before(*p.Organization[j].From) ||
			p.Organization[i].ID < p.Organization[j].ID
	})
}

func (p *Person) SetOrganizationMember(orgMembers ...*OrganizationMember) {
	sort.Slice(orgMembers, func(i, j int) bool {
		return orgMembers[i].From.Before(*orgMembers[j].From) ||
			orgMembers[i].ID < orgMembers[j].ID
	})
	p.Organization = orgMembers
}

func (p *Person) Dup() *Person {
	newP := &Person{
		ID:                  p.ID,
		DateCreated:         copyTime(p.DateCreated),
		DateUpdated:         copyTime(p.DateUpdated),
		Active:              p.Active,
		Name:                p.Name,
		GivenName:           p.GivenName,
		FamilyName:          p.FamilyName,
		Email:               p.Email,
		PreferredGivenName:  p.PreferredGivenName,
		PreferredFamilyName: p.PreferredFamilyName,
		BirthDate:           p.BirthDate,
		HonorificPrefix:     p.HonorificPrefix,
		ExpirationDate:      p.ExpirationDate,
	}
	for _, token := range p.Token {
		newP.Token = append(newP.Token, *token.Dup())
	}
	for _, id := range p.Identifier {
		newP.AddIdentifier(id.PropertyID, id.Value)
	}
	for _, orgMember := range p.Organization {
		newP.Organization = append(newP.Organization, orgMember.Dup())
	}
	if p.Settings != nil {
		newP.Settings = make(map[string]string)
		for key, val := range p.Settings {
			newP.Settings[key] = val
		}
	}
	if p.JobCategory != nil {
		newP.JobCategory = append(newP.JobCategory, p.JobCategory...)
	}
	if p.ObjectClass != nil {
		newP.ObjectClass = append(newP.ObjectClass, p.ObjectClass...)
	}
	if p.Role != nil {
		newP.Role = append(newP.Role, p.Role...)
	}

	return newP
}
