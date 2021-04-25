package bpmn

import (
	"encoding/xml"
)

// ConversationLink from BPMN
type ConversationLink struct {
	BaseElement
	XMLName   xml.Name
	Name      string `xml:"name,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}

// ParseTree of components of ConversationLink.
func (conversationLink *ConversationLink) ParseTree(definitions *Definitions) {
	conversationLink.BaseElement.ParseTree(definitions)

}
