package repository

import (
	"log"
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
	AppendInstance(instance *bpmn.Instance)
	GetId() uuid.UUID
	GetInstance(index int) *bpmn.Instance
	GetInstanceById(id uuid.UUID) *bpmn.Instance
	GetInstanceIds() []uuid.UUID
	GetToken(uuid uuid.UUID) (*bpmn.Token, error)
	GetDefinitions() *bpmn.Definitions
	NextTick() bool
	TraceLog()
	Trace() []string
}

func (request *TRequest) AppendInstance(instance *bpmn.Instance) {
	request.Instances = append(request.Instances, instance)
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

func (request *TRequest) CreateInstance(process *bpmn.Process) (*bpmn.Instance, uuid.UUID) {
	uuid := uuid.New()
	instance := bpmn.Instance{}
	instance.Init(request.Definitions, process)
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

func (request *TRequest) NextTick() (changed bool) {
	MaxNestingLevel := 10
	changed = false
	for i := 0; i < MaxNestingLevel; i++ {
		exit := true
		for j := range request.Instances {
			if request.Instances[j].NextTick(request.Definitions) {
				exit = false
				changed = true
			}
		}
		if exit {
			break
		}
	}
	return
}

func (request *TRequest) TraceLog() {
	log.Println("Request:", request.ID)
	log.Println("======================================================")
	for _, instance := range request.Instances {
		log.Println("------------------------------------------------------")
		log.Println("Instance:", instance.ID)
		log.Println("------------------------------------------------------")
		instance.TraceLog()
	}
}

func (request *TRequest) Trace() (trace []string) {
	trace = make([]string, 0)
	for _, instance := range request.Instances {
		trace = append(trace, "------------------------------------------------------")
		trace = append(trace, "Instance: "+instance.ID.String())
		trace = append(trace, "------------------------------------------------------")
		trace = append(trace, instance.Trace()...)
	}
	return
}
