package bpmn

import (
	"encoding/xml"
)

// DataOutputAssociation from BPMN
type DataOutputAssociation struct {
	DataAssociation
	XMLName xml.Name

}

// ParseTree of components of DataOutputAssociation.
func (dataOutputAssociation *DataOutputAssociation) ParseTree (definitions *Definitions) {
	dataOutputAssociation.DataAssociation.ParseTree(definitions)

}

