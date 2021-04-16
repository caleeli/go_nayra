package bpmn

import (
	"encoding/xml"
)

// CorrelationPropertyRetrievalExpression from BPMN
type CorrelationPropertyRetrievalExpression struct {
	BaseElement
	XMLName     xml.Name
	MessagePath []FormalExpression `xml:"messagePath"`
	MessageRef  string             `xml:"messageRef,attr"`

}

// ParseTree of components of CorrelationPropertyRetrievalExpression.
func (correlationPropertyRetrievalExpression *CorrelationPropertyRetrievalExpression) ParseTree (definitions *Definitions) {
	correlationPropertyRetrievalExpression.BaseElement.ParseTree(definitions)

	for i := 0; i < len(correlationPropertyRetrievalExpression.MessagePath); i++ {
		correlationPropertyRetrievalExpression.MessagePath[i].ParseTree(definitions)
	}

}

// GetFormalExpression by ID
func (correlationPropertyRetrievalExpression *CorrelationPropertyRetrievalExpression) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(correlationPropertyRetrievalExpression.MessagePath); i++ {
		if correlationPropertyRetrievalExpression.MessagePath[i].ID == ID {
			return &correlationPropertyRetrievalExpression.MessagePath[i]

		}
	}

	return nil
}

