package tables

import (
	"github.com/dipshit/compiler/transducer"
)

type Table interface {
	Transducer(*transducer.Transducer)
	Transducer() *transducer.Transducer
	Run()
}

type TableWithTransitions interface {
	ReportError() error
}

type AcceptTable struct{}
type ScannerReadaheadTable struct{}
type ReadaheadTable struct{}
type ReadbackTable struct{}
type SemanticTable struct{}
type ReduceTable struct{}
type ShiftbackTable struct{}
