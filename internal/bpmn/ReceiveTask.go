package bpmn

import (
	"encoding/xml"
)

// ReceiveTask from BPMN
type ReceiveTask struct {
	Task
	XMLName        xml.Name
	Implementation string   `xml:"implementation,attr"`
	Instantiate    bool     `xml:"instantiate,attr"`
	MessageRef     string   `xml:"messageRef,attr"`
	OperationRef   string   `xml:"operationRef,attr"`

}

// ParseTree of components of ReceiveTask.
func (receiveTask *ReceiveTask) ParseTree (definitions *Definitions) {
	receiveTask.Task.ParseTree(definitions)

}

