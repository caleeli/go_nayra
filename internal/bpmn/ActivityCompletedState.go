package bpmn

import "fmt"

type ActivityCompletedState struct {
	State
}

// OnTokenLeaves is activated when a token leaves the Complete state
func (state *ActivityCompletedState) OnTokenLeaves(token *Token) {
	state.State.OnTokenLeaves(token)
	token.removeFromThreadData("taskId")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~ENTRO1!!")
	NotifyEvent("ACTIVITY_COMPLETED", struct {
		TaskId string
	}{
		TaskId: token.ID.String(),
	})
}
