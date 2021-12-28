package fmt

import (
	"github.com/MontFerret/ferret/pkg/parser/fql"
	"github.com/MontFerret/fmt/internal"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
)

type Formatter struct {
	opts *Options
}

func New(setters ...Option) *Formatter {
	opts := DefaultOptions()

	for _, setter := range setters {
		setter(opts)
	}

	return &Formatter{opts}
}

func (fmt *Formatter) Format(query string) (output string, err error) {
	if query == "" {
		return "", nil
	}

	defer func() {
		if r := recover(); r != nil {
			// find out exactly what the error was and set err
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}

			output = ""
		}
	}()

	input := antlr.NewInputStream(query)
	lexer := fql.NewFqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := fql.NewFqlParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(&internal.ErrorListener{})

	w := internal.NewWriter(*fmt.opts)
	l := internal.NewVisitor(w)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Program())

	return w.String(), nil
}

func (fmt *Formatter) MustFormat(query string) string {
	out, err := fmt.Format(query)

	if err != nil {
		panic(err)
	}

	return out
}
