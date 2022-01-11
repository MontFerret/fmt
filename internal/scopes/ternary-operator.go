package scopes

import "github.com/MontFerret/fmt/internal/core"

type TernaryOperatorScope struct {
	*OperatorScope
}

func NewTernaryOperatorScope(opts Options) core.Scope {
	return &TernaryOperatorScope{
		OperatorScope: &OperatorScope{
			Scope:   NewScope(opts),
			variant: core.QuestionMark,
		},
	}
}

func (s *TernaryOperatorScope) Read(out core.Output) {
	cond := s.buff[0]

	cond.WriteTo(out)
	out.WriteWhiteSpace().Write(s.variant).WriteWhiteSpace()

	if len(s.buff) == 3 {
		left := s.buff[1]
		left.WriteTo(out)
		out.WriteWhiteSpace().Write(core.Colon).WriteWhiteSpace()

		right := s.buff[2]
		right.WriteTo(out)
	} else if len(s.buff) == 2 {
		out.Write(core.Colon).WriteWhiteSpace()

		right := s.buff[1]
		right.WriteTo(out)
	}
}
