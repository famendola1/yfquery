package yfquery

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/xmlquery"
)

func parse(xmlDoc, expr string, out any) error {
	doc, err := xmlquery.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, expr)
	if err != nil {
		return err
	}

	return xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(out)
}
