package bpmn

import (
	"encoding/xml"
)

// TerminateEventDefinition from BPMN
type TerminateEventDefinition struct {
	EventDefinition
	XMLName xml.Name

}

// ParseTree of components of TerminateEventDefinition.
func (terminateEventDefinition *TerminateEventDefinition) ParseTree (definitions *Definitions) {
	terminateEventDefinition.EventDefinition.ParseTree(definitions)

}

