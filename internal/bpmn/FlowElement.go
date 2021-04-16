package bpmn

import (
	"encoding/xml"
)

// FlowElement from BPMN
type FlowElement struct {
	BaseElement
	XMLName          xml.Name
	Auditing         Auditing   `xml:"auditing"`
	Monitoring       Monitoring `xml:"monitoring"`
	CategoryValueRef []string   `xml:"categoryValueRef"`
	Name             string     `xml:"name,attr"`

}

// ParseTree of components of FlowElement.
func (flowElement *FlowElement) ParseTree (definitions *Definitions) {
	flowElement.BaseElement.ParseTree(definitions)

	flowElement.Auditing.ParseTree(definitions) // Auditing.

	flowElement.Monitoring.ParseTree(definitions) // Monitoring.

}

