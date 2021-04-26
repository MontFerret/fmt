package scopes

import "bytes"

type ArrayScope struct {
	*BaseScope
	bracketSpacing bool
}

func NewArrayScope(out *bytes.Buffer, bracketSpacing bool) Scope {
	return &ArrayScope{
		BaseScope: newBaseScope(out),
		bracketSpacing: bracketSpacing,
	}
}

func (s *ArrayScope) Start() Scope {
	s.out.WriteString("[")

	return s
}

func (s *ArrayScope) Write(item ScopeItem) Scope {
	s.BaseScope.Write(item)

	return s
}

func (s *ArrayScope) End() Scope {
	size := len(s.items)
	last := size - 1

	//if s.bracketSpacing && size > 0 {
	//	s.out.WriteString(" ")
	//}

	for i, item := range s.items {
		s.out.WriteString(item.String())

		if i != last {
			s.out.WriteString(", ")
		}
	}

	//if s.bracketSpacing && size > 0 {
	//	s.out.WriteString(" ")
	//}

	s.out.WriteString("]")

	return s
}

func (s *ArrayScope) String() string {
	return s.out.String()
}
