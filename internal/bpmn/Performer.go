package bpmn

import (
	"encoding/xml"
)

// Performer from BPMN
type Performer struct {
	ResourceRole
	XMLName xml.Name
}

// ParseTree of components of Performer.
func (performer *Performer) ParseTree(definitions *Definitions) {
	performer.ResourceRole.ParseTree(definitions)

}
