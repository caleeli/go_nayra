package bpmn

import (
	"encoding/xml"
)

// HumanPerformer from BPMN
type HumanPerformer struct {
	Performer
	XMLName xml.Name

}

// ParseTree of components of HumanPerformer.
func (humanPerformer *HumanPerformer) ParseTree (definitions *Definitions) {
	humanPerformer.Performer.ParseTree(definitions)

}

