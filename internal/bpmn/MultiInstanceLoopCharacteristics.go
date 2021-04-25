package bpmn

import (
	"encoding/xml"
)

// MultiInstanceLoopCharacteristics from BPMN
type MultiInstanceLoopCharacteristics struct {
	LoopCharacteristics
	XMLName                   xml.Name
	LoopCardinality           []Expression                `xml:"loopCardinality"`
	LoopDataInputRef          []string                    `xml:"loopDataInputRef"`
	LoopDataOutputRef         []string                    `xml:"loopDataOutputRef"`
	InputDataItem             []DataInput                 `xml:"inputDataItem"`
	OutputDataItem            []DataOutput                `xml:"outputDataItem"`
	ComplexBehaviorDefinition []ComplexBehaviorDefinition `xml:"complexBehaviorDefinition"`
	CompletionCondition       []Expression                `xml:"completionCondition"`
	IsSequential              bool                        `xml:"isSequential,attr"`
	Behavior                  string                      `xml:"behavior,attr"`
	OneBehaviorEventRef       string                      `xml:"oneBehaviorEventRef,attr"`
	NoneBehaviorEventRef      string                      `xml:"noneBehaviorEventRef,attr"`
}

// ParseTree of components of MultiInstanceLoopCharacteristics.
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) ParseTree(definitions *Definitions) {
	multiInstanceLoopCharacteristics.LoopCharacteristics.ParseTree(definitions)

	for i := 0; i < len(multiInstanceLoopCharacteristics.LoopCardinality); i++ {
		multiInstanceLoopCharacteristics.LoopCardinality[i].ParseTree(definitions)
	}

	for i := 0; i < len(multiInstanceLoopCharacteristics.InputDataItem); i++ {
		multiInstanceLoopCharacteristics.InputDataItem[i].ParseTree(definitions)
	}

	for i := 0; i < len(multiInstanceLoopCharacteristics.OutputDataItem); i++ {
		multiInstanceLoopCharacteristics.OutputDataItem[i].ParseTree(definitions)
	}

	for i := 0; i < len(multiInstanceLoopCharacteristics.ComplexBehaviorDefinition); i++ {
		multiInstanceLoopCharacteristics.ComplexBehaviorDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(multiInstanceLoopCharacteristics.CompletionCondition); i++ {
		multiInstanceLoopCharacteristics.CompletionCondition[i].ParseTree(definitions)
	}

}

// GetExpression by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetExpression(ID string) *Expression {

	for i := 0; i < len(multiInstanceLoopCharacteristics.LoopCardinality); i++ {
		if multiInstanceLoopCharacteristics.LoopCardinality[i].ID == ID {
			return &multiInstanceLoopCharacteristics.LoopCardinality[i]

		}
	}

	for i := 0; i < len(multiInstanceLoopCharacteristics.CompletionCondition); i++ {
		if multiInstanceLoopCharacteristics.CompletionCondition[i].ID == ID {
			return &multiInstanceLoopCharacteristics.CompletionCondition[i]

		}
	}

	return nil
}

// GetDataInput by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetDataInput(ID string) *DataInput {

	for i := 0; i < len(multiInstanceLoopCharacteristics.InputDataItem); i++ {
		if multiInstanceLoopCharacteristics.InputDataItem[i].ID == ID {
			return &multiInstanceLoopCharacteristics.InputDataItem[i]

		}
	}

	return nil
}

// GetDataOutput by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetDataOutput(ID string) *DataOutput {

	for i := 0; i < len(multiInstanceLoopCharacteristics.OutputDataItem); i++ {
		if multiInstanceLoopCharacteristics.OutputDataItem[i].ID == ID {
			return &multiInstanceLoopCharacteristics.OutputDataItem[i]

		}
	}

	return nil
}

// GetComplexBehaviorDefinition by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetComplexBehaviorDefinition(ID string) *ComplexBehaviorDefinition {

	for i := 0; i < len(multiInstanceLoopCharacteristics.ComplexBehaviorDefinition); i++ {
		if multiInstanceLoopCharacteristics.ComplexBehaviorDefinition[i].ID == ID {
			return &multiInstanceLoopCharacteristics.ComplexBehaviorDefinition[i]

		}
	}

	return nil
}

// GetFormalExpression by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(multiInstanceLoopCharacteristics.ComplexBehaviorDefinition); i++ {
		FormalExpression := multiInstanceLoopCharacteristics.ComplexBehaviorDefinition[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(multiInstanceLoopCharacteristics.ComplexBehaviorDefinition); i++ {
		ImplicitThrowEvent := multiInstanceLoopCharacteristics.ComplexBehaviorDefinition[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}
