package bpmn

import (
	"encoding/xml"
)

// StandardLoopCharacteristics from BPMN
type StandardLoopCharacteristics struct {
	LoopCharacteristics
	XMLName       xml.Name
	LoopCondition Expression `xml:"loopCondition"`
	TestBefore    bool       `xml:"testBefore,attr"`
	LoopMaximum   int        `xml:"loopMaximum,attr"`
}

// ParseTree of components of StandardLoopCharacteristics.
func (standardLoopCharacteristics *StandardLoopCharacteristics) ParseTree(definitions *Definitions) {
	standardLoopCharacteristics.LoopCharacteristics.ParseTree(definitions)

	standardLoopCharacteristics.LoopCondition.ParseTree(definitions) // Expression.

}
