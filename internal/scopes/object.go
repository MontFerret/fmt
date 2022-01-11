package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type ObjectScope struct {
	*Scope
}

func NewObjectScope(opts Options) core.Scope {
	return &ObjectScope{
		Scope: NewScope(opts),
	}
}

func (s *ObjectScope) Read(out core.Output) {
	out.Write(core.CBracketOpen)

	size := len(s.buff)
	last := size - 1

	withNewLine := s.useNewLine()
	withNewSpace := withNewLine == false && size > 0

	if withNewLine {
		out.AddTab().NewLine()
	} else if withNewSpace {
		out.WriteWhiteSpace()
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
	} else if withNewSpace {
		out.WriteWhiteSpace()
	}

	out.Write(core.CBracketClose)
}
