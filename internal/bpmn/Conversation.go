package bpmn

import (
	"encoding/xml"
)

// Conversation from BPMN
type Conversation struct {
	ConversationNode
	XMLName xml.Name
}

// ParseTree of components of Conversation.
func (conversation *Conversation) ParseTree(definitions *Definitions) {
	conversation.ConversationNode.ParseTree(definitions)

}
