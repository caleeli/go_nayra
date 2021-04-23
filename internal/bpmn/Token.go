package bpmn

import (
	"github.com/google/uuid"
)

// Token from bpmn
type Token struct {
	ID         uuid.UUID
	Instance   *Instance
	Owner      *State
	Transition string
	Active     bool
}
