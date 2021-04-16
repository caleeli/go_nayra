package bpmn

import (
	"encoding/xml"
)

// DataInput from BPMN
type DataInput struct {
	BaseElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	Name           string    `xml:"name,attr"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	IsCollection   bool      `xml:"isCollection,attr"`

}

// ParseTree of components of DataInput.
func (dataInput *DataInput) ParseTree (definitions *Definitions) {
	dataInput.BaseElement.ParseTree(definitions)

	dataInput.DataState.ParseTree(definitions) // DataState.

}

