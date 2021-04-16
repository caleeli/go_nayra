package bpmn

import (
	"encoding/xml"
)

// EventBasedGateway from BPMN
type EventBasedGateway struct {
	Gateway
	XMLName          xml.Name
	Instantiate      bool     `xml:"instantiate,attr"`
	EventGatewayType string   `xml:"eventGatewayType,attr"`

}

// ParseTree of components of EventBasedGateway.
func (eventBasedGateway *EventBasedGateway) ParseTree (definitions *Definitions) {
	eventBasedGateway.Gateway.ParseTree(definitions)

}

