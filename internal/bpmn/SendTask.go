package bpmn

import (
	"encoding/xml"
)

// SendTask from BPMN
type SendTask struct {
	Task
	XMLName        xml.Name
	Implementation string   `xml:"implementation,attr"`
	MessageRef     string   `xml:"messageRef,attr"`
	OperationRef   string   `xml:"operationRef,attr"`

}

// ParseTree of components of SendTask.
func (sendTask *SendTask) ParseTree (definitions *Definitions) {
	sendTask.Task.ParseTree(definitions)

}

