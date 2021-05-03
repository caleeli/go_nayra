package bpmn

type CompletingTransition struct {
	Transition
}

const (
	TOKEN_TRANSITION_COMPLETE = "COMPLETE"
)

// Init transition
func (transition *CompletingTransition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	definitions.Transitions = append(definitions.Transitions, transition)
}

// Execute transition
func (transition CompletingTransition) Execute(instance *Instance) bool {
	inputTokens := transition.SelectInputTokens(instance)
	if inputTokens == nil {
		return false
	}
	if !transition.Condition(instance, inputTokens) {
		return false
	}
	transition.CreateOutputTokens(instance, inputTokens)
	transition.RemoveInputTokens(instance, inputTokens)
	return true
}

// Condition to activate the completing transition
func (transition *CompletingTransition) Condition(instance *Instance, inputTokens []*Token) bool {
	return inputTokens[0].Transition == TOKEN_TRANSITION_COMPLETE
}

// Complete an Activity instance by token
func (activity *Activity) Complete(token *Token) {
	token.Transition = TOKEN_TRANSITION_COMPLETE
}
