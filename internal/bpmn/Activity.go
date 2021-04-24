package bpmn

import (
	"encoding/xml"
)

// Activity from BPMN
type Activity struct {
	ActivityTrait `bson:"-"`
	FlowNode
	XMLName               xml.Name
	IoSpecification       InputOutputSpecification `xml:"ioSpecification"`
	Property              []Property               `xml:"property"`
	DataInputAssociation  []DataInputAssociation   `xml:"dataInputAssociation"`
	DataOutputAssociation []DataOutputAssociation  `xml:"dataOutputAssociation"`
	Performer             []Performer              `xml:"performer"`
	IsForCompensation     bool                     `xml:"isForCompensation,attr"`
	StartQuantity         int                      `xml:"startQuantity,attr"`
	CompletionQuantity    int                      `xml:"completionQuantity,attr"`
	Default               string                   `xml:"default,attr"`

}

// ParseTree of components of Activity.
func (activity *Activity) ParseTree (definitions *Definitions) {
	activity.FlowNode.ParseTree(definitions)

	activity.Init(definitions)

	activity.IoSpecification.ParseTree(definitions) // InputOutputSpecification.

	for i := 0; i < len(activity.Property); i++ {
		activity.Property[i].ParseTree(definitions)
	}

	for i := 0; i < len(activity.DataInputAssociation); i++ {
		activity.DataInputAssociation[i].ParseTree(definitions)
	}

	for i := 0; i < len(activity.DataOutputAssociation); i++ {
		activity.DataOutputAssociation[i].ParseTree(definitions)
	}

	for i := 0; i < len(activity.Performer); i++ {
		activity.Performer[i].ParseTree(definitions)
	}

}

// GetProperty by ID
func (activity *Activity) GetProperty(ID string) *Property {

	for i := 0; i < len(activity.Property); i++ {
		if activity.Property[i].ID == ID {
			return &activity.Property[i]

		}
	}

	return nil
}

// GetDataInputAssociation by ID
func (activity *Activity) GetDataInputAssociation(ID string) *DataInputAssociation {

	for i := 0; i < len(activity.DataInputAssociation); i++ {
		if activity.DataInputAssociation[i].ID == ID {
			return &activity.DataInputAssociation[i]

		}
	}

	return nil
}

// GetDataOutputAssociation by ID
func (activity *Activity) GetDataOutputAssociation(ID string) *DataOutputAssociation {

	for i := 0; i < len(activity.DataOutputAssociation); i++ {
		if activity.DataOutputAssociation[i].ID == ID {
			return &activity.DataOutputAssociation[i]

		}
	}

	return nil
}

// GetPerformer by ID
func (activity *Activity) GetPerformer(ID string) *Performer {

	for i := 0; i < len(activity.Performer); i++ {
		if activity.Performer[i].ID == ID {
			return &activity.Performer[i]

		}
	}

	return nil
}

