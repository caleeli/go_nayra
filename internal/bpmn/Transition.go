package bpmn

// Transition from Nayra
type Transition struct {
	Node
	Incoming        []StateInterface
	Outgoing        []StateInterface
	RemoveAllTokens bool
}

type TransitionInterface interface {
	NodeInterface
	Init(definitions *Definitions, owner NamedElementInterface, name string)
	Connect(target StateInterface)
	Execute(instance *Instance) bool
	SelectInputTokens(instance *Instance) []*Token
	RemoveInputTokens(instance *Instance, inputTokens []*Token)
	CreateOutputTokens(instance *Instance, inputTokens []*Token) []*Token
	Condition(instance *Instance, inputTokens []*Token) bool
	AppendIncoming(state StateInterface)
	GetName() string
	GetOwner() NamedElementInterface
	GetIncoming() []StateInterface
	GetOutgoing() []StateInterface
}

// Init transition
func (transition *Transition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	definitions.Transitions = append(definitions.Transitions, transition)
}

// Connect transition to state
func (transition *Transition) Connect(target StateInterface) {
	transition.Outgoing = append(transition.Outgoing, target)
	target.AddIncoming(transition)
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
		token := Outgoing.CreateToken(instance)
		Outgoing.OnTokenArrived(token, inputTokens)
		outputTokens = append(outputTokens, token)
	}
	return outputTokens
}

// Condition to activate the transition
func (transition *Transition) Condition(instance *Instance, inputTokens []*Token) bool {
	return true
}

// Condition to activate the transition
func (transition *Transition) AppendIncoming(state StateInterface) {
	transition.Incoming = append(transition.Incoming, state)
}

func (transition *Transition) GetName() string {
	return transition.Name
}

func (transition *Transition) GetOwner() NamedElementInterface {
	return transition.Owner
}

func (transition *Transition) GetIncoming() []StateInterface {
	return transition.Incoming
}

func (transition *Transition) GetOutgoing() []StateInterface {
	return transition.Outgoing
}
