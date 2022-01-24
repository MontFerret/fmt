package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type SortClauseExpressionScope struct {
	*Scope
	direction string
}

func NewSortClauseExpressionScope(opts Options, direction string) core.Scope {
	return &SortClauseExpressionScope{
		Scope:     NewScope(opts),
		direction: direction,
	}
}

func (s *SortClauseExpressionScope) Read(out core.Output) {
	s.Scope.Read(out)

	if s.direction != "" {
		out.WriteWhiteSpace().WriteKeyword(s.direction)
	}
}
