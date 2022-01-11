package scopes

import "github.com/MontFerret/fmt/internal/core"

type FunctionScope struct {
	*Scope
	namespace []string
	name      string
	errSup    bool
}

func NewFunctionScope(opts Options, namespace []string, name string, errSup bool) core.Scope {
	return &FunctionScope{
		Scope:     NewScope(opts),
		namespace: namespace,
		name:      name,
		errSup:    errSup,
	}
}

func (s *FunctionScope) Read(out core.Output) {
	for _, seg := range s.namespace {
		out.WriteAs(seg)
	}

	out.WriteAs(s.name).Write(core.RBracketOpen)

	size := len(s.buff)
	last := size - 1

	withNewLine := s.useNewLine()

	if withNewLine {
		out.AddTab().NewLine()
	}

	for i, item := range s.buff {
		item.WriteTo(out)

		if i != last {
			out.Write(core.Comma)

			if withNewLine {
				out.NewLine()
			} else {
				out.WriteWhiteSpace()
			}
		}
	}

	if withNewLine {
		out.RemoveTab().NewLine()
	}

	out.Write(core.RBracketClose)

	if s.errSup {
		out.Write(core.QuestionMark)
	}
}
