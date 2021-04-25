package bpmn

import (
	"encoding/xml"
)

// BaseElementWithMixedContent from BPMN
type BaseElementWithMixedContent struct {
	XMLName           xml.Name
	Documentation     []Documentation   `xml:"documentation"`
	ExtensionElements ExtensionElements `xml:"extensionElements"`
	ID                string            `xml:"id,attr"`
}

// ParseTree of components of BaseElementWithMixedContent.
func (baseElementWithMixedContent *BaseElementWithMixedContent) ParseTree(definitions *Definitions) {

	for i := 0; i < len(baseElementWithMixedContent.Documentation); i++ {
		baseElementWithMixedContent.Documentation[i].ParseTree(definitions)
	}

	baseElementWithMixedContent.ExtensionElements.ParseTree(definitions) // ExtensionElements.

}

// GetDocumentation by ID
func (baseElementWithMixedContent *BaseElementWithMixedContent) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(baseElementWithMixedContent.Documentation); i++ {
		if baseElementWithMixedContent.Documentation[i].ID == ID {
			return &baseElementWithMixedContent.Documentation[i]

		}
	}

	return nil
}
