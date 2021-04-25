package bpmn

import (
	"encoding/xml"
)

// CallableElement from BPMN
type CallableElement struct {
	RootElement
	XMLName               xml.Name
	SupportedInterfaceRef []string                 `xml:"supportedInterfaceRef"`
	IoSpecification       InputOutputSpecification `xml:"ioSpecification"`
	IoBinding             []InputOutputBinding     `xml:"ioBinding"`
	Name                  string                   `xml:"name,attr"`
}

// ParseTree of components of CallableElement.
func (callableElement *CallableElement) ParseTree(definitions *Definitions) {
	callableElement.RootElement.ParseTree(definitions)

	callableElement.IoSpecification.ParseTree(definitions) // InputOutputSpecification.

	for i := 0; i < len(callableElement.IoBinding); i++ {
		callableElement.IoBinding[i].ParseTree(definitions)
	}

}

// GetInputOutputBinding by ID
func (callableElement *CallableElement) GetInputOutputBinding(ID string) *InputOutputBinding {

	for i := 0; i < len(callableElement.IoBinding); i++ {
		if callableElement.IoBinding[i].ID == ID {
			return &callableElement.IoBinding[i]

		}
	}

	return nil
}
