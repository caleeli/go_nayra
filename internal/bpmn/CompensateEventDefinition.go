package bpmn

import (
	"encoding/xml"
)

// CompensateEventDefinition from BPMN
type CompensateEventDefinition struct {
	EventDefinition
	XMLName           xml.Name
	WaitForCompletion bool     `xml:"waitForCompletion,attr"`
	ActivityRef       string   `xml:"activityRef,attr"`

}

// ParseTree of components of CompensateEventDefinition.
func (compensateEventDefinition *CompensateEventDefinition) ParseTree (definitions *Definitions) {
	compensateEventDefinition.EventDefinition.ParseTree(definitions)

}

