package bpmn

import (
	"encoding/xml"
)

// SequenceFlow from BPMN
type SequenceFlow struct {
	SequenceFlowTrait `bson:"-"`
	FlowElement
	XMLName             xml.Name
	ConditionExpression []Expression `xml:"conditionExpression"`
	SourceRef           string       `xml:"sourceRef,attr"`
	TargetRef           string       `xml:"targetRef,attr"`
	IsImmediate         bool         `xml:"isImmediate,attr"`
}

// ParseTree of components of SequenceFlow.
func (sequenceFlow *SequenceFlow) ParseTree(definitions *Definitions) {
	sequenceFlow.FlowElement.ParseTree(definitions)

	sequenceFlow.Init(definitions)

	for i := 0; i < len(sequenceFlow.ConditionExpression); i++ {
		sequenceFlow.ConditionExpression[i].ParseTree(definitions)
	}

}

// GetExpression by ID
func (sequenceFlow *SequenceFlow) GetExpression(ID string) *Expression {

	for i := 0; i < len(sequenceFlow.ConditionExpression); i++ {
		if sequenceFlow.ConditionExpression[i].ID == ID {
			return &sequenceFlow.ConditionExpression[i]

		}
	}

	return nil
}
