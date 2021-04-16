package bpmn

import (
	"encoding/xml"
)

// ResourceAssignmentExpression from BPMN
type ResourceAssignmentExpression struct {
	BaseElement
	XMLName          xml.Name
	FormalExpression FormalExpression `xml:"formalExpression"`

}

// ParseTree of components of ResourceAssignmentExpression.
func (resourceAssignmentExpression *ResourceAssignmentExpression) ParseTree (definitions *Definitions) {
	resourceAssignmentExpression.BaseElement.ParseTree(definitions)

	resourceAssignmentExpression.FormalExpression.ParseTree(definitions) // FormalExpression.

}

