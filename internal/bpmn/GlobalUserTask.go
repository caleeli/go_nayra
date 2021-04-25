package bpmn

import (
	"encoding/xml"
)

// GlobalUserTask from BPMN
type GlobalUserTask struct {
	GlobalTask
	XMLName        xml.Name
	Rendering      []Rendering `xml:"rendering"`
	Implementation string      `xml:"implementation,attr"`
}

// ParseTree of components of GlobalUserTask.
func (globalUserTask *GlobalUserTask) ParseTree(definitions *Definitions) {
	globalUserTask.GlobalTask.ParseTree(definitions)

	for i := 0; i < len(globalUserTask.Rendering); i++ {
		globalUserTask.Rendering[i].ParseTree(definitions)
	}

}

// GetRendering by ID
func (globalUserTask *GlobalUserTask) GetRendering(ID string) *Rendering {

	for i := 0; i < len(globalUserTask.Rendering); i++ {
		if globalUserTask.Rendering[i].ID == ID {
			return &globalUserTask.Rendering[i]

		}
	}

	return nil
}
