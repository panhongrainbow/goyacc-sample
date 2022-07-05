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

// three dimensionals
var arrayMap3 = make(map[int][][][]int)

// temporary
// var planty = []int{}
var array = []int{}
var twoDarray = [][]int{}
var threeDarray = [][][]int{}

// calculate
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
	slice []int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number
%type <slice> plenty

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
        		arrayMap[$1] = array
        	}
        |    LETTER '=' twoDarray
                {
                	arrayMap2[$1] = twoDarray
                }
        |    LETTER '=' threeDarray
                {
                	arrayMap3[$1] = threeDarray
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
        |    '[' '[' '[' LETTER ']' ']' ']'
                { fmt.Println(arrayMap3[$4]) }
	|    number
	;

// array : (three dimensionals)
threeDarray:    '[' threeLplanty ']'
	|    '[' twoDarray ']'
	    { threeDarray = append(threeDarray, twoDarray); twoDarray = [][]int{} }
        ;

threeLplanty	:    twoDarray ','
                { threeDarray = append(threeDarray, twoDarray); twoDarray = [][]int{} }
	|    threeLplanty twoDarray
		{ threeDarray = append(threeDarray, twoDarray); twoDarray = [][]int{} }
	|    threeLplanty ',' twoDarray
		{ threeDarray = append(threeDarray, twoDarray); twoDarray = [][]int{} }
	;

// array : (two dimensionals)
twoDarray:    '[' twoLplanty ']'
	|    '[' array ']'
	    { twoDarray = append(twoDarray, array); array = []int{} }
        ;

twoLplanty	:    array ','
                { twoDarray = append(twoDarray, array); array = []int{} }
	|    twoLplanty array
		{ twoDarray = append(twoDarray, array); array = []int{} }
	|    twoLplanty ',' array
		{ twoDarray = append(twoDarray, array); array = []int{} }
	;

// array : (one dimensional)
array	:    '[' plenty ']'
		{ array = $2; $2 = []int{} }
	|    '[' number ']'
             	{ array = []int{$2} }
	;

plenty	:    plenty ',' number
		{ $$ = append($$, $3) }
	|    number ',' number
		{ $$ = append($$, $1, $3) }
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