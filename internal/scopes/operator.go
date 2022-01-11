package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type OperatorScope struct {
	*Scope
	variant core.Token
}

func NewOperatorScope(opts Options, variant core.Token) core.Scope {
	return &OperatorScope{
		Scope:   NewScope(opts),
		variant: variant,
	}
}

func (s *OperatorScope) Read(out core.Output) {
	if len(s.buff) == 2 {
		left := s.buff[0]
		right := s.buff[1]

		left.WriteTo(out)
		out.WriteWhiteSpace()
		out.WriteKeyword(s.variant.String())
		out.WriteWhiteSpace()
		right.WriteTo(out)
	} else if len(s.buff) == 1 {
		out.Write(s.variant)
		arg := s.buff[0]
		arg.WriteTo(out)
	}
}
