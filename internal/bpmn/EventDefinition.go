package bpmn

import (
	"encoding/xml"
)

// EventDefinition from BPMN
type EventDefinition struct {
	RootElement
	XMLName xml.Name
}

// ParseTree of components of EventDefinition.
func (eventDefinition *EventDefinition) ParseTree(definitions *Definitions) {
	eventDefinition.RootElement.ParseTree(definitions)

}
