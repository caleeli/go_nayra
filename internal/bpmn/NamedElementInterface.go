package bpmn

// NamedElementInterface from BPMN
type NamedElementInterface interface {
	GetName() string
	GetId() string
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

// GetName of the BaseElement (@todo split to a BaseElementInterface)
func (baseElement *BaseElement) GetName() string {
	return ""
}

// GetName of the FlowElement
func (flowElement *FlowElement) GetId() string {
	return flowElement.ID
}

// GetName of the CallableElement
func (callableElement *CallableElement) GetId() string {
	return callableElement.ID
}

// GetName of the CallableElement
func (callableElement *SequenceFlow) GetId() string {
	return callableElement.ID
}

// GetName of the BaseElement (@todo split to a BaseElementInterface)
func (baseElement *BaseElement) GetId() string {
	return baseElement.ID
}
