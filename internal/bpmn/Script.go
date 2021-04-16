package bpmn

import (
	"encoding/xml"
)

// Script from BPMN
type Script struct {
	XMLName xml.Name
	Body    string   `xml:",chardata"`

}

// ParseTree of components of Script.
func (script *Script) ParseTree (definitions *Definitions) {

}

