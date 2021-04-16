package bpmn

// Transition from Nayra
type Transition struct {
	Node
	Incoming        []*State
	Outgoing        []*State
	RemoveAllTokens bool
}

type TransitionInterface interface {
	NodeInterface
	Init(definitions *Definitions, owner NamedElementInterface, name string)
	Connect(target *State)
	Execute(instance *Instance) bool
	SelectInputTokens(instance *Instance) []*Token
	RemoveInputTokens(instance *Instance, inputTokens []*Token)
	CreateOutputTokens(instance *Instance, inputTokens []*Token) []*Token
	Condition(instance *Instance, inputTokens []*Token) bool
	AppendIncoming(state *State)
	GetName() string
	GetOwner() NamedElementInterface
}

// Init transition
func (transition *Transition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	definitions.Transitions = append(definitions.Transitions, transition)
}

// Connect transition to state
func (transition *Transition) Connect(target *State) {
	transition.Outgoing = append(transition.Outgoing, target)
	target.Incoming = append(target.Incoming, transition)
}

// Execute transition
func (transition Transition) Execute(instance *Instance) bool {
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

// SelectInputTokens it gets the required input tokens
func (transition *Transition) SelectInputTokens(instance *Instance) []*Token {
	inputTokens := []*Token{}
	for _, Incoming := range transition.Incoming {
		if selected := Incoming.Select(instance, transition.RemoveAllTokens); len(selected) > 0 {
			inputTokens = append(inputTokens, selected...)
		} else {
			return nil
		}
	}
	return inputTokens
}

// RemoveInputTokens it consumes the required input tokens
func (transition *Transition) RemoveInputTokens(instance *Instance, inputTokens []*Token) {
	for _, token := range inputTokens {
		token.Owner.RemoveToken(token)
	}
}

// CreateOutputTokens it creates tokens in output places
func (transition *Transition) CreateOutputTokens(instance *Instance, inputTokens []*Token) []*Token {
	outputTokens := []*Token{}
	for _, Outgoing := range transition.Outgoing {
		outputTokens = append(outputTokens, Outgoing.CreateToken(instance))
	}
	return outputTokens
}

// Condition to activate the transition
func (transition *Transition) Condition(instance *Instance, inputTokens []*Token) bool {
	return true
}

// Condition to activate the transition
func (transition *Transition) AppendIncoming(state *State) {
	transition.Incoming = append(transition.Incoming, state)
}

func (transition *Transition) GetName() string {
	return transition.Name
}

func (transition *Transition) GetOwner() NamedElementInterface {
	return transition.Owner
}
