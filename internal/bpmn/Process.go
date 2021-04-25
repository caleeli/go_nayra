package bpmn

import (
	"encoding/xml"
)

// Process from BPMN
type Process struct {
	ProcessTrait `bson:"-"`
	CallableElement
	XMLName                      xml.Name
	Auditing                     Auditing                  `xml:"auditing"`
	Monitoring                   Monitoring                `xml:"monitoring"`
	Property                     []Property                `xml:"property"`
	LaneSet                      []LaneSet                 `xml:"laneSet"`
	AdHocSubProcess              []AdHocSubProcess         `xml:"adHocSubProcess"`
	BoundaryEvent                []BoundaryEvent           `xml:"boundaryEvent"`
	BusinessRuleTask             []BusinessRuleTask        `xml:"businessRuleTask"`
	CallActivity                 []CallActivity            `xml:"callActivity"`
	CallChoreography             []CallChoreography        `xml:"callChoreography"`
	ChoreographyTask             []ChoreographyTask        `xml:"choreographyTask"`
	ComplexGateway               []ComplexGateway          `xml:"complexGateway"`
	DataObject                   []DataObject              `xml:"dataObject"`
	DataObjectReference          []DataObjectReference     `xml:"dataObjectReference"`
	DataStoreReference           []DataStoreReference      `xml:"dataStoreReference"`
	EndEvent                     []EndEvent                `xml:"endEvent"`
	EventBasedGateway            []EventBasedGateway       `xml:"eventBasedGateway"`
	ExclusiveGateway             []ExclusiveGateway        `xml:"exclusiveGateway"`
	ImplicitThrowEvent           []ImplicitThrowEvent      `xml:"implicitThrowEvent"`
	InclusiveGateway             []InclusiveGateway        `xml:"inclusiveGateway"`
	IntermediateCatchEvent       []IntermediateCatchEvent  `xml:"intermediateCatchEvent"`
	IntermediateThrowEvent       []IntermediateThrowEvent  `xml:"intermediateThrowEvent"`
	ManualTask                   []ManualTask              `xml:"manualTask"`
	ParallelGateway              []ParallelGateway         `xml:"parallelGateway"`
	ReceiveTask                  []ReceiveTask             `xml:"receiveTask"`
	ScriptTask                   []ScriptTask              `xml:"scriptTask"`
	SendTask                     []SendTask                `xml:"sendTask"`
	SequenceFlow                 []SequenceFlow            `xml:"sequenceFlow"`
	ServiceTask                  []ServiceTask             `xml:"serviceTask"`
	StartEvent                   []StartEvent              `xml:"startEvent"`
	SubChoreography              []SubChoreography         `xml:"subChoreography"`
	SubProcess                   []SubProcess              `xml:"subProcess"`
	Task                         []Task                    `xml:"task"`
	Transaction                  []Transaction             `xml:"transaction"`
	UserTask                     []UserTask                `xml:"userTask"`
	Association                  []Association             `xml:"association"`
	Group                        []Group                   `xml:"group"`
	TextAnnotation               []TextAnnotation          `xml:"textAnnotation"`
	Performer                    []Performer               `xml:"performer"`
	CorrelationSubscription      []CorrelationSubscription `xml:"correlationSubscription"`
	Supports                     []string                  `xml:"supports"`
	ProcessType                  string                    `xml:"processType,attr"`
	IsClosed                     bool                      `xml:"isClosed,attr"`
	IsExecutable                 bool                      `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string                    `xml:"definitionalCollaborationRef,attr"`
}

// ParseTree of components of Process.
func (process *Process) ParseTree(definitions *Definitions) {
	process.CallableElement.ParseTree(definitions)

	process.Init(definitions)

	process.Auditing.ParseTree(definitions) // Auditing.

	process.Monitoring.ParseTree(definitions) // Monitoring.

	for i := 0; i < len(process.Property); i++ {
		process.Property[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.LaneSet); i++ {
		process.LaneSet[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.AdHocSubProcess); i++ {
		process.AdHocSubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.BoundaryEvent); i++ {
		process.BoundaryEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.BusinessRuleTask); i++ {
		process.BusinessRuleTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.CallActivity); i++ {
		process.CallActivity[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.CallChoreography); i++ {
		process.CallChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ChoreographyTask); i++ {
		process.ChoreographyTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ComplexGateway); i++ {
		process.ComplexGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.DataObject); i++ {
		process.DataObject[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.DataObjectReference); i++ {
		process.DataObjectReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.DataStoreReference); i++ {
		process.DataStoreReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.EndEvent); i++ {
		process.EndEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.EventBasedGateway); i++ {
		process.EventBasedGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ExclusiveGateway); i++ {
		process.ExclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ImplicitThrowEvent); i++ {
		process.ImplicitThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.InclusiveGateway); i++ {
		process.InclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.IntermediateCatchEvent); i++ {
		process.IntermediateCatchEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.IntermediateThrowEvent); i++ {
		process.IntermediateThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ManualTask); i++ {
		process.ManualTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ParallelGateway); i++ {
		process.ParallelGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ReceiveTask); i++ {
		process.ReceiveTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ScriptTask); i++ {
		process.ScriptTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.SendTask); i++ {
		process.SendTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.SequenceFlow); i++ {
		process.SequenceFlow[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.ServiceTask); i++ {
		process.ServiceTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.StartEvent); i++ {
		process.StartEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		process.SubChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.SubProcess); i++ {
		process.SubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.Task); i++ {
		process.Task[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.Transaction); i++ {
		process.Transaction[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.UserTask); i++ {
		process.UserTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.Association); i++ {
		process.Association[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.Group); i++ {
		process.Group[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.TextAnnotation); i++ {
		process.TextAnnotation[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.Performer); i++ {
		process.Performer[i].ParseTree(definitions)
	}

	for i := 0; i < len(process.CorrelationSubscription); i++ {
		process.CorrelationSubscription[i].ParseTree(definitions)
	}

}

// GetProperty by ID
func (process *Process) GetProperty(ID string) *Property {

	for i := 0; i < len(process.Property); i++ {
		if process.Property[i].ID == ID {
			return &process.Property[i]

		}
	}

	return nil
}

// GetLaneSet by ID
func (process *Process) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(process.LaneSet); i++ {
		if process.LaneSet[i].ID == ID {
			return &process.LaneSet[i]

		}
	}

	for i := 0; i < len(process.LaneSet); i++ {
		LaneSet := process.LaneSet[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		LaneSet := process.SubChoreography[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		LaneSet := process.SubProcess[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}

// GetLane by ID
func (process *Process) GetLane(ID string) *Lane {

	for i := 0; i < len(process.LaneSet); i++ {
		Lane := process.LaneSet[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		Lane := process.SubChoreography[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Lane := process.SubProcess[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}

// GetBaseElement by ID
func (process *Process) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(process.LaneSet); i++ {
		BaseElement := process.LaneSet[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		BaseElement := process.SubChoreography[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		BaseElement := process.SubProcess[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (process *Process) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(process.LaneSet); i++ {
		Documentation := process.LaneSet[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		Documentation := process.SubChoreography[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Documentation := process.SubProcess[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetAdHocSubProcess by ID
func (process *Process) GetAdHocSubProcess(ID string) *AdHocSubProcess {

	for i := 0; i < len(process.AdHocSubProcess); i++ {
		if process.AdHocSubProcess[i].ID == ID {
			return &process.AdHocSubProcess[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		AdHocSubProcess := process.SubChoreography[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		AdHocSubProcess := process.SubProcess[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	return nil
}

// GetExpression by ID
func (process *Process) GetExpression(ID string) *Expression {

	for i := 0; i < len(process.AdHocSubProcess); i++ {
		Expression := process.AdHocSubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(process.ComplexGateway); i++ {
		Expression := process.ComplexGateway[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(process.SequenceFlow); i++ {
		Expression := process.SequenceFlow[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		Expression := process.SubChoreography[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Expression := process.SubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}

// GetBoundaryEvent by ID
func (process *Process) GetBoundaryEvent(ID string) *BoundaryEvent {

	for i := 0; i < len(process.BoundaryEvent); i++ {
		if process.BoundaryEvent[i].ID == ID {
			return &process.BoundaryEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		BoundaryEvent := process.SubChoreography[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		BoundaryEvent := process.SubProcess[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	return nil
}

// GetBusinessRuleTask by ID
func (process *Process) GetBusinessRuleTask(ID string) *BusinessRuleTask {

	for i := 0; i < len(process.BusinessRuleTask); i++ {
		if process.BusinessRuleTask[i].ID == ID {
			return &process.BusinessRuleTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		BusinessRuleTask := process.SubChoreography[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		BusinessRuleTask := process.SubProcess[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	return nil
}

// GetCallActivity by ID
func (process *Process) GetCallActivity(ID string) *CallActivity {

	for i := 0; i < len(process.CallActivity); i++ {
		if process.CallActivity[i].ID == ID {
			return &process.CallActivity[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		CallActivity := process.SubChoreography[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		CallActivity := process.SubProcess[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	return nil
}

// GetCallChoreography by ID
func (process *Process) GetCallChoreography(ID string) *CallChoreography {

	for i := 0; i < len(process.CallChoreography); i++ {
		if process.CallChoreography[i].ID == ID {
			return &process.CallChoreography[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		CallChoreography := process.SubChoreography[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		CallChoreography := process.SubProcess[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (process *Process) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(process.CallChoreography); i++ {
		ParticipantAssociation := process.CallChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ParticipantAssociation := process.SubChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ParticipantAssociation := process.SubProcess[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetChoreographyTask by ID
func (process *Process) GetChoreographyTask(ID string) *ChoreographyTask {

	for i := 0; i < len(process.ChoreographyTask); i++ {
		if process.ChoreographyTask[i].ID == ID {
			return &process.ChoreographyTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ChoreographyTask := process.SubChoreography[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ChoreographyTask := process.SubProcess[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	return nil
}

// GetComplexGateway by ID
func (process *Process) GetComplexGateway(ID string) *ComplexGateway {

	for i := 0; i < len(process.ComplexGateway); i++ {
		if process.ComplexGateway[i].ID == ID {
			return &process.ComplexGateway[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ComplexGateway := process.SubChoreography[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ComplexGateway := process.SubProcess[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	return nil
}

// GetDataObject by ID
func (process *Process) GetDataObject(ID string) *DataObject {

	for i := 0; i < len(process.DataObject); i++ {
		if process.DataObject[i].ID == ID {
			return &process.DataObject[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		DataObject := process.SubChoreography[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		DataObject := process.SubProcess[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	return nil
}

// GetDataObjectReference by ID
func (process *Process) GetDataObjectReference(ID string) *DataObjectReference {

	for i := 0; i < len(process.DataObjectReference); i++ {
		if process.DataObjectReference[i].ID == ID {
			return &process.DataObjectReference[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		DataObjectReference := process.SubChoreography[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		DataObjectReference := process.SubProcess[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	return nil
}

// GetDataStoreReference by ID
func (process *Process) GetDataStoreReference(ID string) *DataStoreReference {

	for i := 0; i < len(process.DataStoreReference); i++ {
		if process.DataStoreReference[i].ID == ID {
			return &process.DataStoreReference[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		DataStoreReference := process.SubChoreography[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		DataStoreReference := process.SubProcess[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	return nil
}

// GetEndEvent by ID
func (process *Process) GetEndEvent(ID string) *EndEvent {

	for i := 0; i < len(process.EndEvent); i++ {
		if process.EndEvent[i].ID == ID {
			return &process.EndEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		EndEvent := process.SubChoreography[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		EndEvent := process.SubProcess[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	return nil
}

// GetEventBasedGateway by ID
func (process *Process) GetEventBasedGateway(ID string) *EventBasedGateway {

	for i := 0; i < len(process.EventBasedGateway); i++ {
		if process.EventBasedGateway[i].ID == ID {
			return &process.EventBasedGateway[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		EventBasedGateway := process.SubChoreography[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		EventBasedGateway := process.SubProcess[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	return nil
}

// GetExclusiveGateway by ID
func (process *Process) GetExclusiveGateway(ID string) *ExclusiveGateway {

	for i := 0; i < len(process.ExclusiveGateway); i++ {
		if process.ExclusiveGateway[i].ID == ID {
			return &process.ExclusiveGateway[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ExclusiveGateway := process.SubChoreography[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ExclusiveGateway := process.SubProcess[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (process *Process) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(process.ImplicitThrowEvent); i++ {
		if process.ImplicitThrowEvent[i].ID == ID {
			return &process.ImplicitThrowEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ImplicitThrowEvent := process.SubChoreography[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ImplicitThrowEvent := process.SubProcess[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}

// GetInclusiveGateway by ID
func (process *Process) GetInclusiveGateway(ID string) *InclusiveGateway {

	for i := 0; i < len(process.InclusiveGateway); i++ {
		if process.InclusiveGateway[i].ID == ID {
			return &process.InclusiveGateway[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		InclusiveGateway := process.SubChoreography[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		InclusiveGateway := process.SubProcess[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	return nil
}

// GetIntermediateCatchEvent by ID
func (process *Process) GetIntermediateCatchEvent(ID string) *IntermediateCatchEvent {

	for i := 0; i < len(process.IntermediateCatchEvent); i++ {
		if process.IntermediateCatchEvent[i].ID == ID {
			return &process.IntermediateCatchEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		IntermediateCatchEvent := process.SubChoreography[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		IntermediateCatchEvent := process.SubProcess[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	return nil
}

// GetIntermediateThrowEvent by ID
func (process *Process) GetIntermediateThrowEvent(ID string) *IntermediateThrowEvent {

	for i := 0; i < len(process.IntermediateThrowEvent); i++ {
		if process.IntermediateThrowEvent[i].ID == ID {
			return &process.IntermediateThrowEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		IntermediateThrowEvent := process.SubChoreography[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		IntermediateThrowEvent := process.SubProcess[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	return nil
}

// GetManualTask by ID
func (process *Process) GetManualTask(ID string) *ManualTask {

	for i := 0; i < len(process.ManualTask); i++ {
		if process.ManualTask[i].ID == ID {
			return &process.ManualTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ManualTask := process.SubChoreography[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ManualTask := process.SubProcess[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	return nil
}

// GetParallelGateway by ID
func (process *Process) GetParallelGateway(ID string) *ParallelGateway {

	for i := 0; i < len(process.ParallelGateway); i++ {
		if process.ParallelGateway[i].ID == ID {
			return &process.ParallelGateway[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ParallelGateway := process.SubChoreography[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ParallelGateway := process.SubProcess[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	return nil
}

// GetReceiveTask by ID
func (process *Process) GetReceiveTask(ID string) *ReceiveTask {

	for i := 0; i < len(process.ReceiveTask); i++ {
		if process.ReceiveTask[i].ID == ID {
			return &process.ReceiveTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ReceiveTask := process.SubChoreography[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ReceiveTask := process.SubProcess[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	return nil
}

// GetScriptTask by ID
func (process *Process) GetScriptTask(ID string) *ScriptTask {

	for i := 0; i < len(process.ScriptTask); i++ {
		if process.ScriptTask[i].ID == ID {
			return &process.ScriptTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ScriptTask := process.SubChoreography[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ScriptTask := process.SubProcess[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	return nil
}

// GetSendTask by ID
func (process *Process) GetSendTask(ID string) *SendTask {

	for i := 0; i < len(process.SendTask); i++ {
		if process.SendTask[i].ID == ID {
			return &process.SendTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		SendTask := process.SubChoreography[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		SendTask := process.SubProcess[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	return nil
}

// GetSequenceFlow by ID
func (process *Process) GetSequenceFlow(ID string) *SequenceFlow {

	for i := 0; i < len(process.SequenceFlow); i++ {
		if process.SequenceFlow[i].ID == ID {
			return &process.SequenceFlow[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		SequenceFlow := process.SubChoreography[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		SequenceFlow := process.SubProcess[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	return nil
}

// GetServiceTask by ID
func (process *Process) GetServiceTask(ID string) *ServiceTask {

	for i := 0; i < len(process.ServiceTask); i++ {
		if process.ServiceTask[i].ID == ID {
			return &process.ServiceTask[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		ServiceTask := process.SubChoreography[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		ServiceTask := process.SubProcess[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	return nil
}

// GetStartEvent by ID
func (process *Process) GetStartEvent(ID string) *StartEvent {

	for i := 0; i < len(process.StartEvent); i++ {
		if process.StartEvent[i].ID == ID {
			return &process.StartEvent[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		StartEvent := process.SubChoreography[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		StartEvent := process.SubProcess[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	return nil
}

// GetSubChoreography by ID
func (process *Process) GetSubChoreography(ID string) *SubChoreography {

	for i := 0; i < len(process.SubChoreography); i++ {
		if process.SubChoreography[i].ID == ID {
			return &process.SubChoreography[i]

		}
	}

	for i := 0; i < len(process.SubChoreography); i++ {
		SubChoreography := process.SubChoreography[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		SubChoreography := process.SubProcess[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	return nil
}

// GetSubProcess by ID
func (process *Process) GetSubProcess(ID string) *SubProcess {

	for i := 0; i < len(process.SubChoreography); i++ {
		SubProcess := process.SubChoreography[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		if process.SubProcess[i].ID == ID {
			return &process.SubProcess[i]

		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		SubProcess := process.SubProcess[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	return nil
}

// GetTask by ID
func (process *Process) GetTask(ID string) *Task {

	for i := 0; i < len(process.SubChoreography); i++ {
		Task := process.SubChoreography[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Task := process.SubProcess[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(process.Task); i++ {
		if process.Task[i].ID == ID {
			return &process.Task[i]

		}
	}

	return nil
}

// GetTransaction by ID
func (process *Process) GetTransaction(ID string) *Transaction {

	for i := 0; i < len(process.SubChoreography); i++ {
		Transaction := process.SubChoreography[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Transaction := process.SubProcess[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(process.Transaction); i++ {
		if process.Transaction[i].ID == ID {
			return &process.Transaction[i]

		}
	}

	return nil
}

// GetUserTask by ID
func (process *Process) GetUserTask(ID string) *UserTask {

	for i := 0; i < len(process.SubChoreography); i++ {
		UserTask := process.SubChoreography[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		UserTask := process.SubProcess[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(process.UserTask); i++ {
		if process.UserTask[i].ID == ID {
			return &process.UserTask[i]

		}
	}

	return nil
}

// GetRendering by ID
func (process *Process) GetRendering(ID string) *Rendering {

	for i := 0; i < len(process.SubChoreography); i++ {
		Rendering := process.SubChoreography[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Rendering := process.SubProcess[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(process.UserTask); i++ {
		Rendering := process.UserTask[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	return nil
}

// GetAssociation by ID
func (process *Process) GetAssociation(ID string) *Association {

	for i := 0; i < len(process.SubChoreography); i++ {
		Association := process.SubChoreography[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Association := process.SubProcess[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(process.Association); i++ {
		if process.Association[i].ID == ID {
			return &process.Association[i]

		}
	}

	return nil
}

// GetGroup by ID
func (process *Process) GetGroup(ID string) *Group {

	for i := 0; i < len(process.SubChoreography); i++ {
		Group := process.SubChoreography[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		Group := process.SubProcess[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(process.Group); i++ {
		if process.Group[i].ID == ID {
			return &process.Group[i]

		}
	}

	return nil
}

// GetTextAnnotation by ID
func (process *Process) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(process.SubChoreography); i++ {
		TextAnnotation := process.SubChoreography[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(process.SubProcess); i++ {
		TextAnnotation := process.SubProcess[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(process.TextAnnotation); i++ {
		if process.TextAnnotation[i].ID == ID {
			return &process.TextAnnotation[i]

		}
	}

	return nil
}

// GetPerformer by ID
func (process *Process) GetPerformer(ID string) *Performer {

	for i := 0; i < len(process.Performer); i++ {
		if process.Performer[i].ID == ID {
			return &process.Performer[i]

		}
	}

	return nil
}

// GetCorrelationSubscription by ID
func (process *Process) GetCorrelationSubscription(ID string) *CorrelationSubscription {

	for i := 0; i < len(process.CorrelationSubscription); i++ {
		if process.CorrelationSubscription[i].ID == ID {
			return &process.CorrelationSubscription[i]

		}
	}

	return nil
}

// GetCorrelationPropertyBinding by ID
func (process *Process) GetCorrelationPropertyBinding(ID string) *CorrelationPropertyBinding {

	for i := 0; i < len(process.CorrelationSubscription); i++ {
		CorrelationPropertyBinding := process.CorrelationSubscription[i].GetCorrelationPropertyBinding(ID)
		if CorrelationPropertyBinding != nil {
			return CorrelationPropertyBinding
		}
	}

	return nil
}

// GetFormalExpression by ID
func (process *Process) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(process.CorrelationSubscription); i++ {
		FormalExpression := process.CorrelationSubscription[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	return nil
}
