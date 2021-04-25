package bpmn

import (
	"encoding/xml"
)

// CallChoreography from BPMN
type CallChoreography struct {
	ChoreographyActivity
	XMLName                xml.Name
	ParticipantAssociation []ParticipantAssociation `xml:"participantAssociation"`
	CalledChoreographyRef  string                   `xml:"calledChoreographyRef,attr"`
}

// ParseTree of components of CallChoreography.
func (callChoreography *CallChoreography) ParseTree(definitions *Definitions) {
	callChoreography.ChoreographyActivity.ParseTree(definitions)

	for i := 0; i < len(callChoreography.ParticipantAssociation); i++ {
		callChoreography.ParticipantAssociation[i].ParseTree(definitions)
	}

}

// GetParticipantAssociation by ID
func (callChoreography *CallChoreography) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(callChoreography.ParticipantAssociation); i++ {
		if callChoreography.ParticipantAssociation[i].ID == ID {
			return &callChoreography.ParticipantAssociation[i]

		}
	}

	return nil
}
