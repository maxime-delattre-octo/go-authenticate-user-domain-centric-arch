package domain

type Token struct {
	value string
}

func NewToken(value string) *Token {
	return &Token{value: value}
}
