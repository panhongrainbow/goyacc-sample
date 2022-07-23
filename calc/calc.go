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
const THIRDCOMMA = 57348
const SECONDCOMMA = 57349
const FIRSTCOMMA = 57350
const UMINUS = 57351

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
	"','",
	"THIRDCOMMA",
	"SECONDCOMMA",
	"FIRSTCOMMA",
	"UMINUS",
	"'?'",
	"'('",
	"')'",
}

var CalcStatenames = [...]string{}

const CalcEofCode = 1
const CalcErrCode = 2
const CalcInitialStackSize = 16

//line calc.y:305
/*  start  of  programs  */
//line yacctab:1
var CalcExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const CalcPrivate = 57344

const CalcLast = 172

var CalcAct = [...]int8{
	33, 11, 67, 42, 52, 39, 13, 42, 27, 6,
	41, 38, 89, 26, 22, 21, 16, 17, 18, 19,
	20, 41, 66, 40, 11, 88, 57, 31, 5, 13,
	38, 56, 61, 51, 40, 28, 41, 38, 35, 14,
	12, 68, 58, 69, 9, 42, 59, 71, 70, 40,
	10, 15, 73, 72, 24, 38, 38, 7, 8, 75,
	76, 24, 14, 12, 64, 86, 38, 9, 24, 85,
	83, 57, 84, 54, 62, 38, 81, 82, 25, 79,
	76, 8, 35, 77, 57, 14, 12, 14, 12, 87,
	64, 4, 14, 80, 24, 12, 65, 37, 29, 37,
	30, 34, 64, 29, 37, 63, 80, 24, 43, 44,
	45, 46, 47, 48, 49, 50, 14, 12, 14, 12,
	14, 9, 14, 12, 24, 41, 14, 32, 2, 74,
	37, 29, 23, 36, 37, 8, 22, 21, 16, 17,
	18, 19, 20, 21, 16, 17, 18, 19, 20, 14,
	12, 12, 16, 17, 18, 19, 20, 14, 1, 3,
	29, 78, 24, 53, 18, 19, 20, 0, 55, 0,
	0, 60,
}

var CalcPact = [...]int16{
	35, -1000, 37, -1000, 130, 119, 70, 145, 112, 112,
	118, 17, -1000, -10, -1000, -1000, 112, 112, 112, 112,
	112, 112, 112, 58, -1000, 153, 28, 32, 157, 88,
	8, 102, 59, 121, -1000, 89, 81, -1000, 6, -14,
	122, -1000, 122, 154, 154, -1000, -1000, -1000, 144, 136,
	130, 70, 31, 36, 114, 83, -10, 17, -1000, -1000,
	-1000, -1000, 146, -1000, 63, 83, -1000, -1000, 121, 121,
	116, -1000, 55, -1000, 81, 31, -14, 56, 90, 53,
	49, -1000, -1000, 116, -1000, -1000, 9, -10, -4, -1000,
}

var CalcPgo = [...]uint8{
	0, 27, 91, 0, 9, 5, 4, 163, 159, 158,
	128,
}

var CalcR1 = [...]int8{
	0, 9, 9, 9, 8, 8, 8, 10, 10, 10,
	10, 10, 10, 4, 4, 4, 4, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	7, 7, 7, 6, 6, 6, 5, 5, 5, 5,
	5, 3, 3, 1, 1,
}

var CalcR2 = [...]int8{
	0, 0, 2, 1, 3, 3, 3, 1, 3, 1,
	3, 3, 3, 3, 1, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 1, 5, 7, 1,
	2, 3, 2, 2, 3, 2, 2, 3, 3, 3,
	3, 1, 2, 1, 2,
}

var CalcChk = [...]int16{
	-1000, -9, -10, -8, -2, -1, -4, 22, 23, 9,
	15, -3, 5, -5, 4, 14, 8, 9, 10, 11,
	12, 7, 6, 13, 5, 8, -5, -3, -1, 15,
	-2, -1, 15, -3, -2, -1, 15, 16, -3, -5,
	17, 4, 17, -2, -2, -2, -2, -2, -2, -2,
	-2, -4, -6, -7, 15, 15, -5, -3, 14, 14,
	14, 24, 15, 16, -1, 15, 16, 16, -3, -3,
	17, 16, 17, 16, 15, -6, -5, -1, 15, 16,
	-1, -5, -6, 15, 16, 16, 16, -5, 16, 16,
}

