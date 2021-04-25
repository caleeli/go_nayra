package storage

import (
	"nayra/internal/bpmn"
)

type SDefinitions struct {
	Id          string
	Definitions *bpmn.Definitions
}

func MarshalDefinitions(definitions *bpmn.Definitions) *SDefinitions {
	ID := definitions.UUID.String()

	return &SDefinitions{
		Id:          ID,
		Definitions: definitions,
	}
}

func UnmarshalDefinitions(def *SDefinitions) (*bpmn.Definitions, error) {
	return def.Definitions, nil
}

func LoadDefinitions(id string) (*bpmn.Definitions, error) {
	return config.storage.LoadDefinitions(id)
}

func InsertDefinitions(definitions *bpmn.Definitions) (err error) {
	return config.storage.InsertDefinitions(definitions)
}
