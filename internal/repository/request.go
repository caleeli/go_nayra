package repository

import (
	"fmt"
	"nayra/internal/bpmn"
	"nayra/internal/errors"

	"github.com/google/uuid"
)

type TRequest struct {
	ID          uuid.UUID         `json:"id"`
	Definitions *bpmn.Definitions `json:"definitionsId"`
	Instances   []*bpmn.Instance  `json:"instances"`
	// @todo observers   []*tRequest
}

type Request interface {
	GetId() uuid.UUID
	GetInstance(index int) *bpmn.Instance
	GetInstanceById(id uuid.UUID) *bpmn.Instance
	GetInstanceIds() []uuid.UUID
	GetToken(uuid uuid.UUID) (*bpmn.Token, error)
	GetDefinitions() *bpmn.Definitions
	TraceLog()
}

func (request *TRequest) GetId() uuid.UUID {
	return request.ID
}

func (request *TRequest) GetDefinitions() *bpmn.Definitions {
	return request.Definitions
}

func (request *TRequest) GetInstance(index int) *bpmn.Instance {
	return request.Instances[index]
}
func (request *TRequest) GetInstanceById(id uuid.UUID) *bpmn.Instance {
	for i := range request.Instances {
		if request.Instances[i].ID == id {
			return request.Instances[i]
		}
	}
	return nil
}

func (request *TRequest) GetToken(id uuid.UUID) (*bpmn.Token, error) {
	for i := range request.Instances {
		token := request.Instances[i].GetToken(id)
		if token != nil {
			return token, nil
		}
	}
	return nil, errors.WrapStorageError(nil, "Not found token %s", id)
}

func (request *TRequest) CreateInstance() (*bpmn.Instance, uuid.UUID) {
	uuid := uuid.New()
	instance := bpmn.Instance{}
	instance.Init(request.Definitions)
	request.Instances = append(request.Instances, &instance)
	return &instance, uuid
}

func (request *TRequest) GetInstanceIds() (uuids []uuid.UUID) {
	uuids = make([]uuid.UUID, len(request.Instances))
	i := 0
	for k := range request.Instances {
		uuids[i] = request.Instances[k].ID
		i++
	}
	return
}

func (request *TRequest) TraceLog() {
	fmt.Println("Request:", request.ID)
	fmt.Println("======================================================")
	for _, instance := range request.Instances {
		fmt.Println("Instance:", instance.ID)
		fmt.Println("------------------------------------------------------")
		instance.TraceLog()
		fmt.Println("------------------------------------------------------")
		instance.TokensLog()
	}
}