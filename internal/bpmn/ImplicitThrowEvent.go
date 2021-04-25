package bpmn

import (
	"encoding/xml"
)

// ImplicitThrowEvent from BPMN
type ImplicitThrowEvent struct {
	ThrowEvent
	XMLName xml.Name
}

// ParseTree of components of ImplicitThrowEvent.
func (implicitThrowEvent *ImplicitThrowEvent) ParseTree(definitions *Definitions) {
	implicitThrowEvent.ThrowEvent.ParseTree(definitions)

}
