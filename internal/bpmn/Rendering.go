package bpmn

import (
	"encoding/xml"
)

// Rendering from BPMN
type Rendering struct {
	BaseElement
	XMLName xml.Name
}

// ParseTree of components of Rendering.
func (rendering *Rendering) ParseTree(definitions *Definitions) {
	rendering.BaseElement.ParseTree(definitions)

}
