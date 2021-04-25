package bpmn

import (
	"encoding/xml"
)

// SignalEventDefinition from BPMN
type SignalEventDefinition struct {
	EventDefinition
	XMLName   xml.Name
	SignalRef string `xml:"signalRef,attr"`
}

// ParseTree of components of SignalEventDefinition.
func (signalEventDefinition *SignalEventDefinition) ParseTree(definitions *Definitions) {
	signalEventDefinition.EventDefinition.ParseTree(definitions)

}
