package bpmn

import (
	"encoding/xml"
)

// Error from BPMN
type Error struct {
	RootElement
	XMLName      xml.Name
	Name         string `xml:"name,attr"`
	ErrorCode    string `xml:"errorCode,attr"`
	StructureRef string `xml:"structureRef,attr"`
}

// ParseTree of components of Error.
func (error *Error) ParseTree(definitions *Definitions) {
	error.RootElement.ParseTree(definitions)

}
