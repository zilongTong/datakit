// Code generated by goyacc -o io/parser/gram_y.go io/parser/gram.y. DO NOT EDIT.

//line io/parser/gram.y:2
package parser

import __yyfmt__ "fmt"

//line io/parser/gram.y:2

import (
	"time"
)

//line io/parser/gram.y:9
type yySymType struct {
	yys   int
	node  Node
	nodes []Node

	item Item

	strings   []string
	float     float64
	duration  time.Duration
	timestamp time.Time
}

const EQ = 57346
const COLON = 57347
const SEMICOLON = 57348
const COMMA = 57349
const COMMENT = 57350
const DURATION = 57351
const EOF = 57352
const ERROR = 57353
const ID = 57354
const LEFT_BRACE = 57355
const LEFT_BRACKET = 57356
const LEFT_PAREN = 57357
const NUMBER = 57358
const RIGHT_BRACE = 57359
const RIGHT_BRACKET = 57360
const RIGHT_PAREN = 57361
const SPACE = 57362
const STRING = 57363
const QUOTED_STRING = 57364
const NAMESPACE = 57365
const DOT = 57366
const operatorsStart = 57367
const ADD = 57368
const DIV = 57369
const GTE = 57370
const GT = 57371
const LT = 57372
const LTE = 57373
const MOD = 57374
const MUL = 57375
const NEQ = 57376
const POW = 57377
const SUB = 57378
const operatorsEnd = 57379
const keywordsStart = 57380
const AS = 57381
const ASC = 57382
const AUTO = 57383
const BY = 57384
const DESC = 57385
const TRUE = 57386
const FALSE = 57387
const FILTER = 57388
const IDENTIFIER = 57389
const IN = 57390
const NOTIN = 57391
const AND = 57392
const LINK = 57393
const LIMIT = 57394
const SLIMIT = 57395
const OR = 57396
const NIL = 57397
const NULL = 57398
const OFFSET = 57399
const SOFFSET = 57400
const ORDER = 57401
const RE = 57402
const INT = 57403
const FLOAT = 57404
const POINT = 57405
const TIMEZONE = 57406
const WITH = 57407
const keywordsEnd = 57408
const startSymbolsStart = 57409
const START_STMTS = 57410
const START_BINARY_EXPRESSION = 57411
const START_FUNC_EXPRESSION = 57412
const START_WHERE_CONDITION = 57413
const startSymbolsEnd = 57414

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EQ",
	"COLON",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"DURATION",
	"EOF",
	"ERROR",
	"ID",
	"LEFT_BRACE",
	"LEFT_BRACKET",
	"LEFT_PAREN",
	"NUMBER",
	"RIGHT_BRACE",
	"RIGHT_BRACKET",
	"RIGHT_PAREN",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"NAMESPACE",
	"DOT",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"POW",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"AS",
	"ASC",
	"AUTO",
	"BY",
	"DESC",
	"TRUE",
	"FALSE",
	"FILTER",
	"IDENTIFIER",
	"IN",
	"NOTIN",
	"AND",
	"LINK",
	"LIMIT",
	"SLIMIT",
	"OR",
	"NIL",
	"NULL",
	"OFFSET",
	"SOFFSET",
	"ORDER",
	"RE",
	"INT",
	"FLOAT",
	"POINT",
	"TIMEZONE",
	"WITH",
	"keywordsEnd",
	"startSymbolsStart",
	"START_STMTS",
	"START_BINARY_EXPRESSION",
	"START_FUNC_EXPRESSION",
	"START_WHERE_CONDITION",
	"startSymbolsEnd",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line io/parser/gram.y:444

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 11,
	7, 52,
	17, 52,
	-2, 10,
	-1, 12,
	7, 53,
	17, 53,
	-2, 8,
	-1, 20,
	15, 70,
	-2, 12,
	-1, 21,
	15, 71,
	-2, 13,
	-1, 103,
	15, 70,
	-2, 12,
}

