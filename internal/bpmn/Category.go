package bpmn

import (
	"encoding/xml"
)

// Category from BPMN
type Category struct {
	RootElement
	XMLName       xml.Name
	CategoryValue []CategoryValue `xml:"categoryValue"`
	Name          string          `xml:"name,attr"`
}

// ParseTree of components of Category.
func (category *Category) ParseTree(definitions *Definitions) {
	category.RootElement.ParseTree(definitions)

	for i := 0; i < len(category.CategoryValue); i++ {
		category.CategoryValue[i].ParseTree(definitions)
	}

}

// GetCategoryValue by ID
func (category *Category) GetCategoryValue(ID string) *CategoryValue {

	for i := 0; i < len(category.CategoryValue); i++ {
		if category.CategoryValue[i].ID == ID {
			return &category.CategoryValue[i]

		}
	}

	return nil
}
