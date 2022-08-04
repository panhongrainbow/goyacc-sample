// Copyright 2011 Bobby Powers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// based off of Appendix A from http://dinosaur.compilertools.net/yacc/

%{

package calc

import (
	"fmt"
)

// keep variables
var regs = make(map[string]interface{}, 26)

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
 iface interface{}
 test Test
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <str> variable
%type <val> number exprNumber
%type <slice> array
%type <twoDslice> twoDarray
%type <threeDslice> threeDarray
%type <iface> expr exprArray exprVariable
%type <test> unittest

// same for terminals
%token <val> DIGIT LETTER TokenASK
%type <str> ask

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left '='
%left '\n'
%right '['
%left ']'
%left ','
%left THIRDCOMMA
%left SECONDCOMMA
%left FIRSTCOMMA
%right UMINUS /* supplies precedence for unary minus */
%right '?'

%%
root
 : /* empty */
 | root '\n'
 | equal
 | print
 | unittest
 {
  setResult(Calclex, $1)
 }
 ;

equal
 : variable '=' expr
 {
  regs[$1] = $3
 }
 ;

print
 : expr
 {
  fmt.Println($1)
 }
 ;

unittest
 : ask array '?'
 {
  $$.node = "array"
  $$.value = append($$.value, $2)
 }
 | ask number '?'
 {
  $$.node = "number"
  $$.value = append($$.value, $2)
 }
 | ask variable '?'
 {
  $$.node = "variable"
  $$.value = append($$.value, $2)
 }
 ;

expr
 : exprNumber
 {
  setResult(Calclex, $1)
  $$ = $1
 }
 | exprArray
 {
  setResult(Calclex, $1)
  $$ = $1
 }
 | exprVariable
 {
  setResult(Calclex, $1)
  $$ = $1
 }
 // adapt others
 | twoDarray
 {
  $$ = $1
 }
 | threeDarray
 {
  $$ = $1
 }
 ;

exprArray
 : array
 {
  $$ = $1
 }
 | exprArray '+' '[' variable ']'
 {
  tmp := make([]int, len($1.([]int)))
  for i := 0; i < len($1.([]int)); i++ {
   tmp[i] = $1.([]int)[i]+regs[$4].([]int)[i]
  }
  $$ = tmp
 }
 | exprArray '+' array
 {
  tmp := make([]int, len($1.([]int)))
  for i := 0; i < len($1.([]int)); i++ {
   tmp[i] = $1.([]int)[i]+$3[i]
  }
  $$ = tmp
 }
 ;

exprNumber
 : number
 {
  $$ = $1
 }
 | '(' exprNumber ')'
 {
  $$ = $2
 }
 | exprNumber '+' exprNumber
 {
  $$ = $1 + $3
 }
 | exprNumber '-' exprNumber
 {
  $$ = $1 - $3
 }
 | exprNumber '*' exprNumber
 {
  $$ = $1 * $3
 }
 | exprNumber '/' exprNumber
 {
  $$ = $1 / $3
 }
 | exprNumber '%' exprNumber
 {
  $$ = $1 % $3
 }
 | exprNumber '&' exprNumber
 {
  $$ = $1 & $3
 }
 | exprNumber '|' exprNumber
 {
  $$ = $1 | $3
 }
 | '-' exprNumber %prec UMINUS
 {
  $$ = -$2
 }
 ;

exprVariable
 : variable
 {
  $$ = regs[$1]
 }
 ;

// array : (three dimensionals)
threeDarray
 : '[' twoDarray
 {
  $$ = append($$, $2)
 }
 | threeDarray ',' twoDarray %prec THIRDCOMMA
 {
  $$ = append($$, $3)
 }
 | threeDarray ']'
 {
  $$ = $1
 }
 ;

// array : (two dimensionals)
twoDarray
 : '[' array
 {
  $$ = append($$, $2)
 }
 | twoDarray ',' array %prec SECONDCOMMA
 {
  $$ = append($$, $3)
 }
 | twoDarray ']'
 {
  $$ = $1
 }
 ;

// array : (one dimensional)
// prefer to shift over reduce
array
 : '[' ']'
  {
   $$ = []int{}
  }
 | '[' number
 {
  $$ = append($$, $2)
 }
 | array ',' number %prec FIRSTCOMMA
 {
  $$ = append($$, $3)
 }
 | array ']'
 {
  $$ = $1
 }
 ;

// number
// test number not digit because of the rule, "number : DIGIT"
number
 : DIGIT
 {
  $$ = $1;
  if $1==0 {
   base = 8
  } else {
   base = 10
  }
 }
 | number DIGIT
 {
  $$ = base * $1 + $2
 }
 ;

// variables
/* Reference: https://installmd.com/c/113/go/convert-int-to-string
 Warning: If we use plain int to string conversion, the integer value is interpreted as a Unicode code point.
 And the resulting string will contain the character represented by the code point, encoded in UTF-8.
*/
variable
 : LETTER
 {
  $$ = fmt.Sprintf("%c", $1)
 }
 | variable LETTER
 {
  $$ = $1 + fmt.Sprintf("%c", $2)
 }
 ;

// keyword
ask: 'A' 'S' 'K' {};

%%      /*  start  of  programs  */