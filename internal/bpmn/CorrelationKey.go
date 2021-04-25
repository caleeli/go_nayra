package bpmn

import (
	"encoding/xml"
)

// CorrelationKey from BPMN
type CorrelationKey struct {
	BaseElement
	XMLName                xml.Name
	CorrelationPropertyRef []string `xml:"correlationPropertyRef"`
	Name                   string   `xml:"name,attr"`
}

// ParseTree of components of CorrelationKey.
func (correlationKey *CorrelationKey) ParseTree(definitions *Definitions) {
	correlationKey.BaseElement.ParseTree(definitions)

}
