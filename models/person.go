package models

import (
	"sort"
	"strings"
	"time"
)

type Person struct {
	ID                  string               `json:"id,omitempty"`
	Active              bool                 `json:"active,omitempty"`
	DateCreated         *time.Time           `json:"date_created,omitempty"`
	DateUpdated         *time.Time           `json:"date_updated,omitempty"`
	Name                string               `json:"name,omitempty"`
	GivenName           string               `json:"given_name,omitempty"`
	FamilyName          string               `json:"family_name,omitempty"`
	Email               string               `json:"email,omitempty"`
	Token               []URN                `json:"token"`
	PreferredGivenName  string               `json:"preferred_given_name,omitempty"`
	PreferredFamilyName string               `json:"preferred_family_name,omitempty"`
	BirthDate           string               `json:"birth_date,omitempty"`
	HonorificPrefix     string               `json:"honorific_prefix,omitempty"`
	Identifier          []URN                `json:"identifier,omitempty"`
	Organization        []OrganizationMember `json:"organization,omitempty"`
	JobCategory         []string             `json:"job_category,omitempty"`
	Role                []string             `json:"role,omitempty"`
	Settings            map[string]string    `json:"settings,omitempty"`
	ObjectClass         []string             `json:"object_class,omitempty"`
	ExpirationDate      string               `json:"expiration_date,omitempty"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func NewPerson() *Person {
	p := &Person{}
	return p
}

func NewOrganizationMember(id string) OrganizationMember {
	return OrganizationMember{
		ID:    id,
		From:  &BeginningOfTime,
		Until: nil,
	}
}

func (p *Person) SetEmail(email string) {
	p.Email = strings.ToLower(email)
}

func (p *Person) AddIdentifier(urn URN) {
	p.Identifier = append(p.Identifier, urn)
	sort.Sort(ByURN(p.Identifier))
}

func (p *Person) SetIdentifier(ids ...URN) {
	sort.Sort(ByURN(ids))
	p.Identifier = ids
}

func (p *Person) ClearIdentifier() {
	p.Identifier = nil
}

func (p *Person) GetIdentifierQualifiedValues() []string {
	ids := make([]string, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		ids = append(ids, id.String())
	}
	return ids
}

func (p *Person) GetIdentifierValues() []string {
	ids := make([]string, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		ids = append(ids, id.Value)
	}
	return ids
}

func (p *Person) GetIdentifierByNS(ns string) []URN {
	urns := []URN{}
	for _, id := range p.Identifier {
		if id.Namespace == ns {
			urns = append(urns, id)
		}
	}
	return urns
}

func (p *Person) GetIdentifierValuesByNS(ns string) []string {
	vals := make([]string, 0, len(p.Identifier))
	for _, id := range p.Identifier {
		if id.Namespace == ns {
			vals = append(vals, id.Value)
		}
	}
	return vals
}

func (p *Person) AddToken(ns string, value string) {
	p.Token = append(p.Token, NewURN(ns, value))
	sort.Sort(ByURN(p.Token))
}

func (p *Person) SetToken(tokens ...URN) {
	sort.Sort(ByURN(tokens))
	p.Token = tokens
}

func (p *Person) ClearToken() {
	p.Token = nil
}

func (p *Person) GetTokenValues(ns string) []string {
	vals := make([]string, 0, len(p.Token))
	for _, token := range p.Token {
		if token.Namespace == ns {
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

func (p *Person) AddOrganizationMember(orgMembers ...OrganizationMember) {
	p.Organization = append(p.Organization, orgMembers...)
	sort.Sort(ByOrganizationMember(p.Organization))
}

func (p *Person) SetOrganizationMember(orgMembers ...OrganizationMember) {
	sort.Sort(ByOrganizationMember(orgMembers))
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
		newP.AddIdentifier(NewURN(id.Namespace, id.Value))
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
