package bpmn

import (
	"encoding/xml"
)

// Relationship from BPMN
type Relationship struct {
	BaseElement
	XMLName   xml.Name
	Source    []string `xml:"source"`
	Target    []string `xml:"target"`
	Type      string   `xml:"type,attr"`
	Direction string   `xml:"direction,attr"`
}

// ParseTree of components of Relationship.
func (relationship *Relationship) ParseTree(definitions *Definitions) {
	relationship.BaseElement.ParseTree(definitions)

}
