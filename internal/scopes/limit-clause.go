package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type LimitClauseScope struct {
	*Scope
	keyword string
}

func NewLimitClauseScope(opts Options, keyword string) core.Scope {
	return &LimitClauseScope{
		Scope:   NewScope(opts),
		keyword: keyword,
	}
}

func (s *LimitClauseScope) Read(out core.Output) {
	out.WriteKeyword(s.keyword).WriteWhiteSpace()

	if len(s.buff) > 1 {
		offset := s.buff[0]
		offset.WriteTo(out)
		out.Write(core.Comma).WriteWhiteSpace()

		count := s.buff[1]
		count.WriteTo(out)
	} else {
		count := s.buff[0]
		count.WriteTo(out)
	}
}
