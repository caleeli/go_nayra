package bpmn

import (
	"encoding/xml"
)

// EndPoint from BPMN
type EndPoint struct {
	RootElement
	XMLName xml.Name
}

// ParseTree of components of EndPoint.
func (endPoint *EndPoint) ParseTree(definitions *Definitions) {
	endPoint.RootElement.ParseTree(definitions)

}
