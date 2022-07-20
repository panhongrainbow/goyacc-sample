// Code generated by goyacc -o calc.go -p Calc calc.y. DO NOT EDIT.

//line calc.y:8

package calc

import __yyfmt__ "fmt"

//line calc.y:9

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
	node  string
	value []interface{}
}

//line calc.y:40
type CalcSymType struct {
	yys         int
	val         int
	str         string
	slice       []int
	twoDslice   [][]int
	threeDslice [][][]int
	// iface []interface{}
	test Test
}

const DIGIT = 57346
const LETTER = 57347
const UMINUS = 57348

var CalcToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"DIGIT",
	"LETTER",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'='",
	"'\\n'",
	"'['",
	"']'",
	"UMINUS",
	"'?'",
	"'('",
	"')'",
	"','",
}

var CalcStatenames = [...]string{}

const CalcEofCode = 1
const CalcErrCode = 2
const CalcInitialStackSize = 16

//line calc.y:330
/*  start  of  programs  */
//line yacctab:1
var CalcExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const CalcPrivate = 57344

const CalcLast = 150

var CalcAct = [...]int8{
	33, 5, 11, 85, 69, 13, 51, 6, 30, 84,
	29, 36, 27, 38, 23, 22, 17, 18, 19, 20,
	21, 82, 73, 59, 58, 41, 4, 72, 62, 13,
	40, 57, 50, 31, 38, 32, 35, 71, 64, 80,
	37, 25, 70, 66, 42, 43, 44, 45, 46, 47,
	48, 49, 83, 41, 36, 12, 38, 74, 40, 38,
	68, 25, 14, 12, 56, 65, 76, 9, 64, 16,
	38, 26, 81, 10, 25, 78, 7, 8, 76, 79,
	38, 14, 12, 14, 12, 75, 9, 39, 9, 14,
	12, 12, 53, 25, 34, 39, 8, 60, 8, 23,
	22, 17, 18, 19, 20, 21, 14, 12, 14, 12,
	22, 17, 18, 19, 20, 21, 25, 77, 2, 67,
	14, 12, 14, 12, 19, 20, 21, 63, 14, 1,
	25, 37, 25, 31, 17, 18, 19, 20, 21, 61,
	24, 15, 3, 55, 52, 54, 0, 0, 0, 28,
}

var CalcPact = [...]int16{
	58, -1000, 55, -1000, 93, 127, 63, 118, 79, 79,
	116, 91, -1000, -1000, -1000, 37, -1000, 79, 79, 79,
	79, 79, 79, 79, 77, -1000, 49, 10, 9, 83,
	125, 124, 8, 88, 25, -1000, 111, 50, 91, -1000,
	124, -1000, 114, 114, -1000, -1000, -1000, 126, 103, 93,
	63, -1000, -1000, 104, 21, 6, 85, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 69, 86, 91, 102, -1000, -1000,
	18, -1000, 24, -1000, 56, 5, 36, 85, -1000, -1000,
	18, -1000, -1000, -7, -13, -1000,
}

var CalcPgo = [...]uint8{
	0, 0, 26, 2, 7, 4, 141, 6, 145, 144,
	143, 142, 129, 118,
}

var CalcR1 = [...]int8{
	0, 12, 12, 12, 11, 11, 11, 11, 13, 13,
	13, 13, 13, 13, 4, 4, 4, 4, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 10, 10, 9, 8, 8, 7, 6, 6, 6,
	5, 3, 3, 1, 1,
}

var CalcR2 = [...]int8{
	0, 0, 2, 1, 3, 3, 3, 3, 1, 3,
	1, 3, 3, 3, 3, 1, 5, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 1, 5, 7,
	1, 2, 3, 2, 2, 3, 2, 1, 2, 3,
	2, 1, 2, 1, 2,
}

