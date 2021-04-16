package bpmn

import (
	"encoding/xml"
)

// GlobalTask from BPMN
type GlobalTask struct {
	CallableElement
	XMLName   xml.Name
	Performer []Performer `xml:"performer"`

}

// ParseTree of components of GlobalTask.
func (globalTask *GlobalTask) ParseTree (definitions *Definitions) {
	globalTask.CallableElement.ParseTree(definitions)

	for i := 0; i < len(globalTask.Performer); i++ {
		globalTask.Performer[i].ParseTree(definitions)
	}

}

// GetPerformer by ID
func (globalTask *GlobalTask) GetPerformer(ID string) *Performer {

	for i := 0; i < len(globalTask.Performer); i++ {
		if globalTask.Performer[i].ID == ID {
			return &globalTask.Performer[i]

		}
	}

	return nil
}

