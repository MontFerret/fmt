package scopes

import "github.com/MontFerret/fmt/internal/core"

type (
	Options struct {
		MaxLineLen uint64
	}

	Scope struct {
		opts    Options
		buff    []core.OutputWriter
		lineLen uint64
	}
)

func NewScope(opts Options) *Scope {
	return &Scope{
		opts: opts,
		buff: make([]core.OutputWriter, 0, 10),
	}
}

func (s *Scope) Len() int {
	return int(s.lineLen)
}

func (s *Scope) WriteToken(token core.Token) {
	s.write(&tokenToOutput{token}, token.Len())
}

func (s *Scope) WriteScope(scope core.Scope) {
	s.write(&scopeToOutput{scope}, scope.Len())
}

func (s *Scope) Read(out core.Output) {
	for _, tkn := range s.buff {
		tkn.WriteTo(out)
	}
}

func (s *Scope) write(item core.OutputWriter, length int) {
	s.buff = append(s.buff, item)
	s.lineLen += uint64(length) + 2 // len + comma + white space
}

func (s *Scope) useNewLine() bool {
	return s.opts.MaxLineLen > 0 && s.lineLen >= s.opts.MaxLineLen
}
