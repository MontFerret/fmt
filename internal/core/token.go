package core

import "strings"

type (
	Measurable interface {
		Len() int
	}

	Token interface {
		Measurable
		String() string
	}

	TokenWriter interface {
		WriteToken(token Token)
	}

	StringToken string

	GroupToken struct {
		tokens []Token
	}
)

const (
	WhiteSpace = StringToken(" ")

	NewLine = StringToken("\n")

	Comma = StringToken(",")

	Dot = StringToken(".")

	Range = StringToken("..")

	Param = StringToken("@")

	Colon = StringToken(":")

	QuestionMark = StringToken("?")

	SBracketOpen = StringToken("[")

	SBracketClose = StringToken("]")

	CBracketOpen = StringToken("{")

	CBracketClose = StringToken("}")

	RBracketOpen = StringToken("(")

	RBracketClose = StringToken(")")
)

func (t StringToken) Len() int {
	return len(t)
}

func (t StringToken) String() string {
	return string(t)
}

func NewGroupToken(tokens ...Token) Token {
	return &GroupToken{tokens}
}

func (g *GroupToken) Len() int {
	size := len(g.tokens)

	for _, token := range g.tokens {
		size += token.Len()
	}

	return size
}

func (g *GroupToken) String() string {
	var b strings.Builder

	for i, token := range g.tokens {
		if i > 0 {
			b.WriteString(WhiteSpace.String())
		}

		b.WriteString(token.String())
	}

	return b.String()
}
