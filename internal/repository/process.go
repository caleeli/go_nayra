package repository

import (
	"encoding/xml"
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

// func LoadBpmn(id string) (definitions *bpmn.Definitions, err error) {
// 	uuid, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	definitions, err = storage.LoadProcess(uuid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	definitions.ParseTree()
// 	return
// }
//
func ParseBpmn(content []byte) (*bpmn.Definitions, error) {
	definitions := &bpmn.Definitions{}
	xml.Unmarshal(content, definitions)
	definitions.UUID = uuid.New()
	return definitions, nil
}
