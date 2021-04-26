package bpmn

type ActivityReadyState struct {
	State
}

// OnTokenArrived is triggered when a token arrives to the Ready state
func (state *ActivityReadyState) OnTokenArrived(token *Token, inputTokens []*Token) {
	state.State.OnTokenArrived(token, inputTokens)
	token.addToThreadData("taskId", token.ID)
}
