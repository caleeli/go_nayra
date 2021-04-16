package bpmn

import (
	"encoding/xml"
)

// DataObjectReference from BPMN
type DataObjectReference struct {
	FlowElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	DataObjectRef  string    `xml:"dataObjectRef,attr"`

}

// ParseTree of components of DataObjectReference.
func (dataObjectReference *DataObjectReference) ParseTree (definitions *Definitions) {
	dataObjectReference.FlowElement.ParseTree(definitions)

	dataObjectReference.DataState.ParseTree(definitions) // DataState.

}

