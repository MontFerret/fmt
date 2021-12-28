package core

type CaseMode uint64

const (
	CaseModeIgnore CaseMode = iota
	CaseModeUpper
	CaseModeLower
)
