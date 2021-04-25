package bpmn

import (
	"encoding/xml"
)

// GlobalBusinessRuleTask from BPMN
type GlobalBusinessRuleTask struct {
	GlobalTask
	XMLName        xml.Name
	Implementation string `xml:"implementation,attr"`
}

// ParseTree of components of GlobalBusinessRuleTask.
func (globalBusinessRuleTask *GlobalBusinessRuleTask) ParseTree(definitions *Definitions) {
	globalBusinessRuleTask.GlobalTask.ParseTree(definitions)

}
