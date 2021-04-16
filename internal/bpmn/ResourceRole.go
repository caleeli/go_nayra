package bpmn

import (
	"encoding/xml"
)

// ResourceRole from BPMN
type ResourceRole struct {
	BaseElement
	XMLName xml.Name
	Name    string   `xml:"name,attr"`

}

// ParseTree of components of ResourceRole.
func (resourceRole *ResourceRole) ParseTree (definitions *Definitions) {
	resourceRole.BaseElement.ParseTree(definitions)

}

