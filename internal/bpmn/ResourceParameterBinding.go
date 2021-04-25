package bpmn

import (
	"encoding/xml"
)

// ResourceParameterBinding from BPMN
type ResourceParameterBinding struct {
	BaseElement
	XMLName          xml.Name
	FormalExpression FormalExpression `xml:"formalExpression"`
	ParameterRef     string           `xml:"parameterRef,attr"`
}

// ParseTree of components of ResourceParameterBinding.
func (resourceParameterBinding *ResourceParameterBinding) ParseTree(definitions *Definitions) {
	resourceParameterBinding.BaseElement.ParseTree(definitions)

	resourceParameterBinding.FormalExpression.ParseTree(definitions) // FormalExpression.

}
