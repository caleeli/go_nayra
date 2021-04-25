package bpmn

import (
	"encoding/xml"
)

// ManualTask from BPMN
type ManualTask struct {
	Task
	XMLName xml.Name
}

// ParseTree of components of ManualTask.
func (manualTask *ManualTask) ParseTree(definitions *Definitions) {
	manualTask.Task.ParseTree(definitions)

}
