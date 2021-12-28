package internal

import (
	"github.com/MontFerret/ferret/pkg/parser/fql"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitor struct {
	writer *Writer
}

func NewVisitor(writer *Writer) fql.FqlParserListener {
	return &Visitor{writer}
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) {
	//TODO implement me

}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) {
	//TODO implement me

}

func (v *Visitor) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me

}

func (v *Visitor) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//TODO implement me

}

func (v *Visitor) EnterProgram(c *fql.ProgramContext) {
	//TODO implement me

}

func (v *Visitor) EnterHead(c *fql.HeadContext) {
	//TODO implement me

}

func (v *Visitor) EnterUseExpression(c *fql.UseExpressionContext) {
	v.writer.StartUse(
		c.Use().GetChild(0).(antlr.TerminalNode).GetText(),
		c.Use().GetChild(1).(*fql.NamespaceIdentifierContext).GetText(),
	)
}

func (v *Visitor) EnterUse(c *fql.UseContext) {
}

func (v *Visitor) EnterBody(c *fql.BodyContext) {
	//TODO implement me

}

func (v *Visitor) EnterBodyStatement(c *fql.BodyStatementContext) {
	//TODO implement me
}

func (v *Visitor) EnterBodyExpression(c *fql.BodyExpressionContext) {
	//TODO implement me
}

func (v *Visitor) EnterVariableDeclaration(c *fql.VariableDeclarationContext) {
	var name string

	if id := c.Identifier(); id != nil {
		name = id.GetText()
	} else {
		name = c.IgnoreIdentifier().GetText()
	}

	v.writer.StartVariableDeclaration(c.Let().GetText(), name)
}

func (v *Visitor) EnterReturnExpression(c *fql.ReturnExpressionContext) {
	v.writer.StartReturn(c.Return().GetText())
}

func (v *Visitor) EnterForExpression(c *fql.ForExpressionContext) {
	//TODO implement me
}

func (v *Visitor) EnterForExpressionSource(c *fql.ForExpressionSourceContext) {
	//TODO implement me
}

func (v *Visitor) EnterForExpressionClause(c *fql.ForExpressionClauseContext) {
	//TODO implement me
}

func (v *Visitor) EnterForExpressionStatement(c *fql.ForExpressionStatementContext) {
	//TODO implement me

}

func (v *Visitor) EnterForExpressionBody(c *fql.ForExpressionBodyContext) {
	//TODO implement me

}

func (v *Visitor) EnterForExpressionReturn(c *fql.ForExpressionReturnContext) {
	//TODO implement me

}

