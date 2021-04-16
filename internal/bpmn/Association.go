package bpmn

import (
	"encoding/xml"
)

// Association from BPMN
type Association struct {
	Artifact
	XMLName              xml.Name
	SourceRef            string   `xml:"sourceRef,attr"`
	TargetRef            string   `xml:"targetRef,attr"`
	AssociationDirection string   `xml:"associationDirection,attr"`

}

// ParseTree of components of Association.
func (association *Association) ParseTree (definitions *Definitions) {
	association.Artifact.ParseTree(definitions)

}

