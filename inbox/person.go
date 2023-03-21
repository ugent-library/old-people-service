package inbox

import (
	"time"

	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/models"
)

/*
	See also https://github.com/ugent-library/soap-bridge/blob/main/main.go
*/

type Attribute struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
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
