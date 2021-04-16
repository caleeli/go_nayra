package bpmn

import (
	"encoding/xml"
)

// Auditing from BPMN
type Auditing struct {
	BaseElement
	XMLName xml.Name

}

// ParseTree of components of Auditing.
func (auditing *Auditing) ParseTree (definitions *Definitions) {
	auditing.BaseElement.ParseTree(definitions)

}

