package repository

import (
	"encoding/xml"
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

func ParseBpmn(content []byte) (*bpmn.Definitions, error) {
	definitions := &bpmn.Definitions{}
	xml.Unmarshal(content, definitions)
	definitions.UUID = uuid.New()
	return definitions, nil
}
