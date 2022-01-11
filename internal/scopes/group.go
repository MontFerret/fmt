package scopes

import "github.com/MontFerret/fmt/internal/core"

type GroupScope struct {
	*OperatorScope
}

func NewGroupScope(opts Options) core.Scope {
	return &GroupScope{
		OperatorScope: &OperatorScope{
			Scope:   NewScope(opts),
			variant: core.Range,
		},
	}
}

func (s *GroupScope) Read(out core.Output) {
	out.Write(core.RBracketOpen)

	for _, w := range s.buff {
		w.WriteTo(out)
	}

	out.Write(core.RBracketClose)
}
