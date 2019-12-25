// Code generated by goyacc -p expr expr.y. DO NOT EDIT.

//line expr.y:4

package main

import __yyfmt__ "fmt"

//line expr.y:5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"

	"github.com/pschlump/godebug"
)

// "math/big"

//line expr.y:23
type exprSymType struct {
	yys  int
	tree *SyntaxTree
}

const NUM = 57346
const ID = 57347

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"';'",
	"'='",
	"'i'",
	"'d'",
	"NUM",
	"ID",
}
var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line expr.y:97

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

var line_no = 1

// The parser uses the type <prefix>Lex as a lexer. It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type exprLex struct {
	line []byte
	peek rune
}

// The parser calls this method to get each new token. This
// implementation returns operators and NUM.
func (x *exprLex) Lex(yylval *exprSymType) int {
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, yylval)
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j':
			return x.id(c, yylval)
		case '+', '-', '*', '/', '(', ')', ';', '=':
			return int(c)

		// Recognize Unicode multiplication and division
		// symbols, returning what the parser expects.
		case '×':
			return '*'
		case '÷':
			return '/'

		case ' ', '\t', '\r':
		case '\n':
			line_no++

		default:
			log.Printf("unrecognized character %q", c)
		}
	}
}

// Lex a number.
func (x *exprLex) num(c rune, yylval *exprSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
L:
	for {
		c = x.next()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			add(&b, c)
		default:
			break L
		}
	}
	if c != eof {
		x.peek = c
	}
	yylval.tree = NewAst(OpNum, nil, nil, line_no)
	yylval.tree.SValue = b.String()
	// xyzzy - convert to int for other things

	//	yylval.num = &big.Rat{}
	//	_, ok := yylval.num.SetString(b.String())
	//	if !ok {
	//		log.Printf("bad number %q", b.String())
	//		return eof
	//	}

	return NUM
}

// Lex an ID
func (x *exprLex) id(c rune, yylval *exprSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
	yylval.tree = NewAst(OpID, nil, nil, line_no)
	yylval.tree.SValue = b.String()
	return ID
}

// Return the next rune for the lexer.
func (x *exprLex) next() rune {
	if x.peek != eof {
		r := x.peek
		x.peek = eof
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return x.next()
	}
	return c
}

// The parser calls this method on a parse error.
func (x *exprLex) Error(s string) {
	log.Printf("parse error: %s", s)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		exprParse(&exprLex{line: line})
	}
}

//line yacctab:1
var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 41

var exprAct = [...]int{

	11, 8, 9, 17, 16, 13, 3, 14, 32, 5,
	6, 12, 21, 10, 15, 20, 22, 31, 8, 9,
	25, 26, 13, 1, 29, 30, 5, 6, 12, 2,
	13, 7, 27, 28, 23, 24, 12, 21, 18, 19,
	4,
}
var exprPact = [...]int{

	14, -1000, -4, 4, -1000, -11, -12, 34, -3, -3,
	28, -1000, -1000, -3, -3, -1000, -1000, -1000, 22, 22,
	-1000, -1000, -1000, 22, 22, 8, -2, 28, 28, -1000,
	-1000, -1000, -1000,
}
var exprPgo = [...]int{

	0, 6, 40, 31, 13, 0, 23,
}
var exprR1 = [...]int{

	0, 6, 6, 1, 1, 1, 2, 2, 2, 3,
	3, 3, 4, 4, 4, 5, 5, 5,
}
var exprR2 = [...]int{

	0, 4, 2, 1, 2, 2, 1, 2, 2, 1,
	3, 3, 1, 3, 3, 1, 1, 3,
}
var exprChk = [...]int{

	-1000, -6, 15, -1, -2, 12, 13, -3, 4, 5,
	-4, -5, 14, 8, 11, 10, 15, 15, 4, 5,
	-1, 15, -1, 6, 7, -1, -1, -4, -4, -5,
	-5, 9, 10,
}
var exprDef = [...]int{

	0, -2, 16, 0, 3, 0, 0, 6, 0, 0,
	9, 12, 15, 0, 0, 2, 4, 5, 0, 0,
	7, 16, 8, 0, 0, 0, 0, 10, 11, 13,
	14, 17, 1,
}
var exprTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 6, 4, 3, 5, 3, 7, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 10,
	3, 11, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 3, 3, 3, 3, 12,
}
var exprTok2 = [...]int{

	2, 3, 14, 15,
}
var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line expr.y:38
		{
			tmp := NewAst(OpAssign, exprDollar[1].tree, exprDollar[3].tree, line_no)
			fmt.Printf("Tree: %s\n", godebug.SVarI(tmp))
		}
	case 4:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:47
		{
			exprVAL.tree = NewAst(OpIncr, exprDollar[2].tree, nil, line_no)
		}
	case 5:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:51
		{
			exprVAL.tree = NewAst(OpDecr, exprDollar[2].tree, nil, line_no)
		}
	case 7:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:58
		{
			exprVAL.tree = exprDollar[2].tree
		}
	case 8:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:62
		{
			exprVAL.tree = NewAst(OpUMinus, exprDollar[2].tree, nil, line_no)
		}
	case 10:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:69
		{
			exprVAL.tree = NewAst(OpAdd, exprDollar[1].tree, exprDollar[3].tree, line_no)
		}
	case 11:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:73
		{
			exprVAL.tree = NewAst(OpSub, exprDollar[1].tree, exprDollar[3].tree, line_no)
		}
	case 13:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:80
		{
			exprVAL.tree = NewAst(OpMul, exprDollar[1].tree, exprDollar[3].tree, line_no)
		}
	case 14:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:84
		{
			exprVAL.tree = NewAst(OpDiv, exprDollar[1].tree, exprDollar[3].tree, line_no)
		}
	case 17:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:92
		{
			exprVAL.tree = exprDollar[2].tree
		}
	}
	goto exprstack /* stack new state and value */
}