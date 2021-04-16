package bpmn

import (
	"encoding/xml"
)

// SubConversation from BPMN
type SubConversation struct {
	ConversationNode
	XMLName          xml.Name
	CallConversation []CallConversation `xml:"callConversation"`
	Conversation     []Conversation     `xml:"conversation"`
	SubConversation  []SubConversation  `xml:"subConversation"`

}

// ParseTree of components of SubConversation.
func (subConversation *SubConversation) ParseTree (definitions *Definitions) {
	subConversation.ConversationNode.ParseTree(definitions)

	for i := 0; i < len(subConversation.CallConversation); i++ {
		subConversation.CallConversation[i].ParseTree(definitions)
	}

	for i := 0; i < len(subConversation.Conversation); i++ {
		subConversation.Conversation[i].ParseTree(definitions)
	}

	for i := 0; i < len(subConversation.SubConversation); i++ {
		subConversation.SubConversation[i].ParseTree(definitions)
	}

}

// GetCallConversation by ID
func (subConversation *SubConversation) GetCallConversation(ID string) *CallConversation {

	for i := 0; i < len(subConversation.CallConversation); i++ {
		if subConversation.CallConversation[i].ID == ID {
			return &subConversation.CallConversation[i]

		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (subConversation *SubConversation) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(subConversation.CallConversation); i++ {
		ParticipantAssociation := subConversation.CallConversation[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetConversation by ID
func (subConversation *SubConversation) GetConversation(ID string) *Conversation {

	for i := 0; i < len(subConversation.Conversation); i++ {
		if subConversation.Conversation[i].ID == ID {
			return &subConversation.Conversation[i]

		}
	}

	return nil
}

// GetSubConversation by ID
func (subConversation *SubConversation) GetSubConversation(ID string) *SubConversation {

	for i := 0; i < len(subConversation.SubConversation); i++ {
		if subConversation.SubConversation[i].ID == ID {
			return &subConversation.SubConversation[i]

		}
	}

	return nil
}

