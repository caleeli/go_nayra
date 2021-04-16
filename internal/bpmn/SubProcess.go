package bpmn

import (
	"encoding/xml"
)

// SubProcess from BPMN
type SubProcess struct {
	Activity
	XMLName                xml.Name
	LaneSet                []LaneSet                `xml:"laneSet"`
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
	Association            []Association            `xml:"association"`
	Group                  []Group                  `xml:"group"`
	TextAnnotation         []TextAnnotation         `xml:"textAnnotation"`
	TriggeredByEvent       bool                     `xml:"triggeredByEvent,attr"`

}

// ParseTree of components of SubProcess.
func (subProcess *SubProcess) ParseTree (definitions *Definitions) {
	subProcess.Activity.ParseTree(definitions)

	for i := 0; i < len(subProcess.LaneSet); i++ {
		subProcess.LaneSet[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.AdHocSubProcess); i++ {
		subProcess.AdHocSubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.BoundaryEvent); i++ {
		subProcess.BoundaryEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.BusinessRuleTask); i++ {
		subProcess.BusinessRuleTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.CallActivity); i++ {
		subProcess.CallActivity[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.CallChoreography); i++ {
		subProcess.CallChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ChoreographyTask); i++ {
		subProcess.ChoreographyTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ComplexGateway); i++ {
		subProcess.ComplexGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.DataObject); i++ {
		subProcess.DataObject[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.DataObjectReference); i++ {
		subProcess.DataObjectReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.DataStoreReference); i++ {
		subProcess.DataStoreReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.EndEvent); i++ {
		subProcess.EndEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.EventBasedGateway); i++ {
		subProcess.EventBasedGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ExclusiveGateway); i++ {
		subProcess.ExclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ImplicitThrowEvent); i++ {
		subProcess.ImplicitThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.InclusiveGateway); i++ {
		subProcess.InclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.IntermediateCatchEvent); i++ {
		subProcess.IntermediateCatchEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.IntermediateThrowEvent); i++ {
		subProcess.IntermediateThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ManualTask); i++ {
		subProcess.ManualTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ParallelGateway); i++ {
		subProcess.ParallelGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ReceiveTask); i++ {
		subProcess.ReceiveTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ScriptTask); i++ {
		subProcess.ScriptTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.SendTask); i++ {
		subProcess.SendTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.SequenceFlow); i++ {
		subProcess.SequenceFlow[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.ServiceTask); i++ {
		subProcess.ServiceTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.StartEvent); i++ {
		subProcess.StartEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		subProcess.SubChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.SubProcess); i++ {
		subProcess.SubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.Task); i++ {
		subProcess.Task[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.Transaction); i++ {
		subProcess.Transaction[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.UserTask); i++ {
		subProcess.UserTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.Association); i++ {
		subProcess.Association[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.Group); i++ {
		subProcess.Group[i].ParseTree(definitions)
	}

	for i := 0; i < len(subProcess.TextAnnotation); i++ {
		subProcess.TextAnnotation[i].ParseTree(definitions)
	}

}

// GetLaneSet by ID
func (subProcess *SubProcess) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(subProcess.LaneSet); i++ {
		if subProcess.LaneSet[i].ID == ID {
			return &subProcess.LaneSet[i]

		}
	}

	for i := 0; i < len(subProcess.LaneSet); i++ {
		LaneSet := subProcess.LaneSet[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}

// GetLane by ID
func (subProcess *SubProcess) GetLane(ID string) *Lane {

	for i := 0; i < len(subProcess.LaneSet); i++ {
		Lane := subProcess.LaneSet[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}

// GetBaseElement by ID
func (subProcess *SubProcess) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(subProcess.LaneSet); i++ {
		BaseElement := subProcess.LaneSet[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (subProcess *SubProcess) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(subProcess.LaneSet); i++ {
		Documentation := subProcess.LaneSet[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetAdHocSubProcess by ID
func (subProcess *SubProcess) GetAdHocSubProcess(ID string) *AdHocSubProcess {

	for i := 0; i < len(subProcess.AdHocSubProcess); i++ {
		if subProcess.AdHocSubProcess[i].ID == ID {
			return &subProcess.AdHocSubProcess[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		AdHocSubProcess := subProcess.SubChoreography[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	return nil
}

// GetExpression by ID
func (subProcess *SubProcess) GetExpression(ID string) *Expression {

	for i := 0; i < len(subProcess.AdHocSubProcess); i++ {
		Expression := subProcess.AdHocSubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subProcess.ComplexGateway); i++ {
		Expression := subProcess.ComplexGateway[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subProcess.SequenceFlow); i++ {
		Expression := subProcess.SequenceFlow[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Expression := subProcess.SubChoreography[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}

// GetBoundaryEvent by ID
func (subProcess *SubProcess) GetBoundaryEvent(ID string) *BoundaryEvent {

	for i := 0; i < len(subProcess.BoundaryEvent); i++ {
		if subProcess.BoundaryEvent[i].ID == ID {
			return &subProcess.BoundaryEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		BoundaryEvent := subProcess.SubChoreography[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	return nil
}

// GetBusinessRuleTask by ID
func (subProcess *SubProcess) GetBusinessRuleTask(ID string) *BusinessRuleTask {

	for i := 0; i < len(subProcess.BusinessRuleTask); i++ {
		if subProcess.BusinessRuleTask[i].ID == ID {
			return &subProcess.BusinessRuleTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		BusinessRuleTask := subProcess.SubChoreography[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	return nil
}

// GetCallActivity by ID
func (subProcess *SubProcess) GetCallActivity(ID string) *CallActivity {

	for i := 0; i < len(subProcess.CallActivity); i++ {
		if subProcess.CallActivity[i].ID == ID {
			return &subProcess.CallActivity[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		CallActivity := subProcess.SubChoreography[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	return nil
}

// GetCallChoreography by ID
func (subProcess *SubProcess) GetCallChoreography(ID string) *CallChoreography {

	for i := 0; i < len(subProcess.CallChoreography); i++ {
		if subProcess.CallChoreography[i].ID == ID {
			return &subProcess.CallChoreography[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		CallChoreography := subProcess.SubChoreography[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (subProcess *SubProcess) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(subProcess.CallChoreography); i++ {
		ParticipantAssociation := subProcess.CallChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ParticipantAssociation := subProcess.SubChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetChoreographyTask by ID
func (subProcess *SubProcess) GetChoreographyTask(ID string) *ChoreographyTask {

	for i := 0; i < len(subProcess.ChoreographyTask); i++ {
		if subProcess.ChoreographyTask[i].ID == ID {
			return &subProcess.ChoreographyTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ChoreographyTask := subProcess.SubChoreography[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	return nil
}

// GetComplexGateway by ID
func (subProcess *SubProcess) GetComplexGateway(ID string) *ComplexGateway {

	for i := 0; i < len(subProcess.ComplexGateway); i++ {
		if subProcess.ComplexGateway[i].ID == ID {
			return &subProcess.ComplexGateway[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ComplexGateway := subProcess.SubChoreography[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	return nil
}

// GetDataObject by ID
func (subProcess *SubProcess) GetDataObject(ID string) *DataObject {

	for i := 0; i < len(subProcess.DataObject); i++ {
		if subProcess.DataObject[i].ID == ID {
			return &subProcess.DataObject[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		DataObject := subProcess.SubChoreography[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	return nil
}

// GetDataObjectReference by ID
func (subProcess *SubProcess) GetDataObjectReference(ID string) *DataObjectReference {

	for i := 0; i < len(subProcess.DataObjectReference); i++ {
		if subProcess.DataObjectReference[i].ID == ID {
			return &subProcess.DataObjectReference[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		DataObjectReference := subProcess.SubChoreography[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	return nil
}

// GetDataStoreReference by ID
func (subProcess *SubProcess) GetDataStoreReference(ID string) *DataStoreReference {

	for i := 0; i < len(subProcess.DataStoreReference); i++ {
		if subProcess.DataStoreReference[i].ID == ID {
			return &subProcess.DataStoreReference[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		DataStoreReference := subProcess.SubChoreography[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	return nil
}

// GetEndEvent by ID
func (subProcess *SubProcess) GetEndEvent(ID string) *EndEvent {

	for i := 0; i < len(subProcess.EndEvent); i++ {
		if subProcess.EndEvent[i].ID == ID {
			return &subProcess.EndEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		EndEvent := subProcess.SubChoreography[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	return nil
}

// GetEventBasedGateway by ID
func (subProcess *SubProcess) GetEventBasedGateway(ID string) *EventBasedGateway {

	for i := 0; i < len(subProcess.EventBasedGateway); i++ {
		if subProcess.EventBasedGateway[i].ID == ID {
			return &subProcess.EventBasedGateway[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		EventBasedGateway := subProcess.SubChoreography[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	return nil
}

// GetExclusiveGateway by ID
func (subProcess *SubProcess) GetExclusiveGateway(ID string) *ExclusiveGateway {

	for i := 0; i < len(subProcess.ExclusiveGateway); i++ {
		if subProcess.ExclusiveGateway[i].ID == ID {
			return &subProcess.ExclusiveGateway[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ExclusiveGateway := subProcess.SubChoreography[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (subProcess *SubProcess) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(subProcess.ImplicitThrowEvent); i++ {
		if subProcess.ImplicitThrowEvent[i].ID == ID {
			return &subProcess.ImplicitThrowEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ImplicitThrowEvent := subProcess.SubChoreography[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}

// GetInclusiveGateway by ID
func (subProcess *SubProcess) GetInclusiveGateway(ID string) *InclusiveGateway {

	for i := 0; i < len(subProcess.InclusiveGateway); i++ {
		if subProcess.InclusiveGateway[i].ID == ID {
			return &subProcess.InclusiveGateway[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		InclusiveGateway := subProcess.SubChoreography[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	return nil
}

// GetIntermediateCatchEvent by ID
func (subProcess *SubProcess) GetIntermediateCatchEvent(ID string) *IntermediateCatchEvent {

	for i := 0; i < len(subProcess.IntermediateCatchEvent); i++ {
		if subProcess.IntermediateCatchEvent[i].ID == ID {
			return &subProcess.IntermediateCatchEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		IntermediateCatchEvent := subProcess.SubChoreography[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	return nil
}

// GetIntermediateThrowEvent by ID
func (subProcess *SubProcess) GetIntermediateThrowEvent(ID string) *IntermediateThrowEvent {

	for i := 0; i < len(subProcess.IntermediateThrowEvent); i++ {
		if subProcess.IntermediateThrowEvent[i].ID == ID {
			return &subProcess.IntermediateThrowEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		IntermediateThrowEvent := subProcess.SubChoreography[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	return nil
}

// GetManualTask by ID
func (subProcess *SubProcess) GetManualTask(ID string) *ManualTask {

	for i := 0; i < len(subProcess.ManualTask); i++ {
		if subProcess.ManualTask[i].ID == ID {
			return &subProcess.ManualTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ManualTask := subProcess.SubChoreography[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	return nil
}

// GetParallelGateway by ID
func (subProcess *SubProcess) GetParallelGateway(ID string) *ParallelGateway {

	for i := 0; i < len(subProcess.ParallelGateway); i++ {
		if subProcess.ParallelGateway[i].ID == ID {
			return &subProcess.ParallelGateway[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ParallelGateway := subProcess.SubChoreography[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	return nil
}

// GetReceiveTask by ID
func (subProcess *SubProcess) GetReceiveTask(ID string) *ReceiveTask {

	for i := 0; i < len(subProcess.ReceiveTask); i++ {
		if subProcess.ReceiveTask[i].ID == ID {
			return &subProcess.ReceiveTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ReceiveTask := subProcess.SubChoreography[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	return nil
}

// GetScriptTask by ID
func (subProcess *SubProcess) GetScriptTask(ID string) *ScriptTask {

	for i := 0; i < len(subProcess.ScriptTask); i++ {
		if subProcess.ScriptTask[i].ID == ID {
			return &subProcess.ScriptTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ScriptTask := subProcess.SubChoreography[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	return nil
}

// GetSendTask by ID
func (subProcess *SubProcess) GetSendTask(ID string) *SendTask {

	for i := 0; i < len(subProcess.SendTask); i++ {
		if subProcess.SendTask[i].ID == ID {
			return &subProcess.SendTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		SendTask := subProcess.SubChoreography[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	return nil
}

// GetSequenceFlow by ID
func (subProcess *SubProcess) GetSequenceFlow(ID string) *SequenceFlow {

	for i := 0; i < len(subProcess.SequenceFlow); i++ {
		if subProcess.SequenceFlow[i].ID == ID {
			return &subProcess.SequenceFlow[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		SequenceFlow := subProcess.SubChoreography[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	return nil
}

// GetServiceTask by ID
func (subProcess *SubProcess) GetServiceTask(ID string) *ServiceTask {

	for i := 0; i < len(subProcess.ServiceTask); i++ {
		if subProcess.ServiceTask[i].ID == ID {
			return &subProcess.ServiceTask[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		ServiceTask := subProcess.SubChoreography[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	return nil
}

// GetStartEvent by ID
func (subProcess *SubProcess) GetStartEvent(ID string) *StartEvent {

	for i := 0; i < len(subProcess.StartEvent); i++ {
		if subProcess.StartEvent[i].ID == ID {
			return &subProcess.StartEvent[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		StartEvent := subProcess.SubChoreography[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	return nil
}

// GetSubChoreography by ID
func (subProcess *SubProcess) GetSubChoreography(ID string) *SubChoreography {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		if subProcess.SubChoreography[i].ID == ID {
			return &subProcess.SubChoreography[i]

		}
	}

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		SubChoreography := subProcess.SubChoreography[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	return nil
}

// GetSubProcess by ID
func (subProcess *SubProcess) GetSubProcess(ID string) *SubProcess {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		SubProcess := subProcess.SubChoreography[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	for i := 0; i < len(subProcess.SubProcess); i++ {
		if subProcess.SubProcess[i].ID == ID {
			return &subProcess.SubProcess[i]

		}
	}

	return nil
}

// GetTask by ID
func (subProcess *SubProcess) GetTask(ID string) *Task {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Task := subProcess.SubChoreography[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(subProcess.Task); i++ {
		if subProcess.Task[i].ID == ID {
			return &subProcess.Task[i]

		}
	}

	return nil
}

// GetTransaction by ID
func (subProcess *SubProcess) GetTransaction(ID string) *Transaction {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Transaction := subProcess.SubChoreography[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(subProcess.Transaction); i++ {
		if subProcess.Transaction[i].ID == ID {
			return &subProcess.Transaction[i]

		}
	}

	return nil
}

// GetUserTask by ID
func (subProcess *SubProcess) GetUserTask(ID string) *UserTask {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		UserTask := subProcess.SubChoreography[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(subProcess.UserTask); i++ {
		if subProcess.UserTask[i].ID == ID {
			return &subProcess.UserTask[i]

		}
	}

	return nil
}

// GetRendering by ID
func (subProcess *SubProcess) GetRendering(ID string) *Rendering {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Rendering := subProcess.SubChoreography[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(subProcess.UserTask); i++ {
		Rendering := subProcess.UserTask[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	return nil
}

// GetAssociation by ID
func (subProcess *SubProcess) GetAssociation(ID string) *Association {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Association := subProcess.SubChoreography[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(subProcess.Association); i++ {
		if subProcess.Association[i].ID == ID {
			return &subProcess.Association[i]

		}
	}

	return nil
}

// GetGroup by ID
func (subProcess *SubProcess) GetGroup(ID string) *Group {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		Group := subProcess.SubChoreography[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(subProcess.Group); i++ {
		if subProcess.Group[i].ID == ID {
			return &subProcess.Group[i]

		}
	}

	return nil
}

// GetTextAnnotation by ID
func (subProcess *SubProcess) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(subProcess.SubChoreography); i++ {
		TextAnnotation := subProcess.SubChoreography[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(subProcess.TextAnnotation); i++ {
		if subProcess.TextAnnotation[i].ID == ID {
			return &subProcess.TextAnnotation[i]

		}
	}

	return nil
}

