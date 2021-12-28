package core

type (
	Scope interface {
		Measurable
		TokenWriter
		ScopeWriter
		Read(output Output)
	}

	ScopeWriter interface {
		WriteScope(scope Scope)
	}
)
