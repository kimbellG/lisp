package parser

import (
	"unicode"
)

//ParseCmd is parsing command on lexem
func ParseCmd(input string) ([]string, error) {
	return nil, nil
}

func isLex(r rune) bool {
	switch {
	case unicode.IsLetter(r):
		return false
	case unicode.IsNumber(r):
		return false
	case isOperator(r):
		return true
	case unicode.IsSpace(r):
		return true
	default:
		panic("Incorrect character")
	}
}

func isOperator(r rune) bool {
	switch r {
	case '+', '-', '/', '*':
		return true

	default:
		return false
	}
}
