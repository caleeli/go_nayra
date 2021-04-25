package bpmn

import (
	"encoding/xml"
)

// EndEvent from BPMN
type EndEvent struct {
	EndEventTrait `bson:"-"`
	ThrowEvent
	XMLName xml.Name
}

// ParseTree of components of EndEvent.
func (endEvent *EndEvent) ParseTree(definitions *Definitions) {
	endEvent.ThrowEvent.ParseTree(definitions)

	endEvent.Init(definitions)

}
