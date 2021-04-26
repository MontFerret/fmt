package scopes

import "bytes"

type (
	ObjectScopeItem struct {
		name string
		isShorthand bool
	}

	ObjectScope struct {
		*BaseScope
		bracketSpacing bool
	}
)

func NewObjectScope(out *bytes.Buffer, bracketSpacing bool) Scope {
	return &ObjectScope{
		BaseScope: newBaseScope(out),
		bracketSpacing: bracketSpacing,
	}
}

func NewObjectScopeItem(name string, isShorthand bool) ScopeItem {
	return &ObjectScopeItem{name, isShorthand}
}

func (i *ObjectScopeItem) String() string {
	if i.isShorthand {
		return i.name
	}

	return i.name + ": "
}

func (s *ObjectScope) Start() Scope {
	s.out.WriteString("{")

	return s
}

func (s *ObjectScope) Write(item ScopeItem) Scope {
	s.BaseScope.Write(item)

	return s
}

func (s *ObjectScope) End() Scope {
	size := len(s.items)
	last := size - 1

	if s.bracketSpacing && size > 0 {
		s.out.WriteString(" ")
	}

	for i, item := range s.items {
		s.out.WriteString(item.String())

		if i != last {
			if _, ok := item.(*ObjectScopeItem); !ok {
				s.out.WriteString(", ")
			}
		}
	}

	if s.bracketSpacing && size > 0 {
		s.out.WriteString(" ")
	}

	s.out.WriteString("}")

	return s
}

func (s *ObjectScope) String() string {
	return s.out.String()
}
