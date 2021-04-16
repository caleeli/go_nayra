package bpmn

import (
	"encoding/xml"
)

// LoopCharacteristics from BPMN
type LoopCharacteristics struct {
	BaseElement
	XMLName xml.Name

}

// ParseTree of components of LoopCharacteristics.
func (loopCharacteristics *LoopCharacteristics) ParseTree (definitions *Definitions) {
	loopCharacteristics.BaseElement.ParseTree(definitions)

}

