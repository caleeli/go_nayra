package bpmn

import (
	"encoding/xml"
)

// TextAnnotation from BPMN
type TextAnnotation struct {
	Artifact
	XMLName    xml.Name
	Text       Text   `xml:"text"`
	TextFormat string `xml:"textFormat,attr"`
}

// ParseTree of components of TextAnnotation.
func (textAnnotation *TextAnnotation) ParseTree(definitions *Definitions) {
	textAnnotation.Artifact.ParseTree(definitions)

	textAnnotation.Text.ParseTree(definitions) // Text.

}
