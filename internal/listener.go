package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func (d *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	panic(errors.Errorf("%s at %d:%d", msg, line, column))
}
