package storage

import (
	"fmt"
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

type tRequest struct {
	definitions *bpmn.Definitions
	instances   map[uuid.UUID]*bpmn.Instance
	observers   []*tRequest
}

type Request interface {
	GetInstance(uuid uuid.UUID) *bpmn.Instance
	GetInstanceUuids() []uuid.UUID
	TraceLog()
}

func NewRequest(definitions *bpmn.Definitions, instances int) (Request, error) {
	request := tRequest{
		definitions: definitions,
		instances:   make(map[uuid.UUID]*bpmn.Instance),
	}
	for i := 0; i < instances; i++ {
		request.CreateInstance()
	}
	return &request, nil
}

func (request tRequest) GetInstance(uuid uuid.UUID) *bpmn.Instance {
	return request.instances[uuid]
}

func (request tRequest) CreateInstance() (*bpmn.Instance, uuid.UUID) {
	uuid := uuid.New()
	instance := bpmn.Instance{}
	request.instances[uuid] = &instance
	return request.instances[uuid], uuid
}

func (request tRequest) GetInstanceUuids() (uuids []uuid.UUID) {
	uuids = make([]uuid.UUID, len(request.instances))
	i := 0
	for k := range request.instances {
		uuids[i] = k
		i++
	}
	return
}

func (request tRequest) TraceLog() {
	for uuid, instance := range request.instances {
		fmt.Println("Request Instance:", uuid)
		fmt.Println("======================================================")
		instance.TraceLog()
	}
}
