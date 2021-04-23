package storage

import (
	"fmt"
	"nayra/internal/bpmn"
	"nayra/internal/errors"

	"github.com/google/uuid"
)

type tRequest struct {
	ID          uuid.UUID         `json:"id"`
	Definitions *bpmn.Definitions `json:"definitionsId"`
	Instances   []*bpmn.Instance  `json:"instances"`
	observers   []*tRequest
}

var config struct {
	storage StorageService
}

func SetupStorageService(storage StorageService) {
	config.storage = storage
}

type Request interface {
	GetId() uuid.UUID
	GetInstance(index int) *bpmn.Instance
	GetInstanceById(id uuid.UUID) *bpmn.Instance
	GetInstanceIds() []uuid.UUID
	GetToken(uuid uuid.UUID) (*bpmn.Token, error)
	GetDefinitions() *bpmn.Definitions
	TraceLog()
	MarshalJSON() ([]byte, error)
}

func NewRequest(definitions *bpmn.Definitions, instances int) (Request, error) {
	request := tRequest{
		ID:          uuid.New(),
		Definitions: definitions,
		Instances:   []*bpmn.Instance{},
	}
	for i := 0; i < instances; i++ {
		request.CreateInstance()
	}
	return &request, nil
}

func LoadRequest(requestId uuid.UUID) (Request, error) {
	return config.storage.LoadRequest(requestId)
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

func (request *tRequest) GetDefinitions() *bpmn.Definitions {
	return request.Definitions
}

func (request *tRequest) GetInstance(index int) *bpmn.Instance {
	return request.Instances[index]
}
func (request *tRequest) GetInstanceById(id uuid.UUID) *bpmn.Instance {
	for i := range request.Instances {
		if request.Instances[i].ID == id {
			return request.Instances[i]
		}
	}
	return nil
}

func (request *tRequest) GetToken(id uuid.UUID) (*bpmn.Token, error) {
	for i := range request.Instances {
		token := request.Instances[i].GetToken(id)
		if token != nil {
			return token, nil
		}
	}
	return nil, errors.WrapStorageError(nil, "Not found token %s", id)
}

func (request *tRequest) CreateInstance() (*bpmn.Instance, uuid.UUID) {
	uuid := uuid.New()
	instance := bpmn.Instance{}
	instance.Init(request.Definitions)
	request.Instances = append(request.Instances, &instance)
	return &instance, uuid
}

func (request *tRequest) GetInstanceIds() (uuids []uuid.UUID) {
	uuids = make([]uuid.UUID, len(request.Instances))
	i := 0
	for k := range request.Instances {
		uuids[i] = request.Instances[k].ID
		i++
	}
	return
}

func (request *tRequest) TraceLog() {
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
