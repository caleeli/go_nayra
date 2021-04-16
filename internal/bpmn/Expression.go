package bpmn

import (
	"encoding/xml"
)

// Expression from BPMN
type Expression struct {
	BaseElementWithMixedContent
	XMLName xml.Name

}

// ParseTree of components of Expression.
func (expression *Expression) ParseTree (definitions *Definitions) {
	expression.BaseElementWithMixedContent.ParseTree(definitions)

}

