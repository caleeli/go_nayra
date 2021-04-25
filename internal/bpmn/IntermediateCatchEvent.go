package bpmn

import (
	"encoding/xml"
)

// IntermediateCatchEvent from BPMN
type IntermediateCatchEvent struct {
	CatchEvent
	XMLName xml.Name
}

// ParseTree of components of IntermediateCatchEvent.
func (intermediateCatchEvent *IntermediateCatchEvent) ParseTree(definitions *Definitions) {
	intermediateCatchEvent.CatchEvent.ParseTree(definitions)

}
