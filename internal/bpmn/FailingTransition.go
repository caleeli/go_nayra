package bpmn

type FailingTransition struct {
	Transition
}

// Init transition
func (transition *FailingTransition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	definitions.Transitions = append(definitions.Transitions, transition)
}

// Execute transition
func (transition FailingTransition) Execute(instance *Instance) bool {
	inputTokens := transition.SelectInputTokens(instance)
	if inputTokens == nil {
		return false
	}
	if !transition.Condition(instance, inputTokens) {
		return false
	}
	transition.RemoveInputTokens(instance, inputTokens)
	transition.CreateOutputTokens(instance, inputTokens)
	return true
}

// Condition to activate the completing transition
func (transition *FailingTransition) Condition(instance *Instance, inputTokens []*Token) bool {
	return inputTokens[0].Transition == "FAIL"
}
