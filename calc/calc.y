// Copyright 2011 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// based off of Appendix A from http://dinosaur.compilertools.net/yacc/

%{

package calc

import (
	"fmt"
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

// unit test
type Test struct {
    node string
    value []interface{}
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
	str string
	slice []int
	twoDslice [][]int
	threeDslice [][][]int
	// iface []interface{}
	test Test
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <str> variable
%type <val> expr number
%type <slice> exprArray array array2
%type <twoDslice> twoDarray twoDarray2
%type <threeDslice> threeDarray threeDarray2
%type <test> unittest

// same for terminals
%token <val> DIGIT LETTER

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left '='
%left '\n'
%right '['
%left ']'
%right UMINUS      /*  supplies  precedence  for  unary  minus  */
%right '?'

%%

list	: /* empty */
	| stat '\n'
	 {}
        | unittest
       	 {
       	  setResult(Calclex, $1)
       	 }
	;

unittest
 : '?' array2 '\n'
 {
  $$.node = "array2"
  $$.value = append($$.value, $2)
 }
 | '?' variable '\n'
 {
  $$.node = "variable"
  $$.value = append($$.value, $2)
 }
 | '?' LETTER '\n'
 {
  $$.node = "LETTER"
  $$.value = append($$.value, $2)
 }
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
  /* Reference: https://installmd.com/c/113/go/convert-int-to-string
     Warning: If we use plain int to string conversion, the integer value is interpreted as a Unicode code point.
     And the resulting string will contain the character represented by the code point, encoded in UTF-8.
  */
  $$ = fmt.Sprintf("%c", $1)
 }
 | variable LETTER
 {
  $$ = $1 + fmt.Sprintf("%c", $2)
 }
;

%%      /*  start  of  programs  */