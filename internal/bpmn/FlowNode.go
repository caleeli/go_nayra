package bpmn

import (
	"encoding/xml"
)

// FlowNode from BPMN
type FlowNode struct {
	FlowElement
	XMLName  xml.Name
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`

}

// ParseTree of components of FlowNode.
func (flowNode *FlowNode) ParseTree (definitions *Definitions) {
	flowNode.FlowElement.ParseTree(definitions)

}

