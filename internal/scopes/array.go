package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type ArrayScope struct {
	*baseScope
}

func NewArrayScope(opts Options) core.Scope {
	return &ArrayScope{
		baseScope: newBaseScope(opts),
	}
}

func (s *ArrayScope) Read(out core.Output) {
	out.Write(core.SBracketOpen)

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

	out.Write(core.SBracketClose)
}
