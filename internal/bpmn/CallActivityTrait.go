package bpmn

// CallActivityTrait from BPMN
type CallActivityTrait struct {
}

type InstanceCreatedEvent struct {
	Process  *Process
	Instance *Instance
}

type InstanceCompletedEvent struct {
	Process  *Process
	Instance *Instance
}

// Init Activity state definition
func (node *CallActivity) Init(definitions *Definitions) {
	InstanceCompleted := registerCallback(definitions, func(observer *Observer, observable *Observable, body interface{}) {
		node.Complete(observer.SourceToken)
	})
	calledElement := definitions.GetProcess(node.CalledElement)
	if calledElement == nil {
		return
	}
	SubscribeEvent("ACTIVITY_ACTIVATED", func(event string, body interface{}) {
		payload := body.(ActivityActivatedEvent)
		isCallActivity := payload.Token.Owner.GetOwner() == &node.Activity
		if !isCallActivity {
			return
		}
		instance := CallProcess(definitions, calledElement)
		observable := &Observable{
			TargetInstance: instance,
			TargetElement:  calledElement,
			TargetEvent:    PROCESS_INSTANCE_COMPLETED,
		}
		observer := &Observer{
			SourceInstance: payload.Token.Instance,
			SourceToken:    payload.Token,
			Callback:       InstanceCompleted,
		}
		instance.AttachObserver(definitions, observable, observer)
	})
}
