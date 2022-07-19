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

//line calc.y:312
/*  start  of  programs  */
//line yacctab:1
var CalcExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const CalcPrivate = 57344

const CalcLast = 144

var CalcAct = [...]int8{
	32, 5, 11, 83, 50, 67, 13, 71, 29, 82,
	28, 35, 70, 37, 23, 22, 17, 18, 19, 20,
	21, 80, 69, 16, 57, 4, 6, 68, 60, 25,
	13, 39, 56, 37, 31, 34, 40, 62, 30, 26,
	81, 39, 64, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 25, 35, 25, 37, 72, 66, 37, 14,
	12, 12, 12, 79, 74, 73, 62, 78, 37, 36,
	75, 63, 14, 12, 76, 77, 74, 9, 37, 14,
	12, 14, 12, 10, 9, 25, 7, 8, 14, 12,
	52, 25, 65, 9, 8, 25, 61, 55, 38, 33,
	59, 14, 12, 8, 23, 22, 17, 18, 19, 20,
	21, 22, 17, 18, 19, 20, 21, 14, 12, 14,
	12, 38, 17, 18, 19, 20, 21, 15, 36, 25,
	30, 58, 19, 20, 21, 27, 14, 24, 2, 1,
	3, 54, 51, 53,
}

var CalcPact = [...]int16{
	68, -1000, 9, -1000, 98, 124, 31, 115, 84, 84,
	113, 94, -1000, -1000, -1000, 20, -1000, 84, 84, 84,
	84, 84, 84, 84, 75, -1000, 82, 10, 117, 86,
	132, 8, 90, 54, -1000, 80, 56, 94, -1000, 132,
	-1000, 122, 122, -1000, -1000, -1000, 114, 104, 98, 31,
	-1000, -1000, 77, 6, -9, 97, -1000, -1000, -1000, -1000,
	-1000, -1000, 49, 57, 94, 55, -1000, -1000, 23, -1000,
	52, -1000, 47, 5, 24, 97, -1000, -1000, 23, -1000,
	-1000, -7, -13, -1000,
}

var CalcPgo = [...]uint8{
	0, 0, 25, 2, 26, 5, 127, 4, 143, 142,
	141, 140, 139, 138,
}

var CalcR1 = [...]int8{
	0, 12, 12, 12, 11, 11, 11, 13, 13, 13,
	13, 13, 13, 4, 4, 4, 4, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	10, 10, 9, 8, 8, 7, 6, 6, 5, 3,
	3, 1, 1,
}

var CalcR2 = [...]int8{
	0, 0, 2, 1, 3, 3, 3, 1, 3, 1,
	3, 3, 3, 3, 1, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 1, 5, 7, 1,
	2, 3, 2, 2, 3, 2, 2, 3, 2, 1,
	2, 1, 2,
}

var CalcChk = [...]int16{
	-1000, -12, -13, -11, -2, -1, -4, 18, 19, 9,
	15, -3, 5, -5, 4, -6, 14, 8, 9, 10,
	11, 12, 7, 6, 13, 5, 8, -6, -3, -1,
	15, -2, -1, 15, -2, -1, 15, -3, 4, 21,
	16, -2, -2, -2, -2, -2, -2, -2, -2, -4,
	-7, -9, 15, -8, -10, 15, -5, 14, 14, 14,
	20, 16, -1, 15, -3, 15, -7, -5, 21, 16,
	21, 16, -1, 16, -1, 15, -5, -7, 15, 16,
	16, 16, 16, 16,
}

