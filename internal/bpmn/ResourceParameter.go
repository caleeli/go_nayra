package bpmn

import (
	"encoding/xml"
)

// ResourceParameter from BPMN
type ResourceParameter struct {
	BaseElement
	XMLName    xml.Name
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	IsRequired bool   `xml:"isRequired,attr"`
}

// ParseTree of components of ResourceParameter.
func (resourceParameter *ResourceParameter) ParseTree(definitions *Definitions) {
	resourceParameter.BaseElement.ParseTree(definitions)

}
