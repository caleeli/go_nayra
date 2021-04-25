package bpmn

import (
	"encoding/xml"
)

// ConversationNode from BPMN
type ConversationNode struct {
	BaseElement
	XMLName        xml.Name
	ParticipantRef []string         `xml:"participantRef"`
	MessageFlowRef []string         `xml:"messageFlowRef"`
	CorrelationKey []CorrelationKey `xml:"correlationKey"`
	Name           string           `xml:"name,attr"`
}

// ParseTree of components of ConversationNode.
func (conversationNode *ConversationNode) ParseTree(definitions *Definitions) {
	conversationNode.BaseElement.ParseTree(definitions)

	for i := 0; i < len(conversationNode.CorrelationKey); i++ {
		conversationNode.CorrelationKey[i].ParseTree(definitions)
	}

}

// GetCorrelationKey by ID
func (conversationNode *ConversationNode) GetCorrelationKey(ID string) *CorrelationKey {

	for i := 0; i < len(conversationNode.CorrelationKey); i++ {
		if conversationNode.CorrelationKey[i].ID == ID {
			return &conversationNode.CorrelationKey[i]

		}
	}

	return nil
}