var CalcChk = [...]int16{
	-1000, -12, -13, -11, -2, -1, -4, 18, 19, 9,
	15, -3, 5, -5, 4, -6, 14, 8, 9, 10,
	11, 12, 7, 6, 13, 5, 8, -5, -6, -3,
	-1, 15, -2, -1, 15, -2, -1, 15, -3, 4,
	21, 16, -2, -2, -2, -2, -2, -2, -2, -2,
	-4, -7, -9, 15, -8, -10, 15, -5, 14, 14,
	14, 14, 20, 16, -1, 15, -3, 15, -7, -5,
	21, 16, 21, 16, -1, 16, -1, 15, -5, -7,
	15, 16, 16, 16, 16, 16,
}

var CalcDef = [...]int8{
	1, -2, 0, 3, 8, 27, 10, 0, 0, 0,
	37, 30, 43, 15, 41, 0, 2, 0, 0, 0,
	0, 0, 0, 0, 0, 44, 0, 0, 0, 0,
	0, 37, 0, 27, 0, 26, 0, 0, 38, 42,
	0, 40, 19, 20, 21, 22, 23, 24, 25, 9,
	11, 12, 13, 37, 0, 0, 37, 17, 4, 5,
	6, 7, 18, 14, 0, 0, 39, 37, 31, 34,
	0, 36, 0, 33, 0, 0, 0, 37, 35, 32,
	0, 16, 28, 0, 0, 29,
}

var CalcTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 12, 7, 3,
	19, 20, 10, 8, 21, 9, 3, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 13, 3, 18, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 16, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6,
}

var CalcTok2 = [...]int8{
	2, 3, 4, 5, 17,
}

var CalcTok3 = [...]int8{
	0,
}

var CalcErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	CalcDebug        = 0
	CalcErrorVerbose = false
)

type CalcLexer interface {
	Lex(lval *CalcSymType) int
	Error(s string)
}

type CalcParser interface {
	Parse(CalcLexer) int
	Lookahead() int
}

type CalcParserImpl struct {
	lval  CalcSymType
	stack [CalcInitialStackSize]CalcSymType
	char  int
}

func (p *CalcParserImpl) Lookahead() int {
	return p.char
}

func CalcNewParser() CalcParser {
	return &CalcParserImpl{}
}

const CalcFlag = -1000

