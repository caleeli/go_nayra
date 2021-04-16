package bpmn

import (
	"encoding/xml"
)

// ChoreographyTask from BPMN
type ChoreographyTask struct {
	ChoreographyActivity
	XMLName        xml.Name
	MessageFlowRef []string `xml:"messageFlowRef"`

}

// ParseTree of components of ChoreographyTask.
func (choreographyTask *ChoreographyTask) ParseTree (definitions *Definitions) {
	choreographyTask.ChoreographyActivity.ParseTree(definitions)

}

