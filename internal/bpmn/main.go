package bpmn

func prepare(node NodeInterface, definitions *Definitions, owner NamedElementInterface, name string) {
	node.Init(definitions, owner, name)
	state, isState := node.(StateInterface)
	if isState {
		state.SetIndex(len(definitions.States))
		definitions.States = append(definitions.States, state)
	} else {
		definitions.Transitions = append(definitions.Transitions, node.(TransitionInterface))
	}
}

func connect(source NodeInterface, target NodeInterface) {
	source.AppendOutgoing(target)
	target.AppendIncoming(source)
}

// CreateToken into the State
func CreateToken(instance *Instance, state StateInterface) *Token {
	token := instance.CreateToken(state)
	state.AppendToken(token)
	return token
}
