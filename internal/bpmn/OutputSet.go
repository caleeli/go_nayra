package bpmn

import (
	"encoding/xml"
)

// OutputSet from BPMN
type OutputSet struct {
	BaseElement
	XMLName                  xml.Name
	DataOutputRefs           []string `xml:"dataOutputRefs"`
	OptionalOutputRefs       []string `xml:"optionalOutputRefs"`
	WhileExecutingOutputRefs []string `xml:"whileExecutingOutputRefs"`
	InputSetRefs             []string `xml:"inputSetRefs"`
	Name                     string   `xml:"name,attr"`

}

// ParseTree of components of OutputSet.
func (outputSet *OutputSet) ParseTree (definitions *Definitions) {
	outputSet.BaseElement.ParseTree(definitions)

}

