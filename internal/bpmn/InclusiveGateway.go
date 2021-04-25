package bpmn

import (
	"encoding/xml"
)

// InclusiveGateway from BPMN
type InclusiveGateway struct {
	Gateway
	XMLName xml.Name
	Default string `xml:"default,attr"`
}

// ParseTree of components of InclusiveGateway.
func (inclusiveGateway *InclusiveGateway) ParseTree(definitions *Definitions) {
	inclusiveGateway.Gateway.ParseTree(definitions)

}
