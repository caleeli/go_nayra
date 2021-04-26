package bpmn

// ActivityTrait from BPMN
type ActivityTrait struct {
	Ready       ActivityReadyState
	LoadingData LoadingDataTransition
	Active      State
	Completed   State
	Terminating TerminatingTransition
	Failing     FailingTransition
	Failed      State
	Terminated  State
	Cancelled   State
	Cancelling  CancellingTransition
	Completing  CompletingTransition
}

// Init Activity state definition
func (node *Activity) Init(definitions *Definitions) {
	node.Ready.Init(definitions, node, "Ready")
	node.LoadingData.Init(definitions, node, "LoadingData")
	node.Active.Init(definitions, node, "Active")
	node.Completed.Init(definitions, node, "Completed")
	node.Terminating.Init(definitions, node, "Terminating")
	node.Failing.Init(definitions, node, "Failing")
	node.Failed.Init(definitions, node, "Failed")
	node.Terminated.Init(definitions, node, "Terminated")
	node.Cancelled.Init(definitions, node, "Cancelled")
	node.Cancelling.Init(definitions, node, "Cancelling")
	node.Completing.Init(definitions, node, "Completing")
	node.Ready.Connect(&node.LoadingData)
	node.LoadingData.Connect(&node.Active)
	node.Active.Connect(&node.Failing)
	node.Active.Connect(&node.Terminating)
	node.Active.Connect(&node.Cancelling)
	node.Cancelling.Connect(&node.Cancelled)
	node.Terminating.Connect(&node.Terminated)
	node.Failing.Connect(&node.Failed)
	node.Active.Connect(&node.Completing)
	node.Completing.Connect(&node.Completed)
	for i := 0; i < len(node.Outgoing); i++ {
		Outgoing := &node.Outgoing[i]
		node.Completed.Connect(&definitions.GetSequenceFlow(*Outgoing).Transit)
	}
	for i := 0; i < len(node.Incoming); i++ {
		Incoming := &node.Incoming[i]
		definitions.GetSequenceFlow(*Incoming).Transit.Connect(&node.Ready)
	}

}
