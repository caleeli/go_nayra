package bpmn

import (
	"encoding/xml"
)

// ComplexGateway from BPMN
type ComplexGateway struct {
	Gateway
	XMLName             xml.Name
	ActivationCondition []Expression `xml:"activationCondition"`
	Default             string       `xml:"default,attr"`
}

// ParseTree of components of ComplexGateway.
func (complexGateway *ComplexGateway) ParseTree(definitions *Definitions) {
	complexGateway.Gateway.ParseTree(definitions)

	for i := 0; i < len(complexGateway.ActivationCondition); i++ {
		complexGateway.ActivationCondition[i].ParseTree(definitions)
	}

}

// GetExpression by ID
func (complexGateway *ComplexGateway) GetExpression(ID string) *Expression {

	for i := 0; i < len(complexGateway.ActivationCondition); i++ {
		if complexGateway.ActivationCondition[i].ID == ID {
			return &complexGateway.ActivationCondition[i]

		}
	}

	return nil
}
