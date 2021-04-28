package bpmn

// SequenceFlowTrait from BPMN
type SequenceFlowTrait struct {
	Transit Transition
}

// Init SequenceFlow state definition
func (node *SequenceFlow) Init(definitions *Definitions) {
	prepare(&node.Transit, definitions, node, "Transit")

}
