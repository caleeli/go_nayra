package bpmn

import (
	"encoding/xml"
)

// LaneSet from BPMN
type LaneSet struct {
	BaseElement
	XMLName xml.Name
	Lane    []Lane `xml:"lane"`
	Name    string `xml:"name,attr"`
}

// ParseTree of components of LaneSet.
func (laneSet *LaneSet) ParseTree(definitions *Definitions) {
	laneSet.BaseElement.ParseTree(definitions)

	for i := 0; i < len(laneSet.Lane); i++ {
		laneSet.Lane[i].ParseTree(definitions)
	}

}

// GetLane by ID
func (laneSet *LaneSet) GetLane(ID string) *Lane {

	for i := 0; i < len(laneSet.Lane); i++ {
		if laneSet.Lane[i].ID == ID {
			return &laneSet.Lane[i]

		}
	}

	return nil
}

// GetBaseElement by ID
func (laneSet *LaneSet) GetBaseElement(ID string) *BaseElement {

	for i := 0; i < len(laneSet.Lane); i++ {
		BaseElement := laneSet.Lane[i].GetBaseElement(ID)
		if BaseElement != nil {
			return BaseElement
		}
	}

	return nil
}

// GetDocumentation by ID
func (laneSet *LaneSet) GetDocumentation(ID string) *Documentation {

	for i := 0; i < len(laneSet.Lane); i++ {
		Documentation := laneSet.Lane[i].GetDocumentation(ID)
		if Documentation != nil {
			return Documentation
		}
	}

	return nil
}

// GetLaneSet by ID
func (laneSet *LaneSet) GetLaneSet(ID string) *LaneSet {

	for i := 0; i < len(laneSet.Lane); i++ {
		LaneSet := laneSet.Lane[i].GetLaneSet(ID)
		if LaneSet != nil {
			return LaneSet
		}
	}

	return nil
}
