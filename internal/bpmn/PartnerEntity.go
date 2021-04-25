package bpmn

import (
	"encoding/xml"
)

// PartnerEntity from BPMN
type PartnerEntity struct {
	RootElement
	XMLName        xml.Name
	ParticipantRef []string `xml:"participantRef"`
	Name           string   `xml:"name,attr"`
}

// ParseTree of components of PartnerEntity.
func (partnerEntity *PartnerEntity) ParseTree(definitions *Definitions) {
	partnerEntity.RootElement.ParseTree(definitions)

}