const yyPrivate = 57344

const yyLast = 363

var yyAct = [...]int{
	20, 13, 14, 29, 16, 105, 3, 21, 23, 99,
	59, 60, 56, 30, 67, 18, 62, 61, 63, 10,
	116, 45, 46, 66, 12, 65, 11, 53, 54, 46,
	56, 57, 64, 34, 53, 54, 117, 56, 31, 71,
	66, 34, 97, 112, 111, 70, 69, 73, 74, 75,
	76, 77, 78, 79, 80, 81, 82, 83, 84, 85,
	86, 12, 114, 11, 72, 91, 91, 94, 95, 68,
	103, 101, 92, 92, 113, 2, 118, 96, 88, 104,
	90, 93, 87, 7, 4, 1, 118, 127, 108, 108,
	107, 107, 106, 106, 110, 109, 109, 122, 118, 26,
	118, 19, 44, 108, 6, 107, 8, 106, 115, 120,
	109, 119, 43, 42, 22, 103, 101, 24, 123, 108,
	25, 107, 17, 125, 121, 108, 109, 107, 100, 106,
	126, 29, 109, 124, 15, 32, 5, 98, 9, 28,
	34, 30, 33, 0, 0, 40, 0, 0, 0, 0,
	0, 0, 39, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 0, 37, 38, 0, 31, 29, 0, 102,
	15, 32, 0, 0, 35, 36, 34, 30, 0, 27,
	0, 40, 0, 0, 0, 0, 0, 0, 39, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 0, 37,
	38, 0, 31, 29, 0, 0, 15, 32, 0, 0,
	35, 36, 34, 30, 0, 27, 0, 40, 0, 0,
	0, 0, 0, 0, 39, 0, 0, 41, 0, 0,
	58, 0, 0, 0, 0, 37, 38, 0, 31, 0,
	0, 0, 0, 0, 0, 89, 35, 36, 0, 0,
	0, 27, 45, 46, 47, 48, 51, 52, 53, 54,
	55, 56, 57, 29, 0, 0, 0, 32, 0, 0,
	0, 0, 34, 30, 0, 0, 49, 40, 0, 0,
	50, 0, 0, 0, 39, 0, 58, 41, 0, 0,
	0, 0, 0, 0, 0, 37, 38, 58, 31, 0,
	0, 0, 0, 0, 0, 0, 35, 36, 45, 46,
	47, 48, 51, 52, 53, 54, 55, 56, 57, 45,
	46, 47, 48, 51, 52, 53, 54, 55, 56, 57,
	58, 0, 49, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 49, 0, 0, 0, 0, 0, 0,
	0, 0, 45, 46, 47, 48, 51, 52, 53, 54,
	55, 56, 57,
}

var yyPact = [...]int{
	4, 74, 70, -1000, -1000, 100, -1000, 191, 70, 95,
	-1000, -1000, -1000, 282, -38, 191, -1000, -1000, 8, 1,
	-1, -10, -1000, -1000, -1000, -1000, -1000, 54, 31, -1000,
	-1000, 30, -1000, 23, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 191, 191, 191, 191, 191, 191,
	191, 191, 191, 191, 191, 191, 191, 191, 191, 68,
	64, 226, -1000, -1000, -9, -9, -9, -9, 20, 155,
	12, -1000, -1000, 2, -23, -5, -5, 326, 293, -5,
	-5, -23, -23, -5, -23, 2, -5, 251, 251, -1000,
	-1000, -1, -10, -1000, -1000, -1000, 25, 24, 55, -1000,
	-1000, 282, 251, 16, 17, 93, -1000, -1000, -1, -10,
	91, -1000, -1000, -1000, 155, 79, 119, -1000, 251, -1000,
	-1000, -1000, -1000, 282, 251, -1000, 69, -1000,
}

var yyPgo = [...]int{
	0, 142, 139, 0, 138, 137, 136, 104, 4, 5,
	7, 18, 1, 9, 15, 128, 16, 19, 122, 2,
	120, 8, 117, 114, 101, 99, 85,
}

