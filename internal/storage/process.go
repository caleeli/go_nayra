package storage

import (
	"encoding/xml"
	"io/ioutil"
	"nayra/internal/bpmn"
	"nayra/internal/errors"
	"os"

	"github.com/google/uuid"
)

func LoadBpmn(id string) (definitions *bpmn.Definitions, err error) {
	// @todo load from DB
	path := "data/" + id + ".bpmn"
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, errors.WrapStorageError(err, "Unable to open file %s", path)
	}
	defer xmlFile.Close()

	content, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, errors.WrapStorageError(err, "Unable to read file")
	}
	definitions = &bpmn.Definitions{}
	xml.Unmarshal(content, definitions)
	definitions.UUID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	definitions.ParseTree()

	return
}
