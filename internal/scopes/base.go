package scopes

import "github.com/MontFerret/fmt/internal/core"

type (
	Options struct {
		MaxLineLen uint64
	}

	baseScope struct {
		opts    Options
		buff    []core.OutputWriter
		lineLen uint64
	}
)

func newBaseScope(opts Options) *baseScope {
	return &baseScope{
		opts: opts,
		buff: make([]core.OutputWriter, 0, 10),
	}
}

func (b *baseScope) Len() int {
	return int(b.lineLen)
}

func (b *baseScope) WriteToken(token core.Token) {
	b.write(&tokenToOutput{token}, token.Len())
}

func (b *baseScope) WriteScope(scope core.Scope) {
	b.write(&scopeToOutput{scope}, scope.Len())
}

func (b *baseScope) Read(_ core.Output) {
	//TODO implement me
	panic("implement me")
}

func (b *baseScope) write(item core.OutputWriter, length int) {
	b.buff = append(b.buff, item)
	b.lineLen += uint64(length) + 2 // len + comma + white space
}

func (b *baseScope) useNewLine() bool {
	return b.opts.MaxLineLen > 0 && b.lineLen >= b.opts.MaxLineLen
}
