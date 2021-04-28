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
	NodeInterface
	AppendToken(token *Token)
	SetIndex(index int)
	GetIndex() int
	GetName() string
	GetOwner() NamedElementInterface
	GetTokens() []*Token
	OnTokenArrived(token *Token, inputTokens []*Token)
	OnTokenLeaves(token *Token)
	RemoveToken(token *Token) bool
	Select(instance *Instance, selectAll bool) []*Token
}

// Init transition
func (state *State) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	state.Node.Init(definitions, owner, name)
	//state.Index = len(definitions.States)
	//definitions.States = append(definitions.States, state)
}

// Select tokens from state
func (state *State) Select(instance *Instance, selectAll bool) []*Token {
	selected := []*Token{}
	for i := range state.Tokens {
		if state.Tokens[i].Instance == instance {
			selected = append(selected, state.Tokens[i])
			if !selectAll {
				return selected
			}
		}
	}
	return selected
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

// AppendIncoming transition to state
func (state *State) AppendIncoming(transition NodeInterface) {
	state.Incoming = append(state.Incoming, transition.(TransitionInterface))
}

// AppendIncoming transition to state
func (state *State) AppendOutgoing(transition NodeInterface) {
	state.Outgoing = append(state.Outgoing, transition.(TransitionInterface))
}

// GetName of the state
func (state *State) GetName() string {
	return state.Name
}

// GetIndex of the state
func (state *State) GetIndex() int {
	return state.Index
}

// GetIndex of the state
func (state *State) SetIndex(index int) {
	state.Index = index
}

// GetTokens in the state
func (state *State) GetTokens() []*Token {
	return state.Tokens
}

// GetTokens in the state
func (state *State) AppendToken(token *Token) {
	state.Tokens = append(state.Tokens, token)
}

// GetOwner of the state
func (state *State) GetOwner() NamedElementInterface {
	return state.Owner
}

// OnTokenArrived is triggered when a token arrives to the state
func (state *State) OnTokenArrived(token *Token, inputTokens []*Token) {
	for i := range inputTokens {
		token.copyThreadDataFrom(inputTokens[i])
	}
}

// OnTokenLeaves is activated when a token leaves the state
func (state *State) OnTokenLeaves(token *Token) {}
