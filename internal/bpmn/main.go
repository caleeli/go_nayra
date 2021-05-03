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

func registerCallback(definitions *Definitions, callback CallbackSign) *CallbackObs {
	index := len(definitions.Callbacks)
	cb := &CallbackObs{Callback: callback, Index: index}
	definitions.Callbacks = append(definitions.Callbacks, cb)
	return cb
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

func CallProcess(definitions *Definitions, process *Process) *Instance {
	instance := NewInstance(definitions, process)
	CreateToken(instance, &process.Execute)
	NotifyEvent("PROCESS_INSTANCE_CREATED", InstanceCreatedEvent{
		process,
		instance,
	})
	return instance
}
