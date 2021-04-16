package bpmn

import (
	"encoding/xml"
)

// Text from BPMN
type Text struct {
	XMLName xml.Name
	Body    string   `xml:",chardata"`

}

// ParseTree of components of Text.
func (text *Text) ParseTree (definitions *Definitions) {

}

