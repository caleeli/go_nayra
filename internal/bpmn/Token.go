package bpmn

import (
	"github.com/google/uuid"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Token from bpmn
type Token struct {
	ID         uuid.UUID
	Instance   *Instance
	Owner      *State
	Transition string
	Active     bool
}

// MarshalJSON supports json.Marshaler interface
func (token *Token) MarshalJSON() ([]byte, error) {
	out := jwriter.Writer{}
	out.RawByte('{')
	{
		const prefix string = "\"id\":"
		out.RawString(prefix)
		out.String(token.ID.String())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(token.Owner.Name)
	}
	{
		const prefix string = ",\"transition\":"
		out.RawString(prefix)
		out.String(token.Transition)
	}
	out.RawByte('}')
	return out.Buffer.BuildBytes(), out.Error
}

// UnmarshalJSON supports json.Unmarshaler interface
func (token *Token) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	return r.Error()
}
