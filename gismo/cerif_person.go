package gismo

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/ugent-library/cerifutil"
	"github.com/ugent-library/people-service/models"
)

func ParsePersonMessage(buf []byte) (*models.Message, error) {
	doc, err := cerifutil.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

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

	persNode := xmlquery.FindOne(doc, "//Body/persons")

	if persNode == nil {
		return nil, fmt.Errorf("%w: unabel to find node //Body/persons", ErrNonCompliantXml)
	}

	msg := &models.Message{
		ID:       strings.TrimSpace(idNode.InnerText()),
		Language: strings.TrimSpace(langNode.InnerText()),
		Date:     persNode.SelectAttr("date"),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Source = "gismo.person.delete"
	} else {
		msg.Source = "gismo.person.update"

		for _, nameNode := range cerifutil.NodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/officiele-naam") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "family_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "given_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "ugTitles"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "honorific_prefix",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}
		for _, nameNode := range cerifutil.NodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/voorkeursweergave") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "preferred_family_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "preferred_given_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}

		for _, v := range cerifutil.ValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/ugent-id", "cfFedId") {
			startDate := v.StartDate
			endDate := v.EndDate
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "ugent_id",
				Value:     v.Value,
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}
		for _, v := range cerifutil.ValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/orcid", "cfFedId") {
			startDate := v.StartDate
			endDate := v.EndDate
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "orcid",
				Value:     v.Value,
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}
		for _, v := range cerifutil.ValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/memorialis", "cfFedId") {
			startDate := v.StartDate
			endDate := v.EndDate
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "ugent_memorialis_id",
				Value:     v.Value,
				StartDate: &startDate,
				EndDate:   &endDate,
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
			msg.Attributes = append(msg.Attributes, models.Attribute{
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
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "organization_id",
				Value:     strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfOrgUnitId").InnerText()),
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}
	}

	return msg, nil
}
