package calc

import (
	"bytes"
	"errors"
	"fmt"
	"unicode"
)

var identified = map[string]int{
	"ASK": TokenASK,
	"ask": TokenASK,
}

type CalcLex struct {
	S         string
	pos       int
	result    interface{}
	err       error
	pickUp    string
	lexBuffer []rune
	finish    bool
}

func (l *CalcLex) Lex(lval *CalcSymType) int {
	c1 := rune(-1)

LOOP1:
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.S) {
			l.finish = true
			goto LOOP2
			// return 0
		}
		c = rune(l.S[l.pos])
		l.pos += 1
	}

	// fmt.Println("c", c)

LOOP2:
	for {
		l.lexBuffer = append(l.lexBuffer, c)
		if len(l.lexBuffer) < 3 && l.finish == false {
			goto LOOP1
		}
		if len(l.lexBuffer) == 3 && l.finish == false {
			c1 = l.lexBuffer[0]
			l.lexBuffer[0] = l.lexBuffer[1]
			l.lexBuffer[1] = l.lexBuffer[2]
			l.lexBuffer = l.lexBuffer[:2]
			break LOOP2
		}
		if l.finish == true && len(l.lexBuffer) > 0 {
			c1 = l.lexBuffer[0]
			l.lexBuffer = l.lexBuffer[1:]
			break LOOP2
		}
		if l.finish == true && l.lexBuffer[0] == 10 {
			return 0
		}
	}

	// fmt.Println(">>>>> >>>>> >>>>> c1", c1)

	if unicode.IsDigit(c1) {
		lval.val = int(c1) - '0'
		return DIGIT
	} else if unicode.IsLower(c1) {
		lval.val = int(c1)
		return LETTER
	}
	return int(c1)
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
	fmt.Println(identified)
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
