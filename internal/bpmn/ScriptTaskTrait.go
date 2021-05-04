package bpmn

// ScriptTaskTrait from BPMN
type ScriptTaskTrait struct {
}

// Init ScriptTaskTrait state definition
func (node *ScriptTask) Init(definitions *Definitions) {
	SubscribeEvent("ACTIVITY_ACTIVATED", func(event string, body interface{}) {
		payload := body.(ActivityActivatedEvent)
		itIsMe := payload.Token.Owner.GetOwner() == &node.Activity
		if !itIsMe {
			return
		}
		NotifyEvent("RUN_SCRIPT", struct {
			RequestId string
			TokenId   string
			Script    string
		}{
			RequestId: payload.Token.Instance.RequestId,
			TokenId:   payload.Token.ID.String(),
			Script:    node.Script.Body,
		})
	})
}
