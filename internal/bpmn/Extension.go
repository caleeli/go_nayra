package bpmn

import (
	"encoding/xml"
)

// Extension from BPMN
type Extension struct {
	XMLName        xml.Name
	Documentation  []Documentation `xml:"documentation"`
	Definition     string          `xml:"definition,attr"`
	MustUnderstand bool            `xml:"mustUnderstand,attr"`
}

// ParseTree of components of Extension.
func (extension *Extension) ParseTree(definitions *Definitions) {

	for i := 0; i < len(extension.Documentation); i++ {
		extension.Documentation[i].ParseTree(definitions)
	}

}

// GetDocumentation by ID
func (extension *Extension) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(extension.Documentation); i++ {
		if extension.Documentation[i].ID == ID {
			return &extension.Documentation[i]

		}
	}

	return nil
}
