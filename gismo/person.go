package gismo

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/ugent-library/person-service/inbox"
)

func ParsePersonMessage(buf []byte) (*inbox.Message, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	removeNamespace(doc)

	node := xmlquery.FindOne(doc, "//cfPers")

	if node == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers", ErrNonCompliantXml)
	}

	idNode := xmlquery.FindOne(node, "cfPersId")

	if idNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers/cfPersId", ErrNonCompliantXml)
	}

	langNode := xmlquery.FindOne(node, "cfPers_Lang/cfLangCode")

	if langNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers/cfPers_Lang/cfLangCode", ErrNonCompliantXml)
	}

	msg := &inbox.Message{
		ID:       strings.TrimSpace(idNode.InnerText()),
		Language: strings.TrimSpace(langNode.InnerText()),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Subject = "person.delete"
	} else {
		msg.Subject = "person.update"

		for _, nameNode := range cerifNodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/officiele-naam") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, inbox.Attribute{
					Name:      "last_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, inbox.Attribute{
					Name:      "first_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "ugTitles"); n != nil {
				msg.Attributes = append(msg.Attributes, inbox.Attribute{
					Name:      "title",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}
		for _, nameNode := range cerifNodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/voorkeursweergave") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, inbox.Attribute{
					Name:      "preferred_last_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, inbox.Attribute{
					Name:      "preferred_first_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}

		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/ugent-id", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "ugent_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/orcid", "cfFedId") {
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "orcid",
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

		for _, persAddrNode := range xmlquery.Find(node, "cfPers_EAddr") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			addrID := strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEAddrId").InnerText())
			addrNode := xmlquery.FindOne(doc, "//cfEAddr[contains(cfEAddrId, '"+addrID+"')]")
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "email",
				Value:     strings.TrimSpace(xmlquery.FindOne(addrNode, "cfURI").InnerText()),
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}

		for _, persAddrNode := range xmlquery.Find(node, "cfPers_OrgUnit") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			msg.Attributes = append(msg.Attributes, inbox.Attribute{
				Name:      "organization_id",
				Value:     strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfOrgUnitId").InnerText()),
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}
	}

	return msg, nil
}
