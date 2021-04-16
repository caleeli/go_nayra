package bpmn

import (
	"encoding/xml"
)

// LinkEventDefinition from BPMN
type LinkEventDefinition struct {
	EventDefinition
	XMLName xml.Name
	Source  []string `xml:"source"`
	Target  []string `xml:"target"`
	Name    string   `xml:"name,attr"`

}

// ParseTree of components of LinkEventDefinition.
func (linkEventDefinition *LinkEventDefinition) ParseTree (definitions *Definitions) {
	linkEventDefinition.EventDefinition.ParseTree(definitions)

}

