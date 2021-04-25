package bpmn

import (
	"encoding/xml"
)

// Transaction from BPMN
type Transaction struct {
	SubProcess
	XMLName xml.Name
	Method  string `xml:"method,attr"`
}

// ParseTree of components of Transaction.
func (transaction *Transaction) ParseTree(definitions *Definitions) {
	transaction.SubProcess.ParseTree(definitions)

}