var yyR1 = [...]int{
	0, 26, 26, 26, 6, 6, 12, 12, 12, 12,
	12, 12, 19, 19, 10, 10, 1, 1, 21, 22,
	22, 20, 20, 16, 14, 24, 24, 5, 5, 5,
	5, 9, 9, 9, 8, 8, 8, 8, 8, 8,
	25, 13, 13, 13, 15, 15, 7, 7, 4, 4,
	4, 4, 17, 17, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	2, 2, 23, 23, 18, 18, 3, 3, 3,
}

var yyR2 = [...]int{
	0, 2, 2, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 3, 4, 3, 3, 3, 2, 1,
	0, 3, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 5, 3, 0, 1, 3,
	2, 0, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 5,
	1, 1, 1, 2, 4, 4, 1, 1, 4,
}

var yyChk = [...]int{
	-1000, -26, 71, 2, 10, -6, -7, 13, 6, -4,
	-17, -11, -16, -12, -19, 15, -8, -18, -14, -24,
	-3, -10, -23, -21, -22, -20, -25, 60, -2, 12,
	22, 47, 16, -1, 21, 55, 56, 44, 45, 33,
	26, 36, -7, 17, 7, 26, 27, 28, 29, 50,
	54, 30, 31, 32, 33, 34, 35, 36, 4, 48,
	49, -12, -16, -11, 24, 24, 24, 24, 15, 15,
	15, 16, -17, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, 14, 14, 19,
	-14, -3, -10, -14, -3, -3, -21, 22, -5, -13,
	-15, -12, 14, -3, -21, -9, -8, -19, -3, -10,
	-9, 19, 19, 19, 7, -9, 4, 19, 7, 18,
	18, -13, 18, -12, 14, -8, -9, 18,
}

