package inbox

import (
	"fmt"
	"time"

	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
	"github.com/ugent-library/people/validation"
)

/*
	See also https://github.com/ugent-library/soap-bridge/blob/main/main.go
*/

type Attribute struct {
	Name      string     `json:"name"`
	Value     string     `json:"value"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type Message struct {
	ID         string      `json:"id,omitempty"`
	Language   string      `json:"language"`
	Attributes []Attribute `json:"attributes"`
}

type InboxMessage struct {
	Subject string   `json:"subject,omitempty"`
	Message *Message `json:"person_message"`
}

type InboxErrorMessage struct {
	InboxMessage *InboxMessage       `json:"inbox_message,omitempty"`
	Errors       []*validation.Error `json:"errors,omitempty"`
}

type PersonChangeError struct {
	OldPerson *models.Person      `json:"old_person,omitempty"`
	NewPerson *models.Person      `json:"new_person,omitempty"`
	Errors    []*validation.Error `json:"errors,omitempty"`
}

func (s *InboxMessage) UpdatePersonAttr(person *models.Person) *models.Person {

	person.ID = s.Message.ID

	now := time.Now()

	for _, attr := range s.Message.Attributes {

		if !(attr.StartDate.Before(now) && attr.EndDate.After(now)) {
			continue
		}

		switch attr.Name {
		case "first_name":
			person.FirstName = attr.Value
		case "last_name":
			person.LastName = attr.Value
		case "ugent_id":
			person.OtherID = append(person.OtherID, schema.IdRef{
				Type: "ugent_id",
				ID:   attr.Value,
			})
		case "title":
			person.JobTitle = attr.Value
		case "organization_id":
			person.OrganizationID = append(person.OrganizationID, attr.Value)
		}
	}

	// cleanup other_id
	{
		values := make([]schema.IdRef, 0, len(person.OtherID))
		for _, id := range person.OtherID {
			var found bool = false
			for _, val := range values {
				if val.ID == id.ID && val.Type == id.Type {
					found = true
					break
				}
			}
			if !found {
				values = append(values, id)
			}
		}
		person.OtherID = values
	}

	// cleanup organization_id
	{
		values := make([]string, 0, len(person.OrganizationID))
		for _, oid := range person.OrganizationID {
			var found bool = false
			for _, val := range values {
				if oid == val {
					found = true
					break
				}
			}
			if !found {
				values = append(values, oid)
			}
		}
		person.OrganizationID = values
	}

	return person
}

func (s *InboxMessage) Validate() validation.Errors {
	var errs validation.Errors

	if s.Subject == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/subject",
			Code:    "subject.required",
		})
	}
	if s.Message == nil {
		errs = append(errs, &validation.Error{
			Pointer: "/message",
			Code:    "message.required",
		})
	} else if mErrs := s.Message.Validate(); mErrs != nil {
		for _, err := range mErrs {
			errs = append(errs, &validation.Error{
				Pointer: "/message" + err.Pointer,
				Code:    "message." + err.Code,
			})
		}
	}

	return errs
}

func (m *Message) Validate() validation.Errors {
	var errs validation.Errors

	if m.ID == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/id",
			Code:    "id.required",
		})
	}

	if m.Language == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/language",
			Code:    "language.required",
		})
	}

	for i, attr := range m.Attributes {
		if attrErrs := attr.Validate(); attrErrs != nil {
			for _, attrErr := range attrErrs {
				errs = append(errs, &validation.Error{
					Pointer: fmt.Sprintf("/attributes/%d%s", i, attrErr.Pointer),
					Code:    "attributes." + attrErr.Code,
				})
			}
		}
	}

	return errs
}

func (attr *Attribute) Validate() validation.Errors {
	var errs validation.Errors

	if attr.Name == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/name",
			Code:    "name.required",
		})
	}
	if attr.Value == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/value",
			Code:    "value.required",
		})
	}

	if attr.StartDate == nil {
		errs = append(errs, &validation.Error{
			Pointer: "/start_date",
			Code:    "start_date.required",
		})
	}

	if attr.EndDate == nil {
		errs = append(errs, &validation.Error{
			Pointer: "/end_date",
			Code:    "end_date.required",
		})
	}
	return errs
}
