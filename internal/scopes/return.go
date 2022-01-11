package scopes

import "github.com/MontFerret/fmt/internal/core"

type ReturnExpressionScope struct {
	*Scope

	keyword string
}

func NewReturnExpressionScope(opts Options, keyword string) core.Scope {
	return &ReturnExpressionScope{
		Scope:   NewScope(opts),
		keyword: keyword,
	}
}

func (s *ReturnExpressionScope) Read(out core.Output) {
	out.WriteKeyword(s.keyword)

	for _, tkn := range s.buff {
		out.WriteWhiteSpace()
		tkn.WriteTo(out)
	}
}