var yyDef = [...]int{
	0, -2, 47, 3, 2, 1, 4, 51, 47, 0,
	48, -2, -2, 0, 36, 0, 6, 7, 9, 11,
	-2, -2, 34, 35, 37, 38, 39, 0, 0, 76,
	77, 0, 72, 0, 18, 19, 20, 21, 22, 40,
	16, 17, 5, 46, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 8, 10, 0, 0, 0, 0, 0, 30,
	0, 73, 49, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 33, 33, 23,
	25, 70, 71, 26, 14, 15, 0, 0, 0, 29,
	41, 42, 33, -2, 0, 0, 32, 36, 12, 13,
	0, 74, 75, 24, 28, 0, 0, 78, 0, 68,
	69, 27, 43, 44, 33, 31, 0, 45,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line io/parser/gram.y:100
		{
			yylex.(*parser).parseResult = yyDollar[2].node
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:105
		{
			yylex.(*parser).unexpected("", "")
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:111
		{
			yyVAL.node = WhereConditions{yyDollar[1].node}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:115
		{
			if yyDollar[3].node != nil {
				arr := yyDollar[1].node.(WhereConditions)
				arr = append(arr, yyDollar[3].node)
				yyVAL.node = arr
			} else {
				yyVAL.node = yyDollar[1].node
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:131
		{
			yyVAL.node = &Identifier{Name: yyDollar[1].item.Val}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:135
		{
			yyVAL.node = yyDollar[1].node
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:141
		{
			yyVAL.node = &AttrExpr{
				Obj:  &Identifier{Name: yyDollar[1].item.Val},
				Attr: &Identifier{Name: yyDollar[3].item.Val},
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:148
		{
			yyVAL.node = &AttrExpr{
				Obj:  yyDollar[1].node.(*AttrExpr),
				Attr: &Identifier{Name: yyDollar[3].item.Val},
			}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:161
		{
			yyVAL.node = &StringLiteral{Val: yylex.(*parser).unquoteString(yyDollar[1].item.Val)}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:167
		{
			yyVAL.node = &NilLiteral{}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:171
		{
			yyVAL.node = &NilLiteral{}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:177
		{
			yyVAL.node = &BoolLiteral{Val: true}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:181
		{
			yyVAL.node = &BoolLiteral{Val: false}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:187
		{
			yyVAL.node = &ParenExpr{Param: yyDollar[2].node}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line io/parser/gram.y:193
		{
			yyVAL.node = yylex.(*parser).newFunc(yyDollar[1].item.Val, yyDollar[3].nodes)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:199
		{
			yyVAL.node = &CascadeFunctions{Funcs: []*FuncExpr{yyDollar[1].node.(*FuncExpr), yyDollar[3].node.(*FuncExpr)}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:203
		{
			fc := yyDollar[1].node.(*CascadeFunctions)
			fc.Funcs = append(fc.Funcs, yyDollar[3].node.(*FuncExpr))
			yyVAL.node = fc
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:211
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:216
		{
			yyVAL.nodes = []Node{yyDollar[1].node}
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line io/parser/gram.y:220
		{
			yyVAL.nodes = nil
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:226
		{
			nl := yyVAL.node.(NodeList)
			nl = append(nl, yyDollar[3].node)
			yyVAL.node = nl
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:232
		{
			yyVAL.node = NodeList{yyDollar[1].node}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
//line io/parser/gram.y:236
		{
			yyVAL.node = NodeList{}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:250
		{
			yyVAL.node = &Star{}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:258
		{
			yyVAL.node = getFuncArgList(yyDollar[2].node.(NodeList))
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:264
		{
			yyVAL.node = &FuncArg{ArgName: yyDollar[1].item.Val, ArgVal: yyDollar[3].node}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
//line io/parser/gram.y:268
		{
			yyVAL.node = &FuncArg{
				ArgName: yyDollar[1].item.Val,
				ArgVal:  getFuncArgList(yyDollar[4].node.(NodeList)),
			}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:277
		{
			yyVAL.node = yylex.(*parser).newWhereConditions(yyDollar[2].nodes)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
//line io/parser/gram.y:281
		{
			yyVAL.node = nil
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:288
		{
			yyVAL.nodes = []Node{yyDollar[1].node}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:292
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
//line io/parser/gram.y:297
		{
			yyVAL.nodes = nil
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:304
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:308
		{
			yyVAL.node = yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:312
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:318
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:324
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:330
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:336
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:342
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:348
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:353
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:358
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:364
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:369
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			yyVAL.node = bexpr
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line io/parser/gram.y:374
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
//line io/parser/gram.y:380
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
//line io/parser/gram.y:386
		{
			bexpr := yylex.(*parser).newBinExpr(yyDollar[1].node, yyDollar[4].node, yyDollar[2].item)
			bexpr.ReturnBool = true
			yyVAL.node = bexpr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:395
		{
			yyVAL.item = yyDollar[1].item
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:399
		{
			yyVAL.item = Item{Val: yyDollar[1].node.(*AttrExpr).String()}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:406
		{
			yyVAL.node = yylex.(*parser).number(yyDollar[1].item.Val)
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line io/parser/gram.y:410
		{
			num := yylex.(*parser).number(yyDollar[2].item.Val)
			switch yyDollar[1].item.Typ {
			case ADD: // pass
			case SUB:
				if num.IsInt {
					num.Int = -num.Int
				} else {
					num.Float = -num.Float
				}
			}
			yyVAL.node = num
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line io/parser/gram.y:426
		{
			yyVAL.node = &Regex{Regex: yyDollar[3].node.(*StringLiteral).Val}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
//line io/parser/gram.y:430
		{
			yyVAL.node = &Regex{Regex: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line io/parser/gram.y:437
		{
			yyVAL.item.Val = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
//line io/parser/gram.y:441
		{
			yyVAL.item.Val = yyDollar[3].node.(*StringLiteral).Val
		}
	}
	goto yystack /* stack new state and value */
}