func CalcTokname(c int) string {
	if c >= 1 && c-1 < len(CalcToknames) {
		if CalcToknames[c-1] != "" {
			return CalcToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func CalcStatname(s int) string {
	if s >= 0 && s < len(CalcStatenames) {
		if CalcStatenames[s] != "" {
			return CalcStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func CalcErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !CalcErrorVerbose {
		return "syntax error"
	}

	for _, e := range CalcErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + CalcTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(CalcPact[state])
	for tok := TOKSTART; tok-1 < len(CalcToknames); tok++ {
		if n := base + tok; n >= 0 && n < CalcLast && int(CalcChk[int(CalcAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if CalcDef[state] == -2 {
		i := 0
		for CalcExca[i] != -1 || int(CalcExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; CalcExca[i] >= 0; i += 2 {
			tok := int(CalcExca[i])
			if tok < TOKSTART || CalcExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if CalcExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += CalcTokname(tok)
	}
	return res
}

func Calclex1(lex CalcLexer, lval *CalcSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(CalcTok1[0])
		goto out
	}
	if char < len(CalcTok1) {
		token = int(CalcTok1[char])
		goto out
	}
	if char >= CalcPrivate {
		if char < CalcPrivate+len(CalcTok2) {
			token = int(CalcTok2[char-CalcPrivate])
			goto out
		}
	}
	for i := 0; i < len(CalcTok3); i += 2 {
		token = int(CalcTok3[i+0])
		if token == char {
			token = int(CalcTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(CalcTok2[1]) /* unknown char */
	}
	if CalcDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", CalcTokname(token), uint(char))
	}
	return char, token
}

func CalcParse(Calclex CalcLexer) int {
	return CalcNewParser().Parse(Calclex)
}

func (Calcrcvr *CalcParserImpl) Parse(Calclex CalcLexer) int {
	var Calcn int
	var CalcVAL CalcSymType
	var CalcDollar []CalcSymType
	_ = CalcDollar // silence set and not used
	CalcS := Calcrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Calcstate := 0
	Calcrcvr.char = -1
	Calctoken := -1 // Calcrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Calcstate = -1
		Calcrcvr.char = -1
		Calctoken = -1
	}()
	Calcp := -1
	goto Calcstack

ret0:
	return 0

ret1:
	return 1

Calcstack:
	/* put a state and value onto the stack */
	if CalcDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", CalcTokname(Calctoken), CalcStatname(Calcstate))
	}

	Calcp++
	if Calcp >= len(CalcS) {
		nyys := make([]CalcSymType, len(CalcS)*2)
		copy(nyys, CalcS)
		CalcS = nyys
	}
	CalcS[Calcp] = CalcVAL
	CalcS[Calcp].yys = Calcstate

Calcnewstate:
	Calcn = int(CalcPact[Calcstate])
	if Calcn <= CalcFlag {
		goto Calcdefault /* simple state */
	}
	if Calcrcvr.char < 0 {
		Calcrcvr.char, Calctoken = Calclex1(Calclex, &Calcrcvr.lval)
	}
	Calcn += Calctoken
	if Calcn < 0 || Calcn >= CalcLast {
		goto Calcdefault
	}
	Calcn = int(CalcAct[Calcn])
	if int(CalcChk[Calcn]) == Calctoken { /* valid shift */
		Calcrcvr.char = -1
		Calctoken = -1
		CalcVAL = Calcrcvr.lval
		Calcstate = Calcn
		if Errflag > 0 {
			Errflag--
		}
		goto Calcstack
	}

Calcdefault:
	/* default state action */
	Calcn = int(CalcDef[Calcstate])
	if Calcn == -2 {
		if Calcrcvr.char < 0 {
			Calcrcvr.char, Calctoken = Calclex1(Calclex, &Calcrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if CalcExca[xi+0] == -1 && int(CalcExca[xi+1]) == Calcstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Calcn = int(CalcExca[xi+0])
			if Calcn < 0 || Calcn == Calctoken {
				break
			}
		}
		Calcn = int(CalcExca[xi+1])
		if Calcn < 0 {
			goto ret0
		}
	}
	if Calcn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Calclex.Error(CalcErrorMessage(Calcstate, Calctoken))
			Nerrs++
			if CalcDebug >= 1 {
				__yyfmt__.Printf("%s", CalcStatname(Calcstate))
				__yyfmt__.Printf(" saw %s\n", CalcTokname(Calctoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Calcp >= 0 {
				Calcn = int(CalcPact[CalcS[Calcp].yys]) + CalcErrCode
				if Calcn >= 0 && Calcn < CalcLast {
					Calcstate = int(CalcAct[Calcn]) /* simulate a shift of "error" */
					if int(CalcChk[Calcstate]) == CalcErrCode {
						goto Calcstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if CalcDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", CalcS[Calcp].yys)
				}
				Calcp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if CalcDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", CalcTokname(Calctoken))
			}
			if Calctoken == CalcEofCode {
				goto ret1
			}
			Calcrcvr.char = -1
			Calctoken = -1
			goto Calcnewstate /* try again in the same state */
		}
	}

	/* reduction by production Calcn */
	if CalcDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Calcn, CalcStatname(Calcstate))
	}

	Calcnt := Calcn
	Calcpt := Calcp
	_ = Calcpt // guard against "declared and not used"

	Calcp -= int(CalcR2[Calcn])
	// Calcp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Calcp+1 >= len(CalcS) {
		nyys := make([]CalcSymType, len(CalcS)*2)
		copy(nyys, CalcS)
		CalcS = nyys
	}
	CalcVAL = CalcS[Calcp+1]

	/* consult goto table to find next state */
	Calcn = int(CalcR1[Calcn])
	Calcg := int(CalcPgo[Calcn])
	Calcj := Calcg + CalcS[Calcp].yys + 1

	if Calcj >= CalcLast {
		Calcstate = int(CalcAct[Calcg])
	} else {
		Calcstate = int(CalcAct[Calcj])
		if int(CalcChk[Calcstate]) != -Calcn {
			Calcstate = int(CalcAct[Calcg])
		}
	}
	// dummy call; replaced with literal code
	switch Calcnt {

	case 2:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:77
		{
		}
	case 3:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:79
		{
			setResult(Calclex, CalcDollar[1].test)
		}
	case 4:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:86
		{
			CalcVAL.test.node = "array"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].slice)
		}
	case 5:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:91
		{
			CalcVAL.test.node = "array2"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].slice)
		}
	case 6:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:96
		{
			CalcVAL.test.node = "number"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].val)
		}
	case 7:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:101
		{
			CalcVAL.test.node = "variable"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].str)
		}
	case 9:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:116
		{
			regs[CalcDollar[1].str] = CalcDollar[3].val
		}
	case 10:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:120
		{
			fmt.Println(CalcDollar[1].slice)
		}
	case 11:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:124
		{
			arrayMap[CalcDollar[1].str] = CalcDollar[3].slice
		}
	case 12:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:128
		{
			arrayMap2[CalcDollar[1].str] = CalcDollar[3].twoDslice
		}
	case 13:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:132
		{
			arrayMap3[CalcDollar[1].str] = CalcDollar[3].threeDslice
		}
	case 14:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:139
		{
			CalcVAL.slice = arrayMap[CalcDollar[2].str]
		}
	case 15:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:143
		{
			CalcVAL.slice = CalcDollar[1].slice
		}
	case 16:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:147
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + arrayMap[CalcDollar[4].str][i]
			}
		}
	case 17:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:154
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + CalcDollar[3].slice[i]
			}
		}
	case 18:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:164
		{
			CalcVAL.val = CalcDollar[2].val
		}
	case 19:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:168
		{
			CalcVAL.val = CalcDollar[1].val + CalcDollar[3].val
		}
	case 20:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:172
		{
			CalcVAL.val = CalcDollar[1].val - CalcDollar[3].val
		}
	case 21:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:176
		{
			CalcVAL.val = CalcDollar[1].val * CalcDollar[3].val
		}
	case 22:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:180
		{
			CalcVAL.val = CalcDollar[1].val / CalcDollar[3].val
		}
	case 23:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:184
		{
			CalcVAL.val = CalcDollar[1].val % CalcDollar[3].val
		}
	case 24:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:188
		{
			CalcVAL.val = CalcDollar[1].val & CalcDollar[3].val
		}
	case 25:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:192
		{
			CalcVAL.val = CalcDollar[1].val | CalcDollar[3].val
		}
	case 26:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:196
		{
			CalcVAL.val = -CalcDollar[2].val
		}
	case 27:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:200
		{
			fmt.Println(regs[CalcDollar[1].str])
		}
	case 28:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:204
		{
			fmt.Println(arrayMap2[CalcDollar[3].str])
		}
	case 29:
		CalcDollar = CalcS[Calcpt-7 : Calcpt+1]
