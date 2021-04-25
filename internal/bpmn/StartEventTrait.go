package bpmn

// StartEventTrait from BPMN
type StartEventTrait struct {
	Active   State
	Starting Transition
	Execute  State
}

// Init StartEvent state definition
func (node *StartEvent) Init(definitions *Definitions) {
	node.Active.Init(definitions, node, "Active")
	node.Starting.Init(definitions, node, "Starting")
	node.Execute.Init(definitions, node, "Execute")
	for i := 0; i < len(node.Outgoing); i++ {
		Outgoing := &node.Outgoing[i]
		node.Active.Connect(&definitions.GetSequenceFlow(*Outgoing).Transit)
	}
	node.Execute.Connect(&node.Starting)
	node.Starting.Connect(&node.Active)

}
