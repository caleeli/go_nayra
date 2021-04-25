package bpmn

import (
	"encoding/xml"
)

// BaseElement from BPMN
type BaseElement struct {
	XMLName           xml.Name
	Documentation     []Documentation   `xml:"documentation"`
	ExtensionElements ExtensionElements `xml:"extensionElements"`
	ID                string            `xml:"id,attr"`
}

// ParseTree of components of BaseElement.
func (baseElement *BaseElement) ParseTree(definitions *Definitions) {

	for i := 0; i < len(baseElement.Documentation); i++ {
		baseElement.Documentation[i].ParseTree(definitions)
	}

	baseElement.ExtensionElements.ParseTree(definitions) // ExtensionElements.

}

// GetDocumentation by ID
func (baseElement *BaseElement) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(baseElement.Documentation); i++ {
		if baseElement.Documentation[i].ID == ID {
			return &baseElement.Documentation[i]

		}
	}

	return nil
}
