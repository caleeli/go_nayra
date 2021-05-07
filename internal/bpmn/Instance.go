package bpmn

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
)

const (
	DEBUG                    = true
	PROCESS_STATUS_ACTIVE    = "ACTIVE"
	PROCESS_STATUS_COMPLETED = "COMPLETED"
)

type CallbackSign = func(observer *Observer, observable *Observable, body interface{})
type CallbackObs struct {
	Index    int
	Callback CallbackSign
}

type Observable struct {
	TargetInstance *Instance
	TargetElement  NamedElementInterface
	TargetEvent    string
}
type Observer struct {
	SourceInstance *Instance
	SourceToken    *Token
	Callback       *CallbackObs
}
type Observation struct {
	Observable *Observable
	Observer   *Observer
}

// Instance from Nayra
type Instance struct {
	ID           uuid.UUID `json:"id"`
	Process      *Process
	RequestId    string
	Tokens       []*Token
	logs         []string
	Observations []*Observation
	Status       string
}

func NewInstance(definitions *Definitions, process *Process) *Instance {
	instance := Instance{}
	instance.Init(definitions, process)
	return &instance
}

// Init state
func (instance *Instance) Init(definitions *Definitions, process *Process) {
	instance.ID = uuid.New()
	instance.Process = process
	instance.Status = PROCESS_STATUS_ACTIVE
	instance.Tokens = []*Token{}
}

func (instance *Instance) CreateToken(state StateInterface) *Token {
	now := time.Now()
	token := Token{
		ID:         uuid.New(),
		Instance:   instance,
		Owner:      state,
		Active:     true,
		ThreadData: make(map[string]interface{}),
		Timestamp:  now.UnixNano() / 1000,
	}
	instance.Tokens = append(instance.Tokens, &token)
	return &token
}

func (instance *Instance) AttachObserver(definitions *Definitions, observable *Observable, observer *Observer) {
	instance.Observations = append(instance.Observations, &Observation{observable, observer})
}

func (instance *Instance) NotifyObservers(observable *Observable, body interface{}) {
	for _, obs := range instance.Observations {
		if obs.Observable.TargetElement.GetId() == observable.TargetElement.GetId() &&
			obs.Observable.TargetEvent == observable.TargetEvent {
			obs.Observer.Callback.Callback(obs.Observer, observable, body)
		}
	}
}

// GetToken by id
func (instance *Instance) GetToken(id uuid.UUID) *Token {
	for i := range instance.Tokens {
		if instance.Tokens[i].ID == id {
			return instance.Tokens[i]
		}
	}
	return nil
}

// RemoveToken from state
func (instance *Instance) RemoveToken(token *Token) bool {
	token.Active = false
	return true
}

// NextTick execute transits
func (instance *Instance) NextTick(definitions *Definitions) (changed bool) {
	MaxNestingLevel := 256
	changed = false
	for i := 0; i < MaxNestingLevel; i++ {
		exit := true
		for _, transition := range definitions.Transitions {
			if transition.Execute(instance) {
				if DEBUG {
					owner := transition.GetOwner()
					outgoing := transition.GetOutgoing()
					target := make(map[string]int, len(outgoing))
					for _, state := range outgoing {
						target[state.GetName()] = len(state.GetTokens())
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
				changed = true
			}
		}
		if exit {
			break
		}
	}
	return
}

// log an action during execution
func (instance *Instance) log(line string, args ...interface{}) {
	instance.logs = append(instance.logs, fmt.Sprintf(line, args...))
}

// Log an action during execution
func (instance *Instance) TraceLog() {
	// log transitions
	for _, line := range instance.logs {
		log.Println(line)
	}
	log.Println("------------------------------------------------------")
	// log tokens
	for _, token := range instance.Tokens {
		active := "[.]"
		if token.Active {
			active = "[*]"
		}
		log.Printf("%s %s %s (%s)\n", active, token.ID, token.Owner.GetOwner().GetName(), token.Owner.GetName())
	}
}

// Log an action during execution
func (instance *Instance) Trace() (trace []string) {
	trace = make([]string, len(instance.logs)+1+len(instance.Tokens))
	i := 0
	for range instance.logs {
		trace[i] = instance.logs[i]
		i++
	}
	trace[i] = "------------------------------------------------------"
	i++
	for _, token := range instance.Tokens {
		active := "[.]"
		if token.Active {
			active = "[*]"
		}
		trace[i] = fmt.Sprintf("%s %s %s (%s)", active, token.ID, token.Owner.GetOwner().GetName(), token.Owner.GetName())
		i++
	}
	return
}
