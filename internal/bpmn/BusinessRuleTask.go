package bpmn

import (
	"encoding/xml"
)

// BusinessRuleTask from BPMN
type BusinessRuleTask struct {
	Task
	XMLName        xml.Name
	Implementation string   `xml:"implementation,attr"`

}

// ParseTree of components of BusinessRuleTask.
func (businessRuleTask *BusinessRuleTask) ParseTree (definitions *Definitions) {
	businessRuleTask.Task.ParseTree(definitions)

}

