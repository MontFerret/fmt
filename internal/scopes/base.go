package scopes

import "bytes"

type (
	ScopeItem interface {
		String() string
	}

	Scope interface {
		Start() Scope
		Write(item ScopeItem) Scope
		End() Scope
		String() string
	}

	Factory func(out *bytes.Buffer) Scope

	BaseScopeItem string

	BaseScope struct {
		out   *bytes.Buffer
		items []ScopeItem
	}
)

func NewBaseScopeItem(value string) ScopeItem {
	return BaseScopeItem(value)
}

func (b BaseScopeItem) String() string {
	return string(b)
}

func newBaseScope(out *bytes.Buffer) *BaseScope {
	return &BaseScope{
		out:   out,
		items: make([]ScopeItem, 0, 20),
	}
}

func (s *BaseScope) Start() Scope {
	return s
}

func (s *BaseScope) Write(item ScopeItem) Scope {
	s.items = append(s.items, item)

	return s
}

func (s *BaseScope) End() Scope {
	return s
}

func (s *BaseScope) String() string {
	return s.out.String()
}
