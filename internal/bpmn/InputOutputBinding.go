package bpmn

import (
	"encoding/xml"
)

// InputOutputBinding from BPMN
type InputOutputBinding struct {
	BaseElement
	XMLName       xml.Name
	OperationRef  string   `xml:"operationRef,attr"`
	InputDataRef  string   `xml:"inputDataRef,attr"`
	OutputDataRef string   `xml:"outputDataRef,attr"`

}

// ParseTree of components of InputOutputBinding.
func (inputOutputBinding *InputOutputBinding) ParseTree (definitions *Definitions) {
	inputOutputBinding.BaseElement.ParseTree(definitions)

}

