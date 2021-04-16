package bpmn

import (
	"encoding/xml"
)

// Artifact from BPMN
type Artifact struct {
	BaseElement
	XMLName xml.Name

}

// ParseTree of components of Artifact.
func (artifact *Artifact) ParseTree (definitions *Definitions) {
	artifact.BaseElement.ParseTree(definitions)

}

