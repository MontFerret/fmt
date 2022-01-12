package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type LimitClauseScope struct {
	*Scope
	keyword string
	offset  string
	count   string
}

func NewLimitClauseScope(opts Options, keyword, offset, count string) core.Scope {
	return &LimitClauseScope{
		Scope:   NewScope(opts),
		keyword: keyword,
		offset:  offset,
		count:   count,
	}
}

func (s *LimitClauseScope) Read(out core.Output) {
	out.WriteKeyword(s.keyword).WriteWhiteSpace()

	if s.offset != "" {
		out.WriteAs(s.offset).Write(core.Comma).WriteWhiteSpace()
	}

	out.WriteAs(s.count)
}
