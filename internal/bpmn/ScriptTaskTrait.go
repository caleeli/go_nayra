package bpmn

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

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
		// Call Lambda Async
		data := struct {
			RequestId string
			TokenId   string
			Script    string
		}{
			RequestId: payload.Token.Instance.RequestId,
			TokenId:   payload.Token.ID.String(),
			Script:    node.Script.Body,
		}
		dataEvent, _ := json.Marshal(data)
		svc := lambda.New(session.New())
		input := &lambda.InvokeAsyncInput{
			FunctionName: aws.String("arn:aws:lambda:us-east-2:294911624942:function:TestScript"),
			InvokeArgs:   aws.ReadSeekCloser(bytes.NewReader(dataEvent)),
		}
		svc.InvokeAsync(input)
		//NotifyEvent("RUN_SCRIPT", struct {
		//	RequestId string
		//	TokenId   string
		//	Script    string
		//}{
		//	RequestId: payload.Token.Instance.RequestId,
		//	TokenId:   payload.Token.ID.String(),
		//	Script:    node.Script.Body,
		//})
	})
}
