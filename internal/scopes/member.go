package scopes

import "github.com/MontFerret/fmt/internal/core"

type (
	MemberScope struct {
		*baseScope
	}

	PropertyNameToken struct {
		name     string
		optional bool
	}

	ComputedPropertyNameScope struct {
		*baseScope
		optional bool
	}
)

func NewPropertyNameToken(name string, optional bool) core.Token {
	return &PropertyNameToken{name, optional}
}

func (p *PropertyNameToken) Len() int {
	size := len(p.name)

	if p.optional {
		return size + 1
	}

	return size
}

func (p *PropertyNameToken) String() string {
	if p.optional {
		return "?." + p.name
	}

	return "." + p.name
}

func NewComputedPropertyNameScope(opts Options, optional bool) core.Scope {
	return &ComputedPropertyNameScope{
		baseScope: newBaseScope(opts),
		optional:  optional,
	}
}

func (s *ComputedPropertyNameScope) Read(out core.Output) {
	if s.optional {
		out.Write(core.QuestionMark).Write(core.Dot)
	}

	out.Write(core.SBracketOpen)

	for _, item := range s.buff {
		item.WriteTo(out)
	}

	out.Write(core.SBracketClose)
}

func NewMemberScope(opts Options) core.Scope {
	return &MemberScope{
		baseScope: newBaseScope(opts),
	}
}

func (s *MemberScope) Read(out core.Output) {
	for _, item := range s.buff {
		item.WriteTo(out)
	}
}
