package bpmn

import (
	"encoding/xml"
)

// FormalExpression from BPMN
type FormalExpression struct {
	Expression
	XMLName            xml.Name
	Language           string `xml:"language,attr"`
	EvaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}

// ParseTree of components of FormalExpression.
func (formalExpression *FormalExpression) ParseTree(definitions *Definitions) {
	formalExpression.Expression.ParseTree(definitions)

}
