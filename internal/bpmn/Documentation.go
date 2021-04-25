package bpmn

import (
	"encoding/xml"
)

// Documentation from BPMN
type Documentation struct {
	XMLName    xml.Name
	Body       string `xml:",chardata"`
	ID         string `xml:"id,attr"`
	TextFormat string `xml:"textFormat,attr"`
}

// ParseTree of components of Documentation.
func (documentation *Documentation) ParseTree(definitions *Definitions) {

}
