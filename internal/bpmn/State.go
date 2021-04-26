package bpmn

// State from Nayra
type State struct {
	Node
	Incoming []TransitionInterface
	Outgoing []TransitionInterface
	Tokens   []*Token
	Index    int
}

type StateInterface interface {
	AddIncoming(transition *Transition)
	CreateToken(instance *Instance) *Token
	GetName() string
	GetTokens() []*Token
	Select(instance *Instance, selectAll bool) []*Token
	OnTokenArrived(token *Token, inputTokens []*Token)
}

// Init transition
func (state *State) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	state.Node.Init(definitions, owner, name)
	state.Index = len(definitions.States)
	definitions.States = append(definitions.States, state)
}

// Connect transition to state
func (state *State) Connect(target TransitionInterface) {
	state.Outgoing = append(state.Outgoing, target)
	target.AppendIncoming(state)
}

// Select tokens from state
func (state *State) Select(instance *Instance, selectAll bool) []*Token {
	selected := []*Token{}
	for _, token := range state.Tokens {
		if token.Instance == instance {
			selected = append(selected, token)
			if !selectAll {
				return selected
			}
		}
	}
	return selected
}

// CreateToken into the State
func (state *State) CreateToken(instance *Instance) *Token {
	token := instance.CreateToken(state)
	state.Tokens = append(state.Tokens, token)
	return token
}

// RemoveToken from state
func (state *State) RemoveToken(token *Token) bool {
	for i := 0; i < len(state.Tokens); i++ {
		if state.Tokens[i] == token {
			token.Instance.RemoveToken(token)
			state.Tokens = append(state.Tokens[:i], state.Tokens[i+1:]...)
			return true
		}
	}
	return false
}

// AddIncoming transition to state
func (state *State) AddIncoming(transition *Transition) {
	state.Incoming = append(state.Incoming, transition)
}

// GateName of the state
func (state *State) GetName() string {
	return state.Name
}

// GetTokens in the state
func (state *State) GetTokens() []*Token {
	return state.Tokens
}

// OnTokenArrived is triggered when a token arrives to the state
func (state *State) OnTokenArrived(token *Token, inputTokens []*Token) {
	for i := range inputTokens {
		token.copyThreadDataFrom(inputTokens[i])
	}
}
