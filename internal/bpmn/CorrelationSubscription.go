package bpmn

import (
	"encoding/xml"
)

// CorrelationSubscription from BPMN
type CorrelationSubscription struct {
	BaseElement
	XMLName                    xml.Name
	CorrelationPropertyBinding []CorrelationPropertyBinding `xml:"correlationPropertyBinding"`
	CorrelationKeyRef          string                       `xml:"correlationKeyRef,attr"`
}

// ParseTree of components of CorrelationSubscription.
func (correlationSubscription *CorrelationSubscription) ParseTree(definitions *Definitions) {
	correlationSubscription.BaseElement.ParseTree(definitions)

	for i := 0; i < len(correlationSubscription.CorrelationPropertyBinding); i++ {
		correlationSubscription.CorrelationPropertyBinding[i].ParseTree(definitions)
	}

}

// GetCorrelationPropertyBinding by ID
func (correlationSubscription *CorrelationSubscription) GetCorrelationPropertyBinding(ID string) *CorrelationPropertyBinding {

	for i := 0; i < len(correlationSubscription.CorrelationPropertyBinding); i++ {
		if correlationSubscription.CorrelationPropertyBinding[i].ID == ID {
			return &correlationSubscription.CorrelationPropertyBinding[i]

		}
	}

	return nil
}

// GetFormalExpression by ID
func (correlationSubscription *CorrelationSubscription) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(correlationSubscription.CorrelationPropertyBinding); i++ {
		FormalExpression := correlationSubscription.CorrelationPropertyBinding[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	return nil
}
