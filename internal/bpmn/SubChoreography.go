package bpmn

import (
	"encoding/xml"
)

// SubChoreography from BPMN
type SubChoreography struct {
	ChoreographyActivity
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
	Association            []Association            `xml:"association"`
	Group                  []Group                  `xml:"group"`
	TextAnnotation         []TextAnnotation         `xml:"textAnnotation"`

}

// ParseTree of components of SubChoreography.
func (subChoreography *SubChoreography) ParseTree (definitions *Definitions) {
	subChoreography.ChoreographyActivity.ParseTree(definitions)

	for i := 0; i < len(subChoreography.AdHocSubProcess); i++ {
		subChoreography.AdHocSubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.BoundaryEvent); i++ {
		subChoreography.BoundaryEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.BusinessRuleTask); i++ {
		subChoreography.BusinessRuleTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.CallActivity); i++ {
		subChoreography.CallActivity[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.CallChoreography); i++ {
		subChoreography.CallChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ChoreographyTask); i++ {
		subChoreography.ChoreographyTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ComplexGateway); i++ {
		subChoreography.ComplexGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.DataObject); i++ {
		subChoreography.DataObject[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.DataObjectReference); i++ {
		subChoreography.DataObjectReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.DataStoreReference); i++ {
		subChoreography.DataStoreReference[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.EndEvent); i++ {
		subChoreography.EndEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.EventBasedGateway); i++ {
		subChoreography.EventBasedGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ExclusiveGateway); i++ {
		subChoreography.ExclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ImplicitThrowEvent); i++ {
		subChoreography.ImplicitThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.InclusiveGateway); i++ {
		subChoreography.InclusiveGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.IntermediateCatchEvent); i++ {
		subChoreography.IntermediateCatchEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.IntermediateThrowEvent); i++ {
		subChoreography.IntermediateThrowEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ManualTask); i++ {
		subChoreography.ManualTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ParallelGateway); i++ {
		subChoreography.ParallelGateway[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ReceiveTask); i++ {
		subChoreography.ReceiveTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ScriptTask); i++ {
		subChoreography.ScriptTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.SendTask); i++ {
		subChoreography.SendTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.SequenceFlow); i++ {
		subChoreography.SequenceFlow[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.ServiceTask); i++ {
		subChoreography.ServiceTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.StartEvent); i++ {
		subChoreography.StartEvent[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.SubChoreography); i++ {
		subChoreography.SubChoreography[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		subChoreography.SubProcess[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.Task); i++ {
		subChoreography.Task[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.Transaction); i++ {
		subChoreography.Transaction[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.UserTask); i++ {
		subChoreography.UserTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.Association); i++ {
		subChoreography.Association[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.Group); i++ {
		subChoreography.Group[i].ParseTree(definitions)
	}

	for i := 0; i < len(subChoreography.TextAnnotation); i++ {
		subChoreography.TextAnnotation[i].ParseTree(definitions)
	}

}

// GetAdHocSubProcess by ID
func (subChoreography *SubChoreography) GetAdHocSubProcess(ID string) *AdHocSubProcess {

	for i := 0; i < len(subChoreography.AdHocSubProcess); i++ {
		if subChoreography.AdHocSubProcess[i].ID == ID {
			return &subChoreography.AdHocSubProcess[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		AdHocSubProcess := subChoreography.SubProcess[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	return nil
}

// GetExpression by ID
func (subChoreography *SubChoreography) GetExpression(ID string) *Expression {

	for i := 0; i < len(subChoreography.AdHocSubProcess); i++ {
		Expression := subChoreography.AdHocSubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subChoreography.ComplexGateway); i++ {
		Expression := subChoreography.ComplexGateway[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subChoreography.SequenceFlow); i++ {
		Expression := subChoreography.SequenceFlow[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Expression := subChoreography.SubProcess[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}

// GetBoundaryEvent by ID
func (subChoreography *SubChoreography) GetBoundaryEvent(ID string) *BoundaryEvent {

	for i := 0; i < len(subChoreography.BoundaryEvent); i++ {
		if subChoreography.BoundaryEvent[i].ID == ID {
			return &subChoreography.BoundaryEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		BoundaryEvent := subChoreography.SubProcess[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	return nil
}

// GetBusinessRuleTask by ID
func (subChoreography *SubChoreography) GetBusinessRuleTask(ID string) *BusinessRuleTask {

	for i := 0; i < len(subChoreography.BusinessRuleTask); i++ {
		if subChoreography.BusinessRuleTask[i].ID == ID {
			return &subChoreography.BusinessRuleTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		BusinessRuleTask := subChoreography.SubProcess[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	return nil
}

// GetCallActivity by ID
func (subChoreography *SubChoreography) GetCallActivity(ID string) *CallActivity {

	for i := 0; i < len(subChoreography.CallActivity); i++ {
		if subChoreography.CallActivity[i].ID == ID {
			return &subChoreography.CallActivity[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		CallActivity := subChoreography.SubProcess[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	return nil
}

// GetCallChoreography by ID
func (subChoreography *SubChoreography) GetCallChoreography(ID string) *CallChoreography {

	for i := 0; i < len(subChoreography.CallChoreography); i++ {
		if subChoreography.CallChoreography[i].ID == ID {
			return &subChoreography.CallChoreography[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		CallChoreography := subChoreography.SubProcess[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (subChoreography *SubChoreography) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(subChoreography.CallChoreography); i++ {
		ParticipantAssociation := subChoreography.CallChoreography[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ParticipantAssociation := subChoreography.SubProcess[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetChoreographyTask by ID
func (subChoreography *SubChoreography) GetChoreographyTask(ID string) *ChoreographyTask {

	for i := 0; i < len(subChoreography.ChoreographyTask); i++ {
		if subChoreography.ChoreographyTask[i].ID == ID {
			return &subChoreography.ChoreographyTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ChoreographyTask := subChoreography.SubProcess[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	return nil
}

// GetComplexGateway by ID
func (subChoreography *SubChoreography) GetComplexGateway(ID string) *ComplexGateway {

	for i := 0; i < len(subChoreography.ComplexGateway); i++ {
		if subChoreography.ComplexGateway[i].ID == ID {
			return &subChoreography.ComplexGateway[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ComplexGateway := subChoreography.SubProcess[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	return nil
}

// GetDataObject by ID
func (subChoreography *SubChoreography) GetDataObject(ID string) *DataObject {

	for i := 0; i < len(subChoreography.DataObject); i++ {
		if subChoreography.DataObject[i].ID == ID {
			return &subChoreography.DataObject[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		DataObject := subChoreography.SubProcess[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	return nil
}

// GetDataObjectReference by ID
func (subChoreography *SubChoreography) GetDataObjectReference(ID string) *DataObjectReference {

	for i := 0; i < len(subChoreography.DataObjectReference); i++ {
		if subChoreography.DataObjectReference[i].ID == ID {
			return &subChoreography.DataObjectReference[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		DataObjectReference := subChoreography.SubProcess[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	return nil
}

// GetDataStoreReference by ID
func (subChoreography *SubChoreography) GetDataStoreReference(ID string) *DataStoreReference {

	for i := 0; i < len(subChoreography.DataStoreReference); i++ {
		if subChoreography.DataStoreReference[i].ID == ID {
			return &subChoreography.DataStoreReference[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		DataStoreReference := subChoreography.SubProcess[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	return nil
}

// GetEndEvent by ID
func (subChoreography *SubChoreography) GetEndEvent(ID string) *EndEvent {

	for i := 0; i < len(subChoreography.EndEvent); i++ {
		if subChoreography.EndEvent[i].ID == ID {
			return &subChoreography.EndEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		EndEvent := subChoreography.SubProcess[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	return nil
}

// GetEventBasedGateway by ID
func (subChoreography *SubChoreography) GetEventBasedGateway(ID string) *EventBasedGateway {

	for i := 0; i < len(subChoreography.EventBasedGateway); i++ {
		if subChoreography.EventBasedGateway[i].ID == ID {
			return &subChoreography.EventBasedGateway[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		EventBasedGateway := subChoreography.SubProcess[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	return nil
}

// GetExclusiveGateway by ID
func (subChoreography *SubChoreography) GetExclusiveGateway(ID string) *ExclusiveGateway {

	for i := 0; i < len(subChoreography.ExclusiveGateway); i++ {
		if subChoreography.ExclusiveGateway[i].ID == ID {
			return &subChoreography.ExclusiveGateway[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ExclusiveGateway := subChoreography.SubProcess[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (subChoreography *SubChoreography) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(subChoreography.ImplicitThrowEvent); i++ {
		if subChoreography.ImplicitThrowEvent[i].ID == ID {
			return &subChoreography.ImplicitThrowEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ImplicitThrowEvent := subChoreography.SubProcess[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}

// GetInclusiveGateway by ID
func (subChoreography *SubChoreography) GetInclusiveGateway(ID string) *InclusiveGateway {

	for i := 0; i < len(subChoreography.InclusiveGateway); i++ {
		if subChoreography.InclusiveGateway[i].ID == ID {
			return &subChoreography.InclusiveGateway[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		InclusiveGateway := subChoreography.SubProcess[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	return nil
}

// GetIntermediateCatchEvent by ID
func (subChoreography *SubChoreography) GetIntermediateCatchEvent(ID string) *IntermediateCatchEvent {

	for i := 0; i < len(subChoreography.IntermediateCatchEvent); i++ {
		if subChoreography.IntermediateCatchEvent[i].ID == ID {
			return &subChoreography.IntermediateCatchEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		IntermediateCatchEvent := subChoreography.SubProcess[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	return nil
}

// GetIntermediateThrowEvent by ID
func (subChoreography *SubChoreography) GetIntermediateThrowEvent(ID string) *IntermediateThrowEvent {

	for i := 0; i < len(subChoreography.IntermediateThrowEvent); i++ {
		if subChoreography.IntermediateThrowEvent[i].ID == ID {
			return &subChoreography.IntermediateThrowEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		IntermediateThrowEvent := subChoreography.SubProcess[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	return nil
}

// GetManualTask by ID
func (subChoreography *SubChoreography) GetManualTask(ID string) *ManualTask {

	for i := 0; i < len(subChoreography.ManualTask); i++ {
		if subChoreography.ManualTask[i].ID == ID {
			return &subChoreography.ManualTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ManualTask := subChoreography.SubProcess[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	return nil
}

// GetParallelGateway by ID
func (subChoreography *SubChoreography) GetParallelGateway(ID string) *ParallelGateway {

	for i := 0; i < len(subChoreography.ParallelGateway); i++ {
		if subChoreography.ParallelGateway[i].ID == ID {
			return &subChoreography.ParallelGateway[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ParallelGateway := subChoreography.SubProcess[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	return nil
}

// GetReceiveTask by ID
func (subChoreography *SubChoreography) GetReceiveTask(ID string) *ReceiveTask {

	for i := 0; i < len(subChoreography.ReceiveTask); i++ {
		if subChoreography.ReceiveTask[i].ID == ID {
			return &subChoreography.ReceiveTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ReceiveTask := subChoreography.SubProcess[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	return nil
}

// GetScriptTask by ID
func (subChoreography *SubChoreography) GetScriptTask(ID string) *ScriptTask {

	for i := 0; i < len(subChoreography.ScriptTask); i++ {
		if subChoreography.ScriptTask[i].ID == ID {
			return &subChoreography.ScriptTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ScriptTask := subChoreography.SubProcess[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	return nil
}

// GetSendTask by ID
func (subChoreography *SubChoreography) GetSendTask(ID string) *SendTask {

	for i := 0; i < len(subChoreography.SendTask); i++ {
		if subChoreography.SendTask[i].ID == ID {
			return &subChoreography.SendTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		SendTask := subChoreography.SubProcess[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	return nil
}

// GetSequenceFlow by ID
func (subChoreography *SubChoreography) GetSequenceFlow(ID string) *SequenceFlow {

	for i := 0; i < len(subChoreography.SequenceFlow); i++ {
		if subChoreography.SequenceFlow[i].ID == ID {
			return &subChoreography.SequenceFlow[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		SequenceFlow := subChoreography.SubProcess[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	return nil
}

// GetServiceTask by ID
func (subChoreography *SubChoreography) GetServiceTask(ID string) *ServiceTask {

	for i := 0; i < len(subChoreography.ServiceTask); i++ {
		if subChoreography.ServiceTask[i].ID == ID {
			return &subChoreography.ServiceTask[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		ServiceTask := subChoreography.SubProcess[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	return nil
}

// GetStartEvent by ID
func (subChoreography *SubChoreography) GetStartEvent(ID string) *StartEvent {

	for i := 0; i < len(subChoreography.StartEvent); i++ {
		if subChoreography.StartEvent[i].ID == ID {
			return &subChoreography.StartEvent[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		StartEvent := subChoreography.SubProcess[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	return nil
}

// GetSubChoreography by ID
func (subChoreography *SubChoreography) GetSubChoreography(ID string) *SubChoreography {

	for i := 0; i < len(subChoreography.SubChoreography); i++ {
		if subChoreography.SubChoreography[i].ID == ID {
			return &subChoreography.SubChoreography[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		SubChoreography := subChoreography.SubProcess[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	return nil
}

// GetSubProcess by ID
func (subChoreography *SubChoreography) GetSubProcess(ID string) *SubProcess {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		if subChoreography.SubProcess[i].ID == ID {
			return &subChoreography.SubProcess[i]

		}
	}

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		SubProcess := subChoreography.SubProcess[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	return nil
}

// GetLaneSet by ID
func (subChoreography *SubChoreography) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		LaneSet := subChoreography.SubProcess[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}

// GetLane by ID
func (subChoreography *SubChoreography) GetLane(ID string) *Lane {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Lane := subChoreography.SubProcess[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}

// GetBaseElement by ID
func (subChoreography *SubChoreography) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		BaseElement := subChoreography.SubProcess[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (subChoreography *SubChoreography) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Documentation := subChoreography.SubProcess[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetTask by ID
func (subChoreography *SubChoreography) GetTask(ID string) *Task {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Task := subChoreography.SubProcess[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	for i := 0; i < len(subChoreography.Task); i++ {
		if subChoreography.Task[i].ID == ID {
			return &subChoreography.Task[i]

		}
	}

	return nil
}

// GetTransaction by ID
func (subChoreography *SubChoreography) GetTransaction(ID string) *Transaction {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Transaction := subChoreography.SubProcess[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	for i := 0; i < len(subChoreography.Transaction); i++ {
		if subChoreography.Transaction[i].ID == ID {
			return &subChoreography.Transaction[i]

		}
	}

	return nil
}

// GetUserTask by ID
func (subChoreography *SubChoreography) GetUserTask(ID string) *UserTask {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		UserTask := subChoreography.SubProcess[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	for i := 0; i < len(subChoreography.UserTask); i++ {
		if subChoreography.UserTask[i].ID == ID {
			return &subChoreography.UserTask[i]

		}
	}

	return nil
}

// GetRendering by ID
func (subChoreography *SubChoreography) GetRendering(ID string) *Rendering {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Rendering := subChoreography.SubProcess[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(subChoreography.UserTask); i++ {
		Rendering := subChoreography.UserTask[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	return nil
}

// GetAssociation by ID
func (subChoreography *SubChoreography) GetAssociation(ID string) *Association {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Association := subChoreography.SubProcess[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(subChoreography.Association); i++ {
		if subChoreography.Association[i].ID == ID {
			return &subChoreography.Association[i]

		}
	}

	return nil
}

// GetGroup by ID
func (subChoreography *SubChoreography) GetGroup(ID string) *Group {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		Group := subChoreography.SubProcess[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(subChoreography.Group); i++ {
		if subChoreography.Group[i].ID == ID {
			return &subChoreography.Group[i]

		}
	}

	return nil
}

// GetTextAnnotation by ID
func (subChoreography *SubChoreography) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(subChoreography.SubProcess); i++ {
		TextAnnotation := subChoreography.SubProcess[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(subChoreography.TextAnnotation); i++ {
		if subChoreography.TextAnnotation[i].ID == ID {
			return &subChoreography.TextAnnotation[i]

		}
	}

	return nil
}

