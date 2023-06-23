package subscribers

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/inbox"
	"github.com/ugent-library/person-service/models"
)

type GismoOrganizationSubscriber struct {
	BaseSubscriber
	organizationService models.OrganizationService
}

func NewGismoOrganizationSubscriber(subject string, organizationService models.OrganizationService, subOpts ...nats.SubOpt) *GismoOrganizationSubscriber {
	os := &GismoOrganizationSubscriber{
		BaseSubscriber:      NewBaseSubscriber(subject),
		organizationService: organizationService,
	}
	os.subOpts = append(os.subOpts, subOpts...)
	return os
}

func (os *GismoOrganizationSubscriber) Listen(msg *nats.Msg) (*inbox.Message, error) {
	ctx := context.Background()
	iMsg, err := gismo.ParseOrganizationMessage(msg.Data)

	if err != nil {
		return nil, fmt.Errorf("%w: unable to process malformed message: %s", models.ErrNonFatal, err)
	}

	org, err := os.organizationService.GetOrganization(ctx, iMsg.ID)

	if err != nil && err == models.ErrNotFound {
		org = models.NewOrganization()
	} else if err != nil {
		return iMsg, fmt.Errorf("%w: unable to fetch organization record: %s", models.ErrFatal, err)
	}

	if iMsg.Subject == "organization.update" {
		iMsg.UpdateOrganizationAttr(org)

		if org.IsStored() {
			_, err = os.organizationService.UpdateOrganization(ctx, org)
		} else {
			_, err = os.organizationService.CreateOrganization(ctx, org)
		}

		if err != nil {
			return iMsg, fmt.Errorf("%w: unable to store organization record: %s", models.ErrFatal, err)
		}
	} else if iMsg.Subject == "organization.delete" {
		if org.IsStored() {
			if err := os.organizationService.DeleteOrganization(ctx, org.Id); err != nil {
				return iMsg, fmt.Errorf("%w: unable to delete organization record: %s", models.ErrFatal, err)
			}
		}
	}

	return iMsg, nil
}
