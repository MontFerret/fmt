package internal

import (
	"bytes"
	"strings"

	"github.com/MontFerret/fmt/internal/scopes"
)

const (
	SingleQuote       = "'"
	DoubleQuote       = "\""
	EscapeSingleQuote = "\\'"
	EscapeDoubleQuote = "\\\""
)

type Writer struct {
	opts  *Options
	buf   *bytes.Buffer
	lines uint64
	chars uint64
	scope scopes.Scope
}

func NewWriter(opts *Options) *Writer {
	return &Writer{
		opts: opts,
		buf:  &bytes.Buffer{},
	}
}

func (w *Writer) StartScope() {}

func (w *Writer) EndScope() {}

func (w *Writer) StartArray() *Writer {
	return w.startScope(func(out *bytes.Buffer) scopes.Scope {
		return scopes.NewArrayScope(out, w.opts.BracketSpacing)
	})
}

func (w *Writer) EndArray() *Writer {
	return w.endScope()
}

func (w *Writer) StartObject() *Writer {
	return w.startScope(func(out *bytes.Buffer) scopes.Scope {
		return scopes.NewObjectScope(out, w.opts.BracketSpacing)
	})
}

func (w *Writer) EndObject() *Writer {
	return w.endScope()
}

func (w *Writer) WritePropertyName(input string, isShorthand bool) *Writer {
	return w.writeToken(scopes.NewObjectScopeItem(input, isShorthand))
}

func (w *Writer) WriteBoolean(input string) *Writer {
	return w.Write(input)
}

func (w *Writer) WriteString(input string) *Writer {
	var targetQuote string
	var currentEscape string
	var targetEscape string

	if w.opts.SingleQuote {
		targetQuote = SingleQuote
		targetEscape = EscapeSingleQuote
		currentEscape = EscapeDoubleQuote
	} else {
		targetQuote = DoubleQuote
		targetEscape = EscapeDoubleQuote
		currentEscape = EscapeSingleQuote
	}

	b := []rune(input)

	// TODO: Use regexp
	// Replace opening and closing quotes
	result := targetQuote + string(b[1:len(b)-1]) + targetQuote
	// Remove redundant escapes
	result = strings.ReplaceAll(result, targetEscape, "")
	// Replace nested escapes
	result = strings.ReplaceAll(result, currentEscape, targetEscape)

	return w.Write(result)
}

func (w *Writer) WriteInteger(input string) *Writer {
	return w.Write(input)
}

func (w *Writer) WriteFloat(input string) *Writer {
	return w.Write(input)
}

func (w *Writer) WriteNone(input string) *Writer {
	return w.Write(input)
}

func (w *Writer) WriteVariableDeclaration(keyword, name string) *Writer {
	return w.
		NewLine().
		WriteKeyword(keyword).
		WriteWhiteSpace().
		Write(name).
		WriteWhiteSpace().
		Write("=").
		WriteWhiteSpace()
}

func (w *Writer) WriteVariable(name string) *Writer {
	return w.Write(name)
}

func (w *Writer) WriteReturn(input string) *Writer {
	return w.NewLine().WriteKeyword(input).WriteWhiteSpace()
}

func (w *Writer) WriteLine(input string) *Writer {
	return w.NewLine().Write(input)
}

func (w *Writer) NewLine() *Writer {
	if w.buf.Len() == 0 {
		return w
	}

	w.lines++

	w.writeToken(NewLine)

	//for i := uint64(0); i < w.opts.TabWidth; i++ {
	//	w.WriteWhiteSpace()
	//}

	return w
}

func (w *Writer) WriteKeyword(word string) *Writer {
	switch w.opts.CaseMode {
	case CaseModeUpper:
		word = strings.ToUpper(word)
	case CaseModeLower:
		word = strings.ToLower(word)
	default:
		break
	}

	return w.Write(word)
}

func (w *Writer) WriteWhiteSpace() *Writer {
	return w.writeToken(WhiteSpace)
}

func (w *Writer) Write(input string) *Writer {
	return w.writeToken(StringToken(input))
}

func (w *Writer) String() string {
	return w.buf.String()
}

func (w *Writer) writeToken(input Token) *Writer {
	if input != NewLine {
		w.chars += 1
	} else {
		w.chars = 0
	}

	if w.scope == nil {
		w.buf.WriteString(input.String())
	} else {
		w.scope.Write(input)
	}

	return w
}

func (w *Writer) startScope(f scopes.Factory) *Writer {
	w.scope = f(w.buf)
	w.scope.Start()

	return w
}

func (w *Writer) endScope() *Writer {
	w.scope.End()
	w.scope = nil

	return w
}
