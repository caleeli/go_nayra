package storage

import (
	"nayra/internal/bpmn"

	"github.com/google/uuid"
)

type sRequest struct {
	id            uuid.UUID
	definitionsId uuid.UUID
	instances     []sInstance
}

type sInstance struct {
	id     uuid.UUID
	tokens []sToken
}

type sToken struct {
	id         uuid.UUID
	status     string
	transition string
}

func marshalRequest(req *tRequest) sRequest {
	output := sRequest{
		id:            req.ID,
		definitionsId: req.Definitions.UUID,
		instances:     make([]sInstance, len(req.Instances)),
	}
	for i := range req.Instances {
		marshalInstance(req.Instances[i])
	}
	return output
}

func marshalInstance(instance *bpmn.Instance) sInstance {
	output := sInstance{
		id:     instance.ID,
		tokens: make([]sToken, len(instance.Tokens)),
	}
	for i := range instance.Tokens {
		marshalToken(instance.Tokens[i])
	}
	return output
}

func marshalToken(token *bpmn.Token) sToken {
	return sToken{
		id:         token.ID,
		status:     token.Owner.Name,
		transition: token.Transition,
	}
}
