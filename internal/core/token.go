package core

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
)

const (
	WhiteSpace = StringToken(" ")

	NewLine = StringToken("\n")

	Comma = StringToken(",")

	Dot = StringToken(".")

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
