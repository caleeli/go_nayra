package bpmn

import (
	"encoding/xml"
)

// ConditionalEventDefinition from BPMN
type ConditionalEventDefinition struct {
	EventDefinition
	XMLName   xml.Name
	Condition Expression `xml:"condition"`
}

// ParseTree of components of ConditionalEventDefinition.
func (conditionalEventDefinition *ConditionalEventDefinition) ParseTree(definitions *Definitions) {
	conditionalEventDefinition.EventDefinition.ParseTree(definitions)

	conditionalEventDefinition.Condition.ParseTree(definitions) // Expression.

}
