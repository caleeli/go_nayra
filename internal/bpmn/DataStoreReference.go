package bpmn

import (
	"encoding/xml"
)

// DataStoreReference from BPMN
type DataStoreReference struct {
	FlowElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	DataStoreRef   string    `xml:"dataStoreRef,attr"`

}

// ParseTree of components of DataStoreReference.
func (dataStoreReference *DataStoreReference) ParseTree (definitions *Definitions) {
	dataStoreReference.FlowElement.ParseTree(definitions)

	dataStoreReference.DataState.ParseTree(definitions) // DataState.

}

