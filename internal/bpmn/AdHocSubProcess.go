package bpmn

import (
	"encoding/xml"
)

// AdHocSubProcess from BPMN
type AdHocSubProcess struct {
	SubProcess
	XMLName                  xml.Name
	CompletionCondition      []Expression `xml:"completionCondition"`
	CancelRemainingInstances bool         `xml:"cancelRemainingInstances,attr"`
	Ordering                 string       `xml:"ordering,attr"`

}

// ParseTree of components of AdHocSubProcess.
func (adHocSubProcess *AdHocSubProcess) ParseTree (definitions *Definitions) {
	adHocSubProcess.SubProcess.ParseTree(definitions)

	for i := 0; i < len(adHocSubProcess.CompletionCondition); i++ {
		adHocSubProcess.CompletionCondition[i].ParseTree(definitions)
	}

}

// GetExpression by ID
func (adHocSubProcess *AdHocSubProcess) GetExpression(ID string) *Expression {

	for i := 0; i < len(adHocSubProcess.CompletionCondition); i++ {
		if adHocSubProcess.CompletionCondition[i].ID == ID {
			return &adHocSubProcess.CompletionCondition[i]

		}
	}

	return nil
}

