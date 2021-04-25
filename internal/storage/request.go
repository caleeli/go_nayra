package storage

import (
	"nayra/internal/bpmn"
	"nayra/internal/repository"

	"github.com/google/uuid"
)

type SRequest struct {
	Id            string
	DefinitionsId string
	Instances     []sInstance
}

type sInstance struct {
	Id     string
	Tokens []sToken
}

type sToken struct {
	ID         string
	Status     string
	StateIndex int
	Transition string
	Active     bool
}

func MarshalRequest(request repository.Request) SRequest {
	req := request.(*repository.TRequest)
	output := SRequest{
		Id:            req.ID.String(),
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
		Id:     instance.ID.String(),
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

func UnmarshalRequest(request *SRequest) (repository.Request, error) {
	ID, _ := uuid.Parse(request.Id)
	definitions, err := LoadDefinitions(request.DefinitionsId)
	if err != nil {
		return nil, err
	}
	output := &repository.TRequest{
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
	ID, err := uuid.Parse(instance.Id)
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
	if output.Active {
		owner.Tokens = append(owner.Tokens, output)
	}
	return output, nil
}

func NewRequest(definitions *bpmn.Definitions, instances int) (repository.Request, error) {
	request := repository.TRequest{
		ID:          uuid.New(),
		Definitions: definitions,
		Instances:   []*bpmn.Instance{},
	}
	for i := 0; i < instances; i++ {
		request.CreateInstance()
	}
	return &request, nil
}

func LoadRequest(requestId uuid.UUID) (repository.Request, error) {
	return config.storage.LoadRequest(requestId)
}

func InsertRequest(request repository.Request) error {
	return config.storage.InsertRequest(request)
}

func UpdateRequest(request repository.Request) error {
	return config.storage.UpdateRequest(request)
}
