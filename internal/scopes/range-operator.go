package scopes

import "github.com/MontFerret/fmt/internal/core"

type RangeOperatorScope struct {
	*OperatorScope
}

func NewRangeOperatorScope(opts Options) core.Scope {
	return &RangeOperatorScope{
		OperatorScope: &OperatorScope{
			Scope:   NewScope(opts),
			variant: core.Range,
		},
	}
}

func (s *RangeOperatorScope) Read(out core.Output) {
	left := s.buff[0]
	right := s.buff[1]

	left.WriteTo(out)
	out.Write(s.variant)
	right.WriteTo(out)
}
