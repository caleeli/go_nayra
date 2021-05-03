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
	Id           string
	ProcessId    string
	Tokens       []sToken
	Observations []sObservation
	Status       string
}

type sToken struct {
	ID         string
	Status     string
	StateIndex int
	ElementId  string
	ThreadData map[string]interface{}
	Transition string
	Active     bool
}

type sObservation struct {
	TargetInstanceId string
	TargetElementId  string
	TargetEvent      string
	CallbackIndex    int
	SourceInstanceId string
	SourceTokenId    string
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
		Id:           instance.ID.String(),
		ProcessId:    instance.Process.ID,
		Tokens:       make([]sToken, len(instance.Tokens)),
		Observations: make([]sObservation, len(instance.Observations)),
		Status:       instance.Status,
	}
	for i := range instance.Tokens {
		output.Tokens[i] = marshalToken(instance.Tokens[i])
	}
	for i := range instance.Observations {
		output.Observations[i] = marshalObservation(instance.Observations[i])
	}
	return output
}

func marshalToken(token *bpmn.Token) sToken {
	return sToken{
		ID:         token.ID.String(),
		Status:     token.Owner.GetName(),
		StateIndex: token.Owner.GetIndex(),
		ElementId:  token.Owner.GetOwner().GetId(),
		ThreadData: token.ThreadData,
		Transition: token.Transition,
		Active:     token.Active,
	}
}

func marshalObservation(observation *bpmn.Observation) sObservation {
	return sObservation{
		TargetInstanceId: observation.Observable.TargetInstance.ID.String(),
		TargetElementId:  observation.Observable.TargetElement.GetId(),
		TargetEvent:      observation.Observable.TargetEvent,
		CallbackIndex:    observation.Observer.Callback.Index,
		SourceInstanceId: observation.Observer.SourceInstance.ID.String(),
		SourceTokenId:    observation.Observer.SourceToken.ID.String(),
	}
}

func UnmarshalRequest(request *SRequest) (repository.Request, error) {
	ID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}
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
		instance := &request.Instances[i]
		ID, err := uuid.Parse(instance.Id)
		if err != nil {
			return nil, err
		}
		output.Instances[i] = &bpmn.Instance{
			ID:           ID,
			Process:      definitions.GetProcess(instance.ProcessId),
			Tokens:       make([]*bpmn.Token, len(instance.Tokens)),
			Observations: make([]*bpmn.Observation, len(instance.Observations)),
			Status:       instance.Status,
		}
	}
	for i := range request.Instances {
		output.Instances[i], err = unmarshalInstance(definitions, output, &request.Instances[i], output.Instances[i])
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func unmarshalInstance(definitions *bpmn.Definitions, request repository.Request, instance *sInstance, output *bpmn.Instance) (*bpmn.Instance, error) {
	var err error
	for i := range instance.Tokens {
		output.Tokens[i], err = unmarshalToken(definitions, output, &instance.Tokens[i])
		if err != nil {
			return nil, err
		}
	}
	for i := range instance.Observations {
		output.Observations[i], err = unmarshalObservation(definitions, request, &instance.Observations[i])
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
		ThreadData: token.ThreadData,
		Transition: token.Transition,
		Active:     token.Active,
	}
	if output.Active {
		owner.AppendToken(output)
	}
	return output, nil
}

func unmarshalObservation(definitions *bpmn.Definitions, request repository.Request, observation *sObservation) (*bpmn.Observation, error) {
	targetId, err := uuid.Parse(observation.TargetInstanceId)
	if err != nil {
		return nil, err
	}
	sourceId, err := uuid.Parse(observation.SourceInstanceId)
	if err != nil {
		return nil, err
	}
	sourceTokenId, err := uuid.Parse(observation.SourceTokenId)
	if err != nil {
		return nil, err
	}
	callback := definitions.Callbacks[observation.CallbackIndex]
	// @todo change to getbase
	targetElement := definitions.GetProcess(observation.TargetElementId)
	output := &bpmn.Observation{
		Observable: &bpmn.Observable{
			TargetInstance: request.GetInstanceById(targetId),
			TargetElement:  targetElement,
			TargetEvent:    observation.TargetEvent,
		},
		Observer: &bpmn.Observer{
			SourceInstance: request.GetInstanceById(sourceId),
			SourceToken:    request.GetInstanceById(sourceId).GetToken(sourceTokenId),
			Callback:       callback,
		},
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
		//request.CreateInstance()
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
