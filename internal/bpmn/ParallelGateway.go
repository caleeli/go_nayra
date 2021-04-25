package bpmn

import (
	"encoding/xml"
)

// ParallelGateway from BPMN
type ParallelGateway struct {
	Gateway
	XMLName xml.Name
}

// ParseTree of components of ParallelGateway.
func (parallelGateway *ParallelGateway) ParseTree(definitions *Definitions) {
	parallelGateway.Gateway.ParseTree(definitions)

}
