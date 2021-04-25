package bpmn

import (
	"encoding/xml"
)

// TimerEventDefinition from BPMN
type TimerEventDefinition struct {
	EventDefinition
	XMLName xml.Name
}

// ParseTree of components of TimerEventDefinition.
func (timerEventDefinition *TimerEventDefinition) ParseTree(definitions *Definitions) {
	timerEventDefinition.EventDefinition.ParseTree(definitions)

}
