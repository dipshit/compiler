package tables

import (
	"github.com/dipshit/compiler/transducer"
)

type Table interface {
	// returns next table to use
	// can panic if its a transition table and an error occurs
	Run() int
}

// may be unused
type TableWithTransitions interface {
	ReportError() error
}

// builds an array of tables
func NewTableSet(transducer *transducer.Transducer) []Table {
	tableSlice := make([]Table, 7)
	tableSlice[0] = AcceptTable{transducer}
	tableSlice[1] = SemanticTable{transducer}
	tableSlice[2] = ShiftbackTable{transducer}
	tableSlice[3] = ReadaheadTable{transducer}
	tableSlice[4] = ReadbackTable{transducer}
	tableSlice[4] = ReduceTable{transducer}
	tableSlice[5] = ScannerReadaheadTable{transducer}
	return tableSlice
}

type AcceptTable struct {
	transducer *transducer.Transducer
}

type ScannerReadaheadTable struct {
	transducer *transducer.Transducer
}

type ReadaheadTable struct {
	transducer *transducer.Transducer
}

type ReadbackTable struct {
	transducer *transducer.Transducer
}

type SemanticTable struct {
	transducer *transducer.Transducer
}

type ReduceTable struct {
	transducer *transducer.Transducer
}

type ShiftbackTable struct {
	transducer *transducer.Transducer
}

// nothing left to parse
func (a AcceptTable) Run() int {
	return -1
}

func (a ScannerReadaheadTable) Run() int {
	return -1
}

func (a ReadaheadTable) Run() int {
	return -1
}

func (a ReadbackTable) Run() int {
	return -1
}

func (a SemanticTable) Run() int {
	return -1
}

func (a ReduceTable) Run() int {
	return -1
}

func (a ShiftbackTable) Run() int {
	return -1
}
