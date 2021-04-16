package bpmn

import (
	"encoding/xml"
)

// CatchEvent from BPMN
type CatchEvent struct {
	Event
	XMLName                    xml.Name
	DataOutput                 []DataOutput                 `xml:"dataOutput"`
	DataOutputAssociation      []DataOutputAssociation      `xml:"dataOutputAssociation"`
	OutputSet                  OutputSet                    `xml:"outputSet"`
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
	ParallelMultiple           bool                         `xml:"parallelMultiple,attr"`

}

// ParseTree of components of CatchEvent.
func (catchEvent *CatchEvent) ParseTree (definitions *Definitions) {
	catchEvent.Event.ParseTree(definitions)

	for i := 0; i < len(catchEvent.DataOutput); i++ {
		catchEvent.DataOutput[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.DataOutputAssociation); i++ {
		catchEvent.DataOutputAssociation[i].ParseTree(definitions)
	}

	catchEvent.OutputSet.ParseTree(definitions) // OutputSet.

	for i := 0; i < len(catchEvent.CancelEventDefinition); i++ {
		catchEvent.CancelEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.CompensateEventDefinition); i++ {
		catchEvent.CompensateEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.ConditionalEventDefinition); i++ {
		catchEvent.ConditionalEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.ErrorEventDefinition); i++ {
		catchEvent.ErrorEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.EscalationEventDefinition); i++ {
		catchEvent.EscalationEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.LinkEventDefinition); i++ {
		catchEvent.LinkEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.MessageEventDefinition); i++ {
		catchEvent.MessageEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.SignalEventDefinition); i++ {
		catchEvent.SignalEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.TerminateEventDefinition); i++ {
		catchEvent.TerminateEventDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(catchEvent.TimerEventDefinition); i++ {
		catchEvent.TimerEventDefinition[i].ParseTree(definitions)
	}

}

// GetDataOutput by ID
func (catchEvent *CatchEvent) GetDataOutput(ID string) *DataOutput {

	for i := 0; i < len(catchEvent.DataOutput); i++ {
		if catchEvent.DataOutput[i].ID == ID {
			return &catchEvent.DataOutput[i]

		}
	}

	return nil
}

// GetDataOutputAssociation by ID
func (catchEvent *CatchEvent) GetDataOutputAssociation(ID string) *DataOutputAssociation {

	for i := 0; i < len(catchEvent.DataOutputAssociation); i++ {
		if catchEvent.DataOutputAssociation[i].ID == ID {
			return &catchEvent.DataOutputAssociation[i]

		}
	}

	return nil
}

// GetCancelEventDefinition by ID
func (catchEvent *CatchEvent) GetCancelEventDefinition(ID string) *CancelEventDefinition {

	for i := 0; i < len(catchEvent.CancelEventDefinition); i++ {
		if catchEvent.CancelEventDefinition[i].ID == ID {
			return &catchEvent.CancelEventDefinition[i]

		}
	}

	return nil
}

// GetCompensateEventDefinition by ID
func (catchEvent *CatchEvent) GetCompensateEventDefinition(ID string) *CompensateEventDefinition {

	for i := 0; i < len(catchEvent.CompensateEventDefinition); i++ {
		if catchEvent.CompensateEventDefinition[i].ID == ID {
			return &catchEvent.CompensateEventDefinition[i]

		}
	}

	return nil
}

// GetConditionalEventDefinition by ID
func (catchEvent *CatchEvent) GetConditionalEventDefinition(ID string) *ConditionalEventDefinition {

	for i := 0; i < len(catchEvent.ConditionalEventDefinition); i++ {
		if catchEvent.ConditionalEventDefinition[i].ID == ID {
			return &catchEvent.ConditionalEventDefinition[i]

		}
	}

	return nil
}

// GetErrorEventDefinition by ID
func (catchEvent *CatchEvent) GetErrorEventDefinition(ID string) *ErrorEventDefinition {

	for i := 0; i < len(catchEvent.ErrorEventDefinition); i++ {
		if catchEvent.ErrorEventDefinition[i].ID == ID {
			return &catchEvent.ErrorEventDefinition[i]

		}
	}

	return nil
}

// GetEscalationEventDefinition by ID
func (catchEvent *CatchEvent) GetEscalationEventDefinition(ID string) *EscalationEventDefinition {

	for i := 0; i < len(catchEvent.EscalationEventDefinition); i++ {
		if catchEvent.EscalationEventDefinition[i].ID == ID {
			return &catchEvent.EscalationEventDefinition[i]

		}
	}

	return nil
}

// GetLinkEventDefinition by ID
func (catchEvent *CatchEvent) GetLinkEventDefinition(ID string) *LinkEventDefinition {

	for i := 0; i < len(catchEvent.LinkEventDefinition); i++ {
		if catchEvent.LinkEventDefinition[i].ID == ID {
			return &catchEvent.LinkEventDefinition[i]

		}
	}

	return nil
}

// GetMessageEventDefinition by ID
func (catchEvent *CatchEvent) GetMessageEventDefinition(ID string) *MessageEventDefinition {

	for i := 0; i < len(catchEvent.MessageEventDefinition); i++ {
		if catchEvent.MessageEventDefinition[i].ID == ID {
			return &catchEvent.MessageEventDefinition[i]

		}
	}

	return nil
}

// GetSignalEventDefinition by ID
func (catchEvent *CatchEvent) GetSignalEventDefinition(ID string) *SignalEventDefinition {

	for i := 0; i < len(catchEvent.SignalEventDefinition); i++ {
		if catchEvent.SignalEventDefinition[i].ID == ID {
			return &catchEvent.SignalEventDefinition[i]

		}
	}

	return nil
}

// GetTerminateEventDefinition by ID
func (catchEvent *CatchEvent) GetTerminateEventDefinition(ID string) *TerminateEventDefinition {

	for i := 0; i < len(catchEvent.TerminateEventDefinition); i++ {
		if catchEvent.TerminateEventDefinition[i].ID == ID {
			return &catchEvent.TerminateEventDefinition[i]

		}
	}

	return nil
}

// GetTimerEventDefinition by ID
func (catchEvent *CatchEvent) GetTimerEventDefinition(ID string) *TimerEventDefinition {

	for i := 0; i < len(catchEvent.TimerEventDefinition); i++ {
		if catchEvent.TimerEventDefinition[i].ID == ID {
			return &catchEvent.TimerEventDefinition[i]

		}
	}

	return nil
}

