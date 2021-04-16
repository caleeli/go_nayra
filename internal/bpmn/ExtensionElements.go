package bpmn

import (
	"encoding/xml"
)

// ExtensionElements from BPMN
type ExtensionElements struct {
	XMLName xml.Name

}

// ParseTree of components of ExtensionElements.
func (extensionElements *ExtensionElements) ParseTree (definitions *Definitions) {

}