var CalcDef = [...]int8{
	1, -2, 0, 3, 7, 26, 9, 0, 0, 0,
	0, 29, 41, 14, 39, 0, 2, 0, 0, 0,
	0, 0, 0, 0, 0, 42, 0, 0, 0, 0,
	0, 0, 26, 0, 25, 0, 0, 36, 40, 0,
	38, 18, 19, 20, 21, 22, 23, 24, 8, 10,
	11, 12, 0, 0, 0, 0, 16, 4, 5, 6,
	17, 13, 0, 0, 37, 0, 30, 33, 0, 35,
	0, 32, 0, 0, 0, 0, 34, 31, 0, 15,
	27, 0, 0, 28,
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
			CalcVAL.test.node = "array2"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].slice)
		}
	case 5:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:91
		{
			CalcVAL.test.node = "number"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].val)
		}
	case 6:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:96
		{
			CalcVAL.test.node = "variable"
			CalcVAL.test.value = append(CalcVAL.test.value, CalcDollar[2].str)
		}
	case 8:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:110
		{
			regs[CalcDollar[1].str] = CalcDollar[3].val
		}
	case 10:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:115
		{
			arrayMap[CalcDollar[1].str] = CalcDollar[3].slice
		}
	case 11:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:119
		{
			arrayMap2[CalcDollar[1].str] = CalcDollar[3].twoDslice
		}
	case 12:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:123
		{
			arrayMap3[CalcDollar[1].str] = CalcDollar[3].threeDslice
		}
	case 13:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:129
		{
			fmt.Println(arrayMap[CalcDollar[2].str])
		}
	case 14:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:131
		{
			CalcVAL.slice = CalcDollar[1].slice
		}
	case 15:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:133
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + arrayMap[CalcDollar[4].str][i]
			}
		}
	case 16:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:140
		{
			CalcVAL.slice = make([]int, len(CalcDollar[1].slice))
			for i := 0; i < len(CalcDollar[1].slice); i++ {
				CalcVAL.slice[i] = CalcDollar[1].slice[i] + CalcDollar[3].slice[i]
			}
		}
	case 17:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:150
		{
			CalcVAL.val = CalcDollar[2].val
		}
	case 18:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:154
		{
			CalcVAL.val = CalcDollar[1].val + CalcDollar[3].val
		}
	case 19:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:158
		{
			CalcVAL.val = CalcDollar[1].val - CalcDollar[3].val
		}
	case 20:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:162
		{
			CalcVAL.val = CalcDollar[1].val * CalcDollar[3].val
		}
	case 21:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:166
		{
			CalcVAL.val = CalcDollar[1].val / CalcDollar[3].val
		}
	case 22:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:170
		{
			CalcVAL.val = CalcDollar[1].val % CalcDollar[3].val
		}
	case 23:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:174
		{
			CalcVAL.val = CalcDollar[1].val & CalcDollar[3].val
		}
	case 24:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:178
		{
			CalcVAL.val = CalcDollar[1].val | CalcDollar[3].val
		}
	case 25:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:182
		{
			CalcVAL.val = -CalcDollar[2].val
		}
	case 26:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:186
		{
			fmt.Println(regs[CalcDollar[1].str])
		}
	case 27:
		CalcDollar = CalcS[Calcpt-5 : Calcpt+1]
//line calc.y:190
		{
			fmt.Println(arrayMap2[CalcDollar[3].str])
		}
	case 28:
		CalcDollar = CalcS[Calcpt-7 : Calcpt+1]
//line calc.y:194
		{
			fmt.Println(arrayMap3[CalcDollar[4].str])
		}
	case 30:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:203
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[2].twoDslice)
		}
	case 31:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:207
		{
			CalcVAL.threeDslice = append(CalcVAL.threeDslice, CalcDollar[3].twoDslice)
		}
	case 32:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:212
		{
			CalcVAL.threeDslice = CalcDollar[1].threeDslice
		}
	case 33:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:220
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[2].slice)
		}
	case 34:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:224
		{
			CalcVAL.twoDslice = append(CalcVAL.twoDslice, CalcDollar[3].slice)
		}
	case 35:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:230
		{
			CalcVAL.twoDslice = CalcDollar[1].twoDslice
		}
	case 36:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:263
		{
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[2].val)
		}
	case 37:
		CalcDollar = CalcS[Calcpt-3 : Calcpt+1]
//line calc.y:267
		{
			CalcVAL.slice = append(CalcVAL.slice, CalcDollar[3].val)
		}
	case 38:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:273
		{
			CalcVAL.slice = CalcDollar[1].slice
		}
	case 39:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:282
		{
			CalcVAL.val = CalcDollar[1].val
			if CalcDollar[1].val == 0 {
				base = 8
			} else {
				base = 10
			}
		}
	case 40:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:291
		{
			CalcVAL.val = base*CalcDollar[1].val + CalcDollar[2].val
		}
	case 41:
		CalcDollar = CalcS[Calcpt-1 : Calcpt+1]
//line calc.y:303
		{
			CalcVAL.str = fmt.Sprintf("%c", CalcDollar[1].val)
		}
	case 42:
		CalcDollar = CalcS[Calcpt-2 : Calcpt+1]
//line calc.y:307
		{
			CalcVAL.str = CalcDollar[1].str + fmt.Sprintf("%c", CalcDollar[2].val)
		}
	}
	goto Calcstack /* stack new state and value */
}
