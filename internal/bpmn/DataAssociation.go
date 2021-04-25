package bpmn

import (
	"encoding/xml"
)

// DataAssociation from BPMN
type DataAssociation struct {
	BaseElement
	XMLName        xml.Name
	SourceRef      []string           `xml:"sourceRef"`
	TargetRef      []string           `xml:"targetRef"`
	Transformation []FormalExpression `xml:"transformation"`
	Assignment     []Assignment       `xml:"assignment"`
}

// ParseTree of components of DataAssociation.
func (dataAssociation *DataAssociation) ParseTree(definitions *Definitions) {
	dataAssociation.BaseElement.ParseTree(definitions)

	for i := 0; i < len(dataAssociation.Transformation); i++ {
		dataAssociation.Transformation[i].ParseTree(definitions)
	}

	for i := 0; i < len(dataAssociation.Assignment); i++ {
		dataAssociation.Assignment[i].ParseTree(definitions)
	}

}

// GetFormalExpression by ID
func (dataAssociation *DataAssociation) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(dataAssociation.Transformation); i++ {
		if dataAssociation.Transformation[i].ID == ID {
			return &dataAssociation.Transformation[i]

		}
	}

	return nil
}

// GetAssignment by ID
func (dataAssociation *DataAssociation) GetAssignment(ID string) *Assignment {

	for i := 0; i < len(dataAssociation.Assignment); i++ {
		if dataAssociation.Assignment[i].ID == ID {
			return &dataAssociation.Assignment[i]

		}
	}

	return nil
}

// GetExpression by ID
func (dataAssociation *DataAssociation) GetExpression(ID string) *Expression {

	for i := 0; i < len(dataAssociation.Assignment); i++ {
		Expression := dataAssociation.Assignment[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}
