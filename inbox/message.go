package inbox

import (
	"fmt"
	"time"

	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/validation"
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
	Date       string      `json:"date,omitempty"`
	Language   string      `json:"language"`
	Attributes []Attribute `json:"attributes"`
	Source     string      `json:"source"`
}

func (m *Message) Validate() validation.Errors {
	var errs validation.Errors

	if m.ID == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/id",
			Code:    "id.required",
		})
	}

	if m.Source == "" {
		errs = append(errs, &validation.Error{
			Pointer: "/subject",
			Code:    "subject.required",
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

func (m *Message) GetAttributesAt(t time.Time) []Attribute {
	attrs := make([]Attribute, 0, len(m.Attributes))
	for _, attr := range m.Attributes {
		if !attr.ValidAt(t) {
			continue
		}
		attrs = append(attrs, attr)
	}
	return attrs
}

func (m *Message) GetAttributeAt(name string, t time.Time) (string, error) {
	for _, attr := range m.GetAttributesAt(t) {
		if attr.Name == name {
			return attr.Value, nil
		}
	}
	return "", models.ErrNotFound
}

func (attr *Attribute) ValidAt(t time.Time) bool {
	return attr.StartDate.Before(t) && attr.EndDate.After(t)
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
