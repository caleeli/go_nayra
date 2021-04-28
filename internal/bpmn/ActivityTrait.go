package bpmn

// ActivityTrait from BPMN
type ActivityTrait struct {
	Ready       State
	LoadingData LoadingDataTransition
	Active      ActivityActiveState
	Completed   ActivityCompletedState
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
	prepare(&node.Ready, definitions, node, "Ready")
	prepare(&node.LoadingData, definitions, node, "LoadingData")
	prepare(&node.Active, definitions, node, "Active")
	prepare(&node.Completed, definitions, node, "Completed")
	prepare(&node.Terminating, definitions, node, "Terminating")
	prepare(&node.Failing, definitions, node, "Failing")
	prepare(&node.Failed, definitions, node, "Failed")
	prepare(&node.Terminated, definitions, node, "Terminated")
	prepare(&node.Cancelled, definitions, node, "Cancelled")
	prepare(&node.Cancelling, definitions, node, "Cancelling")
	prepare(&node.Completing, definitions, node, "Completing")
	connect(&node.Ready, &node.LoadingData)
	connect(&node.LoadingData, &node.Active)
	connect(&node.Active, &node.Failing)
	connect(&node.Active, &node.Terminating)
	connect(&node.Active, &node.Cancelling)
	connect(&node.Cancelling, &node.Cancelled)
	connect(&node.Terminating, &node.Terminated)
	connect(&node.Failing, &node.Failed)
	connect(&node.Active, &node.Completing)
	connect(&node.Completing, &node.Completed)
	for i := 0; i < len(node.Outgoing); i++ {
		Outgoing := &node.Outgoing[i]
		connect(&node.Completed, &definitions.GetSequenceFlow(*Outgoing).Transit)
	}
	for i := 0; i < len(node.Incoming); i++ {
		Incoming := &node.Incoming[i]
		connect(&definitions.GetSequenceFlow(*Incoming).Transit, &node.Ready)
	}

}
