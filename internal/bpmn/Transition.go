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
	Condition(instance *Instance, inputTokens []*Token) bool
	Connect(target StateInterface)
	CreateOutputTokens(instance *Instance, inputTokens []*Token) []*Token
	Execute(instance *Instance) bool
	GetIncoming() []StateInterface
	GetName() string
	GetOutgoing() []StateInterface
	GetOwner() NamedElementInterface
	Init(definitions *Definitions, owner NamedElementInterface, name string)
	RemoveInputTokens(instance *Instance, inputTokens []*Token)
	SelectInputTokens(instance *Instance) []*Token
}

// Init transition
func (transition *Transition) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	transition.Node.Init(definitions, owner, name)
	//definitions.Transitions = append(definitions.Transitions, transition)
}

// Connect transition to state
func (transition *Transition) Connect(target StateInterface) {
	transition.Outgoing = append(transition.Outgoing, target)
	target.AppendIncoming(transition)
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
	transition.CreateOutputTokens(instance, inputTokens)
	transition.RemoveInputTokens(instance, inputTokens)
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
	for i := range inputTokens {
		inputTokens[i].Owner.OnTokenLeaves(inputTokens[i])
		inputTokens[i].Owner.RemoveToken(inputTokens[i])
	}
}

// CreateOutputTokens it creates tokens in output places
func (transition *Transition) CreateOutputTokens(instance *Instance, inputTokens []*Token) []*Token {
	outputTokens := []*Token{}
	for _, Outgoing := range transition.Outgoing {
		token := CreateToken(instance, Outgoing)
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
func (transition *Transition) AppendIncoming(state NodeInterface) {
	transition.Incoming = append(transition.Incoming, state.(StateInterface))
}

// AppendIncoming transition to state
func (transition *Transition) AppendOutgoing(state NodeInterface) {
	transition.Outgoing = append(transition.Outgoing, state.(StateInterface))
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
