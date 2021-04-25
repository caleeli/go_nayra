package bpmn

import (
	"encoding/xml"
)

// InputOutputSpecification from BPMN
type InputOutputSpecification struct {
	BaseElement
	XMLName    xml.Name
	DataInput  []DataInput  `xml:"dataInput"`
	DataOutput []DataOutput `xml:"dataOutput"`
	InputSet   []InputSet   `xml:"inputSet"`
	OutputSet  []OutputSet  `xml:"outputSet"`
}

// ParseTree of components of InputOutputSpecification.
func (inputOutputSpecification *InputOutputSpecification) ParseTree(definitions *Definitions) {
	inputOutputSpecification.BaseElement.ParseTree(definitions)

	for i := 0; i < len(inputOutputSpecification.DataInput); i++ {
		inputOutputSpecification.DataInput[i].ParseTree(definitions)
	}

	for i := 0; i < len(inputOutputSpecification.DataOutput); i++ {
		inputOutputSpecification.DataOutput[i].ParseTree(definitions)
	}

	for i := 0; i < len(inputOutputSpecification.InputSet); i++ {
		inputOutputSpecification.InputSet[i].ParseTree(definitions)
	}

	for i := 0; i < len(inputOutputSpecification.OutputSet); i++ {
		inputOutputSpecification.OutputSet[i].ParseTree(definitions)
	}

}

// GetDataInput by ID
func (inputOutputSpecification *InputOutputSpecification) GetDataInput(ID string) *DataInput {

	for i := 0; i < len(inputOutputSpecification.DataInput); i++ {
		if inputOutputSpecification.DataInput[i].ID == ID {
			return &inputOutputSpecification.DataInput[i]

		}
	}

	return nil
}

// GetDataOutput by ID
func (inputOutputSpecification *InputOutputSpecification) GetDataOutput(ID string) *DataOutput {

	for i := 0; i < len(inputOutputSpecification.DataOutput); i++ {
		if inputOutputSpecification.DataOutput[i].ID == ID {
			return &inputOutputSpecification.DataOutput[i]

		}
	}

	return nil
}

// GetInputSet by ID
func (inputOutputSpecification *InputOutputSpecification) GetInputSet(ID string) *InputSet {

	for i := 0; i < len(inputOutputSpecification.InputSet); i++ {
		if inputOutputSpecification.InputSet[i].ID == ID {
			return &inputOutputSpecification.InputSet[i]

		}
	}

	return nil
}

// GetOutputSet by ID
func (inputOutputSpecification *InputOutputSpecification) GetOutputSet(ID string) *OutputSet {

	for i := 0; i < len(inputOutputSpecification.OutputSet); i++ {
		if inputOutputSpecification.OutputSet[i].ID == ID {
			return &inputOutputSpecification.OutputSet[i]

		}
	}

	return nil
}
