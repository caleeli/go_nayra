package bpmn

type ActivityActiveState struct {
	State
}

// OnTokenArrived is triggered when a token arrives to the Ready state
func (state *ActivityActiveState) OnTokenArrived(token *Token, inputTokens []*Token) {
	state.State.OnTokenArrived(token, inputTokens)
	token.addToThreadData("taskId", token.ID.ID())
}
