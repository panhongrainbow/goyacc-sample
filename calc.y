// Copyright 2011 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// based off of Appendix A from http://dinosaur.compilertools.net/yacc/

%{

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// a number
var regs = make([]int, 26)

// one dimensional
var arrayMap = make(map[int][]int)

// two dimensionals
var arrayMap2 = make(map[int][][]int)

// temporary
var tmp = []int{}
var tmp1 = []int{}
var tmp2 = [][]int{}
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number plenty

// same for terminals
%token <val> DIGIT LETTER

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%

list	: /* empty */
	| stat '\n'
	;

stat	:    expr
	|    LETTER '=' expr
		{
			regs[$1] = $3
		}
	|    LETTER '=' array
        	{
        		arrayMap[$1] = tmp
        	}
        |    LETTER '=' array2
                {
                	arrayMap2[$1] = tmp2
                }
	;

expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$  =  $1 + $3 }
	|    expr '-' expr
		{ $$  =  $1 - $3 }
	|    expr '*' expr
		{ $$  =  $1 * $3 }
	|    expr '/' expr
		{ $$  =  $1 / $3 }
	|    expr '%' expr
		{ $$  =  $1 % $3 }
	|    expr '&' expr
		{ $$  =  $1 & $3 }
	|    expr '|' expr
		{ $$  =  $1 | $3 }
	|    '-'  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    LETTER
		{ fmt.Println(regs[$1]) }
	|    '[' LETTER ']'
                { fmt.Println(arrayMap[$2]) }
        |    '[' '[' LETTER ']' ']'
                { fmt.Println(arrayMap2[$3]) }
	|    number
	;

// array : (two dimensionals)
array2	:    '[' plenty2 ']'

plenty2	:    array ','
                { tmp2 = append(tmp2, tmp1); tmp1 = []int{}}
	|    plenty2 array
		{ tmp2 = append(tmp2, tmp1); tmp1 = []int{}}
	|    plenty2 ',' array
		{ tmp2 = append(tmp2, tmp1); tmp1 = []int{}}

// array : (one dimensional)
array	:    '[' plenty ']'
		{ tmp1 = tmp; tmp = []int{} }
	;

plenty	:    plenty ',' number
		{ tmp = append(tmp, $3) }
	|    number ',' number
		{ tmp = append(tmp, $1, $3) }
	;

// number
number	:    DIGIT
		{
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}
	|    number DIGIT
		{ $$ = base * $1 + $2 }
	;

%%      /*  start  of  programs  */

type CalcLex struct {
	s string
	pos int
}


func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c) - 'a'
		return LETTER
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			CalcParse(&CalcLex{s: eqn})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}
