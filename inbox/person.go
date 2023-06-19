package inbox

import (
	"strings"
	"time"

	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/models"
	"github.com/ugent-library/people/validation"
)

func (s *InboxMessage) UpdatePersonAttr(person *models.Person) *models.Person {

	person.Id = s.Message.ID

	now := time.Now()

	// clear previous values
	person.BirthDate = ""
	person.Email = ""
	person.FirstName = ""
	person.LastName = ""
	person.FullName = ""
	person.Orcid = ""
	person.OrcidToken = ""
	person.OtherId = make([]*v1.IdRef, 0)
	person.OtherOrganizationId = make([]string, 0)
	person.OrganizationId = make([]string, 0)
	person.JobCategory = make([]string, 0)
	person.PreferredFirstName = ""
	person.PreferredLastName = ""

	for _, attr := range s.Message.Attributes {

		if !(attr.StartDate.Before(now) && attr.EndDate.After(now)) {
			continue
		}

		switch attr.Name {
		case "birth_date":
			person.BirthDate = attr.Value
		case "email":
			person.Email = strings.ToLower(attr.Value)
		case "first_name":
			person.FirstName = attr.Value
		case "full_name":
			person.FullName = attr.Value
		case "job_category":
			person.JobCategory = append(person.JobCategory, attr.Value)
		case "last_name":
			person.LastName = attr.Value
		case "ugent_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_id",
				Id:   attr.Value,
			})
		case "historic_ugent_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "historic_ugent_id",
				Id:   attr.Value,
			})
		case "ugent_barcode":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_barcode",
				Id:   attr.Value,
			})
		case "ugent_username":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_username",
				Id:   attr.Value,
			})
		case "ugent_memorialis_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_memorialis_id",
				Id:   attr.Value,
			})
		case "uzgent_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "uzgent_id",
				Id:   attr.Value,
			})
		case "title":
			person.Title = attr.Value
		case "organization_id":
			person.OrganizationId = append(person.OrganizationId, attr.Value)
		case "preferred_first_name":
			person.PreferredFirstName = attr.Value
		case "preferred_last_name":
			person.PreferredLastName = attr.Value
		}
	}

	// TODO: cleanup job_category
	person.JobCategory = validation.Uniq(person.JobCategory)

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
	// person.OtherOrganizationId is filled during storing
	person.OrganizationId = validation.Uniq(person.OrganizationId)

	return person
}
