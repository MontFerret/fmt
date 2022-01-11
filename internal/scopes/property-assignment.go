package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type PropertyAssignmentScope struct {
	*Scope
	name        string
	isShorthand bool
}

func NewPropertyAssignmentScope(opts Options, name string, isShorthand bool) core.Scope {
	return &PropertyAssignmentScope{
		Scope:       NewScope(opts),
		name:        name,
		isShorthand: isShorthand,
	}
}

func (s *PropertyAssignmentScope) Len() int {
	return len(s.name) + s.Scope.Len()
}

func (s *PropertyAssignmentScope) WriteToken(token core.Token) {
	if s.isShorthand {
		return
	}

	s.Scope.WriteToken(token)
}

func (s *PropertyAssignmentScope) WriteScope(scope core.Scope) {
	if s.isShorthand {
		return
	}

	s.Scope.WriteScope(scope)
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
