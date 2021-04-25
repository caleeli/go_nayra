package bpmn

import (
	"encoding/xml"
)

// DataOutput from BPMN
type DataOutput struct {
	BaseElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	Name           string    `xml:"name,attr"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	IsCollection   bool      `xml:"isCollection,attr"`
}

// ParseTree of components of DataOutput.
func (dataOutput *DataOutput) ParseTree(definitions *Definitions) {
	dataOutput.BaseElement.ParseTree(definitions)

	dataOutput.DataState.ParseTree(definitions) // DataState.

}
