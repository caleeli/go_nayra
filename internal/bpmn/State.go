package bpmn

//easyjson:skip
// State from Nayra
type State struct {
	Node
	Incoming []TransitionInterface
	Outgoing []TransitionInterface
	Tokens   []*Token
}

type StateInterface interface {
}

// Init transition
func (state *State) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	state.Node.Init(definitions, owner, name)
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
			state.Tokens = append(state.Tokens[:i], state.Tokens[i+1:]...)
			return true
		}
	}
	return false
}