//line calc.y:208
		{
			fmt.Println(arrayMap3[CalcDollar[4].str])
		}
	case 31:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:217
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[2].twoDslice)
		}
	case 32:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:221
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[3].twoDslice)
		}
	case 33:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:226
		{
			CalcVAL.threeDslice = CalcDollar[1].threeDslice
		}
	case 34:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:234
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[2].slice)
		}
	case 35:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:238
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[3].slice)
		}
	case 36:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:244
		{
			CalcVAL.twoDslice = CalcDollar[1].twoDslice
		}
	case 37:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:277
		{
			CalcVAL.slice = []int{}
		}
	case 38:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:281
		{
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[2].val)
		}
	case 39:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:285
		{
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[3].val)
		}
	case 40:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:291
		{
			CalcVAL.slice = CalcDollar[1].slice
		}
	case 41:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:300
		{
			CalcVAL.val = CalcDollar[1].val
			if CalcDollar[1].val == 0 {
				base = 8
			} else {
				base = 10
			}
		}
	case 42:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:309
		{
			CalcVAL.val = base*CalcDollar[1].val + CalcDollar[2].val
		}
	case 43:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:321
		{
			CalcVAL.str = fmt.Sprintf("%c", CalcDollar[1].val)
		}
	case 44:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:325
		{
			CalcVAL.str = CalcDollar[1].str + fmt.Sprintf("%c", CalcDollar[2].val)
		}
	}
	goto Calcstack /* stack new state and value */
}
