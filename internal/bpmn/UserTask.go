package bpmn

import (
	"encoding/xml"
)

// UserTask from BPMN
type UserTask struct {
	Task
	XMLName        xml.Name
	Rendering      []Rendering `xml:"rendering"`
	Implementation string      `xml:"implementation,attr"`
}

// ParseTree of components of UserTask.
func (userTask *UserTask) ParseTree(definitions *Definitions) {
	userTask.Task.ParseTree(definitions)

	for i := 0; i < len(userTask.Rendering); i++ {
		userTask.Rendering[i].ParseTree(definitions)
	}

}

// GetRendering by ID
func (userTask *UserTask) GetRendering(ID string) *Rendering {

	for i := 0; i < len(userTask.Rendering); i++ {
		if userTask.Rendering[i].ID == ID {
			return &userTask.Rendering[i]

		}
	}

	return nil
}
