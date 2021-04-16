package bpmn

// NamedElementInterface from BPMN
type NamedElementInterface interface {
	GetName() string
}

// GetName of the FlowElement
func (flowElement *FlowElement) GetName() string {
	return flowElement.Name
}

// GetName of the CallableElement
func (callableElement *CallableElement) GetName() string {
	return callableElement.Name
}

// GetName of the CallableElement
func (callableElement *SequenceFlow) GetName() string {
	return callableElement.Name
}
