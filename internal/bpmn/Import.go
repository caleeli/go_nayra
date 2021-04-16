package bpmn

import (
	"encoding/xml"
)

// Import from BPMN
type Import struct {
	XMLName    xml.Name
	Namespace  string   `xml:"namespace,attr"`
	Location   string   `xml:"location,attr"`
	ImportType string   `xml:"importType,attr"`

}

// ParseTree of components of Import.
func (_import *Import) ParseTree (definitions *Definitions) {

}

