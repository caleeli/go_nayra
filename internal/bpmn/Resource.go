package bpmn

import (
	"encoding/xml"
)

// Resource from BPMN
type Resource struct {
	RootElement
	XMLName           xml.Name
	ResourceParameter []ResourceParameter `xml:"resourceParameter"`
	Name              string              `xml:"name,attr"`

}

// ParseTree of components of Resource.
func (resource *Resource) ParseTree (definitions *Definitions) {
	resource.RootElement.ParseTree(definitions)

	for i := 0; i < len(resource.ResourceParameter); i++ {
		resource.ResourceParameter[i].ParseTree(definitions)
	}

}

// GetResourceParameter by ID
func (resource *Resource) GetResourceParameter(ID string) *ResourceParameter {

	for i := 0; i < len(resource.ResourceParameter); i++ {
		if resource.ResourceParameter[i].ID == ID {
			return &resource.ResourceParameter[i]

		}
	}

	return nil
}

