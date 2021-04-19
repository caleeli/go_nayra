package bpmn

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

const (
	DEBUG = true
)

// Instance from Nayra
type Instance struct {
	ID     uuid.UUID `json:"id"`
	Tokens []*Token
	logs   []string
}

// Init state
func (instance *Instance) Init(definitions *Definitions) {
	instance.ID = uuid.New()
	instance.Tokens = []*Token{}
}

func (instance *Instance) CreateToken(state *State) *Token {
	token := Token{ID: uuid.New(), Instance: instance, Owner: state}
	instance.Tokens = append(instance.Tokens, &token)
	return &token
}

// NextTick execute transits
func (instance *Instance) NextTick(definitions *Definitions) {
	MaxNestingLevel := 256
	for i := 0; i < MaxNestingLevel; i++ {
		exit := true
		for _, transition := range definitions.Transitions {
			if transition.Execute(instance) {
				if DEBUG {
					owner := transition.GetOwner()
					outgoing := transition.GetOutgoing()
					target := make(map[string]int, len(outgoing))
					for _, state := range outgoing {
						target[state.Name] = len(state.Tokens)
					}
					instance.log(
						"(%s) %s: %s [%v]",
						reflect.TypeOf(owner),
						owner.GetName(),
						transition.GetName(),
						target,
					)
				}
				exit = false
			}
		}
		if exit {
			break
		}
	}
}

// log an action during execution
func (instance *Instance) log(line string, args ...interface{}) {
	instance.logs = append(instance.logs, fmt.Sprintf(line, args...))
}

// Log an action during execution
func (instance *Instance) TraceLog() {
	for _, line := range instance.logs {
		fmt.Println(line)
	}
}
