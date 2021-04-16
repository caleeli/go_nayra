package bpmn

import (
	"encoding/xml"
)

// Group from BPMN
type Group struct {
	Artifact
	XMLName          xml.Name
	CategoryValueRef string   `xml:"categoryValueRef,attr"`

}

// ParseTree of components of Group.
func (group *Group) ParseTree (definitions *Definitions) {
	group.Artifact.ParseTree(definitions)

}

