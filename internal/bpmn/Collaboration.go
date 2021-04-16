package bpmn

import (
	"encoding/xml"
)

// Collaboration from BPMN
type Collaboration struct {
	RootElement
	XMLName                 xml.Name
	Participant             []Participant             `xml:"participant"`
	MessageFlow             []MessageFlow             `xml:"messageFlow"`
	Association             []Association             `xml:"association"`
	Group                   []Group                   `xml:"group"`
	TextAnnotation          []TextAnnotation          `xml:"textAnnotation"`
	CallConversation        []CallConversation        `xml:"callConversation"`
	Conversation            []Conversation            `xml:"conversation"`
	SubConversation         []SubConversation         `xml:"subConversation"`
	ConversationAssociation []ConversationAssociation `xml:"conversationAssociation"`
	ParticipantAssociation  []ParticipantAssociation  `xml:"participantAssociation"`
	MessageFlowAssociation  []MessageFlowAssociation  `xml:"messageFlowAssociation"`
	CorrelationKey          []CorrelationKey          `xml:"correlationKey"`
	ChoreographyRef         []string                  `xml:"choreographyRef"`
	ConversationLink        []ConversationLink        `xml:"conversationLink"`
	Name                    string                    `xml:"name,attr"`
	IsClosed                bool                      `xml:"isClosed,attr"`

}

// ParseTree of components of Collaboration.
func (collaboration *Collaboration) ParseTree (definitions *Definitions) {
	collaboration.RootElement.ParseTree(definitions)

	for i := 0; i < len(collaboration.Participant); i++ {
		collaboration.Participant[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.MessageFlow); i++ {
		collaboration.MessageFlow[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.Association); i++ {
		collaboration.Association[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.Group); i++ {
		collaboration.Group[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.TextAnnotation); i++ {
		collaboration.TextAnnotation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.CallConversation); i++ {
		collaboration.CallConversation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.Conversation); i++ {
		collaboration.Conversation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.SubConversation); i++ {
		collaboration.SubConversation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.ConversationAssociation); i++ {
		collaboration.ConversationAssociation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.ParticipantAssociation); i++ {
		collaboration.ParticipantAssociation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.MessageFlowAssociation); i++ {
		collaboration.MessageFlowAssociation[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.CorrelationKey); i++ {
		collaboration.CorrelationKey[i].ParseTree(definitions)
	}

	for i := 0; i < len(collaboration.ConversationLink); i++ {
		collaboration.ConversationLink[i].ParseTree(definitions)
	}

}

// GetParticipant by ID
func (collaboration *Collaboration) GetParticipant(ID string) *Participant {

	for i := 0; i < len(collaboration.Participant); i++ {
		if collaboration.Participant[i].ID == ID {
			return &collaboration.Participant[i]

		}
	}

	return nil
}

// GetMessageFlow by ID
func (collaboration *Collaboration) GetMessageFlow(ID string) *MessageFlow {

	for i := 0; i < len(collaboration.MessageFlow); i++ {
		if collaboration.MessageFlow[i].ID == ID {
			return &collaboration.MessageFlow[i]

		}
	}

	return nil
}

// GetAssociation by ID
func (collaboration *Collaboration) GetAssociation(ID string) *Association {

	for i := 0; i < len(collaboration.Association); i++ {
		if collaboration.Association[i].ID == ID {
			return &collaboration.Association[i]

		}
	}

	return nil
}

// GetGroup by ID
func (collaboration *Collaboration) GetGroup(ID string) *Group {

	for i := 0; i < len(collaboration.Group); i++ {
		if collaboration.Group[i].ID == ID {
			return &collaboration.Group[i]

		}
	}

	return nil
}

// GetTextAnnotation by ID
func (collaboration *Collaboration) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(collaboration.TextAnnotation); i++ {
		if collaboration.TextAnnotation[i].ID == ID {
			return &collaboration.TextAnnotation[i]

		}
	}

	return nil
}

// GetCallConversation by ID
func (collaboration *Collaboration) GetCallConversation(ID string) *CallConversation {

	for i := 0; i < len(collaboration.CallConversation); i++ {
		if collaboration.CallConversation[i].ID == ID {
			return &collaboration.CallConversation[i]

		}
	}

	for i := 0; i < len(collaboration.SubConversation); i++ {
		CallConversation := collaboration.SubConversation[i].GetCallConversation(ID)
		if CallConversation != nil {
			return CallConversation
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (collaboration *Collaboration) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(collaboration.CallConversation); i++ {
		ParticipantAssociation := collaboration.CallConversation[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(collaboration.SubConversation); i++ {
		ParticipantAssociation := collaboration.SubConversation[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(collaboration.ParticipantAssociation); i++ {
		if collaboration.ParticipantAssociation[i].ID == ID {
			return &collaboration.ParticipantAssociation[i]

		}
	}

	return nil
}

// GetConversation by ID
func (collaboration *Collaboration) GetConversation(ID string) *Conversation {

	for i := 0; i < len(collaboration.Conversation); i++ {
		if collaboration.Conversation[i].ID == ID {
			return &collaboration.Conversation[i]

		}
	}

	for i := 0; i < len(collaboration.SubConversation); i++ {
		Conversation := collaboration.SubConversation[i].GetConversation(ID)
		if Conversation != nil {
			return Conversation
		}
	}

	return nil
}

// GetSubConversation by ID
func (collaboration *Collaboration) GetSubConversation(ID string) *SubConversation {

	for i := 0; i < len(collaboration.SubConversation); i++ {
		if collaboration.SubConversation[i].ID == ID {
			return &collaboration.SubConversation[i]

		}
	}

	for i := 0; i < len(collaboration.SubConversation); i++ {
		SubConversation := collaboration.SubConversation[i].GetSubConversation(ID)
		if SubConversation != nil {
			return SubConversation
		}
	}

	return nil
}

// GetConversationAssociation by ID
func (collaboration *Collaboration) GetConversationAssociation(ID string) *ConversationAssociation {

	for i := 0; i < len(collaboration.ConversationAssociation); i++ {
		if collaboration.ConversationAssociation[i].ID == ID {
			return &collaboration.ConversationAssociation[i]

		}
	}

	return nil
}

// GetMessageFlowAssociation by ID
func (collaboration *Collaboration) GetMessageFlowAssociation(ID string) *MessageFlowAssociation {

	for i := 0; i < len(collaboration.MessageFlowAssociation); i++ {
		if collaboration.MessageFlowAssociation[i].ID == ID {
			return &collaboration.MessageFlowAssociation[i]

		}
	}

	return nil
}

// GetCorrelationKey by ID
func (collaboration *Collaboration) GetCorrelationKey(ID string) *CorrelationKey {

	for i := 0; i < len(collaboration.CorrelationKey); i++ {
		if collaboration.CorrelationKey[i].ID == ID {
			return &collaboration.CorrelationKey[i]

		}
	}

	return nil
}

// GetConversationLink by ID
func (collaboration *Collaboration) GetConversationLink(ID string) *ConversationLink {

	for i := 0; i < len(collaboration.ConversationLink); i++ {
		if collaboration.ConversationLink[i].ID == ID {
			return &collaboration.ConversationLink[i]

		}
	}

	return nil
}

