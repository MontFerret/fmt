package internal

type (
	CaseMode uint64

	Options struct {
		PrintWidth     uint64
		TabWidth       uint64
		SingleQuote    bool
		BracketSpacing bool
		CaseMode       CaseMode
	}
)

const (
	CaseModeIgnore CaseMode = iota
	CaseModeUpper
	CaseModeLower
)
