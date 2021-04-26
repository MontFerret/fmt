package fmt

import "github.com/MontFerret/fmt/internal"

type (
	Options = internal.Options

	CaseMode = internal.CaseMode

	Option func(opts *Options)
)

const (
	CaseModeIgnore = internal.CaseModeIgnore
	CaseModeUpper  = internal.CaseModeUpper
	CaseModeLower  = internal.CaseModeLower
)

func DefaultOptions() *Options {
	return &Options{
		PrintWidth:     80, // Characters
		TabWidth:       4,  // Spaces
		SingleQuote:    false,
		BracketSpacing: true,
		CaseMode:       CaseModeIgnore,
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
