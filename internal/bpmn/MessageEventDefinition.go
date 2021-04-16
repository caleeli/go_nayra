package bpmn

import (
	"encoding/xml"
)

// MessageEventDefinition from BPMN
type MessageEventDefinition struct {
	EventDefinition
	XMLName      xml.Name
	OperationRef []string `xml:"operationRef"`
	MessageRef   string   `xml:"messageRef,attr"`

}

// ParseTree of components of MessageEventDefinition.
func (messageEventDefinition *MessageEventDefinition) ParseTree (definitions *Definitions) {
	messageEventDefinition.EventDefinition.ParseTree(definitions)

}

