package bpmn

import (
	"encoding/xml"
)

// ParticipantMultiplicity from BPMN
type ParticipantMultiplicity struct {
	BaseElement
	XMLName xml.Name
	Minimum int `xml:"minimum,attr"`
	Maximum int `xml:"maximum,attr"`
}

// ParseTree of components of ParticipantMultiplicity.
func (participantMultiplicity *ParticipantMultiplicity) ParseTree(definitions *Definitions) {
	participantMultiplicity.BaseElement.ParseTree(definitions)

}
