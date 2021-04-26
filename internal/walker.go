package internal

import (
	"github.com/MontFerret/ferret/pkg/parser/fql"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Walker struct {
	writer *Writer
}

func NewWalker(writer *Writer) *Walker {
	return &Walker{writer}
}

func (w *Walker) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Println("Temrminal node:", node.GetText())
}

func (w *Walker) VisitErrorNode(node antlr.ErrorNode) {

}

func (w *Walker) EnterEveryRule(ctx antlr.ParserRuleContext) {

}

func (w *Walker) ExitEveryRule(ctx antlr.ParserRuleContext) {

}

func (w *Walker) EnterProgram(c *fql.ProgramContext) {

}

func (w *Walker) EnterHead(c *fql.HeadContext) {

}

func (w *Walker) EnterUseExpression(c *fql.UseExpressionContext) {

}

func (w *Walker) EnterUse(c *fql.UseContext) {

}

func (w *Walker) EnterBody(c *fql.BodyContext) {

}

func (w *Walker) EnterBodyStatement(c *fql.BodyStatementContext) {

}

func (w *Walker) EnterBodyExpression(c *fql.BodyExpressionContext) {

}

func (w *Walker) EnterReturnExpression(c *fql.ReturnExpressionContext) {
	w.writer.WriteReturn(c.Return().GetSymbol().GetText())
}

func (w *Walker) EnterForExpression(c *fql.ForExpressionContext) {

}

func (w *Walker) EnterForExpressionValueVariable(c *fql.ForExpressionValueVariableContext) {

}

func (w *Walker) EnterForExpressionKeyVariable(c *fql.ForExpressionKeyVariableContext) {

}

func (w *Walker) EnterForExpressionSource(c *fql.ForExpressionSourceContext) {

}

func (w *Walker) EnterForExpressionClause(c *fql.ForExpressionClauseContext) {

}

func (w *Walker) EnterForExpressionStatement(c *fql.ForExpressionStatementContext) {

}

func (w *Walker) EnterForExpressionBody(c *fql.ForExpressionBodyContext) {

}

func (w *Walker) EnterForExpressionReturn(c *fql.ForExpressionReturnContext) {

}

func (w *Walker) EnterFilterClause(c *fql.FilterClauseContext) {

}

func (w *Walker) EnterLimitClause(c *fql.LimitClauseContext) {

}

func (w *Walker) EnterLimitClauseValue(c *fql.LimitClauseValueContext) {

}

func (w *Walker) EnterSortClause(c *fql.SortClauseContext) {

}

func (w *Walker) EnterSortClauseExpression(c *fql.SortClauseExpressionContext) {

}

func (w *Walker) EnterCollectClause(c *fql.CollectClauseContext) {

}

func (w *Walker) EnterCollectSelector(c *fql.CollectSelectorContext) {

}

func (w *Walker) EnterCollectGrouping(c *fql.CollectGroupingContext) {

}

func (w *Walker) EnterCollectAggregator(c *fql.CollectAggregatorContext) {

}

func (w *Walker) EnterCollectAggregateSelector(c *fql.CollectAggregateSelectorContext) {

}

func (w *Walker) EnterCollectGroupVariable(c *fql.CollectGroupVariableContext) {

}

func (w *Walker) EnterCollectCounter(c *fql.CollectCounterContext) {

}

func (w *Walker) EnterVariableDeclaration(c *fql.VariableDeclarationContext) {
	w.writer.WriteVariableDeclaration(c.Let().GetText(), c.Identifier().GetText())
}

func (w *Walker) EnterParam(c *fql.ParamContext) {

}

func (w *Walker) EnterVariable(c *fql.VariableContext) {
	w.writer.WriteVariable(c.Identifier().GetText())
}

func (w *Walker) EnterRangeOperator(c *fql.RangeOperatorContext) {

}

func (w *Walker) EnterArrayLiteral(c *fql.ArrayLiteralContext) {
	w.writer.StartArray()
}

func (w *Walker) EnterObjectLiteral(c *fql.ObjectLiteralContext) {
	w.writer.StartObject()
}

func (w *Walker) EnterBooleanLiteral(c *fql.BooleanLiteralContext) {
	w.writer.WriteBoolean(c.GetText())
}

func (w *Walker) EnterStringLiteral(c *fql.StringLiteralContext) {
	w.writer.WriteString(c.GetText())
}

func (w *Walker) EnterIntegerLiteral(c *fql.IntegerLiteralContext) {
	w.writer.WriteInteger(c.GetText())
}

