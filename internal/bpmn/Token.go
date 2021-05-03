package bpmn

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Token from bpmn
type Token struct {
	ID         uuid.UUID
	Instance   *Instance
	Owner      StateInterface
	ThreadData map[string]interface{}
	Transition string
	Active     bool
}

// AddToThreadData adds data to the tokens thread
func (token *Token) addToThreadData(name string, value interface{}) {
	token.ThreadData[name] = value
}

// RemoveFromThread removes data from the tokens thread
func (token *Token) removeFromThreadData(name string) {
	delete(token.ThreadData, name)
}

// getFromThread gets data from the tokens thread
func (token *Token) getFromThreadData(name string) interface{} {
	return token.ThreadData[name]
}

func (token *Token) copyThreadDataFrom(source *Token) (err error) {
	jsonStr, err := json.Marshal(source.ThreadData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonStr, &token.ThreadData)
	if err != nil {
		return err
	}
	return nil
}
