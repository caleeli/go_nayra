package bpmn

type ActivityCompletedState struct {
	State
}

// OnTokenLeaves is activated when a token leaves the Complete state
func (state *ActivityCompletedState) OnTokenLeaves(token *Token) {
	state.State.OnTokenLeaves(token)
	token.removeFromThreadData("taskId")
}
