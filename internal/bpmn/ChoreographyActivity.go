package bpmn

import (
	"encoding/xml"
)

// ChoreographyActivity from BPMN
type ChoreographyActivity struct {
	FlowNode
	XMLName                  xml.Name
	ParticipantRef           []string         `xml:"participantRef"`
	CorrelationKey           []CorrelationKey `xml:"correlationKey"`
	InitiatingParticipantRef string           `xml:"initiatingParticipantRef,attr"`
	LoopType                 string           `xml:"loopType,attr"`
}

// ParseTree of components of ChoreographyActivity.
func (choreographyActivity *ChoreographyActivity) ParseTree(definitions *Definitions) {
	choreographyActivity.FlowNode.ParseTree(definitions)

	for i := 0; i < len(choreographyActivity.CorrelationKey); i++ {
		choreographyActivity.CorrelationKey[i].ParseTree(definitions)
	}

}

// GetCorrelationKey by ID
func (choreographyActivity *ChoreographyActivity) GetCorrelationKey(ID string) *CorrelationKey {

	for i := 0; i < len(choreographyActivity.CorrelationKey); i++ {
		if choreographyActivity.CorrelationKey[i].ID == ID {
			return &choreographyActivity.CorrelationKey[i]

		}
	}

	return nil
}