func (w *Walker) EnterFloatLiteral(c *fql.FloatLiteralContext) {
	w.writer.WriteFloat(c.GetText())
}

func (w *Walker) EnterNoneLiteral(c *fql.NoneLiteralContext) {
	w.writer.WriteNone(c.GetText())
}

func (w *Walker) EnterArrayElementList(c *fql.ArrayElementListContext) {
}

func (w *Walker) EnterPropertyAssignment(c *fql.PropertyAssignmentContext) {
	pn := c.PropertyName()

	if pn != nil {
		w.writer.WritePropertyName(pn.GetText(), false)
	}
}

func (w *Walker) EnterShorthandPropertyName(c *fql.ShorthandPropertyNameContext) {
	// w.writer.WritePropertyName(c.Variable().GetText(), true)
}

func (w *Walker) EnterComputedPropertyName(c *fql.ComputedPropertyNameContext) {

}

func (w *Walker) EnterPropertyName(c *fql.PropertyNameContext) {

}

func (w *Walker) EnterExpressionGroup(c *fql.ExpressionGroupContext) {

}

func (w *Walker) EnterNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {

}

func (w *Walker) EnterNamespace(c *fql.NamespaceContext) {

}

func (w *Walker) EnterFunctionIdentifier(c *fql.FunctionIdentifierContext) {

}

func (w *Walker) EnterFunctionCallExpression(c *fql.FunctionCallExpressionContext) {

}

func (w *Walker) EnterMember(c *fql.MemberContext) {

}

func (w *Walker) EnterMemberPath(c *fql.MemberPathContext) {

}

func (w *Walker) EnterMemberExpression(c *fql.MemberExpressionContext) {

}

func (w *Walker) EnterArguments(c *fql.ArgumentsContext) {

}

func (w *Walker) EnterExpression(c *fql.ExpressionContext) {

}

func (w *Walker) EnterForTernaryExpression(c *fql.ForTernaryExpressionContext) {

}

func (w *Walker) EnterArrayOperator(c *fql.ArrayOperatorContext) {

}

func (w *Walker) EnterInOperator(c *fql.InOperatorContext) {

}

func (w *Walker) EnterLikeOperator(c *fql.LikeOperatorContext) {

}

func (w *Walker) EnterEqualityOperator(c *fql.EqualityOperatorContext) {

}

func (w *Walker) EnterRegexpOperator(c *fql.RegexpOperatorContext) {

}

func (w *Walker) EnterLogicalAndOperator(c *fql.LogicalAndOperatorContext) {

}

func (w *Walker) EnterLogicalOrOperator(c *fql.LogicalOrOperatorContext) {

}

func (w *Walker) EnterMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {

}

func (w *Walker) EnterAdditiveOperator(c *fql.AdditiveOperatorContext) {

}

func (w *Walker) EnterUnaryOperator(c *fql.UnaryOperatorContext) {

}

func (w *Walker) ExitProgram(c *fql.ProgramContext) {

}

func (w *Walker) ExitHead(c *fql.HeadContext) {

}

func (w *Walker) ExitUseExpression(c *fql.UseExpressionContext) {

}

func (w *Walker) ExitUse(c *fql.UseContext) {

}

func (w *Walker) ExitBody(c *fql.BodyContext) {

}

func (w *Walker) ExitBodyStatement(c *fql.BodyStatementContext) {

}

func (w *Walker) ExitBodyExpression(c *fql.BodyExpressionContext) {

}

func (w *Walker) ExitReturnExpression(c *fql.ReturnExpressionContext) {

}

func (w *Walker) ExitForExpression(c *fql.ForExpressionContext) {

}

func (w *Walker) ExitForExpressionValueVariable(c *fql.ForExpressionValueVariableContext) {

}

func (w *Walker) ExitForExpressionKeyVariable(c *fql.ForExpressionKeyVariableContext) {

}

func (w *Walker) ExitForExpressionSource(c *fql.ForExpressionSourceContext) {

}

func (w *Walker) ExitForExpressionClause(c *fql.ForExpressionClauseContext) {

}

func (w *Walker) ExitForExpressionStatement(c *fql.ForExpressionStatementContext) {

}

func (w *Walker) ExitForExpressionBody(c *fql.ForExpressionBodyContext) {

}

func (w *Walker) ExitForExpressionReturn(c *fql.ForExpressionReturnContext) {

}

func (w *Walker) ExitFilterClause(c *fql.FilterClauseContext) {

}

