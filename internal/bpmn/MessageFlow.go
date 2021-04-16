package bpmn

import (
	"encoding/xml"
)

// MessageFlow from BPMN
type MessageFlow struct {
	BaseElement
	XMLName    xml.Name
	Name       string   `xml:"name,attr"`
	SourceRef  string   `xml:"sourceRef,attr"`
	TargetRef  string   `xml:"targetRef,attr"`
	MessageRef string   `xml:"messageRef,attr"`

}

// ParseTree of components of MessageFlow.
func (messageFlow *MessageFlow) ParseTree (definitions *Definitions) {
	messageFlow.BaseElement.ParseTree(definitions)

}

