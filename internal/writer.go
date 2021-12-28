package internal

import (
	"github.com/MontFerret/fmt/internal/core"
	"github.com/MontFerret/fmt/internal/scopes"
	"strings"
)

type Writer struct {
	opts core.Options
	out  core.Output
}

func NewWriter(opts core.Options) *Writer {
	return &Writer{
		opts: opts,
		out:  core.NewOutput(&strings.Builder{}, opts),
	}
}

func (w *Writer) StartArray() *Writer {
	w.out.StartScope(scopes.NewArrayScope(
		w.opts.PrintWidth,
	))

	return w
}

func (w *Writer) EndArray() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartUse(keyword, namespace string) *Writer {
	w.out.WriteKeyword(keyword).WriteWhiteSpace().Write(core.StringToken(namespace))

	return w
}

func (w *Writer) EndUse() *Writer {
	w.out.NewLine()

	return w
}

func (w *Writer) StartVariableDeclaration(keyword, name string) *Writer {
	w.
		out.
		NewLine().
		WriteKeyword(keyword).
		WriteWhiteSpace().
		WriteAs(name).
		WriteWhiteSpace().
		WriteAs("=").
		WriteWhiteSpace()

	return w
}

func (w *Writer) EndVariableDeclaration() *Writer {
	w.out.NewLine()

	return w
}

func (w *Writer) StartReturn(keyword string) *Writer {
	if w.out.Lines() > 0 {
		w.out.NewLine()
	}

	w.out.WriteKeyword(keyword).WriteWhiteSpace()

	return w
}

func (w *Writer) EndReturn() *Writer {
	//w.out.NewLine()

	return w
}

func (w *Writer) WriteString(input string) *Writer {
	w.out.WriteString(input)

	return w
}

func (w *Writer) Write(input string) *Writer {
	w.out.WriteAs(input)

	return w
}

func (w *Writer) String() string {
	return w.out.String()
}
