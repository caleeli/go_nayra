package bpmn

import (
	"encoding/xml"
)

// ComplexBehaviorDefinition from BPMN
type ComplexBehaviorDefinition struct {
	BaseElement
	XMLName   xml.Name
	Condition []FormalExpression   `xml:"condition"`
	Event     []ImplicitThrowEvent `xml:"event"`

}

// ParseTree of components of ComplexBehaviorDefinition.
func (complexBehaviorDefinition *ComplexBehaviorDefinition) ParseTree (definitions *Definitions) {
	complexBehaviorDefinition.BaseElement.ParseTree(definitions)

	for i := 0; i < len(complexBehaviorDefinition.Condition); i++ {
		complexBehaviorDefinition.Condition[i].ParseTree(definitions)
	}

	for i := 0; i < len(complexBehaviorDefinition.Event); i++ {
		complexBehaviorDefinition.Event[i].ParseTree(definitions)
	}

}

// GetFormalExpression by ID
func (complexBehaviorDefinition *ComplexBehaviorDefinition) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(complexBehaviorDefinition.Condition); i++ {
		if complexBehaviorDefinition.Condition[i].ID == ID {
			return &complexBehaviorDefinition.Condition[i]

		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (complexBehaviorDefinition *ComplexBehaviorDefinition) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(complexBehaviorDefinition.Event); i++ {
		if complexBehaviorDefinition.Event[i].ID == ID {
			return &complexBehaviorDefinition.Event[i]

		}
	}

	return nil
}

