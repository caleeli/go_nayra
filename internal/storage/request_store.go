package storage

import (
	"nayra/internal/bpmn"
)

type sRequest struct {
	ID            string      `json:"id" bson:"id"`
	DefinitionsId string      `json:"definitionsId" bson:"definitionsId"`
	Instances     []sInstance `json:"instances" bson:"instances"`
}

type sInstance struct {
	ID     string   `json:"id" bson:"id"`
	Tokens []sToken `json:"tokens"`
}

type sToken struct {
	ID         string
	Status     string
	Transition string
}

func marshalRequest(request Request) sRequest {
	req := request.(*tRequest)
	output := sRequest{
		ID:            req.ID.String(),
		DefinitionsId: req.Definitions.UUID.String(),
		Instances:     make([]sInstance, len(req.Instances)),
	}
	for i := range req.Instances {
		output.Instances[i] = marshalInstance(req.Instances[i])
	}
	return output
}

func marshalInstance(instance *bpmn.Instance) sInstance {
	output := sInstance{
		ID:     instance.ID.String(),
		Tokens: make([]sToken, len(instance.Tokens)),
	}
	for i := range instance.Tokens {
		output.Tokens[i] = marshalToken(instance.Tokens[i])
	}
	return output
}

func marshalToken(token *bpmn.Token) sToken {
	return sToken{
		ID:         token.ID.String(),
		Status:     token.Owner.Name,
		Transition: token.Transition,
	}
}
