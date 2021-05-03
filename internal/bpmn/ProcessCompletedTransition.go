package bpmn

type ProcessCompletedTransition struct {
	Transition
}

const (
	PROCESS_INSTANCE_COMPLETED = "PROCESS_INSTANCE_COMPLETED"
)

// Execute transition
func (transition ProcessCompletedTransition) Execute(instance *Instance) bool {
	if transition.Owner != instance.Process || instance.Status == PROCESS_STATUS_COMPLETED || instance.HasActiveTokens() {
		return false
	}
	instance.Status = PROCESS_STATUS_COMPLETED
	inputTokens := transition.SelectInputTokens(instance)
	if inputTokens == nil {
		return false
	}
	transition.CreateOutputTokens(instance, inputTokens)
	transition.RemoveInputTokens(instance, inputTokens)
	observable := &Observable{
		TargetInstance: instance,
		TargetElement:  transition.Owner,
		TargetEvent:    PROCESS_INSTANCE_COMPLETED,
	}
	instance.NotifyObservers(observable, nil)
	return true
}

func (instance *Instance) HasActiveTokens() bool {
	for _, token := range instance.Tokens {
		if token.Active {
			return true
		}
	}
	return false
}
