package bpmn

import (
	"encoding/xml"
)

// Message from BPMN
type Message struct {
	RootElement
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	ItemRef string `xml:"itemRef,attr"`
}

// ParseTree of components of Message.
func (message *Message) ParseTree(definitions *Definitions) {
	message.RootElement.ParseTree(definitions)

}
