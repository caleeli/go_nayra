package bpmn

import (
	"encoding/xml"
)

// CorrelationProperty from BPMN
type CorrelationProperty struct {
	RootElement
	XMLName                                xml.Name
	CorrelationPropertyRetrievalExpression []CorrelationPropertyRetrievalExpression `xml:"correlationPropertyRetrievalExpression"`
	Name                                   string                                   `xml:"name,attr"`
	Type                                   string                                   `xml:"type,attr"`
}

// ParseTree of components of CorrelationProperty.
func (correlationProperty *CorrelationProperty) ParseTree(definitions *Definitions) {
	correlationProperty.RootElement.ParseTree(definitions)

	for i := 0; i < len(correlationProperty.CorrelationPropertyRetrievalExpression); i++ {
		correlationProperty.CorrelationPropertyRetrievalExpression[i].ParseTree(definitions)
	}

}

// GetCorrelationPropertyRetrievalExpression by ID
func (correlationProperty *CorrelationProperty) GetCorrelationPropertyRetrievalExpression(ID string) *CorrelationPropertyRetrievalExpression {

	for i := 0; i < len(correlationProperty.CorrelationPropertyRetrievalExpression); i++ {
		if correlationProperty.CorrelationPropertyRetrievalExpression[i].ID == ID {
			return &correlationProperty.CorrelationPropertyRetrievalExpression[i]

		}
	}

	return nil
}

// GetFormalExpression by ID
func (correlationProperty *CorrelationProperty) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(correlationProperty.CorrelationPropertyRetrievalExpression); i++ {
		FormalExpression := correlationProperty.CorrelationPropertyRetrievalExpression[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	return nil
}
