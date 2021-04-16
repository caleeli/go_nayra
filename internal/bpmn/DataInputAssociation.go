package bpmn

import (
	"encoding/xml"
)

// DataInputAssociation from BPMN
type DataInputAssociation struct {
	DataAssociation
	XMLName xml.Name

}

// ParseTree of components of DataInputAssociation.
func (dataInputAssociation *DataInputAssociation) ParseTree (definitions *Definitions) {
	dataInputAssociation.DataAssociation.ParseTree(definitions)

}

