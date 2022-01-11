package core

import (
	"fmt"
	"io"
	"strings"
)

const (
	SingleQuote       = "'"
	DoubleQuote       = "\""
	EscapeSingleQuote = "\\'"
	EscapeDoubleQuote = "\\\""
)

type (
	OutputSink interface {
		io.StringWriter
		fmt.Stringer
	}

	Output interface {
		StartScope(scope Scope) Output
		EndScope() Output
		Scopes() uint64
		Len() uint64
		Lines() uint64
		AddTab() Output
		RemoveTab() Output
		NewLine() Output
		WriteWhiteSpace() Output
		WriteString(input string) Output
		WriteLine(input string) Output
		WriteKeyword(word string) Output
		WriteAs(input string) Output
		Write(input Token) Output
		String() string
	}

	OutputWriter interface {
		WriteTo(output Output)
	}

	OutputBuilder struct {
		sink       OutputSink
		scopes     []Scope
		opts       Options
		len        uint64
		lines      uint64
		lineTokens uint64
		tabs       uint64
	}
)

func NewOutput(out OutputSink, opts Options) Output {
	return &OutputBuilder{
		sink:   out,
		scopes: make([]Scope, 0, 5),
		opts:   opts,
	}
}

func (o *OutputBuilder) StartScope(scope Scope) Output {
	o.scopes = append(o.scopes, scope)

	return o
}

func (o *OutputBuilder) EndScope() Output {
	if len(o.scopes) > 0 {
		scope := o.scopes[len(o.scopes)-1]
		o.scopes = o.scopes[0 : len(o.scopes)-1]

		// if it's a nested scope
		if len(o.scopes) > 0 {
			parentScope := o.scopes[len(o.scopes)-1]
			parentScope.WriteScope(scope)
		} else {
			// it's a root scope
			scope.Read(o)
		}
	}

	return o
}

func (o *OutputBuilder) Scopes() uint64 {
	return uint64(len(o.scopes))
}

func (o *OutputBuilder) Len() uint64 {
	return o.len
}

func (o *OutputBuilder) Lines() uint64 {
	return o.lines
}

func (o *OutputBuilder) LineTokens() uint64 {
	return o.lineTokens
}

func (o *OutputBuilder) AddTab() Output {
	o.tabs += 1

	return o
}

func (o *OutputBuilder) RemoveTab() Output {
	if o.tabs > 0 {
		o.tabs -= 1
	}

	return o
}

func (o *OutputBuilder) NewLine() Output {
	if o.len == 0 {
		return o
	}

	o.lines++

	o.Write(NewLine)

	for i := uint64(0); i < o.tabs; i++ {
		for i := uint64(0); i < o.opts.TabWidth; i++ {
			o.WriteWhiteSpace()
		}
	}

	return o
}

func (o *OutputBuilder) WriteWhiteSpace() Output {
	return o.Write(WhiteSpace)
}

func (o *OutputBuilder) WriteString(input string) Output {
	var targetQuote string
	var currentEscape string
	var targetEscape string

	if o.opts.SingleQuote {
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

	return o.Write(StringToken(result))
}

func (o *OutputBuilder) WriteLine(input string) Output {
	return o.NewLine().Write(StringToken(input))
}

func (o *OutputBuilder) WriteKeyword(word string) Output {
	switch o.opts.CaseMode {
	case CaseModeUpper:
		word = strings.ToUpper(word)
	case CaseModeLower:
		word = strings.ToLower(word)
	default:
		break
	}

	return o.Write(StringToken(word))
}

func (o *OutputBuilder) WriteAs(input string) Output {
	return o.Write(StringToken(input))
}

func (o *OutputBuilder) Write(input Token) Output {
	o.len += uint64(input.Len())

	if input != NewLine {
		o.lineTokens += 1
	} else {
		o.lineTokens = 0
	}

	if len(o.scopes) == 0 {
		_, _ = o.sink.WriteString(input.String())
	} else {
		o.scopes[len(o.scopes)-1].WriteToken(input)
	}

	return o
}

func (o *OutputBuilder) String() string {
	return o.sink.String()
}
