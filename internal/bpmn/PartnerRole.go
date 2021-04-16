package bpmn

import (
	"encoding/xml"
)

// PartnerRole from BPMN
type PartnerRole struct {
	RootElement
	XMLName        xml.Name
	ParticipantRef []string `xml:"participantRef"`
	Name           string   `xml:"name,attr"`

}

// ParseTree of components of PartnerRole.
func (partnerRole *PartnerRole) ParseTree (definitions *Definitions) {
	partnerRole.RootElement.ParseTree(definitions)

}

