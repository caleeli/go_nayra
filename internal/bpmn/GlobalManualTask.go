package bpmn

import (
	"encoding/xml"
)

// GlobalManualTask from BPMN
type GlobalManualTask struct {
	GlobalTask
	XMLName xml.Name

}

// ParseTree of components of GlobalManualTask.
func (globalManualTask *GlobalManualTask) ParseTree (definitions *Definitions) {
	globalManualTask.GlobalTask.ParseTree(definitions)

}

