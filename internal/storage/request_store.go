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
		Active:     token.Active,
	}
}

func unmarshalRequest(request *sRequest) (Request, error) {
	ID, _ := uuid.Parse(request.ID)
	definitions, err := LoadBpmn(request.DefinitionsId)
	if err != nil {
		return nil, err
	}
	output := &tRequest{
		ID:          ID,
		Definitions: definitions,
		Instances:   make([]*bpmn.Instance, len(request.Instances)),
	}
	for i := range request.Instances {
		output.Instances[i], err = unmarshalInstance(definitions, &request.Instances[i])
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func unmarshalInstance(definitions *bpmn.Definitions, instance *sInstance) (*bpmn.Instance, error) {
	ID, err := uuid.Parse(instance.ID)
	if err != nil {
		return nil, err
	}
	output := &bpmn.Instance{
		ID:     ID,
		Tokens: make([]*bpmn.Token, len(instance.Tokens)),
	}
	for i := range instance.Tokens {
		output.Tokens[i], err = unmarshalToken(definitions, output, &instance.Tokens[i])
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func unmarshalToken(definitions *bpmn.Definitions, instance *bpmn.Instance, token *sToken) (*bpmn.Token, error) {
	ID, err := uuid.Parse(token.ID)
	if err != nil {
		return nil, err
	}
	owner := definitions.States[token.StateIndex]
	output := &bpmn.Token{
		ID:         ID,
		Instance:   instance,
		Owner:      owner,
		Transition: token.Transition,
		Active:     token.Active,
	}
	owner.Tokens = append(owner.Tokens, output)
	return output, nil
}
