package bpmn

import (
	"encoding/xml"
)

// Monitoring from BPMN
type Monitoring struct {
	BaseElement
	XMLName xml.Name
}

// ParseTree of components of Monitoring.
func (monitoring *Monitoring) ParseTree(definitions *Definitions) {
	monitoring.BaseElement.ParseTree(definitions)

}
