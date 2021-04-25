package bpmn

import (
	"encoding/xml"
)

// Lane from BPMN
type Lane struct {
	BaseElement
	XMLName             xml.Name
	PartitionElement    []BaseElement `xml:"partitionElement"`
	FlowNodeRef         []string      `xml:"flowNodeRef"`
	ChildLaneSet        []LaneSet     `xml:"childLaneSet"`
	Name                string        `xml:"name,attr"`
	PartitionElementRef string        `xml:"partitionElementRef,attr"`
}

// ParseTree of components of Lane.
func (lane *Lane) ParseTree(definitions *Definitions) {
	lane.BaseElement.ParseTree(definitions)

	for i := 0; i < len(lane.PartitionElement); i++ {
		lane.PartitionElement[i].ParseTree(definitions)
	}

	for i := 0; i < len(lane.ChildLaneSet); i++ {
		lane.ChildLaneSet[i].ParseTree(definitions)
	}

}

// GetBaseElement by ID
func (lane *Lane) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(lane.PartitionElement); i++ {
		if lane.PartitionElement[i].ID == ID {
			return &lane.PartitionElement[i]

		}
	}

	return nil
}

// GetDocumentation by ID
func (lane *Lane) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(lane.PartitionElement); i++ {
		Documentation := lane.PartitionElement[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetLaneSet by ID
func (lane *Lane) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(lane.ChildLaneSet); i++ {
		if lane.ChildLaneSet[i].ID == ID {
			return &lane.ChildLaneSet[i]

		}
	}

	return nil
}

// GetLane by ID
func (lane *Lane) GetLane(ID string) *Lane {

	for i := 0; i < len(lane.ChildLaneSet); i++ {
		Lane := lane.ChildLaneSet[i].GetLane(ID)
		if Lane != nil {
			return Lane
		}
	}

	return nil
}
