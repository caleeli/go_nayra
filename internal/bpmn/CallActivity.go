package bpmn

import (
	"encoding/xml"
)

// CallActivity from BPMN
type CallActivity struct {
	Activity
	CallActivityTrait `bson:"-"`
	XMLName           xml.Name
	CalledElement     string `xml:"calledElement,attr"`
}

// ParseTree of components of CallActivity.
func (callActivity *CallActivity) ParseTree(definitions *Definitions) {
	callActivity.Activity.ParseTree(definitions)
	callActivity.Init(definitions)
}
