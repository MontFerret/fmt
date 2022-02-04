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

func (w *Writer) StartFilterClause(keyword string) *Writer {
	w.out.StartScope(scopes.NewClauseScope(w.toScopeOptions(), keyword))

	return w
}

func (w *Writer) EndFilterClause() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartSortClauseExpression(direction string) *Writer {
	w.out.StartScope(scopes.NewSortClauseExpressionScope(
		w.toScopeOptions(),
		direction,
	))

	return w
}

func (w *Writer) EndSortClauseExpression() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartSortClause(keyword string) *Writer {
	w.out.StartScope(scopes.NewSortClauseScope(
		w.toScopeOptions(),
		keyword,
	))

	return w
}

func (w *Writer) EndSortClause() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartLimitClause(keyword string) *Writer {
	w.out.StartScope(scopes.NewLimitClauseScope(
		w.toScopeOptions(),
		keyword,
	))

	return w
}

func (w *Writer) EndLimitClause() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartForExpressionBody() *Writer {
	w.out.StartScope(scopes.NewScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndForExpressionBody() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartForExpressionSource() *Writer {
	w.out.StartScope(scopes.NewScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndForExpressionSource() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartForExpression(keyword, valVar, keyVar, variant string) *Writer {
	w.out.StartScope(scopes.NewForExpressionScope(
		w.toScopeOptions(),
		keyword,
		valVar,
		keyVar,
		variant,
	))

	return w
}

func (w *Writer) EndForExpression() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartMember() *Writer {
	w.out.StartScope(scopes.NewMemberScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndMember() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartPropertyName(name string, optional bool) *Writer {
	w.out.Write(scopes.NewPropertyNameToken(name, optional))

	return w
}

func (w *Writer) EndPropertyName() *Writer {
	return w
}

func (w *Writer) StartComputedPropertyName(optional bool) *Writer {
	w.out.StartScope(scopes.NewComputedPropertyNameScope(
		w.toScopeOptions(),
		optional,
	))

	return w
}

func (w *Writer) EndComputedPropertyName() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartFunctionCall(namespace []string, name string, errSup bool) *Writer {
	w.out.StartScope(scopes.NewFunctionScope(
		w.toScopeOptions(),
		namespace,
		name,
		errSup,
	))

	return w
}

func (w *Writer) EndFunctionCall() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartArray() *Writer {
	w.out.StartScope(scopes.NewArrayScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndArray() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartObject() *Writer {
	w.out.StartScope(scopes.NewObjectScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndObject() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartPropertyAssignment(name string, isShorthand bool) *Writer {
	w.out.StartScope(scopes.NewPropertyAssignmentScope(
		w.toScopeOptions(),
		name,
		isShorthand,
	))

	return w
}

func (w *Writer) EndPropertyAssignment() *Writer {
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
	if w.out.Scopes() == 0 {
		w.out.NewLine()
	}

	w.out.StartScope(scopes.NewReturnExpressionScope(w.toScopeOptions(), keyword))

	return w
}

func (w *Writer) EndReturn() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartTernaryOperator() *Writer {
	w.out.StartScope(scopes.NewTernaryOperatorScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndTernaryOperator() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartRangeOperator() *Writer {
	w.out.StartScope(scopes.NewRangeOperatorScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndRangeOperator() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartOperator(tokens ...string) *Writer {
	var token core.Token

	if len(tokens) > 1 {
		arr := make([]core.Token, 0, len(tokens))

		for _, t := range tokens {
			arr = append(arr, core.StringToken(t))
		}

		token = core.NewGroupToken(arr, core.WhiteSpace)
	} else {
		token = core.StringToken(tokens[0])
	}

	w.out.StartScope(scopes.NewOperatorScope(w.toScopeOptions(), token))

	return w
}

func (w *Writer) EndOperator() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) StartGroup() *Writer {
	w.out.StartScope(scopes.NewGroupScope(w.toScopeOptions()))

	return w
}

func (w *Writer) EndGroup() *Writer {
	w.out.EndScope()

	return w
}

func (w *Writer) WriteParam(input string) *Writer {
	w.out.Write(core.Param + core.StringToken(input))

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

func (w *Writer) toScopeOptions() scopes.Options {
	return scopes.Options{
		MaxLineLen: w.opts.PrintWidth,
	}
}
