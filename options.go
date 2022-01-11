package fmt

import (
	"github.com/MontFerret/fmt/internal/core"
)

type (
	Options = core.Options

	CaseMode = core.CaseMode

	Option func(opts *Options)
)

const (
	CaseModeIgnore = core.CaseModeIgnore
	CaseModeUpper  = core.CaseModeUpper
	CaseModeLower  = core.CaseModeLower
)

func DefaultOptions() *Options {
	return &Options{
		PrintWidth:     80, // Characters
		TabWidth:       4,  // Spaces
		SingleQuote:    false,
		BracketSpacing: true,
		CaseMode:       CaseModeUpper,
	}
}

func WithPrintWidth(size uint64) Option {
	return func(opts *Options) {
		opts.PrintWidth = size
	}
}

func WithTabWidth(size uint64) Option {
	return func(opts *Options) {
		opts.TabWidth = size
	}
}

func WithSingleQuote(val bool) Option {
	return func(opts *Options) {
		opts.SingleQuote = val
	}
}

func WithBracketSpacing(val bool) Option {
	return func(opts *Options) {
		opts.BracketSpacing = val
	}
}

func WithCaseMode(mode CaseMode) Option {
	return func(opts *Options) {
		opts.CaseMode = mode
	}
}
