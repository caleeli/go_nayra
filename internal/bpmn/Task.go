package bpmn

import (
	"encoding/xml"
)

// Task from BPMN
type Task struct {
	Activity
	XMLName xml.Name
}

// ParseTree of components of Task.
func (task *Task) ParseTree(definitions *Definitions) {
	task.Activity.ParseTree(definitions)

}
