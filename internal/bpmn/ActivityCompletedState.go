package bpmn

type ActivityCompletedState struct {
	State
}

// OnTokenArrived is triggered when a token arrives to the Ready state
func (state *ActivityCompletedState) OnTokenArrived(token *Token, inputTokens []*Token) {
	state.State.OnTokenArrived(token, inputTokens)
	token.removeFromThreadData("taskId")
}
