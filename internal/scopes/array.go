package scopes

import (
	"github.com/MontFerret/fmt/internal/core"
)

type ArrayScope struct {
	buff        []core.OutputWriter
	maxLineLen  int
	currLineLen int
}

func NewArrayScope(maxLineLen uint64) core.Scope {
	return &ArrayScope{
		buff:       make([]core.OutputWriter, 0, 10),
		maxLineLen: int(maxLineLen),
	}
}

func (s *ArrayScope) Len() int {
	return s.currLineLen
}

func (s *ArrayScope) WriteToken(token core.Token) {
	s.buff = append(s.buff, &tokenToOutput{token})
	s.increaseLen(token)
}

func (s *ArrayScope) WriteScope(scope core.Scope) {
	s.buff = append(s.buff, &scopeToOutput{scope})
	s.increaseLen(scope)
}

func (s *ArrayScope) Read(out core.Output) {
	out.Write(core.SBracketOpen)

	size := len(s.buff)
	last := size - 1

	withNewLine := s.maxLineLen > 0 && s.currLineLen >= s.maxLineLen

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

func (s *ArrayScope) increaseLen(m core.Measurable) {
	s.currLineLen += m.Len() + 2 // len + comma + white space
}
