package bpmn

import (
	"encoding/xml"
)

// ScriptTask from BPMN
type ScriptTask struct {
	Task
	ScriptTaskTrait `bson:"-"`
	XMLName         xml.Name
	Script          Script `xml:"script"`
	ScriptFormat    string `xml:"scriptFormat,attr"`
}

// ParseTree of components of ScriptTask.
func (scriptTask *ScriptTask) ParseTree(definitions *Definitions) {
	scriptTask.Task.ParseTree(definitions)

	scriptTask.Script.ParseTree(definitions) // Script.

	scriptTask.Init(definitions)

}
