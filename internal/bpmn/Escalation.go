package bpmn

import (
	"encoding/xml"
)

// Escalation from BPMN
type Escalation struct {
	RootElement
	XMLName        xml.Name
	Name           string   `xml:"name,attr"`
	EscalationCode string   `xml:"escalationCode,attr"`
	StructureRef   string   `xml:"structureRef,attr"`

}

// ParseTree of components of Escalation.
func (escalation *Escalation) ParseTree (definitions *Definitions) {
	escalation.RootElement.ParseTree(definitions)

}

