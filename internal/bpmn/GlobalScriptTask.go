package bpmn

import (
	"encoding/xml"
)

// GlobalScriptTask from BPMN
type GlobalScriptTask struct {
	GlobalTask
	XMLName        xml.Name
	Script         Script   `xml:"script"`
	ScriptLanguage string   `xml:"scriptLanguage,attr"`

}

// ParseTree of components of GlobalScriptTask.
func (globalScriptTask *GlobalScriptTask) ParseTree (definitions *Definitions) {
	globalScriptTask.GlobalTask.ParseTree(definitions)

	globalScriptTask.Script.ParseTree(definitions) // Script.

}