func (w *Walker) ExitLimitClause(c *fql.LimitClauseContext) {

}

func (w *Walker) ExitLimitClauseValue(c *fql.LimitClauseValueContext) {

}

func (w *Walker) ExitSortClause(c *fql.SortClauseContext) {

}

func (w *Walker) ExitSortClauseExpression(c *fql.SortClauseExpressionContext) {

}

func (w *Walker) ExitCollectClause(c *fql.CollectClauseContext) {

}

func (w *Walker) ExitCollectSelector(c *fql.CollectSelectorContext) {

}

func (w *Walker) ExitCollectGrouping(c *fql.CollectGroupingContext) {

}

func (w *Walker) ExitCollectAggregator(c *fql.CollectAggregatorContext) {

}

func (w *Walker) ExitCollectAggregateSelector(c *fql.CollectAggregateSelectorContext) {

}

func (w *Walker) ExitCollectGroupVariable(c *fql.CollectGroupVariableContext) {

}

func (w *Walker) ExitCollectCounter(c *fql.CollectCounterContext) {

}

func (w *Walker) ExitVariableDeclaration(c *fql.VariableDeclarationContext) {

}

func (w *Walker) ExitParam(c *fql.ParamContext) {

}

func (w *Walker) ExitVariable(c *fql.VariableContext) {

}

func (w *Walker) ExitRangeOperator(c *fql.RangeOperatorContext) {

}

func (w *Walker) ExitArrayLiteral(c *fql.ArrayLiteralContext) {
	w.writer.EndArray()
}

func (w *Walker) ExitObjectLiteral(c *fql.ObjectLiteralContext) {
	w.writer.EndObject()
}

func (w *Walker) ExitBooleanLiteral(c *fql.BooleanLiteralContext) {

}

func (w *Walker) ExitStringLiteral(c *fql.StringLiteralContext) {

}

func (w *Walker) ExitIntegerLiteral(c *fql.IntegerLiteralContext) {

}

func (w *Walker) ExitFloatLiteral(c *fql.FloatLiteralContext) {

}

func (w *Walker) ExitNoneLiteral(c *fql.NoneLiteralContext) {

}

func (w *Walker) ExitArrayElementList(c *fql.ArrayElementListContext) {

}

func (w *Walker) ExitPropertyAssignment(c *fql.PropertyAssignmentContext) {

}

func (w *Walker) ExitShorthandPropertyName(c *fql.ShorthandPropertyNameContext) {

}

func (w *Walker) ExitComputedPropertyName(c *fql.ComputedPropertyNameContext) {

}

func (w *Walker) ExitPropertyName(c *fql.PropertyNameContext) {

}

func (w *Walker) ExitExpressionGroup(c *fql.ExpressionGroupContext) {

}

func (w *Walker) ExitNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {

}

func (w *Walker) ExitNamespace(c *fql.NamespaceContext) {

}

func (w *Walker) ExitFunctionIdentifier(c *fql.FunctionIdentifierContext) {

}

func (w *Walker) ExitFunctionCallExpression(c *fql.FunctionCallExpressionContext) {

}

func (w *Walker) ExitMember(c *fql.MemberContext) {

}

func (w *Walker) ExitMemberPath(c *fql.MemberPathContext) {

}

func (w *Walker) ExitMemberExpression(c *fql.MemberExpressionContext) {

}

func (w *Walker) ExitArguments(c *fql.ArgumentsContext) {

}

func (w *Walker) ExitExpression(c *fql.ExpressionContext) {

}

func (w *Walker) ExitForTernaryExpression(c *fql.ForTernaryExpressionContext) {

}

func (w *Walker) ExitArrayOperator(c *fql.ArrayOperatorContext) {

}

func (w *Walker) ExitInOperator(c *fql.InOperatorContext) {

}

func (w *Walker) ExitLikeOperator(c *fql.LikeOperatorContext) {

}

func (w *Walker) ExitEqualityOperator(c *fql.EqualityOperatorContext) {

}

func (w *Walker) ExitRegexpOperator(c *fql.RegexpOperatorContext) {

}

func (w *Walker) ExitLogicalAndOperator(c *fql.LogicalAndOperatorContext) {

}

func (w *Walker) ExitLogicalOrOperator(c *fql.LogicalOrOperatorContext) {

}

func (w *Walker) ExitMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {

}

func (w *Walker) ExitAdditiveOperator(c *fql.AdditiveOperatorContext) {

}

func (w *Walker) ExitUnaryOperator(c *fql.UnaryOperatorContext) {

}