var CalcDef = [...]int8{
	1, -2, 0, 3, 7, 26, 9, 0, 0, 0,
	0, 29, 43, 14, 41, 2, 0, 0, 0, 0,
	0, 0, 0, 0, 44, 0, 0, 0, 0, 0,
	0, 26, 0, 29, 25, 0, 0, 36, 0, 0,
	0, 42, 0, 18, 19, 20, 21, 22, 23, 24,
	8, 10, 11, 12, 0, 0, 16, 0, 4, 5,
	6, 17, 0, 13, 0, 0, 37, 40, 38, 39,
	0, 35, 0, 32, 0, 30, 33, 0, 0, 0,
	0, 34, 31, 0, 15, 27, 0, 33, 0, 28,
}

var CalcTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 12, 7, 3,
	23, 24, 10, 8, 17, 9, 3, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 13, 3, 22, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 16, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6,
}

var CalcTok2 = [...]int8{
	2, 3, 4, 5, 18, 19, 20, 21,
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
//line calc.y:81
		{
		}
	case 3:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:83
		{
			setResult(Calclex, CalcDollar[1].test)
		}
	case 4:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:90
		{
			CalcVAL.test.node = "array"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].slice)
		}
	case 5:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:95
		{
			CalcVAL.test.node = "number"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].val)
		}
	case 6:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:100
		{
			CalcVAL.test.node = "variable"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].str)
		}
	case 8:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:115
		{
			regs[CalcDollar[1].str] = CalcDollar[3].val
		}
	case 9:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:119
		{
			fmt.Println(CalcDollar[1].slice)
		}
	case 10:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:123
		{
			arrayMap[CalcDollar[1].str] = CalcDollar[3].slice
		}
	case 11:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:127
		{
			arrayMap2[CalcDollar[1].str] = CalcDollar[3].twoDslice
		}
	case 12:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:131
		{
			arrayMap3[CalcDollar[1].str] = CalcDollar[3].threeDslice
		}
	case 13:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:138
		{
			CalcVAL.slice = arrayMap[CalcDollar[2].str]
		}
	case 14:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:142
		{
			CalcVAL.slice = CalcDollar[1].slice
		}
	case 15:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:146
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + arrayMap[CalcDollar[4].str][i]
			}
		}
	case 16:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:153
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + CalcDollar[3].slice[i]
			}
		}
	case 17:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:163
		{
			CalcVAL.val = CalcDollar[2].val
		}
	case 18:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:167
		{
			CalcVAL.val = CalcDollar[1].val + CalcDollar[3].val
		}
	case 19:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:171
		{
			CalcVAL.val = CalcDollar[1].val - CalcDollar[3].val
		}
	case 20:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:175
		{
			CalcVAL.val = CalcDollar[1].val * CalcDollar[3].val
		}
	case 21:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:179
		{
			CalcVAL.val = CalcDollar[1].val / CalcDollar[3].val
		}
	case 22:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:183
		{
			CalcVAL.val = CalcDollar[1].val % CalcDollar[3].val
		}
	case 23:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:187
		{
			CalcVAL.val = CalcDollar[1].val & CalcDollar[3].val
		}
	case 24:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:191
		{
			CalcVAL.val = CalcDollar[1].val | CalcDollar[3].val
		}
	case 25:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:195
		{
			CalcVAL.val = -CalcDollar[2].val
		}
	case 26:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:199
		{
			fmt.Println(regs[CalcDollar[1].str])
		}
	case 27:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:203
		{
			fmt.Println(arrayMap2[CalcDollar[3].str])
		}
	case 28:
		CalcDollar = CalcS[Calcpt-7 : Calcpt+1]
//line calc.y:207
		{
			fmt.Println(arrayMap3[CalcDollar[4].str])
		}
	case 30:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:216
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[2].twoDslice)
		}
	case 31:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:220
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[3].twoDslice)
		}
	case 32:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:224
		{
			CalcVAL.threeDslice = CalcDollar[1].threeDslice
		}
	case 33:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:232
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[2].slice)
		}
	case 34:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:236
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[3].slice)
		}
	case 35:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:240
		{
			CalcVAL.twoDslice = CalcDollar[1].twoDslice
		}
	case 36:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:249
		{
			CalcVAL.slice = []int{}
		}
	case 37:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:253
		{
			CalcVAL.slice = []int{CalcDollar[2].val}
		}
	case 38:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:257
		{
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[1].val, CalcDollar[3].val)
		}
	case 39:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:261
		{
			CalcVAL.slice = CalcDollar[1].slice
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[3].val)
		}
	case 40:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:266
		{
			CalcVAL.slice = CalcDollar[2].slice
		}
	case 41:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:275
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
//line calc.y:284
		{
			CalcVAL.val = base*CalcDollar[1].val + CalcDollar[2].val
		}
	case 43:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:296
		{
			CalcVAL.str = fmt.Sprintf("%c", CalcDollar[1].val)
		}
	case 44:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:300
		{
			CalcVAL.str = CalcDollar[1].str + fmt.Sprintf("%c", CalcDollar[2].val)
		}
	}
	goto Calcstack /* stack new state and value */
}
