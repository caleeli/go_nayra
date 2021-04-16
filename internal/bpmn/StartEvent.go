package bpmn

import (
	"encoding/xml"
)

// StartEvent from BPMN
type StartEvent struct {
	StartEventTrait
	CatchEvent
	XMLName        xml.Name
	IsInterrupting bool     `xml:"isInterrupting,attr"`

}

// ParseTree of components of StartEvent.
func (startEvent *StartEvent) ParseTree (definitions *Definitions) {
	startEvent.CatchEvent.ParseTree(definitions)

	startEvent.Init(definitions)

}

