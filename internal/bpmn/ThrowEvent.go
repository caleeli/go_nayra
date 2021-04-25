package bpmn

import (
	"encoding/xml"
)

// ThrowEvent from BPMN
type ThrowEvent struct {
	Event
	XMLName                    xml.Name
	DataInput                  []DataInput                  `xml:"dataInput"`
	DataInputAssociation       []DataInputAssociation       `xml:"dataInputAssociation"`
	InputSet                   InputSet                     `xml:"inputSet"`
	CancelEventDefinition      []CancelEventDefinition      `xml:"cancelEventDefinition"`
	CompensateEventDefinition  []CompensateEventDefinition  `xml:"compensateEventDefinition"`
	ConditionalEventDefinition []ConditionalEventDefinition `xml:"conditionalEventDefinition"`
	ErrorEventDefinition       []ErrorEventDefinition       `xml:"errorEventDefinition"`
	EscalationEventDefinition  []EscalationEventDefinition  `xml:"escalationEventDefinition"`
	LinkEventDefinition        []LinkEventDefinition        `xml:"linkEventDefinition"`
	MessageEventDefinition     []MessageEventDefinition     `xml:"messageEventDefinition"`
	SignalEventDefinition      []SignalEventDefinition      `xml:"signalEventDefinition"`
	TerminateEventDefinition   []TerminateEventDefinition   `xml:"terminateEventDefinition"`
	TimerEventDefinition       []TimerEventDefinition       `xml:"timerEventDefinition"`
	EventDefinitionRef         []string                     `xml:"eventDefinitionRef"`
}

// ParseTree of components of ThrowEvent.
func (throwEvent *ThrowEvent) ParseTree(definitions *Definitions) {
	throwEvent.Event.ParseTree(definitions)

	for i := 0; i < len(throwEvent.DataInput); i++ {
		throwEvent.DataInput[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.DataInputAssociation); i++ {
		throwEvent.DataInputAssociation[i].ParseTree(definitions)
	}

	throwEvent.InputSet.ParseTree(definitions) // InputSet.

	for i := 0; i < len(throwEvent.CancelEventDefinition); i++ {
		throwEvent.CancelEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.CompensateEventDefinition); i++ {
		throwEvent.CompensateEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.ConditionalEventDefinition); i++ {
		throwEvent.ConditionalEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.ErrorEventDefinition); i++ {
		throwEvent.ErrorEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.EscalationEventDefinition); i++ {
		throwEvent.EscalationEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.LinkEventDefinition); i++ {
		throwEvent.LinkEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.MessageEventDefinition); i++ {
		throwEvent.MessageEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.SignalEventDefinition); i++ {
		throwEvent.SignalEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.TerminateEventDefinition); i++ {
		throwEvent.TerminateEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(throwEvent.TimerEventDefinition); i++ {
		throwEvent.TimerEventDefinition[i].ParseTree(definitions)
	}

}

// GetDataInput by ID
func (throwEvent *ThrowEvent) GetDataInput(ID string) *DataInput {

	for i := 0; i < len(throwEvent.DataInput); i++ {
		if throwEvent.DataInput[i].ID == ID {
			return &throwEvent.DataInput[i]

		}
	}

	return nil
}

// GetDataInputAssociation by ID
func (throwEvent *ThrowEvent) GetDataInputAssociation(ID string) *DataInputAssociation {

	for i := 0; i < len(throwEvent.DataInputAssociation); i++ {
		if throwEvent.DataInputAssociation[i].ID == ID {
			return &throwEvent.DataInputAssociation[i]

		}
	}

	return nil
}

// GetCancelEventDefinition by ID
func (throwEvent *ThrowEvent) GetCancelEventDefinition(ID string) *CancelEventDefinition {

	for i := 0; i < len(throwEvent.CancelEventDefinition); i++ {
		if throwEvent.CancelEventDefinition[i].ID == ID {
			return &throwEvent.CancelEventDefinition[i]

		}
	}

	return nil
}

// GetCompensateEventDefinition by ID
func (throwEvent *ThrowEvent) GetCompensateEventDefinition(ID string) *CompensateEventDefinition {

	for i := 0; i < len(throwEvent.CompensateEventDefinition); i++ {
		if throwEvent.CompensateEventDefinition[i].ID == ID {
			return &throwEvent.CompensateEventDefinition[i]

		}
	}

	return nil
}

// GetConditionalEventDefinition by ID
func (throwEvent *ThrowEvent) GetConditionalEventDefinition(ID string) *ConditionalEventDefinition {

	for i := 0; i < len(throwEvent.ConditionalEventDefinition); i++ {
		if throwEvent.ConditionalEventDefinition[i].ID == ID {
			return &throwEvent.ConditionalEventDefinition[i]

		}
	}

	return nil
}

// GetErrorEventDefinition by ID
func (throwEvent *ThrowEvent) GetErrorEventDefinition(ID string) *ErrorEventDefinition {

	for i := 0; i < len(throwEvent.ErrorEventDefinition); i++ {
		if throwEvent.ErrorEventDefinition[i].ID == ID {
			return &throwEvent.ErrorEventDefinition[i]

		}
	}

	return nil
}

// GetEscalationEventDefinition by ID
func (throwEvent *ThrowEvent) GetEscalationEventDefinition(ID string) *EscalationEventDefinition {

	for i := 0; i < len(throwEvent.EscalationEventDefinition); i++ {
		if throwEvent.EscalationEventDefinition[i].ID == ID {
			return &throwEvent.EscalationEventDefinition[i]

		}
	}

	return nil
}

// GetLinkEventDefinition by ID
func (throwEvent *ThrowEvent) GetLinkEventDefinition(ID string) *LinkEventDefinition {

	for i := 0; i < len(throwEvent.LinkEventDefinition); i++ {
		if throwEvent.LinkEventDefinition[i].ID == ID {
			return &throwEvent.LinkEventDefinition[i]

		}
	}

	return nil
}

// GetMessageEventDefinition by ID
func (throwEvent *ThrowEvent) GetMessageEventDefinition(ID string) *MessageEventDefinition {

	for i := 0; i < len(throwEvent.MessageEventDefinition); i++ {
		if throwEvent.MessageEventDefinition[i].ID == ID {
			return &throwEvent.MessageEventDefinition[i]

		}
	}

	return nil
}

// GetSignalEventDefinition by ID
func (throwEvent *ThrowEvent) GetSignalEventDefinition(ID string) *SignalEventDefinition {

	for i := 0; i < len(throwEvent.SignalEventDefinition); i++ {
		if throwEvent.SignalEventDefinition[i].ID == ID {
			return &throwEvent.SignalEventDefinition[i]

		}
	}

	return nil
}

// GetTerminateEventDefinition by ID
func (throwEvent *ThrowEvent) GetTerminateEventDefinition(ID string) *TerminateEventDefinition {

	for i := 0; i < len(throwEvent.TerminateEventDefinition); i++ {
		if throwEvent.TerminateEventDefinition[i].ID == ID {
			return &throwEvent.TerminateEventDefinition[i]

		}
	}

	return nil
}

// GetTimerEventDefinition by ID
func (throwEvent *ThrowEvent) GetTimerEventDefinition(ID string) *TimerEventDefinition {

	for i := 0; i < len(throwEvent.TimerEventDefinition); i++ {
		if throwEvent.TimerEventDefinition[i].ID == ID {
			return &throwEvent.TimerEventDefinition[i]

		}
	}

	return nil
}
