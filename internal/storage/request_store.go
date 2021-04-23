package storage

import (
	"nayra/internal/bpmn"

	"github.com/google/uuid"
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
	StateIndex int
	Transition string
	Active     bool
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
		StateIndex: token.Owner.Index,
		Transition: token.Transition,
	}
}

func unmarshalRequest(request *sRequest) Request {
	ID, _ := uuid.Parse(request.ID)
	definitions, _ := LoadBpmn(request.DefinitionsId)
	output := &tRequest{
		ID:          ID,
		Definitions: definitions,
		Instances:   make([]*bpmn.Instance, len(request.Instances)),
	}
	for i := range request.Instances {
		output.Instances[i] = unmarshalInstance(definitions, &request.Instances[i])
	}
	return output
}

func unmarshalInstance(definitions *bpmn.Definitions, instance *sInstance) *bpmn.Instance {
	ID, _ := uuid.Parse(instance.ID)
	output := &bpmn.Instance{
		ID:     ID,
		Tokens: make([]*bpmn.Token, len(instance.Tokens)),
	}
	for i := range instance.Tokens {
		output.Tokens[i] = unmarshalToken(definitions, output, &instance.Tokens[i])
	}
	return output
}

func unmarshalToken(definitions *bpmn.Definitions, instance *bpmn.Instance, token *sToken) *bpmn.Token {
	ID, _ := uuid.Parse(token.ID)
	return &bpmn.Token{
		ID:         ID,
		Instance:   instance,
		Owner:      definitions.States[token.StateIndex],
		Transition: token.Transition,
		Active:     token.Active,
	}
}
