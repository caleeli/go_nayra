package bpmn

import (
	"encoding/xml"
)

// Event from BPMN
type Event struct {
	FlowNode
	XMLName  xml.Name
	Property []Property `xml:"property"`

}

// ParseTree of components of Event.
func (event *Event) ParseTree (definitions *Definitions) {
	event.FlowNode.ParseTree(definitions)

	for i := 0; i < len(event.Property); i++ {
		event.Property[i].ParseTree(definitions)
	}

}

// GetProperty by ID
func (event *Event) GetProperty(ID string) *Property {

	for i := 0; i < len(event.Property); i++ {
		if event.Property[i].ID == ID {
			return &event.Property[i]

		}
	}

	return nil
}

