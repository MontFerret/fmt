package scopes

import "github.com/MontFerret/fmt/internal/core"

type ForExpressionScope struct {
	*Scope
	keyword string
	valVar  string
	keyVar  string
	variant string
}

func NewForExpressionScope(opts Options, keyword, valVar, keyVar, variant string) core.Scope {
	return &ForExpressionScope{
		Scope:   NewScope(opts),
		keyword: keyword,
		valVar:  valVar,
		keyVar:  keyVar,
		variant: variant,
	}
}

func (s *ForExpressionScope) Read(out core.Output) {
	out.WriteKeyword(s.keyword).WriteWhiteSpace().WriteAs(s.valVar)

	if s.keyVar != "" {
		out.Write(core.Comma).WriteWhiteSpace().WriteAs(s.keyVar)
	}

	out.WriteWhiteSpace().WriteKeyword(s.variant).WriteWhiteSpace()

	srcIdx := 0
	src := s.buff[srcIdx:1][0]
	src.WriteTo(out)

	out.AddTab()

	retIdx := len(s.buff) - 1

	hasBody := len(s.buff) > 2

	if hasBody {
		for _, tkn := range s.buff[srcIdx+1 : retIdx] {
			out.NewLine()
			tkn.WriteTo(out)
		}
	}

	ret := s.buff[retIdx:][0]

	out.NewLine()
	ret.WriteTo(out)
	out.RemoveTab()
}
