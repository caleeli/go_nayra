package bpmn

import (
	"encoding/xml"
)

// RootElement from BPMN
type RootElement struct {
	BaseElement
	XMLName xml.Name

}

// ParseTree of components of RootElement.
func (rootElement *RootElement) ParseTree (definitions *Definitions) {
	rootElement.BaseElement.ParseTree(definitions)

}

