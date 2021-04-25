package bpmn

import (
	"encoding/xml"
)

// CorrelationPropertyBinding from BPMN
type CorrelationPropertyBinding struct {
	BaseElement
	XMLName                xml.Name
	DataPath               []FormalExpression `xml:"dataPath"`
	CorrelationPropertyRef string             `xml:"correlationPropertyRef,attr"`
}

// ParseTree of components of CorrelationPropertyBinding.
func (correlationPropertyBinding *CorrelationPropertyBinding) ParseTree(definitions *Definitions) {
	correlationPropertyBinding.BaseElement.ParseTree(definitions)

	for i := 0; i < len(correlationPropertyBinding.DataPath); i++ {
		correlationPropertyBinding.DataPath[i].ParseTree(definitions)
	}

}

// GetFormalExpression by ID
func (correlationPropertyBinding *CorrelationPropertyBinding) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(correlationPropertyBinding.DataPath); i++ {
		if correlationPropertyBinding.DataPath[i].ID == ID {
			return &correlationPropertyBinding.DataPath[i]

		}
	}

	return nil
}
