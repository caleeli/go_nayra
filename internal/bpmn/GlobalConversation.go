package bpmn

import (
	"encoding/xml"
)

// GlobalConversation from BPMN
type GlobalConversation struct {
	Collaboration
	XMLName xml.Name
}

// ParseTree of components of GlobalConversation.
func (globalConversation *GlobalConversation) ParseTree(definitions *Definitions) {
	globalConversation.Collaboration.ParseTree(definitions)

}
