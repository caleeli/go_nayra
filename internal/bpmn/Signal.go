package bpmn

import (
	"encoding/xml"
)

// Signal from BPMN
type Signal struct {
	RootElement
	XMLName      xml.Name
	Name         string `xml:"name,attr"`
	StructureRef string `xml:"structureRef,attr"`
}

// ParseTree of components of Signal.
func (signal *Signal) ParseTree(definitions *Definitions) {
	signal.RootElement.ParseTree(definitions)

}
