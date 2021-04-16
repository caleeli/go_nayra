package bpmn

import (
	"encoding/xml"
)

// Definitions from BPMN
type Definitions struct {
	DefinitionsTrait
	XMLName                xml.Name
	Import                 []Import                 `xml:"import"`
	Extension              []Extension              `xml:"extension"`
	Category               []Category               `xml:"category"`
	Collaboration          []Collaboration          `xml:"collaboration"`
	CorrelationProperty    []CorrelationProperty    `xml:"correlationProperty"`
	DataStore              []DataStore              `xml:"dataStore"`
	EndPoint               []EndPoint               `xml:"endPoint"`
	Error                  []Error                  `xml:"error"`
	Escalation             []Escalation             `xml:"escalation"`
	GlobalBusinessRuleTask []GlobalBusinessRuleTask `xml:"globalBusinessRuleTask"`
	GlobalManualTask       []GlobalManualTask       `xml:"globalManualTask"`
	GlobalScriptTask       []GlobalScriptTask       `xml:"globalScriptTask"`
	GlobalTask             []GlobalTask             `xml:"globalTask"`
	GlobalUserTask         []GlobalUserTask         `xml:"globalUserTask"`
	Interface              []Interface              `xml:"interface"`
	ItemDefinition         []ItemDefinition         `xml:"itemDefinition"`
	Message                []Message                `xml:"message"`
	PartnerEntity          []PartnerEntity          `xml:"partnerEntity"`
	PartnerRole            []PartnerRole            `xml:"partnerRole"`
	Process                []Process                `xml:"process"`
	Resource               []Resource               `xml:"resource"`
	Signal                 []Signal                 `xml:"signal"`
	Relationship           []Relationship           `xml:"relationship"`
	ID                     string                   `xml:"id,attr"`
	Name                   string                   `xml:"name,attr"`
	TargetNamespace        string                   `xml:"targetNamespace,attr"`
	ExpressionLanguage     string                   `xml:"expressionLanguage,attr"`
	TypeLanguage           string                   `xml:"typeLanguage,attr"`
	Exporter               string                   `xml:"exporter,attr"`
	ExporterVersion        string                   `xml:"exporterVersion,attr"`

}