func (v *Visitor) EnterFilterClause(c *fql.FilterClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterLimitClause(c *fql.LimitClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterLimitClauseValue(c *fql.LimitClauseValueContext) {
	//TODO implement me

}

func (v *Visitor) EnterSortClause(c *fql.SortClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterSortClauseExpression(c *fql.SortClauseExpressionContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectClause(c *fql.CollectClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectSelector(c *fql.CollectSelectorContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectGrouping(c *fql.CollectGroupingContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectAggregator(c *fql.CollectAggregatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectAggregateSelector(c *fql.CollectAggregateSelectorContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectGroupVariable(c *fql.CollectGroupVariableContext) {
	//TODO implement me

}

func (v *Visitor) EnterCollectCounter(c *fql.CollectCounterContext) {
	//TODO implement me

}

func (v *Visitor) EnterWaitForExpression(c *fql.WaitForExpressionContext) {
	//TODO implement me

}

func (v *Visitor) EnterWaitForEventName(c *fql.WaitForEventNameContext) {
	//TODO implement me

}

func (v *Visitor) EnterWaitForEventSource(c *fql.WaitForEventSourceContext) {
	//TODO implement me

}

func (v *Visitor) EnterOptionsClause(c *fql.OptionsClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterTimeoutClause(c *fql.TimeoutClauseContext) {
	//TODO implement me

}

func (v *Visitor) EnterParam(c *fql.ParamContext) {
	//TODO implement me

}

func (v *Visitor) EnterVariable(c *fql.VariableContext) {
	v.writer.Write(c.GetText())
}

func (v *Visitor) EnterLiteral(c *fql.LiteralContext) {
	//TODO implement me

}

func (v *Visitor) EnterArrayLiteral(c *fql.ArrayLiteralContext) {
	v.writer.StartArray()
}

func (v *Visitor) EnterObjectLiteral(c *fql.ObjectLiteralContext) {
	//TODO implement me

}

func (v *Visitor) EnterBooleanLiteral(c *fql.BooleanLiteralContext) {
	v.writer.Write(c.BooleanLiteral().GetText())
}

func (v *Visitor) EnterStringLiteral(c *fql.StringLiteralContext) {
	v.writer.WriteString(c.GetText())
}

func (v *Visitor) EnterFloatLiteral(c *fql.FloatLiteralContext) {
	v.writer.Write(c.FloatLiteral().GetText())
}

func (v *Visitor) EnterIntegerLiteral(c *fql.IntegerLiteralContext) {
	v.writer.Write(c.IntegerLiteral().GetText())
}

func (v *Visitor) EnterNoneLiteral(c *fql.NoneLiteralContext) {
	v.writer.Write(c.GetText())
}

func (v *Visitor) EnterPropertyAssignment(c *fql.PropertyAssignmentContext) {
	//TODO implement me

}

func (v *Visitor) EnterComputedPropertyName(c *fql.ComputedPropertyNameContext) {
	//TODO implement me

}

func (v *Visitor) EnterPropertyName(c *fql.PropertyNameContext) {
	//TODO implement me

}

func (v *Visitor) EnterNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {
}

func (v *Visitor) EnterNamespace(c *fql.NamespaceContext) {
}

func (v *Visitor) EnterMemberExpression(c *fql.MemberExpressionContext) {
	//TODO implement me
}

func (v *Visitor) EnterMemberExpressionSource(c *fql.MemberExpressionSourceContext) {
	//TODO implement me

}

func (v *Visitor) EnterFunctionCallExpression(c *fql.FunctionCallExpressionContext) {
	//TODO implement me

}

func (v *Visitor) EnterFunctionCall(c *fql.FunctionCallContext) {
	//TODO implement me

}

func (v *Visitor) EnterFunctionName(c *fql.FunctionNameContext) {
	//TODO implement me

}

func (v *Visitor) EnterArgumentList(c *fql.ArgumentListContext) {
	//TODO implement me

}

func (v *Visitor) EnterMemberExpressionPath(c *fql.MemberExpressionPathContext) {
	//TODO implement me

}

func (v *Visitor) EnterSafeReservedWord(c *fql.SafeReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) EnterUnsafReservedWord(c *fql.UnsafReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) EnterRangeOperator(c *fql.RangeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterRangeOperand(c *fql.RangeOperandContext) {
	//TODO implement me

}

func (v *Visitor) EnterExpression(c *fql.ExpressionContext) {
	//TODO implement me

}

func (v *Visitor) EnterPredicate(c *fql.PredicateContext) {
	//TODO implement me

}

func (v *Visitor) EnterExpressionAtom(c *fql.ExpressionAtomContext) {
	//TODO implement me

}

func (v *Visitor) EnterArrayOperator(c *fql.ArrayOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterEqualityOperator(c *fql.EqualityOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterInOperator(c *fql.InOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterLikeOperator(c *fql.LikeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterUnaryOperator(c *fql.UnaryOperatorContext) {
	v.writer.Write(c.GetText())
}

func (v *Visitor) EnterRegexpOperator(c *fql.RegexpOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterLogicalAndOperator(c *fql.LogicalAndOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterLogicalOrOperator(c *fql.LogicalOrOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterAdditiveOperator(c *fql.AdditiveOperatorContext) {
	//TODO implement me

}

func (v *Visitor) EnterErrorOperator(c *fql.ErrorOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitProgram(c *fql.ProgramContext) {
	//TODO implement me

}

func (v *Visitor) ExitHead(c *fql.HeadContext) {
	//TODO implement me

}

func (v *Visitor) ExitUseExpression(c *fql.UseExpressionContext) {
	v.writer.EndUse()
}

func (v *Visitor) ExitUse(c *fql.UseContext) {

}

func (v *Visitor) ExitBody(c *fql.BodyContext) {
	//TODO implement me

}

func (v *Visitor) ExitBodyStatement(c *fql.BodyStatementContext) {
	//TODO implement me

}

func (v *Visitor) ExitBodyExpression(c *fql.BodyExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitVariableDeclaration(c *fql.VariableDeclarationContext) {
	v.writer.EndVariableDeclaration()
}

func (v *Visitor) ExitReturnExpression(c *fql.ReturnExpressionContext) {
	v.writer.EndReturn()

}

func (v *Visitor) ExitForExpression(c *fql.ForExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitForExpressionSource(c *fql.ForExpressionSourceContext) {
	//TODO implement me

}

func (v *Visitor) ExitForExpressionClause(c *fql.ForExpressionClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitForExpressionStatement(c *fql.ForExpressionStatementContext) {
	//TODO implement me

}

func (v *Visitor) ExitForExpressionBody(c *fql.ForExpressionBodyContext) {
	//TODO implement me

}

func (v *Visitor) ExitForExpressionReturn(c *fql.ForExpressionReturnContext) {
	//TODO implement me

}

func (v *Visitor) ExitFilterClause(c *fql.FilterClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitLimitClause(c *fql.LimitClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitLimitClauseValue(c *fql.LimitClauseValueContext) {
	//TODO implement me

}

func (v *Visitor) ExitSortClause(c *fql.SortClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitSortClauseExpression(c *fql.SortClauseExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectClause(c *fql.CollectClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectSelector(c *fql.CollectSelectorContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectGrouping(c *fql.CollectGroupingContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectAggregator(c *fql.CollectAggregatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectAggregateSelector(c *fql.CollectAggregateSelectorContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectGroupVariable(c *fql.CollectGroupVariableContext) {
	//TODO implement me

}

func (v *Visitor) ExitCollectCounter(c *fql.CollectCounterContext) {
	//TODO implement me

}

func (v *Visitor) ExitWaitForExpression(c *fql.WaitForExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitWaitForEventName(c *fql.WaitForEventNameContext) {
	//TODO implement me

}

func (v *Visitor) ExitWaitForEventSource(c *fql.WaitForEventSourceContext) {
	//TODO implement me

}

func (v *Visitor) ExitOptionsClause(c *fql.OptionsClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitTimeoutClause(c *fql.TimeoutClauseContext) {
	//TODO implement me

}

func (v *Visitor) ExitParam(c *fql.ParamContext) {
	//TODO implement me

}

func (v *Visitor) ExitVariable(c *fql.VariableContext) {
	//TODO implement me

}

func (v *Visitor) ExitLiteral(c *fql.LiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitArrayLiteral(c *fql.ArrayLiteralContext) {
	v.writer.EndArray()
}

func (v *Visitor) ExitObjectLiteral(c *fql.ObjectLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitBooleanLiteral(c *fql.BooleanLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitStringLiteral(c *fql.StringLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitFloatLiteral(c *fql.FloatLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitIntegerLiteral(c *fql.IntegerLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitNoneLiteral(c *fql.NoneLiteralContext) {
	//TODO implement me

}

func (v *Visitor) ExitPropertyAssignment(c *fql.PropertyAssignmentContext) {
	//TODO implement me

}

func (v *Visitor) ExitComputedPropertyName(c *fql.ComputedPropertyNameContext) {
	//TODO implement me

}

func (v *Visitor) ExitPropertyName(c *fql.PropertyNameContext) {
	//TODO implement me

}

func (v *Visitor) ExitNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {
	//TODO implement me

}

func (v *Visitor) ExitNamespace(c *fql.NamespaceContext) {
	//TODO implement me

}

func (v *Visitor) ExitMemberExpression(c *fql.MemberExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitMemberExpressionSource(c *fql.MemberExpressionSourceContext) {
	//TODO implement me

}

func (v *Visitor) ExitFunctionCallExpression(c *fql.FunctionCallExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitFunctionCall(c *fql.FunctionCallContext) {
	//TODO implement me

}

func (v *Visitor) ExitFunctionName(c *fql.FunctionNameContext) {
	//TODO implement me

}

func (v *Visitor) ExitArgumentList(c *fql.ArgumentListContext) {
	//TODO implement me

}

func (v *Visitor) ExitMemberExpressionPath(c *fql.MemberExpressionPathContext) {
	//TODO implement me

}

func (v *Visitor) ExitSafeReservedWord(c *fql.SafeReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) ExitUnsafReservedWord(c *fql.UnsafReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) ExitRangeOperator(c *fql.RangeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitRangeOperand(c *fql.RangeOperandContext) {
	//TODO implement me

}

func (v *Visitor) ExitExpression(c *fql.ExpressionContext) {
	//TODO implement me

}

func (v *Visitor) ExitPredicate(c *fql.PredicateContext) {
	//TODO implement me

}

func (v *Visitor) ExitExpressionAtom(c *fql.ExpressionAtomContext) {
	//TODO implement me

}

func (v *Visitor) ExitArrayOperator(c *fql.ArrayOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitEqualityOperator(c *fql.EqualityOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitInOperator(c *fql.InOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitLikeOperator(c *fql.LikeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitUnaryOperator(c *fql.UnaryOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitRegexpOperator(c *fql.RegexpOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitLogicalAndOperator(c *fql.LogicalAndOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitLogicalOrOperator(c *fql.LogicalOrOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitAdditiveOperator(c *fql.AdditiveOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitErrorOperator(c *fql.ErrorOperatorContext) {
	//TODO implement me

}
