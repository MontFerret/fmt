package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type PropertyAssignmentScope struct {
	*baseScope
	name        string
	isShorthand bool
}

func NewPropertyAssignmentScope(opts Options, name string, isShorthand bool) core.Scope {
	return &PropertyAssignmentScope{
		baseScope:   newBaseScope(opts),
		name:        name,
		isShorthand: isShorthand,
	}
}

func (s *PropertyAssignmentScope) Len() int {
	return len(s.name) + s.baseScope.Len()
}

func (s *PropertyAssignmentScope) WriteToken(token core.Token) {
	if s.isShorthand {
		return
	}

	s.baseScope.WriteToken(token)
}

func (s *PropertyAssignmentScope) WriteScope(scope core.Scope) {
	if s.isShorthand {
		return
	}

	s.baseScope.WriteScope(scope)
}

func (s *PropertyAssignmentScope) Read(out core.Output) {
	out.WriteAs(s.name)

	// shorthand
	if len(s.buff) == 0 {
		return
	}

	out.Write(core.Colon).WriteWhiteSpace()

	s.buff[len(s.buff)-1].WriteTo(out)
}
