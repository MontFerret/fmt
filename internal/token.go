package internal

type (
	Token interface {
		String() string
	}

	StringToken string
)

const (
	WhiteSpace = StringToken(" ")

	NewLine = StringToken("\n")

	Comma = StringToken(",")
)

func (t StringToken) String() string {
	return string(t)
}