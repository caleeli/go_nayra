package bpmn

import (
	"encoding/xml"
)

// ExclusiveGateway from BPMN
type ExclusiveGateway struct {
	Gateway
	XMLName xml.Name
	Default string `xml:"default,attr"`
}

// ParseTree of components of ExclusiveGateway.
func (exclusiveGateway *ExclusiveGateway) ParseTree(definitions *Definitions) {
	exclusiveGateway.Gateway.ParseTree(definitions)

}
