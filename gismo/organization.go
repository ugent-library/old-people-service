package gismo

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/ugent-library/person-service/inbox"
)

func ParseOrganizationMessage(buf []byte) (*inbox.Message, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	removeNamespace(doc)

	node := xmlquery.FindOne(doc, "//cfOrgUnit")

	if node == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfOrgUnit", ErrNonCompliantXml)
	}

	idNode := xmlquery.FindOne(node, "cfOrgUnitId")

	if idNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfOrgUnit/cfOrgUnitId", ErrNonCompliantXml)
	}

	msg := &inbox.Message{
		ID: strings.TrimSpace(idNode.InnerText()),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Source = "gismo.organization.delete"
	} else {
		msg.Source = "gismo.organization.update"

		for _, n := range xmlquery.Find(node, "./cfName") {
			t1, err := time.Parse(time.RFC3339, n.SelectAttr("cfStartDate"))
			if err != nil {
				return nil, err
			}
			t2, err := time.Parse(time.RFC3339, n.SelectAttr("cfEndDate"))
			if err != nil {
				return nil, err
			}
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "name_" + n.SelectAttr("cfLangCode"),
				Value:     strings.TrimSpace(n.InnerText()),
				StartDate: &t1,
				EndDate:   &t2,
			})
		}

		for _, v := range cerifValuesByClassName(doc, "cfOrgUnit_Class", "/be.ugent/organisatie/type/vakgroep", "") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "type",
				Value:     "department",
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfOrgUnit_OrgUnit", "/be.ugent/gismo/organisatie-organisatie/type/kind-van", "cfOrgUnitId1") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "parent_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. 000006047
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/memorialis", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "ugent_memorialis_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. "WE03V"
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/org-code", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "code",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. WE03*
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/biblio-code", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "biblio_code",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}

	}

	return msg, nil
}
