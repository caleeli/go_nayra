package bpmn

import (
	"fmt"

	"github.com/google/uuid"
)

// DefinitionsTrait from BPMN
type DefinitionsTrait struct {
	UUID        uuid.UUID
	Transitions []TransitionInterface
	States      []*State
}

// Print definitions state
func (definitions *Definitions) Print() {
	for _, transition := range definitions.Transitions {
		fmt.Println("Available transition: ", transition.GetName())
	}
	for _, state := range definitions.States {
		fmt.Println("State ", state.Name, ":", len(state.Tokens))
	}
}
