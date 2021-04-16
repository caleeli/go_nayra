package bpmn

// Token from bpmn
type Token struct {
	ID       string
	Instance *Instance
	Owner    *State
	Status   string
}
