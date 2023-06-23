package gismo

import (
	"bytes"
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
		return nil, ErrNonCompliantXml
	}

	idNode := xmlquery.FindOne(node, "cfOrgUnitId")

	if idNode == nil {
		return nil, ErrNonCompliantXml
	}

	msg := &inbox.Message{
		ID: strings.TrimSpace(idNode.InnerText()),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Subject = "gismo.organization.delete"
	} else {
		msg.Subject = "organization.update"

		for _, n := range xmlquery.Find(node, "//cfName") {
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
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/memorialis", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "ugent_memorialis_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/org-code", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "code",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
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
