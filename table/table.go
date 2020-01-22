package tables

import (
	"github.com/dipshit/compiler/token"
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
	return tableSlice
}

func NewScannerTableSet(transducer *transducer.Transducer) []Table {
	table := []Table{
		ScannerReadaheadTable{transducer, 0,
			[]*ScannerRow{
				// EOF row is dummy, should compare to io.EOF
				NewScannerRow("EOF", token.L, 5),
				NewScannerRow(")", token.RK, 7),
				NewScannerRow("*", token.RK, 8),
				NewScannerRow("+", token.RK, 9),
				NewScannerRow(",", token.RK, 10),
				NewScannerRow("0123456789", token.RK, 2),
				NewScannerRow("(", token.RK, 6),
				NewScannerRow(";", token.RK, 12),
				NewScannerRow("=", token.RK, 13),
				NewScannerRow("ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz", token.RK, 3),
				NewScannerRow("\t\n\f\r", token.R, 4),
				NewScannerRow(" ", token.R, 4),
			},
		},
		ScannerReadaheadTable{transducer, 1,
			[]*ScannerRow{
				NewScannerRow("\t\n\f\r", token.L, 11),
				NewScannerRow("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_+*=[]{}()^;#:.$ ", token.L, 11),
				NewScannerRow("0123456789", token.RK, 2),
			},
		},
		ScannerReadaheadTable{transducer, 2,
			[]*ScannerRow{
				NewScannerRow("\t\n\f\r", token.L, 14),
				NewScannerRow("+*=[]{}()^;#:.$ ", token.L, 14),
				NewScannerRow("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz", token.RK, 3),
			},
		},
		ScannerReadaheadTable{transducer, 3,
			[]*ScannerRow{
				NewScannerRow("EOF", token.L, 1),
				NewScannerRow("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_0123456789+*=[]{}()^;#:.$", token.L, 1),
				NewScannerRow("\t\n\f\r", token.R, 4),
				NewScannerRow(" ", token.R, 4),
			},
		},
	}
	return table
}

type AcceptTable struct {
	transducer *transducer.Transducer
}

type ScannerReadaheadTable struct {
	transducer *transducer.Transducer
	index      int
	rows       []*ScannerRow
}

func NewScannerRow(regexpr string, action token.TokenAction, next int) *ScannerRow {
	return &ScannerRow{regexpr, action, next}
}

type ScannerRow struct {
	regexpr string            // some sort of string or array
	action  token.TokenAction // action to do
	next    int               // next row to run
}

type ReadaheadTable struct {
	transducer *transducer.Transducer
}

type ReadbackTable struct {
	transducer *transducer.Transducer
}

type SemanticTable struct {
	transducer *transducer.Transducer
	index      int
	action     string
	tok        token.TokenType
	next       int
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
