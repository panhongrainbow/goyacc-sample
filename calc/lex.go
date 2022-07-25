package calc

import (
	"bytes"
	"errors"
	"unicode"
)

type CalcLex struct {
	S      string
	pos    int
	result interface{}
	err    error
}

func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.S) {
			return 0
		}
		c = rune(l.S[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c)
		return LETTER
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	// fmt.Printf("syntax error: %s\n", s)
	l.err = errors.New("syntax error")
}

func setResult(lexer CalcLexer, value interface{}) {
	lexer.(*CalcLex).result = value
}

// >>>>> CalcLex2

type CalcLex2 struct {
	input  *bytes.Buffer
	result interface{}
	err    error
}

func (l *CalcLex2) Lex(lval *CalcSymType) int {
	for {
		r, _, err := l.input.ReadRune()
		if err != nil {
			return 0
		}
		switch {
		case unicode.IsSpace(r):
			continue
		case unicode.IsDigit(r):
			lval.val = int(r) - '0'
			return DIGIT
		case unicode.IsLower(r):
			lval.val = int(r)
			return LETTER
		default:
			return int(r)
		}
	}
}

func (l *CalcLex2) Error(s string) {
	// fmt.Printf("syntax error: %s\n", s)
	l.err = errors.New("syntax error")
}

func CalcParse2(Calclex2 CalcLexer) int {
	return CalcNewParser().Parse(Calclex2)
}
