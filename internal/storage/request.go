package storage

import (
	"fmt"
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

type tRequest struct {
	ID           uuid.UUID         `json:"id"`
	Definitions  *bpmn.Definitions `json:"definitionsId"`
	instancesMap map[uuid.UUID]*bpmn.Instance
	Instances    []*bpmn.Instance `json:"instances"`
	observers    []*tRequest
}

var config struct {
	storage StorageService
}

func SetupStorageService(storage StorageService) {
	config.storage = storage
}

type Request interface {
	GetId() uuid.UUID
	GetInstance(uuid uuid.UUID) *bpmn.Instance
	GetInstanceUuids() []uuid.UUID
	TraceLog()
	MarshalJSON() ([]byte, error)
}

func NewRequest(definitions *bpmn.Definitions, instances int) (Request, error) {
	request := tRequest{
		ID:           uuid.New(),
		Definitions:  definitions,
		instancesMap: make(map[uuid.UUID]*bpmn.Instance),
		Instances:    []*bpmn.Instance{},
	}
	for i := 0; i < instances; i++ {
		request.CreateInstance()
	}
	return &request, nil
}

func InsertRequest(request Request) error {
	return config.storage.InsertRequest(request)
}

func UpdateRequest(request Request) error {
	return config.storage.UpdateRequest(request)
}

func (request *tRequest) GetId() uuid.UUID {
	return request.ID
}

func (request *tRequest) GetInstance(uuid uuid.UUID) *bpmn.Instance {
	return request.instancesMap[uuid]
}

func (request *tRequest) CreateInstance() (*bpmn.Instance, uuid.UUID) {
	uuid := uuid.New()
	instance := bpmn.Instance{}
	instance.Init(request.Definitions)
	request.instancesMap[uuid] = &instance
	request.Instances = append(request.Instances, &instance)
	return request.instancesMap[uuid], uuid
}

func (request *tRequest) GetInstanceUuids() (uuids []uuid.UUID) {
	uuids = make([]uuid.UUID, len(request.instancesMap))
	i := 0
	for k := range request.instancesMap {
		uuids[i] = k
		i++
	}
	return
}

func (request *tRequest) TraceLog() {
	for uuid, instance := range request.instancesMap {
		fmt.Println("Request Instance:", uuid)
		fmt.Println("======================================================")
		instance.TraceLog()
	}
}
