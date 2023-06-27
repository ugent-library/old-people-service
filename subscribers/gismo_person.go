package subscribers

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/inbox"
	"github.com/ugent-library/person-service/models"
)

type GismoPersonSubscriber struct {
	BaseSubscriber
	personService models.PersonService
}

func NewGismoPersonSubscriber(subject string, personService models.PersonService, subOpts ...nats.SubOpt) *GismoPersonSubscriber {
	sub := &GismoPersonSubscriber{
		BaseSubscriber: NewBaseSubscriber(subject),
		personService:  personService,
	}
	sub.subOpts = append(sub.subOpts, subOpts...)
	return sub
}

func (ps *GismoPersonSubscriber) Listen(msg *nats.Msg) (*inbox.Message, error) {

	ctx := context.Background()
	iMsg, err := gismo.ParsePersonMessage(msg.Data)

	if err != nil {
		return nil, fmt.Errorf("%w: unable to process malformed message: %s", models.ErrNonFatal, err)
	}

	person, err := ps.personService.GetPerson(ctx, iMsg.ID)

	if err != nil && err == models.ErrNotFound {
		person = models.NewPerson()
	} else if err != nil {
		return iMsg, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	if iMsg.Source == "gismo.person.update" {
		iMsg.UpdatePersonAttr(person)
		person.Active = true
	} else if iMsg.Source == "gismo.person.delete" {
		person.Active = false
	}

	if person.IsStored() {
		_, err = ps.personService.UpdatePerson(ctx, person)
	} else {
		_, err = ps.personService.CreatePerson(ctx, person)
	}

	if err != nil {
		return iMsg, fmt.Errorf("%w: unable to store person record: %s", models.ErrFatal, err)
	}

	return iMsg, nil
}
