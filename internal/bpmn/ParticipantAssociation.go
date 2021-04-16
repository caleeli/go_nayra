package bpmn

import (
	"encoding/xml"
)

// ParticipantAssociation from BPMN
type ParticipantAssociation struct {
	BaseElement
	XMLName             xml.Name
	InnerParticipantRef []string `xml:"innerParticipantRef"`
	OuterParticipantRef []string `xml:"outerParticipantRef"`

}

// ParseTree of components of ParticipantAssociation.
func (participantAssociation *ParticipantAssociation) ParseTree (definitions *Definitions) {
	participantAssociation.BaseElement.ParseTree(definitions)

}

