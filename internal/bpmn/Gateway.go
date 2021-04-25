package bpmn

import (
	"encoding/xml"
)

// Gateway from BPMN
type Gateway struct {
	FlowNode
	XMLName          xml.Name
	GatewayDirection string `xml:"gatewayDirection,attr"`
}

// ParseTree of components of Gateway.
func (gateway *Gateway) ParseTree(definitions *Definitions) {
	gateway.FlowNode.ParseTree(definitions)

}
