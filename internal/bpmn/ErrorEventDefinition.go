package bpmn

import (
	"encoding/xml"
)

// ErrorEventDefinition from BPMN
type ErrorEventDefinition struct {
	EventDefinition
	XMLName  xml.Name
	ErrorRef string   `xml:"errorRef,attr"`

}

// ParseTree of components of ErrorEventDefinition.
func (errorEventDefinition *ErrorEventDefinition) ParseTree (definitions *Definitions) {
	errorEventDefinition.EventDefinition.ParseTree(definitions)

}