// ParseTree of components of Definitions.
func (definitions *Definitions) ParseTree () {

	for i := 0; i < len(definitions.Import); i++ {
		definitions.Import[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Extension); i++ {
		definitions.Extension[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Category); i++ {
		definitions.Category[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Collaboration); i++ {
		definitions.Collaboration[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.CorrelationProperty); i++ {
		definitions.CorrelationProperty[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.DataStore); i++ {
		definitions.DataStore[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.EndPoint); i++ {
		definitions.EndPoint[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Error); i++ {
		definitions.Error[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Escalation); i++ {
		definitions.Escalation[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.GlobalBusinessRuleTask); i++ {
		definitions.GlobalBusinessRuleTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.GlobalManualTask); i++ {
		definitions.GlobalManualTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.GlobalScriptTask); i++ {
		definitions.GlobalScriptTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.GlobalTask); i++ {
		definitions.GlobalTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.GlobalUserTask); i++ {
		definitions.GlobalUserTask[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Interface); i++ {
		definitions.Interface[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.ItemDefinition); i++ {
		definitions.ItemDefinition[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Message); i++ {
		definitions.Message[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.PartnerEntity); i++ {
		definitions.PartnerEntity[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.PartnerRole); i++ {
		definitions.PartnerRole[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Process); i++ {
		definitions.Process[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Resource); i++ {
		definitions.Resource[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Signal); i++ {
		definitions.Signal[i].ParseTree(definitions)
	}

	for i := 0; i < len(definitions.Relationship); i++ {
		definitions.Relationship[i].ParseTree(definitions)
	}

}

// GetCategory by ID
func (definitions *Definitions) GetCategory(ID string) *Category {

	for i := 0; i < len(definitions.Category); i++ {
		if definitions.Category[i].ID == ID {
			return &definitions.Category[i]

		}
	}

	return nil
}

// GetCategoryValue by ID
func (definitions *Definitions) GetCategoryValue(ID string) *CategoryValue {

	for i := 0; i < len(definitions.Category); i++ {
		CategoryValue := definitions.Category[i].GetCategoryValue(ID)
		if CategoryValue != nil {
			return CategoryValue
		}
	}

	return nil
}

// GetCollaboration by ID
func (definitions *Definitions) GetCollaboration(ID string) *Collaboration {

	for i := 0; i < len(definitions.Collaboration); i++ {
		if definitions.Collaboration[i].ID == ID {
			return &definitions.Collaboration[i]

		}
	}

	return nil
}

// GetParticipant by ID
func (definitions *Definitions) GetParticipant(ID string) *Participant {

	for i := 0; i < len(definitions.Collaboration); i++ {
		Participant := definitions.Collaboration[i].GetParticipant(ID)
		if Participant != nil {
			return Participant
		}
	}

	return nil
}

// GetMessageFlow by ID
func (definitions *Definitions) GetMessageFlow(ID string) *MessageFlow {

	for i := 0; i < len(definitions.Collaboration); i++ {
		MessageFlow := definitions.Collaboration[i].GetMessageFlow(ID)
		if MessageFlow != nil {
			return MessageFlow
		}
	}

	return nil
}

// GetAssociation by ID
func (definitions *Definitions) GetAssociation(ID string) *Association {

	for i := 0; i < len(definitions.Collaboration); i++ {
		Association := definitions.Collaboration[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		Association := definitions.Process[i].GetAssociation(ID)
		if Association != nil {
			return Association
		}
	}

	return nil
}

// GetGroup by ID
func (definitions *Definitions) GetGroup(ID string) *Group {

	for i := 0; i < len(definitions.Collaboration); i++ {
		Group := definitions.Collaboration[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		Group := definitions.Process[i].GetGroup(ID)
		if Group != nil {
			return Group
		}
	}

	return nil
}

// GetTextAnnotation by ID
func (definitions *Definitions) GetTextAnnotation(ID string) *TextAnnotation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		TextAnnotation := definitions.Collaboration[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		TextAnnotation := definitions.Process[i].GetTextAnnotation(ID)
		if TextAnnotation != nil {
			return TextAnnotation
		}
	}

	return nil
}

// GetCallConversation by ID
func (definitions *Definitions) GetCallConversation(ID string) *CallConversation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		CallConversation := definitions.Collaboration[i].GetCallConversation(ID)
		if CallConversation != nil {
			return CallConversation
		}
	}

	return nil
}

// GetParticipantAssociation by ID
func (definitions *Definitions) GetParticipantAssociation(ID string) *ParticipantAssociation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		ParticipantAssociation := definitions.Collaboration[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		ParticipantAssociation := definitions.Process[i].GetParticipantAssociation(ID)
		if ParticipantAssociation != nil {
			return ParticipantAssociation
		}
	}

	return nil
}

// GetConversation by ID
func (definitions *Definitions) GetConversation(ID string) *Conversation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		Conversation := definitions.Collaboration[i].GetConversation(ID)
		if Conversation != nil {
			return Conversation
		}
	}

	return nil
}

// GetSubConversation by ID
func (definitions *Definitions) GetSubConversation(ID string) *SubConversation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		SubConversation := definitions.Collaboration[i].GetSubConversation(ID)
		if SubConversation != nil {
			return SubConversation
		}
	}

	return nil
}

// GetConversationAssociation by ID
func (definitions *Definitions) GetConversationAssociation(ID string) *ConversationAssociation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		ConversationAssociation := definitions.Collaboration[i].GetConversationAssociation(ID)
		if ConversationAssociation != nil {
			return ConversationAssociation
		}
	}

	return nil
}

// GetMessageFlowAssociation by ID
func (definitions *Definitions) GetMessageFlowAssociation(ID string) *MessageFlowAssociation {

	for i := 0; i < len(definitions.Collaboration); i++ {
		MessageFlowAssociation := definitions.Collaboration[i].GetMessageFlowAssociation(ID)
		if MessageFlowAssociation != nil {
			return MessageFlowAssociation
		}
	}

	return nil
}

// GetCorrelationKey by ID
func (definitions *Definitions) GetCorrelationKey(ID string) *CorrelationKey {

	for i := 0; i < len(definitions.Collaboration); i++ {
		CorrelationKey := definitions.Collaboration[i].GetCorrelationKey(ID)
		if CorrelationKey != nil {
			return CorrelationKey
		}
	}

	return nil
}

// GetConversationLink by ID
func (definitions *Definitions) GetConversationLink(ID string) *ConversationLink {

	for i := 0; i < len(definitions.Collaboration); i++ {
		ConversationLink := definitions.Collaboration[i].GetConversationLink(ID)
		if ConversationLink != nil {
			return ConversationLink
		}
	}

	return nil
}

// GetCorrelationProperty by ID
func (definitions *Definitions) GetCorrelationProperty(ID string) *CorrelationProperty {

	for i := 0; i < len(definitions.CorrelationProperty); i++ {
		if definitions.CorrelationProperty[i].ID == ID {
			return &definitions.CorrelationProperty[i]

		}
	}

	return nil
}

// GetCorrelationPropertyRetrievalExpression by ID
func (definitions *Definitions) GetCorrelationPropertyRetrievalExpression(ID string) *CorrelationPropertyRetrievalExpression {

	for i := 0; i < len(definitions.CorrelationProperty); i++ {
		CorrelationPropertyRetrievalExpression := definitions.CorrelationProperty[i].GetCorrelationPropertyRetrievalExpression(ID)
		if CorrelationPropertyRetrievalExpression != nil {
			return CorrelationPropertyRetrievalExpression
		}
	}

	return nil
}

// GetFormalExpression by ID
func (definitions *Definitions) GetFormalExpression(ID string) *FormalExpression {

	for i := 0; i < len(definitions.CorrelationProperty); i++ {
		FormalExpression := definitions.CorrelationProperty[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		FormalExpression := definitions.Process[i].GetFormalExpression(ID)
		if FormalExpression != nil {
			return FormalExpression
		}
	}

	return nil
}

// GetDataStore by ID
func (definitions *Definitions) GetDataStore(ID string) *DataStore {

	for i := 0; i < len(definitions.DataStore); i++ {
		if definitions.DataStore[i].ID == ID {
			return &definitions.DataStore[i]

		}
	}

	return nil
}

// GetEndPoint by ID
func (definitions *Definitions) GetEndPoint(ID string) *EndPoint {

	for i := 0; i < len(definitions.EndPoint); i++ {
		if definitions.EndPoint[i].ID == ID {
			return &definitions.EndPoint[i]

		}
	}

	return nil
}

// GetError by ID
func (definitions *Definitions) GetError(ID string) *Error {

	for i := 0; i < len(definitions.Error); i++ {
		if definitions.Error[i].ID == ID {
			return &definitions.Error[i]

		}
	}

	return nil
}

// GetEscalation by ID
func (definitions *Definitions) GetEscalation(ID string) *Escalation {

	for i := 0; i < len(definitions.Escalation); i++ {
		if definitions.Escalation[i].ID == ID {
			return &definitions.Escalation[i]

		}
	}

	return nil
}

// GetGlobalBusinessRuleTask by ID
func (definitions *Definitions) GetGlobalBusinessRuleTask(ID string) *GlobalBusinessRuleTask {

	for i := 0; i < len(definitions.GlobalBusinessRuleTask); i++ {
		if definitions.GlobalBusinessRuleTask[i].ID == ID {
			return &definitions.GlobalBusinessRuleTask[i]

		}
	}

	return nil
}

// GetGlobalManualTask by ID
func (definitions *Definitions) GetGlobalManualTask(ID string) *GlobalManualTask {

	for i := 0; i < len(definitions.GlobalManualTask); i++ {
		if definitions.GlobalManualTask[i].ID == ID {
			return &definitions.GlobalManualTask[i]

		}
	}

	return nil
}

// GetGlobalScriptTask by ID
func (definitions *Definitions) GetGlobalScriptTask(ID string) *GlobalScriptTask {

	for i := 0; i < len(definitions.GlobalScriptTask); i++ {
		if definitions.GlobalScriptTask[i].ID == ID {
			return &definitions.GlobalScriptTask[i]

		}
	}

	return nil
}

// GetGlobalTask by ID
func (definitions *Definitions) GetGlobalTask(ID string) *GlobalTask {

	for i := 0; i < len(definitions.GlobalTask); i++ {
		if definitions.GlobalTask[i].ID == ID {
			return &definitions.GlobalTask[i]

		}
	}

	return nil
}

// GetPerformer by ID
func (definitions *Definitions) GetPerformer(ID string) *Performer {

	for i := 0; i < len(definitions.GlobalTask); i++ {
		Performer := definitions.GlobalTask[i].GetPerformer(ID)
		if Performer != nil {
			return Performer
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		Performer := definitions.Process[i].GetPerformer(ID)
		if Performer != nil {
			return Performer
		}
	}

	return nil
}

// GetGlobalUserTask by ID
func (definitions *Definitions) GetGlobalUserTask(ID string) *GlobalUserTask {

	for i := 0; i < len(definitions.GlobalUserTask); i++ {
		if definitions.GlobalUserTask[i].ID == ID {
			return &definitions.GlobalUserTask[i]

		}
	}

	return nil
}

// GetRendering by ID
func (definitions *Definitions) GetRendering(ID string) *Rendering {

	for i := 0; i < len(definitions.GlobalUserTask); i++ {
		Rendering := definitions.GlobalUserTask[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	for i := 0; i < len(definitions.Process); i++ {
		Rendering := definitions.Process[i].GetRendering(ID)
		if Rendering != nil {
			return Rendering
		}
	}

	return nil
}

// GetInterface by ID
func (definitions *Definitions) GetInterface(ID string) *Interface {

	for i := 0; i < len(definitions.Interface); i++ {
		if definitions.Interface[i].ID == ID {
			return &definitions.Interface[i]

		}
	}

	return nil
}

// GetOperation by ID
func (definitions *Definitions) GetOperation(ID string) *Operation {

	for i := 0; i < len(definitions.Interface); i++ {
		Operation := definitions.Interface[i].GetOperation(ID)
		if Operation != nil {
			return Operation
		}
	}

	return nil
}

// GetItemDefinition by ID
func (definitions *Definitions) GetItemDefinition(ID string) *ItemDefinition {

	for i := 0; i < len(definitions.ItemDefinition); i++ {
		if definitions.ItemDefinition[i].ID == ID {
			return &definitions.ItemDefinition[i]

		}
	}

	return nil
}

// GetMessage by ID
func (definitions *Definitions) GetMessage(ID string) *Message {

	for i := 0; i < len(definitions.Message); i++ {
		if definitions.Message[i].ID == ID {
			return &definitions.Message[i]

		}
	}

	return nil
}

// GetPartnerEntity by ID
func (definitions *Definitions) GetPartnerEntity(ID string) *PartnerEntity {

	for i := 0; i < len(definitions.PartnerEntity); i++ {
		if definitions.PartnerEntity[i].ID == ID {
			return &definitions.PartnerEntity[i]

		}
	}

	return nil
}

// GetPartnerRole by ID
func (definitions *Definitions) GetPartnerRole(ID string) *PartnerRole {

	for i := 0; i < len(definitions.PartnerRole); i++ {
		if definitions.PartnerRole[i].ID == ID {
			return &definitions.PartnerRole[i]

		}
	}

	return nil
}

// GetProcess by ID
func (definitions *Definitions) GetProcess(ID string) *Process {

	for i := 0; i < len(definitions.Process); i++ {
		if definitions.Process[i].ID == ID {
			return &definitions.Process[i]

		}
	}

	return nil
}

// GetProperty by ID
func (definitions *Definitions) GetProperty(ID string) *Property {

	for i := 0; i < len(definitions.Process); i++ {
		Property := definitions.Process[i].GetProperty(ID)
		if Property != nil {
			return Property
		}
	}

	return nil
}

// GetLaneSet by ID
func (definitions *Definitions) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(definitions.Process); i++ {
		LaneSet := definitions.Process[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}

// GetLane by ID
func (definitions *Definitions) GetLane(ID string) *Lane {

	for i := 0; i < len(definitions.Process); i++ {
		Lane := definitions.Process[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}

// GetBaseElement by ID
func (definitions *Definitions) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(definitions.Process); i++ {
		BaseElement := definitions.Process[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (definitions *Definitions) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(definitions.Process); i++ {
		Documentation := definitions.Process[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetAdHocSubProcess by ID
func (definitions *Definitions) GetAdHocSubProcess(ID string) *AdHocSubProcess {

	for i := 0; i < len(definitions.Process); i++ {
		AdHocSubProcess := definitions.Process[i].GetAdHocSubProcess(ID)
		if AdHocSubProcess != nil {
			return AdHocSubProcess
		}
	}

	return nil
}

// GetExpression by ID
func (definitions *Definitions) GetExpression(ID string) *Expression {

	for i := 0; i < len(definitions.Process); i++ {
		Expression := definitions.Process[i].GetExpression(ID)
		if Expression != nil {
			return Expression
		}
	}

	return nil
}

// GetBoundaryEvent by ID
func (definitions *Definitions) GetBoundaryEvent(ID string) *BoundaryEvent {

	for i := 0; i < len(definitions.Process); i++ {
		BoundaryEvent := definitions.Process[i].GetBoundaryEvent(ID)
		if BoundaryEvent != nil {
			return BoundaryEvent
		}
	}

	return nil
}

// GetBusinessRuleTask by ID
func (definitions *Definitions) GetBusinessRuleTask(ID string) *BusinessRuleTask {

	for i := 0; i < len(definitions.Process); i++ {
		BusinessRuleTask := definitions.Process[i].GetBusinessRuleTask(ID)
		if BusinessRuleTask != nil {
			return BusinessRuleTask
		}
	}

	return nil
}

// GetCallActivity by ID
func (definitions *Definitions) GetCallActivity(ID string) *CallActivity {

	for i := 0; i < len(definitions.Process); i++ {
		CallActivity := definitions.Process[i].GetCallActivity(ID)
		if CallActivity != nil {
			return CallActivity
		}
	}

	return nil
}

// GetCallChoreography by ID
func (definitions *Definitions) GetCallChoreography(ID string) *CallChoreography {

	for i := 0; i < len(definitions.Process); i++ {
		CallChoreography := definitions.Process[i].GetCallChoreography(ID)
		if CallChoreography != nil {
			return CallChoreography
		}
	}

	return nil
}

// GetChoreographyTask by ID
func (definitions *Definitions) GetChoreographyTask(ID string) *ChoreographyTask {

	for i := 0; i < len(definitions.Process); i++ {
		ChoreographyTask := definitions.Process[i].GetChoreographyTask(ID)
		if ChoreographyTask != nil {
			return ChoreographyTask
		}
	}

	return nil
}

// GetComplexGateway by ID
func (definitions *Definitions) GetComplexGateway(ID string) *ComplexGateway {

	for i := 0; i < len(definitions.Process); i++ {
		ComplexGateway := definitions.Process[i].GetComplexGateway(ID)
		if ComplexGateway != nil {
			return ComplexGateway
		}
	}

	return nil
}

// GetDataObject by ID
func (definitions *Definitions) GetDataObject(ID string) *DataObject {

	for i := 0; i < len(definitions.Process); i++ {
		DataObject := definitions.Process[i].GetDataObject(ID)
		if DataObject != nil {
			return DataObject
		}
	}

	return nil
}

// GetDataObjectReference by ID
func (definitions *Definitions) GetDataObjectReference(ID string) *DataObjectReference {

	for i := 0; i < len(definitions.Process); i++ {
		DataObjectReference := definitions.Process[i].GetDataObjectReference(ID)
		if DataObjectReference != nil {
			return DataObjectReference
		}
	}

	return nil
}

// GetDataStoreReference by ID
func (definitions *Definitions) GetDataStoreReference(ID string) *DataStoreReference {

	for i := 0; i < len(definitions.Process); i++ {
		DataStoreReference := definitions.Process[i].GetDataStoreReference(ID)
		if DataStoreReference != nil {
			return DataStoreReference
		}
	}

	return nil
}

// GetEndEvent by ID
func (definitions *Definitions) GetEndEvent(ID string) *EndEvent {

	for i := 0; i < len(definitions.Process); i++ {
		EndEvent := definitions.Process[i].GetEndEvent(ID)
		if EndEvent != nil {
			return EndEvent
		}
	}

	return nil
}

// GetEventBasedGateway by ID
func (definitions *Definitions) GetEventBasedGateway(ID string) *EventBasedGateway {

	for i := 0; i < len(definitions.Process); i++ {
		EventBasedGateway := definitions.Process[i].GetEventBasedGateway(ID)
		if EventBasedGateway != nil {
			return EventBasedGateway
		}
	}

	return nil
}

// GetExclusiveGateway by ID
func (definitions *Definitions) GetExclusiveGateway(ID string) *ExclusiveGateway {

	for i := 0; i < len(definitions.Process); i++ {
		ExclusiveGateway := definitions.Process[i].GetExclusiveGateway(ID)
		if ExclusiveGateway != nil {
			return ExclusiveGateway
		}
	}

	return nil
}

// GetImplicitThrowEvent by ID
func (definitions *Definitions) GetImplicitThrowEvent(ID string) *ImplicitThrowEvent {

	for i := 0; i < len(definitions.Process); i++ {
		ImplicitThrowEvent := definitions.Process[i].GetImplicitThrowEvent(ID)
		if ImplicitThrowEvent != nil {
			return ImplicitThrowEvent
		}
	}

	return nil
}

// GetInclusiveGateway by ID
func (definitions *Definitions) GetInclusiveGateway(ID string) *InclusiveGateway {

	for i := 0; i < len(definitions.Process); i++ {
		InclusiveGateway := definitions.Process[i].GetInclusiveGateway(ID)
		if InclusiveGateway != nil {
			return InclusiveGateway
		}
	}

	return nil
}

// GetIntermediateCatchEvent by ID
func (definitions *Definitions) GetIntermediateCatchEvent(ID string) *IntermediateCatchEvent {

	for i := 0; i < len(definitions.Process); i++ {
		IntermediateCatchEvent := definitions.Process[i].GetIntermediateCatchEvent(ID)
		if IntermediateCatchEvent != nil {
			return IntermediateCatchEvent
		}
	}

	return nil
}

// GetIntermediateThrowEvent by ID
func (definitions *Definitions) GetIntermediateThrowEvent(ID string) *IntermediateThrowEvent {

	for i := 0; i < len(definitions.Process); i++ {
		IntermediateThrowEvent := definitions.Process[i].GetIntermediateThrowEvent(ID)
		if IntermediateThrowEvent != nil {
			return IntermediateThrowEvent
		}
	}

	return nil
}

// GetManualTask by ID
func (definitions *Definitions) GetManualTask(ID string) *ManualTask {

	for i := 0; i < len(definitions.Process); i++ {
		ManualTask := definitions.Process[i].GetManualTask(ID)
		if ManualTask != nil {
			return ManualTask
		}
	}

	return nil
}

// GetParallelGateway by ID
func (definitions *Definitions) GetParallelGateway(ID string) *ParallelGateway {

	for i := 0; i < len(definitions.Process); i++ {
		ParallelGateway := definitions.Process[i].GetParallelGateway(ID)
		if ParallelGateway != nil {
			return ParallelGateway
		}
	}

	return nil
}

// GetReceiveTask by ID
func (definitions *Definitions) GetReceiveTask(ID string) *ReceiveTask {

	for i := 0; i < len(definitions.Process); i++ {
		ReceiveTask := definitions.Process[i].GetReceiveTask(ID)
		if ReceiveTask != nil {
			return ReceiveTask
		}
	}

	return nil
}

// GetScriptTask by ID
func (definitions *Definitions) GetScriptTask(ID string) *ScriptTask {

	for i := 0; i < len(definitions.Process); i++ {
		ScriptTask := definitions.Process[i].GetScriptTask(ID)
		if ScriptTask != nil {
			return ScriptTask
		}
	}

	return nil
}

// GetSendTask by ID
func (definitions *Definitions) GetSendTask(ID string) *SendTask {

	for i := 0; i < len(definitions.Process); i++ {
		SendTask := definitions.Process[i].GetSendTask(ID)
		if SendTask != nil {
			return SendTask
		}
	}

	return nil
}

// GetSequenceFlow by ID
func (definitions *Definitions) GetSequenceFlow(ID string) *SequenceFlow {

	for i := 0; i < len(definitions.Process); i++ {
		SequenceFlow := definitions.Process[i].GetSequenceFlow(ID)
		if SequenceFlow != nil {
			return SequenceFlow
		}
	}

	return nil
}

// GetServiceTask by ID
func (definitions *Definitions) GetServiceTask(ID string) *ServiceTask {

	for i := 0; i < len(definitions.Process); i++ {
		ServiceTask := definitions.Process[i].GetServiceTask(ID)
		if ServiceTask != nil {
			return ServiceTask
		}
	}

	return nil
}

// GetStartEvent by ID
func (definitions *Definitions) GetStartEvent(ID string) *StartEvent {

	for i := 0; i < len(definitions.Process); i++ {
		StartEvent := definitions.Process[i].GetStartEvent(ID)
		if StartEvent != nil {
			return StartEvent
		}
	}

	return nil
}

// GetSubChoreography by ID
func (definitions *Definitions) GetSubChoreography(ID string) *SubChoreography {

	for i := 0; i < len(definitions.Process); i++ {
		SubChoreography := definitions.Process[i].GetSubChoreography(ID)
		if SubChoreography != nil {
			return SubChoreography
		}
	}

	return nil
}

// GetSubProcess by ID
func (definitions *Definitions) GetSubProcess(ID string) *SubProcess {

	for i := 0; i < len(definitions.Process); i++ {
		SubProcess := definitions.Process[i].GetSubProcess(ID)
		if SubProcess != nil {
			return SubProcess
		}
	}

	return nil
}

// GetTask by ID
func (definitions *Definitions) GetTask(ID string) *Task {

	for i := 0; i < len(definitions.Process); i++ {
		Task := definitions.Process[i].GetTask(ID)
		if Task != nil {
			return Task
		}
	}

	return nil
}

// GetTransaction by ID
func (definitions *Definitions) GetTransaction(ID string) *Transaction {

	for i := 0; i < len(definitions.Process); i++ {
		Transaction := definitions.Process[i].GetTransaction(ID)
		if Transaction != nil {
			return Transaction
		}
	}

	return nil
}

// GetUserTask by ID
func (definitions *Definitions) GetUserTask(ID string) *UserTask {

	for i := 0; i < len(definitions.Process); i++ {
		UserTask := definitions.Process[i].GetUserTask(ID)
		if UserTask != nil {
			return UserTask
		}
	}

	return nil
}

// GetCorrelationSubscription by ID
func (definitions *Definitions) GetCorrelationSubscription(ID string) *CorrelationSubscription {

	for i := 0; i < len(definitions.Process); i++ {
		CorrelationSubscription := definitions.Process[i].GetCorrelationSubscription(ID)
		if CorrelationSubscription != nil {
			return CorrelationSubscription
		}
	}

	return nil
}

// GetCorrelationPropertyBinding by ID
func (definitions *Definitions) GetCorrelationPropertyBinding(ID string) *CorrelationPropertyBinding {

	for i := 0; i < len(definitions.Process); i++ {
		CorrelationPropertyBinding := definitions.Process[i].GetCorrelationPropertyBinding(ID)
		if CorrelationPropertyBinding != nil {
			return CorrelationPropertyBinding
		}
	}

	return nil
}

// GetResource by ID
func (definitions *Definitions) GetResource(ID string) *Resource {

	for i := 0; i < len(definitions.Resource); i++ {
		if definitions.Resource[i].ID == ID {
			return &definitions.Resource[i]

		}
	}

	return nil
}

// GetResourceParameter by ID
func (definitions *Definitions) GetResourceParameter(ID string) *ResourceParameter {

	for i := 0; i < len(definitions.Resource); i++ {
		ResourceParameter := definitions.Resource[i].GetResourceParameter(ID)
		if ResourceParameter != nil {
			return ResourceParameter
		}
	}

	return nil
}

// GetSignal by ID
func (definitions *Definitions) GetSignal(ID string) *Signal {

	for i := 0; i < len(definitions.Signal); i++ {
		if definitions.Signal[i].ID == ID {
			return &definitions.Signal[i]

		}
	}

	return nil
}

// GetRelationship by ID
func (definitions *Definitions) GetRelationship(ID string) *Relationship {

	for i := 0; i < len(definitions.Relationship); i++ {
		if definitions.Relationship[i].ID == ID {
			return &definitions.Relationship[i]

		}
	}

	return nil
}

