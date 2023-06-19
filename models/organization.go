package models

import (
	v1 "github.com/ugent-library/people/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Organization struct {
	*v1.Organization
	OtherParentId string `json:"-"`
}

func (org *Organization) IsStored() bool {
	return org.DateCreated != nil
}

func (org *Organization) Dup() *Organization {
	// *org copies mutex values too..
	no := &Organization{
		Organization: &v1.Organization{
			Id:       org.Id,
			Type:     org.Type,
			NameDut:  org.NameDut,
			NameEng:  org.NameEng,
			ParentId: org.ParentId,
		},
	}

	if org.DateCreated != nil {
		no.DateCreated = timestamppb.New(org.DateCreated.AsTime())
	}
	if org.DateUpdated != nil {
		no.DateUpdated = timestamppb.New(org.DateUpdated.AsTime())
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
