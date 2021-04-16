package bpmn

import (
	"encoding/xml"
)

// CancelEventDefinition from BPMN
type CancelEventDefinition struct {
	EventDefinition
	XMLName xml.Name

}

// ParseTree of components of CancelEventDefinition.
func (cancelEventDefinition *CancelEventDefinition) ParseTree (definitions *Definitions) {
	cancelEventDefinition.EventDefinition.ParseTree(definitions)

}

