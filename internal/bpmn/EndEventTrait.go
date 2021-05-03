package bpmn

// EndEventTrait from BPMN
type EndEventTrait struct {
	Triggered  State
	TokenEnded Transition
}

// Init EndEvent state definition
func (node *EndEvent) Init(definitions *Definitions) {
	prepare(&node.Triggered, definitions, node, "Triggered")
	prepare(&node.TokenEnded, definitions, node, "TokenEnded")
	for i := 0; i < len(node.Incoming); i++ {
		Incoming := &node.Incoming[i]
		connect(&definitions.GetSequenceFlow(*Incoming).Transit, &node.Triggered)
	}
	connect(&node.Triggered, &node.TokenEnded)

}
