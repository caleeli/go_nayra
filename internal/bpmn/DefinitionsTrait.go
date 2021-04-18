package bpmn

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
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

// MarshalJSON supports json.Marshaler interface
func (definitions *Definitions) MarshalJSON() ([]byte, error) {
	out := jwriter.Writer{}
	out.String(definitions.UUID.String())
	return out.Buffer.BuildBytes(), out.Error
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Definitions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	return r.Error()
}
