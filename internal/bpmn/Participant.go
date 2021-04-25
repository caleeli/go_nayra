package bpmn

import (
	"encoding/xml"
)

// Participant from BPMN
type Participant struct {
	BaseElement
	XMLName                 xml.Name
	InterfaceRef            []string                `xml:"interfaceRef"`
	EndPointRef             []string                `xml:"endPointRef"`
	ParticipantMultiplicity ParticipantMultiplicity `xml:"participantMultiplicity"`
	Name                    string                  `xml:"name,attr"`
	ProcessRef              string                  `xml:"processRef,attr"`
}

// ParseTree of components of Participant.
func (participant *Participant) ParseTree(definitions *Definitions) {
	participant.BaseElement.ParseTree(definitions)

	participant.ParticipantMultiplicity.ParseTree(definitions) // ParticipantMultiplicity.

}
