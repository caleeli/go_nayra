package bpmn

// StartEventTrait from BPMN
type StartEventTrait struct {
	Active   State
	Starting Transition
	Execute  State
}

// Init StartEvent state definition
func (node *StartEvent) Init(definitions *Definitions) {
	prepare(&node.Active, definitions, node, "Active")
	prepare(&node.Starting, definitions, node, "Starting")
	prepare(&node.Execute, definitions, node, "Execute")
	for i := 0; i < len(node.Outgoing); i++ {
		Outgoing := &node.Outgoing[i]
		connect(&node.Active, &definitions.GetSequenceFlow(*Outgoing).Transit)
	}
	connect(&node.Execute, &node.Starting)
	connect(&node.Starting, &node.Active)

}
