package bpmn

import (
	"encoding/xml"
)

// Interface from BPMN
type Interface struct {
	RootElement
	XMLName           xml.Name
	Operation         []Operation `xml:"operation"`
	Name              string      `xml:"name,attr"`
	ImplementationRef string      `xml:"implementationRef,attr"`
}

// ParseTree of components of Interface.
func (_interface *Interface) ParseTree(definitions *Definitions) {
	_interface.RootElement.ParseTree(definitions)

	for i := 0; i < len(_interface.Operation); i++ {
		_interface.Operation[i].ParseTree(definitions)
	}

}

// GetOperation by ID
func (_interface *Interface) GetOperation(ID string) *Operation {

	for i := 0; i < len(_interface.Operation); i++ {
		if _interface.Operation[i].ID == ID {
			return &_interface.Operation[i]

		}
	}

	return nil
}
