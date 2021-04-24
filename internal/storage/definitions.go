package storage

import (
	"nayra/internal/bpmn"
)

type sDefinitions struct {
	ID          string            `json:"id"`
	Definitions *bpmn.Definitions `json:"definitions"`
}

func (db *mongoDB) MarshalProcess(definitions *bpmn.Definitions) *sDefinitions {
	ID := definitions.UUID.String()

	return &sDefinitions{
		ID:          ID,
		Definitions: definitions,
	}
}

func (db *mongoDB) UnmarshalProcess(def *sDefinitions) (*bpmn.Definitions, error) {
	return def.Definitions, nil
}

func LoadDefinitions(id string) (*bpmn.Definitions, error) {
	return config.storage.LoadDefinitions(id)
}

func InsertDefinitions(definitions *bpmn.Definitions) (err error) {
	return config.storage.InsertDefinitions(definitions)
}
