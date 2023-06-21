package inbox

import (
	"time"

	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/models"
)

func (s *InboxMessage) UpdateOrganizationAttr(org *models.Organization) *models.Organization {

	org.Id = s.Message.ID

	//clear previous values
	org.NameDut = ""
	org.NameEng = ""
	org.OtherId = make([]*v1.IdRef, 0)
	org.Type = "organization"
	org.ParentId = ""

	now := time.Now()

	for _, attr := range s.Message.Attributes {

		if !(attr.StartDate.Before(now) && attr.EndDate.After(now)) {
			continue
		}

		switch attr.Name {
		case "name_dut":
			org.NameDut = attr.Value
		case "name_eng":
			org.NameEng = attr.Value
		case "type":
			org.Type = attr.Value
		case "parent_id":
			org.ParentId = attr.Value
		case "ugent_memorialis_id":
			org.OtherId = append(org.OtherId, &v1.IdRef{
				Type: "ugent_memorialis_id",
				Id:   attr.Value,
			})
		case "code":
			org.OtherId = append(org.OtherId, &v1.IdRef{
				Type: "ugent_id",
				Id:   attr.Value,
			})
		case "biblio_code":
			org.OtherId = append(org.OtherId, &v1.IdRef{
				Type: "biblio_id",
				Id:   attr.Value,
			})
		}
	}

	return org
}
