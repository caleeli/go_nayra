package bpmn

import (
	"encoding/xml"
)

// PotentialOwner from BPMN
type PotentialOwner struct {
	HumanPerformer
	XMLName xml.Name
}

// ParseTree of components of PotentialOwner.
func (potentialOwner *PotentialOwner) ParseTree(definitions *Definitions) {
	potentialOwner.HumanPerformer.ParseTree(definitions)

}
