package bpmn

import (
	"encoding/xml"
)

// CategoryValue from BPMN
type CategoryValue struct {
	BaseElement
	XMLName xml.Name
	Value   string `xml:"value,attr"`
}

// ParseTree of components of CategoryValue.
func (categoryValue *CategoryValue) ParseTree(definitions *Definitions) {
	categoryValue.BaseElement.ParseTree(definitions)

}
