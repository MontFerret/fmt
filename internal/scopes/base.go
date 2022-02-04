package scopes

import "github.com/MontFerret/fmt/internal/core"

type (
	measurableOutputWriter interface {
		core.OutputWriter
		core.Measurable
	}

	Options struct {
		MaxLineLen uint64
	}

	Scope struct {
		opts    Options
		buff    []measurableOutputWriter
		lineLen uint64
	}
)

func NewScope(opts Options) *Scope {
	return &Scope{
		opts: opts,
		buff: make([]measurableOutputWriter, 0, 10),
	}
}

func (s *Scope) Len() int {
	return int(s.lineLen)
}

func (s *Scope) WriteToken(token core.Token) {
	s.write(&tokenToOutput{token})
}

func (s *Scope) WriteScope(scope core.Scope) {
	s.write(&scopeToOutput{scope})
}

func (s *Scope) Read(out core.Output) {
	for _, tkn := range s.buff {
		tkn.WriteTo(out)
	}
}

func (s *Scope) write(item measurableOutputWriter) {
	s.buff = append(s.buff, item)
	s.lineLen += uint64(item.Len())
}

func (s *Scope) useNewLine() bool {
	return s.opts.MaxLineLen > 0 && s.lineLen >= s.opts.MaxLineLen
}
