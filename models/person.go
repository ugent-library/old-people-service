package models

import (
	v1 "github.com/ugent-library/people/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Person struct {
	*v1.Person
	// unconfirmed organization identifiers
	OtherOrganizationId []string `json:"-"`
}

func (person *Person) IsStored() bool {
	return person.DateCreated != nil
}

func (p *Person) Dup() *Person {
	// *p copies mutex values too..
	np := &Person{
		Person: &v1.Person{
			Id:                 p.Id,
			Active:             p.Active,
			FullName:           p.FullName,
			FirstName:          p.FirstName,
			LastName:           p.LastName,
			Email:              p.Email,
			Orcid:              p.Orcid,
			OrcidToken:         p.OrcidToken,
			JobCategory:        append([]string{}, p.JobCategory...),
			OrganizationId:     append([]string{}, p.OrganizationId...),
			PreferredFirstName: p.PreferredFirstName,
			PreferredLastName:  p.PreferredLastName,
			BirthDate:          p.BirthDate,
			Title:              p.Title,
		},
	}

	if p.DateCreated != nil {
		np.DateCreated = timestamppb.New(p.DateCreated.AsTime())
	}
	if p.DateUpdated != nil {
		np.DateUpdated = timestamppb.New(p.DateUpdated.AsTime())
	}
	np.OtherId = make([]*v1.IdRef, 0, len(p.OtherId))
	for _, oldIdRef := range p.OtherId {
		np.OtherId = append(np.OtherId, &v1.IdRef{
			Id:   oldIdRef.Id,
			Type: oldIdRef.Type,
		})
	}

	return np
}
