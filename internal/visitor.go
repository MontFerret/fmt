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
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) {
}

func (v *Visitor) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (v *Visitor) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (v *Visitor) EnterProgram(c *fql.ProgramContext) {
}

func (v *Visitor) EnterHead(c *fql.HeadContext) {
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
}

func (v *Visitor) EnterBodyStatement(c *fql.BodyStatementContext) {
}

func (v *Visitor) EnterBodyExpression(c *fql.BodyExpressionContext) {
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
	var keyVar string
	var variant string

	valVar := c.GetValueVariable().GetText()

	if keyVarCtx := c.GetCounterVariable(); keyVarCtx != nil {
		keyVar = keyVarCtx.GetText()
	}

	if c.In() != nil {
		variant = "IN"
	} else {
		if c.Do() != nil {
			variant = "DO WHILE"
		} else {
			variant = "WHILE"
		}
	}

	v.writer.StartForExpression(c.For().GetText(), valVar, keyVar, variant)
}

func (v *Visitor) EnterForExpressionSource(c *fql.ForExpressionSourceContext) {
	v.writer.StartForExpressionSource()
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
	v.writer.WriteParam(c.Identifier().GetText())
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
	v.writer.StartObject()
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
	var name string
	var isShorthand bool

	if prop := c.PropertyName(); prop != nil {
		name = prop.GetText()
	} else if computed := c.ComputedPropertyName(); computed != nil {
		name = computed.GetText()
	} else if variable := c.Variable(); variable != nil {
		name = variable.GetText()
		isShorthand = true
	} else {
		name = c.GetText()
	}

	v.writer.StartPropertyAssignment(name, isShorthand)
}

func (v *Visitor) EnterComputedPropertyName(c *fql.ComputedPropertyNameContext) {
}

func (v *Visitor) EnterPropertyName(c *fql.PropertyNameContext) {
}

func (v *Visitor) EnterNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {
}

func (v *Visitor) EnterNamespace(c *fql.NamespaceContext) {
}

func (v *Visitor) EnterMemberExpression(c *fql.MemberExpressionContext) {
	v.writer.StartMember()
}

func (v *Visitor) EnterMemberExpressionSource(c *fql.MemberExpressionSourceContext) {
	if fc := c.FunctionCall(); fc != nil {
		v.startFunctionCall(fc.(*fql.FunctionCallContext), false)
	}
}

func (v *Visitor) EnterFunctionCallExpression(c *fql.FunctionCallExpressionContext) {
	fc := c.FunctionCall().(*fql.FunctionCallContext)

	v.startFunctionCall(fc, c.ErrorOperator() != nil)
}

func (v *Visitor) EnterFunctionCall(c *fql.FunctionCallContext) {}

func (v *Visitor) startFunctionCall(c *fql.FunctionCallContext, errSup bool) {
	ns := c.Namespace().(*fql.NamespaceContext)
	nsSegments := ns.AllNamespaceSegment()
	namespace := make([]string, 0, len(nsSegments))

	for _, seg := range nsSegments {
		namespace = append(namespace, seg.GetText())
	}

	name := c.FunctionName().GetText()

	v.writer.StartFunctionCall(namespace, name, errSup)
}

func (v *Visitor) EnterFunctionName(c *fql.FunctionNameContext) {
}

func (v *Visitor) EnterArgumentList(c *fql.ArgumentListContext) {
}

func (v *Visitor) EnterMemberExpressionPath(c *fql.MemberExpressionPathContext) {
	optional := c.ErrorOperator() != nil

	if prop := c.PropertyName(); prop != nil {
		v.writer.StartPropertyName(prop.(*fql.PropertyNameContext).Identifier().GetText(), optional)
	} else {
		v.writer.StartComputedPropertyName(optional)
	}
}

func (v *Visitor) EnterSafeReservedWord(c *fql.SafeReservedWordContext) {
}

func (v *Visitor) EnterUnsafReservedWord(c *fql.UnsafReservedWordContext) {
}

func (v *Visitor) EnterRangeOperator(c *fql.RangeOperatorContext) {
	v.writer.StartRangeOperator()
}

func (v *Visitor) EnterRangeOperand(c *fql.RangeOperandContext) {
}

func (v *Visitor) EnterExpression(c *fql.ExpressionContext) {
	if op := c.UnaryOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.LogicalAndOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.LogicalOrOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.GetTernaryOperator(); op != nil {
		v.writer.StartTernaryOperator()
	}
}

func (v *Visitor) EnterPredicate(c *fql.PredicateContext) {
	if op := c.EqualityOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.ArrayOperator(); op != nil {
		arrOp := op.(*fql.ArrayOperatorContext)
		var subOpText string
		var opText string

		if subOp := arrOp.GetOperator(); subOp != nil {
			subOpText = subOp.GetText()
		}

		if inOp := arrOp.InOperator(); inOp != nil {
			opText = inOp.GetText()
		} else {
			opText = arrOp.EqualityOperator().GetText()
		}

		if subOpText == "" {
			v.writer.StartOperator(opText)
		} else {
			v.writer.StartOperator(subOpText, opText)
		}
	} else if op := c.InOperator(); op != nil {
		inOp := op.(*fql.InOperatorContext)
		inToken := inOp.In().GetText()

		if not := inOp.Not(); not != nil {
			v.writer.StartOperator(not.GetText(), inToken)
		} else {
			v.writer.StartOperator(inToken)
		}
	} else if op := c.LikeOperator(); op != nil {
		likeOp := op.(*fql.LikeOperatorContext)
		likeToken := likeOp.Like().GetText()

		if not := likeOp.Not(); not != nil {
			v.writer.StartOperator(not.GetText(), likeToken)
		} else {
			v.writer.StartOperator(likeToken)
		}
	}
}

