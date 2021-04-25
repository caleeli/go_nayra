package bpmn

import (
	"encoding/xml"
)

// MessageFlowAssociation from BPMN
type MessageFlowAssociation struct {
	BaseElement
	XMLName             xml.Name
	InnerMessageFlowRef string `xml:"innerMessageFlowRef,attr"`
	OuterMessageFlowRef string `xml:"outerMessageFlowRef,attr"`
}

// ParseTree of components of MessageFlowAssociation.
func (messageFlowAssociation *MessageFlowAssociation) ParseTree(definitions *Definitions) {
	messageFlowAssociation.BaseElement.ParseTree(definitions)

}
