package bpmn

// InitInterface requires Init implementation
type InitInterface interface {
	Init(definitions *Definitions)
}
