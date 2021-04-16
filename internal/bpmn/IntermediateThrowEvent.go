package bpmn

import (
	"encoding/xml"
)

// IntermediateThrowEvent from BPMN
type IntermediateThrowEvent struct {
	ThrowEvent
	XMLName xml.Name

}

// ParseTree of components of IntermediateThrowEvent.
func (intermediateThrowEvent *IntermediateThrowEvent) ParseTree (definitions *Definitions) {
	intermediateThrowEvent.ThrowEvent.ParseTree(definitions)

}

