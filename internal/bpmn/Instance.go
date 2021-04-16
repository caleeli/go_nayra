package bpmn

import (
	"fmt"
	"reflect"
)

const (
	DEBUG = true
)

// Instance from Nayra
type Instance struct {
	Tokens []*Token
	logs   []string
}

// Init state
func (instance *Instance) Init(definitions *Definitions) {}

// NextTick execute transits
func (instance *Instance) NextTick(definitions *Definitions) {
	MaxNestingLevel := 256
	for i := 0; i < MaxNestingLevel; i++ {
		exit := true
		for _, transition := range definitions.Transitions {
			if transition.Execute(instance) {
				if DEBUG {
					owner := transition.GetOwner()
					instance.log("(%s) %s: %s", reflect.TypeOf(owner), owner.GetName(), transition.GetName())
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
