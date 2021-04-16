package bpmn

import (
	"encoding/xml"
)

// ConversationAssociation from BPMN
type ConversationAssociation struct {
	BaseElement
	XMLName                  xml.Name
	InnerConversationNodeRef string   `xml:"innerConversationNodeRef,attr"`
	OuterConversationNodeRef string   `xml:"outerConversationNodeRef,attr"`

}

// ParseTree of components of ConversationAssociation.
func (conversationAssociation *ConversationAssociation) ParseTree (definitions *Definitions) {
	conversationAssociation.BaseElement.ParseTree(definitions)

}

