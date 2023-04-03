package inbox

import (
	"fmt"
	"time"

	v1 "github.com/ugent-library/people/api/v1"
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

	person.Id = s.Message.ID

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
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_id",
				Id:   attr.Value,
			})
		case "title":
			person.Title = attr.Value
		case "organization_id":
			person.OrganizationId = append(person.OrganizationId, attr.Value)
		}
	}

	// cleanup other_id
	{
		values := make([]*v1.IdRef, 0, len(person.OtherId))
		for _, id := range person.OtherId {
			var found bool = false
			for _, val := range values {
				if val.Id == id.Id && val.Type == id.Type {
					found = true
					break
				}
			}
			if !found {
				values = append(values, id)
			}
		}
		person.OtherId = values
	}

	// cleanup organization_id
	{
		values := make([]string, 0, len(person.OrganizationId))
		for _, oid := range person.OrganizationId {
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
		person.OrganizationId = values
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
