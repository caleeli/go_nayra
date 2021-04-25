package bpmn

import (
	"encoding/xml"
)

// ItemDefinition from BPMN
type ItemDefinition struct {
	RootElement
	XMLName      xml.Name
	StructureRef string `xml:"structureRef,attr"`
	IsCollection bool   `xml:"isCollection,attr"`
	ItemKind     string `xml:"itemKind,attr"`
}

// ParseTree of components of ItemDefinition.
func (itemDefinition *ItemDefinition) ParseTree(definitions *Definitions) {
	itemDefinition.RootElement.ParseTree(definitions)

}
