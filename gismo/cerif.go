package gismo

import (
	"errors"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

/*
	unhappily stolen from
	https://github.com/ugent-library/soap-bridge/blob/main/main.go
*/

type cerifValue struct {
	Value     string
	StartDate time.Time
	EndDate   time.Time
}

var ErrNonCompliantXml = errors.New("non compliant cerif xml")

func removeNamespace(n *xmlquery.Node) {
	n.Prefix = ""
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		removeNamespace(child)
	}
}

func cerifClassID(root *xmlquery.Node, name string) string {
	classIDNode := xmlquery.FindOne(root, "//cfClass[contains(cfURI, '"+name+"')]/cfClassId")
	if classIDNode == nil {
		return ""
	}
	return strings.TrimSpace(classIDNode.InnerText())
}

func cerifNodesByClassName(root *xmlquery.Node, tag, name string) []*xmlquery.Node {
	classID := cerifClassID(root, name)
	if classID == "" {
		return nil
	}
	return xmlquery.Find(root, "//"+tag+"[contains(cfClassId, '"+classID+"')]")
}

func cerifValuesByClassName(root *xmlquery.Node, tag, name, valueTag string) []cerifValue {
	nodes := cerifNodesByClassName(root, tag, name)
	vals := make([]cerifValue, 0, len(nodes))
	for _, node := range nodes {
		val := cerifValue{}
		if valueTag != "" {
			if n := xmlquery.FindOne(node, valueTag); n != nil {
				val.Value = strings.TrimSpace(n.InnerText())
			} else {
				continue
			}
		}
		if n := xmlquery.FindOne(node, "cfStartDate"); n != nil {
			t, err := time.Parse(time.RFC3339, strings.TrimSpace(n.InnerText()))
			if err != nil {
				continue
			}
			val.StartDate = t
		} else {
			continue
		}
		if n := xmlquery.FindOne(node, "cfEndDate"); n != nil {
			t, err := time.Parse(time.RFC3339, strings.TrimSpace(n.InnerText()))
			if err != nil {
				continue
			}
			val.EndDate = t
		} else {
			continue
		}
		vals = append(vals, val)
	}
	return vals
}
