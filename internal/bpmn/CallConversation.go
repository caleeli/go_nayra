package bpmn

import (
	"encoding/xml"
)

// CallConversation from BPMN
type CallConversation struct {
	ConversationNode
	XMLName                xml.Name
	ParticipantAssociation []ParticipantAssociation `xml:"participantAssociation"`
	CalledCollaborationRef string                   `xml:"calledCollaborationRef,attr"`
}

// ParseTree of components of CallConversation.
func (callConversation *CallConversation) ParseTree(definitions *Definitions) {
	callConversation.ConversationNode.ParseTree(definitions)

	for i := 0; i < len(callConversation.ParticipantAssociation); i++ {
		callConversation.ParticipantAssociation[i].ParseTree(definitions)
	}

}

// GetParticipantAssociation by ID
func (callConversation *CallConversation) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(callConversation.ParticipantAssociation); i++ {
		if callConversation.ParticipantAssociation[i].ID == ID {
			return &callConversation.ParticipantAssociation[i]

		}
	}

	return nil
}
