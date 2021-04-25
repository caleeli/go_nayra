package bpmn

import (
	"encoding/xml"
)

// DataObject from BPMN
type DataObject struct {
	FlowElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	IsCollection   bool      `xml:"isCollection,attr"`
}

// ParseTree of components of DataObject.
func (dataObject *DataObject) ParseTree(definitions *Definitions) {
	dataObject.FlowElement.ParseTree(definitions)

	dataObject.DataState.ParseTree(definitions) // DataState.

}
