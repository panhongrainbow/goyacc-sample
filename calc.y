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
var regs = make(map[string]int, 26)

// one dimensional
var arrayMap = make(map[string][]int)

// two dimensionals
var arrayMap2 = make(map[string][][]int)

// three dimensionals
var arrayMap3 = make(map[string][][][]int)

// calculate
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
	str string
	slice []int
	twoDslice [][]int
	threeDslice [][][]int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <str> variable
%type <val> expr number
%type <slice> exprArray array array2
%type <twoDslice> twoDarray twoDarray2
%type <threeDslice> threeDarray threeDarray2

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
	|    variable '=' expr
		{
			regs[$1] = $3
		}
	|    exprArray
	|    variable '=' exprArray
        	{
        		arrayMap[$1] = $3
        	}
        |    variable '=' twoDarray
                {
                	arrayMap2[$1] = $3
                }
        |    variable '=' threeDarray
                {
                	arrayMap3[$1] = $3
                }
	;

exprArray : '[' variable ']'
	 { fmt.Println(arrayMap[$2]) }
	| array
         { $$ = $1 }
	| exprArray '+' '[' variable ']'
	 {
	  $$=make([]int, len($1))
	  for i := 0; i < len($1); i++ {
	    $$[i] = $1[i]+arrayMap[$4][i]
	  }
	 }
	| exprArray '+' array
	 {
	  $$=make([]int, len($1))
	  for i := 0; i < len($1); i++ {
	    $$[i] = $1[i]+$3[i]
	  }
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
	|    variable
		{ fmt.Println(regs[$1]) }
        |    '[' '[' variable ']' ']'
                { fmt.Println(arrayMap2[$3]) }
        |    '[' '[' '[' variable ']' ']' ']'
                { fmt.Println(arrayMap3[$4]) }
	|    number
	;

// array : (three dimensionals)
threeDarray2 : '[' twoDarray
	 { $$ = append($$, $2) }
	| threeDarray2 ',' twoDarray
	 { $$ = append($$, $3) }
	;
threeDarray : threeDarray2 ']'
	 { $$ = $1 }
	;

// array : (two dimensionals)
twoDarray2 : '[' array
	 { $$ = append($$, $2) }
	| twoDarray2 ',' array
	 { $$ = append($$, $3) }
	;
twoDarray : twoDarray2 ']'
	 { $$ = $1 }
	;

// array : (one dimensional)
array2 : '[' number
	 { $$ = append($$, $2) }
	| array2 ',' number
	 { $$ = append($$, $3) }
	;
array : array2 ']'
	 { $$ = $1 }
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

// variables
variable : LETTER
	{
	 $$ = string($1)
	}
	| variable LETTER
	 { $$ = $1 + string($2) }
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