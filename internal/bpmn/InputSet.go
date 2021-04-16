package bpmn

import (
	"encoding/xml"
)

// InputSet from BPMN
type InputSet struct {
	BaseElement
	XMLName                 xml.Name
	DataInputRefs           []string `xml:"dataInputRefs"`
	OptionalInputRefs       []string `xml:"optionalInputRefs"`
	WhileExecutingInputRefs []string `xml:"whileExecutingInputRefs"`
	OutputSetRefs           []string `xml:"outputSetRefs"`
	Name                    string   `xml:"name,attr"`

}

// ParseTree of components of InputSet.
func (inputSet *InputSet) ParseTree (definitions *Definitions) {
	inputSet.BaseElement.ParseTree(definitions)

}

