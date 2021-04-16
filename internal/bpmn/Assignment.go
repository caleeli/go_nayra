package bpmn

import (
	"encoding/xml"
)

// Assignment from BPMN
type Assignment struct {
	BaseElement
	XMLName xml.Name
	From    []Expression `xml:"from"`
	To      []Expression `xml:"to"`

}

// ParseTree of components of Assignment.
func (assignment *Assignment) ParseTree (definitions *Definitions) {
	assignment.BaseElement.ParseTree(definitions)

	for i := 0; i < len(assignment.From); i++ {
		assignment.From[i].ParseTree(definitions)
	}

	for i := 0; i < len(assignment.To); i++ {
		assignment.To[i].ParseTree(definitions)
	}

}

// GetExpression by ID
func (assignment *Assignment) GetExpression(ID string) *Expression {

	for i := 0; i < len(assignment.From); i++ {
		if assignment.From[i].ID == ID {
			return &assignment.From[i]

		}
	}

	for i := 0; i < len(assignment.To); i++ {
		if assignment.To[i].ID == ID {
			return &assignment.To[i]

		}
	}

	return nil
}

