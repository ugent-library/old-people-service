package models

import (
	"fmt"

	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/ent/schema"
	"github.com/ugent-library/people/validation"
)

type Organization struct {
	v1.Organization
	OtherParentId string `json:"-"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func (org *Organization) Validate() validation.Errors {
	var errs validation.Errors

	for i, otherId := range org.GetOtherId() {
		if otherId.Id == "" {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/id", i),
				Code:    "other_id.id.required",
			})
		}
		if otherId.Type == "" {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/type", i),
				Code:    "other_id.type.required",
			})
		} else if !validation.InArray(schema.OrganizationIdTypes, otherId.Type) {
			errs = append(errs, &validation.Error{
				Pointer: fmt.Sprintf("/other_id/%d/type", i),
				Code:    "other_id.type.invalid",
			})
		}
	}

	return errs
}

func (org *Organization) Dup() *Organization {
	// *org copies mutex values too..
	no := &Organization{
		Organization: v1.Organization{
			Id:       org.Id,
			Type:     org.Type,
			NameDut:  org.NameDut,
			NameEng:  org.NameEng,
			ParentId: org.ParentId,
		},
	}

	if org.DateCreated != nil {
		t := *org.DateCreated
		no.DateCreated = &t
	}
	if org.DateUpdated != nil {
		t := *org.DateUpdated
		no.DateUpdated = &t
	}
	no.OtherId = make([]*v1.IdRef, 0, len(org.OtherId))
	for _, oldIdRef := range org.OtherId {
		no.OtherId = append(no.OtherId, &v1.IdRef{
			Id:   oldIdRef.Id,
			Type: oldIdRef.Type,
		})
	}

	return no
}
