package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type ClauseScope struct {
	*Scope
	keyword string
}

func NewClauseScope(opts Options, keyword string) core.Scope {
	return &ClauseScope{
		Scope:   NewScope(opts),
		keyword: keyword,
	}
}

func (s *ClauseScope) Read(out core.Output) {
	out.WriteKeyword(s.keyword).WriteWhiteSpace()

	for _, item := range s.buff {
		item.WriteTo(out)
	}
}
