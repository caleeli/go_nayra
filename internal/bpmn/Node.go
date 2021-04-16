package bpmn

// Node from Nayra
type Node struct {
	NodeInterface
	Owner NamedElementInterface
	Name  string
}

// Init transition
func (node *Node) Init(definitions *Definitions, owner NamedElementInterface, name string) {
	node.Owner = owner
	node.Name = name
}
