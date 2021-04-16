package bpmn

import (
	"encoding/xml"
)

// DataStore from BPMN
type DataStore struct {
	RootElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	Name           string    `xml:"name,attr"`
	Capacity       int       `xml:"capacity,attr"`
	IsUnlimited    bool      `xml:"isUnlimited,attr"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`

}

// ParseTree of components of DataStore.
func (dataStore *DataStore) ParseTree (definitions *Definitions) {
	dataStore.RootElement.ParseTree(definitions)

	dataStore.DataState.ParseTree(definitions) // DataState.

}

