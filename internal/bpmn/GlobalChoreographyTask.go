package bpmn

import (
	"encoding/xml"
)

// GlobalChoreographyTask from BPMN
type GlobalChoreographyTask struct {
	Choreography
	XMLName                  xml.Name
	InitiatingParticipantRef string   `xml:"initiatingParticipantRef,attr"`

}

// ParseTree of components of GlobalChoreographyTask.
func (globalChoreographyTask *GlobalChoreographyTask) ParseTree (definitions *Definitions) {
	globalChoreographyTask.Choreography.ParseTree(definitions)

}

