package scopes

import "github.com/MontFerret/fmt/internal/core"

type (
	tokenToOutput struct {
		token core.Token
	}

	scopeToOutput struct {
		scope core.Scope
	}
)

func (t *tokenToOutput) Len() int {
	return t.token.Len()
}

func (t *tokenToOutput) WriteTo(output core.Output) {
	output.Write(t.token)
}

func (t *scopeToOutput) Len() int {
	return t.scope.Len()
}

func (t *scopeToOutput) WriteTo(output core.Output) {
	t.scope.Read(output)
}
