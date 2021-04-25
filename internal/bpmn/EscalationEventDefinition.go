package bpmn

import (
	"encoding/xml"
)

// EscalationEventDefinition from BPMN
type EscalationEventDefinition struct {
	EventDefinition
	XMLName       xml.Name
	EscalationRef string `xml:"escalationRef,attr"`
}

// ParseTree of components of EscalationEventDefinition.
func (escalationEventDefinition *EscalationEventDefinition) ParseTree(definitions *Definitions) {
	escalationEventDefinition.EventDefinition.ParseTree(definitions)

}