func (v *Visitor) EnterExpressionAtom(c *fql.ExpressionAtomContext) {
	if paren := c.OpenParen(); paren != nil {
		v.writer.StartGroup()
	}

	if op := c.MultiplicativeOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.AdditiveOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	} else if op := c.RegexpOperator(); op != nil {
		v.writer.StartOperator(op.GetText())
	}
}

func (v *Visitor) EnterArrayOperator(c *fql.ArrayOperatorContext) {
}

func (v *Visitor) EnterEqualityOperator(c *fql.EqualityOperatorContext) {
}

func (v *Visitor) EnterInOperator(c *fql.InOperatorContext) {
}

func (v *Visitor) EnterLikeOperator(c *fql.LikeOperatorContext) {
}

func (v *Visitor) EnterUnaryOperator(c *fql.UnaryOperatorContext) {
}

func (v *Visitor) EnterRegexpOperator(c *fql.RegexpOperatorContext) {
}

func (v *Visitor) EnterLogicalAndOperator(c *fql.LogicalAndOperatorContext) {
}

func (v *Visitor) EnterLogicalOrOperator(c *fql.LogicalOrOperatorContext) {
}

func (v *Visitor) EnterMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {
}

func (v *Visitor) EnterAdditiveOperator(c *fql.AdditiveOperatorContext) {
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
	v.writer.EndForExpression()
}

func (v *Visitor) ExitForExpressionSource(c *fql.ForExpressionSourceContext) {
	v.writer.EndForExpressionSource()
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
	v.writer.EndObject()
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
	v.writer.EndPropertyAssignment()
}

func (v *Visitor) ExitComputedPropertyName(c *fql.ComputedPropertyNameContext) {
}

func (v *Visitor) ExitPropertyName(c *fql.PropertyNameContext) {
}

func (v *Visitor) ExitNamespaceIdentifier(c *fql.NamespaceIdentifierContext) {
	//TODO implement me

}

func (v *Visitor) ExitNamespace(c *fql.NamespaceContext) {
	//TODO implement me

}

func (v *Visitor) ExitMemberExpression(c *fql.MemberExpressionContext) {
	v.writer.EndMember()
}

func (v *Visitor) ExitMemberExpressionSource(c *fql.MemberExpressionSourceContext) {
	if c.FunctionCall() != nil {
		v.writer.EndFunctionCall()
	}
}

func (v *Visitor) ExitFunctionCallExpression(c *fql.FunctionCallExpressionContext) {
	v.writer.EndFunctionCall()
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
	if c.PropertyName() != nil {
		v.writer.EndPropertyName()
	} else {
		v.writer.EndComputedPropertyName()
	}
}

func (v *Visitor) ExitSafeReservedWord(c *fql.SafeReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) ExitUnsafReservedWord(c *fql.UnsafReservedWordContext) {
	//TODO implement me

}

func (v *Visitor) ExitRangeOperator(c *fql.RangeOperatorContext) {
	v.writer.EndRangeOperator()
}

func (v *Visitor) ExitRangeOperand(c *fql.RangeOperandContext) {
	//TODO implement me

}

func (v *Visitor) ExitExpression(c *fql.ExpressionContext) {
	if op := c.UnaryOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.LogicalAndOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.LogicalOrOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.GetTernaryOperator(); op != nil {
		v.writer.EndOperator()
	}
}

func (v *Visitor) ExitPredicate(c *fql.PredicateContext) {
	if op := c.EqualityOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.ArrayOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.InOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.LikeOperator(); op != nil {
		v.writer.EndOperator()
	}
}

func (v *Visitor) ExitExpressionAtom(c *fql.ExpressionAtomContext) {
	if op := c.MultiplicativeOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.AdditiveOperator(); op != nil {
		v.writer.EndOperator()
	} else if op := c.RegexpOperator(); op != nil {
		v.writer.EndOperator()
	}

	if paren := c.CloseParen(); paren != nil {
		v.writer.EndGroup()
	}
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
}

func (v *Visitor) ExitRegexpOperator(c *fql.RegexpOperatorContext) {
	//TODO implement me

}

func (v *Visitor) ExitLogicalAndOperator(c *fql.LogicalAndOperatorContext) {
}

func (v *Visitor) ExitLogicalOrOperator(c *fql.LogicalOrOperatorContext) {
}

func (v *Visitor) ExitMultiplicativeOperator(c *fql.MultiplicativeOperatorContext) {
}

func (v *Visitor) ExitAdditiveOperator(c *fql.AdditiveOperatorContext) {
}

func (v *Visitor) ExitErrorOperator(c *fql.ErrorOperatorContext) {
}
