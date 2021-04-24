package bpmn

// ProcessTrait from BPMN
type ProcessTrait struct {
	Calling Transition
	Execute State
}

// Init Process state definition
func (node *Process) Init(definitions *Definitions) {
	node.Calling.Init(definitions, node, "Calling")
	node.Execute.Init(definitions, node, "Execute")
	for i := 0; i < len(node.StartEvent); i++ {
		StartEvent := &node.StartEvent[i]
		node.Calling.Connect(&StartEvent.Active)
	}
	node.Execute.Connect(&node.Calling)

}
