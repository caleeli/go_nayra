package bpmn

// EndEventTrait from BPMN
type EndEventTrait struct {
	Triggered State
	Ending    Transition
}

// Init EndEvent state definition
func (node *EndEvent) Init(definitions *Definitions) {
	node.Triggered.Init(definitions, node, "Triggered")
	node.Ending.Init(definitions, node, "Ending")
	for i := 0; i < len(node.Incoming); i++ {
		Incoming := &node.Incoming[i]
		definitions.GetSequenceFlow(*Incoming).Transit.Connect(&node.Triggered)
	}
	node.Triggered.Connect(&node.Ending)
	
}
