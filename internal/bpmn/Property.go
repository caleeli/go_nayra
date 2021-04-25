package bpmn

import (
	"encoding/xml"
)

// Property from BPMN
type Property struct {
	BaseElement
	XMLName        xml.Name
	DataState      DataState `xml:"dataState"`
	Name           string    `xml:"name,attr"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
}

// ParseTree of components of Property.
func (property *Property) ParseTree(definitions *Definitions) {
	property.BaseElement.ParseTree(definitions)

	property.DataState.ParseTree(definitions) // DataState.

}
