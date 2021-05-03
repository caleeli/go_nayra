package bpmn

type ActivityActiveState struct {
	State
}

type ActivityActivatedEvent struct {
	TaskId string
	Token  *Token
}

// OnTokenArrived is triggered when a token arrives to the Ready state
func (state *ActivityActiveState) OnTokenArrived(token *Token, inputTokens []*Token) {
	state.State.OnTokenArrived(token, inputTokens)
	token.addToThreadData("taskId", token.ID.ID())
	NotifyEvent("ACTIVITY_ACTIVATED", ActivityActivatedEvent{
		TaskId: token.ID.String(),
		Token:  token,
	})
}
