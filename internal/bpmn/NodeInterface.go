package bpmn

// NodeInterface from Nayra
type NodeInterface interface {
	//getSource()
	//getTarget()
	Init(definitions *Definitions, owner NamedElementInterface, name string)
	AppendIncoming(node NodeInterface)
	AppendOutgoing(node NodeInterface)
}
