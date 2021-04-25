package bpmn

import (
	"encoding/xml"
)

// Operation from BPMN
type Operation struct {
	BaseElement
	XMLName           xml.Name
	InMessageRef      []string `xml:"inMessageRef"`
	OutMessageRef     []string `xml:"outMessageRef"`
	ErrorRef          []string `xml:"errorRef"`
	Name              string   `xml:"name,attr"`
	ImplementationRef string   `xml:"implementationRef,attr"`
}

// ParseTree of components of Operation.
func (operation *Operation) ParseTree(definitions *Definitions) {
	operation.BaseElement.ParseTree(definitions)

}
