package bpmn

import (
	"encoding/xml"
)

// BoundaryEvent from BPMN
type BoundaryEvent struct {
	CatchEvent
	XMLName        xml.Name
	CancelActivity bool     `xml:"cancelActivity,attr"`
	AttachedToRef  string   `xml:"attachedToRef,attr"`

}

// ParseTree of components of BoundaryEvent.
func (boundaryEvent *BoundaryEvent) ParseTree (definitions *Definitions) {
	boundaryEvent.CatchEvent.ParseTree(definitions)

}

