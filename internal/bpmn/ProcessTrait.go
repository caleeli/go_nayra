package bpmn

// ProcessTrait from BPMN
type ProcessTrait struct {
	Calling          Transition
	Execute          State
	ProcessCompleted ProcessCompletedTransition
}

// Init Process state definition
func (node *Process) Init(definitions *Definitions) {
	prepare(&node.Calling, definitions, node, "Calling")
	prepare(&node.Execute, definitions, node, "Execute")
	prepare(&node.ProcessCompleted, definitions, node, "ProcessCompleted")
	for i := 0; i < len(node.StartEvent); i++ {
		StartEvent := &node.StartEvent[i]
		connect(&node.Calling, &StartEvent.Active)
	}
	connect(&node.Execute, &node.Calling)

}
