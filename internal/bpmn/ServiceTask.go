package bpmn

import (
	"encoding/xml"
)

// ServiceTask from BPMN
type ServiceTask struct {
	Task
	XMLName        xml.Name
	Implementation string   `xml:"implementation,attr"`
	OperationRef   string   `xml:"operationRef,attr"`

}

// ParseTree of components of ServiceTask.
func (serviceTask *ServiceTask) ParseTree (definitions *Definitions) {
	serviceTask.Task.ParseTree(definitions)

}

