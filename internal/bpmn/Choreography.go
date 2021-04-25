package bpmn

import (
	"encoding/xml"
)

// Choreography from BPMN
type Choreography struct {
	Collaboration
	XMLName                xml.Name
	AdHocSubProcess        []AdHocSubProcess        `xml:"adHocSubProcess"`
	BoundaryEvent          []BoundaryEvent          `xml:"boundaryEvent"`
	BusinessRuleTask       []BusinessRuleTask       `xml:"businessRuleTask"`
	CallActivity           []CallActivity           `xml:"callActivity"`
	CallChoreography       []CallChoreography       `xml:"callChoreography"`
	ChoreographyTask       []ChoreographyTask       `xml:"choreographyTask"`
	ComplexGateway         []ComplexGateway         `xml:"complexGateway"`
	DataObject             []DataObject             `xml:"dataObject"`
	DataObjectReference    []DataObjectReference    `xml:"dataObjectReference"`
	DataStoreReference     []DataStoreReference     `xml:"dataStoreReference"`
	EndEvent               []EndEvent               `xml:"endEvent"`
	EventBasedGateway      []EventBasedGateway      `xml:"eventBasedGateway"`
	ExclusiveGateway       []ExclusiveGateway       `xml:"exclusiveGateway"`
	ImplicitThrowEvent     []ImplicitThrowEvent     `xml:"implicitThrowEvent"`
	InclusiveGateway       []InclusiveGateway       `xml:"inclusiveGateway"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	IntermediateThrowEvent []IntermediateThrowEvent `xml:"intermediateThrowEvent"`
	ManualTask             []ManualTask             `xml:"manualTask"`
	ParallelGateway        []ParallelGateway        `xml:"parallelGateway"`
	ReceiveTask            []ReceiveTask            `xml:"receiveTask"`
	ScriptTask             []ScriptTask             `xml:"scriptTask"`
	SendTask               []SendTask               `xml:"sendTask"`
	SequenceFlow           []SequenceFlow           `xml:"sequenceFlow"`
	ServiceTask            []ServiceTask            `xml:"serviceTask"`
	StartEvent             []StartEvent             `xml:"startEvent"`
	SubChoreography        []SubChoreography        `xml:"subChoreography"`
	SubProcess             []SubProcess             `xml:"subProcess"`
	Task                   []Task                   `xml:"task"`
	Transaction            []Transaction            `xml:"transaction"`
	UserTask               []UserTask               `xml:"userTask"`
}

// ParseTree of components of Choreography.
func (choreography *Choreography) ParseTree(definitions *Definitions) {
	choreography.Collaboration.ParseTree(definitions)

	for i := 0; i < len(choreography.AdHocSubProcess); i++ {
		choreography.AdHocSubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.BoundaryEvent); i++ {
		choreography.BoundaryEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.BusinessRuleTask); i++ {
		choreography.BusinessRuleTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.CallActivity); i++ {
		choreography.CallActivity[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.CallChoreography); i++ {
		choreography.CallChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ChoreographyTask); i++ {
		choreography.ChoreographyTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ComplexGateway); i++ {
		choreography.ComplexGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.DataObject); i++ {
		choreography.DataObject[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.DataObjectReference); i++ {
		choreography.DataObjectReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.DataStoreReference); i++ {
		choreography.DataStoreReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.EndEvent); i++ {
		choreography.EndEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.EventBasedGateway); i++ {
		choreography.EventBasedGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ExclusiveGateway); i++ {
		choreography.ExclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ImplicitThrowEvent); i++ {
		choreography.ImplicitThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.InclusiveGateway); i++ {
		choreography.InclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.IntermediateCatchEvent); i++ {
		choreography.IntermediateCatchEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.IntermediateThrowEvent); i++ {
		choreography.IntermediateThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ManualTask); i++ {
		choreography.ManualTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ParallelGateway); i++ {
		choreography.ParallelGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ReceiveTask); i++ {
		choreography.ReceiveTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ScriptTask); i++ {
		choreography.ScriptTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.SendTask); i++ {
		choreography.SendTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.SequenceFlow); i++ {
		choreography.SequenceFlow[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.ServiceTask); i++ {
		choreography.ServiceTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.StartEvent); i++ {
		choreography.StartEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		choreography.SubChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		choreography.SubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.Task); i++ {
		choreography.Task[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.Transaction); i++ {
		choreography.Transaction[i].ParseTree(definitions)
	}

	for i := 0; i < len(choreography.UserTask); i++ {
		choreography.UserTask[i].ParseTree(definitions)
	}

}

// GetAdHocSubProcess by ID
func (choreography *Choreography) GetAdHocSubProcess(ID string) *AdHocSubProcess {

	for i := 0; i < len(choreography.AdHocSubProcess); i++ {
		if choreography.AdHocSubProcess[i].ID == ID {
			return &choreography.AdHocSubProcess[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		AdHocSubProcess := choreography.SubChoreography[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		AdHocSubProcess := choreography.SubProcess[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	return nil
}

// GetExpression by ID
func (choreography *Choreography) GetExpression(ID string) *Expression {

	for i := 0; i < len(choreography.AdHocSubProcess); i++ {
		Expression := choreography.AdHocSubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(choreography.ComplexGateway); i++ {
		Expression := choreography.ComplexGateway[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(choreography.SequenceFlow); i++ {
		Expression := choreography.SequenceFlow[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Expression := choreography.SubChoreography[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Expression := choreography.SubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}

// GetBoundaryEvent by ID
func (choreography *Choreography) GetBoundaryEvent(ID string) *BoundaryEvent {

	for i := 0; i < len(choreography.BoundaryEvent); i++ {
		if choreography.BoundaryEvent[i].ID == ID {
			return &choreography.BoundaryEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		BoundaryEvent := choreography.SubChoreography[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		BoundaryEvent := choreography.SubProcess[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	return nil
}

// GetBusinessRuleTask by ID
func (choreography *Choreography) GetBusinessRuleTask(ID string) *BusinessRuleTask {

	for i := 0; i < len(choreography.BusinessRuleTask); i++ {
		if choreography.BusinessRuleTask[i].ID == ID {
			return &choreography.BusinessRuleTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		BusinessRuleTask := choreography.SubChoreography[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		BusinessRuleTask := choreography.SubProcess[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	return nil
}

// GetCallActivity by ID
func (choreography *Choreography) GetCallActivity(ID string) *CallActivity {

	for i := 0; i < len(choreography.CallActivity); i++ {
		if choreography.CallActivity[i].ID == ID {
			return &choreography.CallActivity[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		CallActivity := choreography.SubChoreography[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		CallActivity := choreography.SubProcess[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	return nil
}

// GetCallChoreography by ID
func (choreography *Choreography) GetCallChoreography(ID string) *CallChoreography {

	for i := 0; i < len(choreography.CallChoreography); i++ {
		if choreography.CallChoreography[i].ID == ID {
			return &choreography.CallChoreography[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		CallChoreography := choreography.SubChoreography[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		CallChoreography := choreography.SubProcess[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (choreography *Choreography) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(choreography.CallChoreography); i++ {
		ParticipantAssociation := choreography.CallChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ParticipantAssociation := choreography.SubChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ParticipantAssociation := choreography.SubProcess[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetChoreographyTask by ID
func (choreography *Choreography) GetChoreographyTask(ID string) *ChoreographyTask {

	for i := 0; i < len(choreography.ChoreographyTask); i++ {
		if choreography.ChoreographyTask[i].ID == ID {
			return &choreography.ChoreographyTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ChoreographyTask := choreography.SubChoreography[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ChoreographyTask := choreography.SubProcess[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	return nil
}

// GetComplexGateway by ID
func (choreography *Choreography) GetComplexGateway(ID string) *ComplexGateway {

	for i := 0; i < len(choreography.ComplexGateway); i++ {
		if choreography.ComplexGateway[i].ID == ID {
			return &choreography.ComplexGateway[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ComplexGateway := choreography.SubChoreography[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ComplexGateway := choreography.SubProcess[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	return nil
}

// GetDataObject by ID
func (choreography *Choreography) GetDataObject(ID string) *DataObject {

	for i := 0; i < len(choreography.DataObject); i++ {
		if choreography.DataObject[i].ID == ID {
			return &choreography.DataObject[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		DataObject := choreography.SubChoreography[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		DataObject := choreography.SubProcess[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	return nil
}

// GetDataObjectReference by ID
func (choreography *Choreography) GetDataObjectReference(ID string) *DataObjectReference {

	for i := 0; i < len(choreography.DataObjectReference); i++ {
		if choreography.DataObjectReference[i].ID == ID {
			return &choreography.DataObjectReference[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		DataObjectReference := choreography.SubChoreography[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		DataObjectReference := choreography.SubProcess[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	return nil
}

// GetDataStoreReference by ID
func (choreography *Choreography) GetDataStoreReference(ID string) *DataStoreReference {

	for i := 0; i < len(choreography.DataStoreReference); i++ {
		if choreography.DataStoreReference[i].ID == ID {
			return &choreography.DataStoreReference[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		DataStoreReference := choreography.SubChoreography[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		DataStoreReference := choreography.SubProcess[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	return nil
}

// GetEndEvent by ID
func (choreography *Choreography) GetEndEvent(ID string) *EndEvent {

	for i := 0; i < len(choreography.EndEvent); i++ {
		if choreography.EndEvent[i].ID == ID {
			return &choreography.EndEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		EndEvent := choreography.SubChoreography[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		EndEvent := choreography.SubProcess[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	return nil
}

// GetEventBasedGateway by ID
func (choreography *Choreography) GetEventBasedGateway(ID string) *EventBasedGateway {

	for i := 0; i < len(choreography.EventBasedGateway); i++ {
		if choreography.EventBasedGateway[i].ID == ID {
			return &choreography.EventBasedGateway[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		EventBasedGateway := choreography.SubChoreography[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		EventBasedGateway := choreography.SubProcess[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	return nil
}

// GetExclusiveGateway by ID
func (choreography *Choreography) GetExclusiveGateway(ID string) *ExclusiveGateway {

	for i := 0; i < len(choreography.ExclusiveGateway); i++ {
		if choreography.ExclusiveGateway[i].ID == ID {
			return &choreography.ExclusiveGateway[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ExclusiveGateway := choreography.SubChoreography[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ExclusiveGateway := choreography.SubProcess[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (choreography *Choreography) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(choreography.ImplicitThrowEvent); i++ {
		if choreography.ImplicitThrowEvent[i].ID == ID {
			return &choreography.ImplicitThrowEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ImplicitThrowEvent := choreography.SubChoreography[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ImplicitThrowEvent := choreography.SubProcess[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}

// GetInclusiveGateway by ID
func (choreography *Choreography) GetInclusiveGateway(ID string) *InclusiveGateway {

	for i := 0; i < len(choreography.InclusiveGateway); i++ {
		if choreography.InclusiveGateway[i].ID == ID {
			return &choreography.InclusiveGateway[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		InclusiveGateway := choreography.SubChoreography[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		InclusiveGateway := choreography.SubProcess[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	return nil
}

// GetIntermediateCatchEvent by ID
func (choreography *Choreography) GetIntermediateCatchEvent(ID string) *IntermediateCatchEvent {

	for i := 0; i < len(choreography.IntermediateCatchEvent); i++ {
		if choreography.IntermediateCatchEvent[i].ID == ID {
			return &choreography.IntermediateCatchEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		IntermediateCatchEvent := choreography.SubChoreography[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		IntermediateCatchEvent := choreography.SubProcess[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	return nil
}

// GetIntermediateThrowEvent by ID
func (choreography *Choreography) GetIntermediateThrowEvent(ID string) *IntermediateThrowEvent {

	for i := 0; i < len(choreography.IntermediateThrowEvent); i++ {
		if choreography.IntermediateThrowEvent[i].ID == ID {
			return &choreography.IntermediateThrowEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		IntermediateThrowEvent := choreography.SubChoreography[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		IntermediateThrowEvent := choreography.SubProcess[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	return nil
}

// GetManualTask by ID
func (choreography *Choreography) GetManualTask(ID string) *ManualTask {

	for i := 0; i < len(choreography.ManualTask); i++ {
		if choreography.ManualTask[i].ID == ID {
			return &choreography.ManualTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ManualTask := choreography.SubChoreography[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ManualTask := choreography.SubProcess[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	return nil
}

// GetParallelGateway by ID
func (choreography *Choreography) GetParallelGateway(ID string) *ParallelGateway {

	for i := 0; i < len(choreography.ParallelGateway); i++ {
		if choreography.ParallelGateway[i].ID == ID {
			return &choreography.ParallelGateway[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ParallelGateway := choreography.SubChoreography[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ParallelGateway := choreography.SubProcess[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	return nil
}

// GetReceiveTask by ID
func (choreography *Choreography) GetReceiveTask(ID string) *ReceiveTask {

	for i := 0; i < len(choreography.ReceiveTask); i++ {
		if choreography.ReceiveTask[i].ID == ID {
			return &choreography.ReceiveTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ReceiveTask := choreography.SubChoreography[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ReceiveTask := choreography.SubProcess[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	return nil
}

// GetScriptTask by ID
func (choreography *Choreography) GetScriptTask(ID string) *ScriptTask {

	for i := 0; i < len(choreography.ScriptTask); i++ {
		if choreography.ScriptTask[i].ID == ID {
			return &choreography.ScriptTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ScriptTask := choreography.SubChoreography[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ScriptTask := choreography.SubProcess[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	return nil
}

// GetSendTask by ID
func (choreography *Choreography) GetSendTask(ID string) *SendTask {

	for i := 0; i < len(choreography.SendTask); i++ {
		if choreography.SendTask[i].ID == ID {
			return &choreography.SendTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		SendTask := choreography.SubChoreography[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		SendTask := choreography.SubProcess[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	return nil
}

// GetSequenceFlow by ID
func (choreography *Choreography) GetSequenceFlow(ID string) *SequenceFlow {

	for i := 0; i < len(choreography.SequenceFlow); i++ {
		if choreography.SequenceFlow[i].ID == ID {
			return &choreography.SequenceFlow[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		SequenceFlow := choreography.SubChoreography[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		SequenceFlow := choreography.SubProcess[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	return nil
}

// GetServiceTask by ID
func (choreography *Choreography) GetServiceTask(ID string) *ServiceTask {

	for i := 0; i < len(choreography.ServiceTask); i++ {
		if choreography.ServiceTask[i].ID == ID {
			return &choreography.ServiceTask[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		ServiceTask := choreography.SubChoreography[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		ServiceTask := choreography.SubProcess[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	return nil
}

// GetStartEvent by ID
func (choreography *Choreography) GetStartEvent(ID string) *StartEvent {

	for i := 0; i < len(choreography.StartEvent); i++ {
		if choreography.StartEvent[i].ID == ID {
			return &choreography.StartEvent[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		StartEvent := choreography.SubChoreography[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		StartEvent := choreography.SubProcess[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	return nil
}

// GetSubChoreography by ID
func (choreography *Choreography) GetSubChoreography(ID string) *SubChoreography {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		if choreography.SubChoreography[i].ID == ID {
			return &choreography.SubChoreography[i]

		}
	}

	for i := 0; i < len(choreography.SubChoreography); i++ {
		SubChoreography := choreography.SubChoreography[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		SubChoreography := choreography.SubProcess[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	return nil
}

// GetSubProcess by ID
func (choreography *Choreography) GetSubProcess(ID string) *SubProcess {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		SubProcess := choreography.SubChoreography[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		if choreography.SubProcess[i].ID == ID {
			return &choreography.SubProcess[i]

		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		SubProcess := choreography.SubProcess[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	return nil
}

// GetLaneSet by ID
func (choreography *Choreography) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		LaneSet := choreography.SubChoreography[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		LaneSet := choreography.SubProcess[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}

// GetLane by ID
func (choreography *Choreography) GetLane(ID string) *Lane {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Lane := choreography.SubChoreography[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Lane := choreography.SubProcess[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}

// GetBaseElement by ID
func (choreography *Choreography) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		BaseElement := choreography.SubChoreography[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		BaseElement := choreography.SubProcess[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (choreography *Choreography) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Documentation := choreography.SubChoreography[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Documentation := choreography.SubProcess[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetTask by ID
func (choreography *Choreography) GetTask(ID string) *Task {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Task := choreography.SubChoreography[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Task := choreography.SubProcess[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(choreography.Task); i++ {
		if choreography.Task[i].ID == ID {
			return &choreography.Task[i]

		}
	}

	return nil
}

// GetTransaction by ID
func (choreography *Choreography) GetTransaction(ID string) *Transaction {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Transaction := choreography.SubChoreography[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Transaction := choreography.SubProcess[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(choreography.Transaction); i++ {
		if choreography.Transaction[i].ID == ID {
			return &choreography.Transaction[i]

		}
	}

	return nil
}

// GetUserTask by ID
func (choreography *Choreography) GetUserTask(ID string) *UserTask {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		UserTask := choreography.SubChoreography[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		UserTask := choreography.SubProcess[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(choreography.UserTask); i++ {
		if choreography.UserTask[i].ID == ID {
			return &choreography.UserTask[i]

		}
	}

	return nil
}

// GetRendering by ID
func (choreography *Choreography) GetRendering(ID string) *Rendering {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Rendering := choreography.SubChoreography[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Rendering := choreography.SubProcess[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(choreography.UserTask); i++ {
		Rendering := choreography.UserTask[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	return nil
}

// GetAssociation by ID
func (choreography *Choreography) GetAssociation(ID string) *Association {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Association := choreography.SubChoreography[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Association := choreography.SubProcess[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	return nil
}

// GetGroup by ID
func (choreography *Choreography) GetGroup(ID string) *Group {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		Group := choreography.SubChoreography[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		Group := choreography.SubProcess[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	return nil
}

// GetTextAnnotation by ID
func (choreography *Choreography) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(choreography.SubChoreography); i++ {
		TextAnnotation := choreography.SubChoreography[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(choreography.SubProcess); i++ {
		TextAnnotation := choreography.SubProcess[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	return nil
}
