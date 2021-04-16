package bpmn

import (
	"encoding/xml"
)

// DataState from BPMN
type DataState struct {
	BaseElement
	XMLName xml.Name
	Name    string   `xml:"name,attr"`

}

// ParseTree of components of DataState.
func (dataState *DataState) ParseTree (definitions *Definitions) {
	dataState.BaseElement.ParseTree(definitions)

}

