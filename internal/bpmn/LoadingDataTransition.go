package bpmn

type LoadingDataTransition struct {
	Transition
}

// Init transition
func (transition *LoadingDataTransition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	definitions.Transitions = append(definitions.Transitions, transition)
}

// Execute transition
func (transition LoadingDataTransition) Execute(instance *Instance) bool {
	inputTokens := transition.SelectInputTokens(instance)
	if inputTokens == nil {
		return false
	}
	transition.RemoveInputTokens(instance, inputTokens)
	times := 1
	for i := 0; i < times; i++ {
		transition.CreateOutputTokens(instance, inputTokens)
	}
	return true
}
