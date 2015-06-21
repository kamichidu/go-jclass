//line ./parser/fd/parser.go.y:2
package fd

import __yyfmt__ "fmt"

//line ./parser/fd/parser.go.y:2
import ()

type FDToken struct {
	Id   int
	Text string
	Pos  int
}

//line ./parser/fd/parser.go.y:24
type fdSymType struct {
	yys   int
	res   *FDInfo
	token FDToken
}

const CLASS_NAME = 57346

var fdToknames = []string{
	"'B'",
	"'C'",
	"'D'",
	"'F'",
	"'I'",
	"'J'",
	"'S'",
	"'Z'",
	"'L'",
	"';'",
	"'['",
	"CLASS_NAME",
}
var fdStatenames = []string{}

const fdEofCode = 1
const fdErrCode = 2
const fdMaxDepth = 200

//line ./parser/fd/parser.go.y:106

//line yacctab:1
var fdExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const fdNprod = 16
const fdPrivate = 57344

var fdTokenNames []string
var fdStates []string

const fdLast = 26

var fdAct = []int{

	6, 7, 8, 9, 10, 11, 12, 13, 14, 2,
	15, 19, 16, 17, 5, 4, 3, 1, 0, 0,
	0, 0, 0, 0, 0, 18,
}
var fdPact = []int{

	-4, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -3, -4, -2, -1000, -1000, -1000,
}
var fdPgo = []int{

	0, 17, 9, 16, 15, 14, 13,
}
var fdR1 = []int{

	0, 1, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 5, 6,
}
var fdR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 2, 1,
}
var fdChk = []int{

	-1000, -1, -2, -3, -4, -5, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 14, 15, -6, -2, 13,
}
var fdDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 0, 0, 0, 14, 15, 13,
}
var fdTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	3, 3, 3, 3, 3, 3, 4, 5, 6, 3,
	7, 3, 3, 8, 9, 3, 12, 3, 3, 3,
	3, 3, 3, 10, 3, 3, 3, 3, 3, 3,
	11, 14,
}
var fdTok2 = []int{

	2, 3, 15,
}
var fdTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var fdDebug = 0

type fdLexer interface {
	Lex(lval *fdSymType) int
	Error(s string)
}

const fdFlag = -1000

func fdTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(fdToknames) {
		if fdToknames[c-4] != "" {
			return fdToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func fdStatname(s int) string {
	if s >= 0 && s < len(fdStatenames) {
		if fdStatenames[s] != "" {
			return fdStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func fdlex1(lex fdLexer, lval *fdSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = fdTok1[0]
		goto out
	}
	if char < len(fdTok1) {
		c = fdTok1[char]
		goto out
	}
	if char >= fdPrivate {
		if char < fdPrivate+len(fdTok2) {
			c = fdTok2[char-fdPrivate]
			goto out
		}
	}
	for i := 0; i < len(fdTok3); i += 2 {
		c = fdTok3[i+0]
		if c == char {
			c = fdTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = fdTok2[1] /* unknown char */
	}
	if fdDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", fdTokname(c), uint(char))
	}
	return c
}

func fdParse(fdlex fdLexer) int {
	var fdn int
	var fdlval fdSymType
	var fdVAL fdSymType
	fdS := make([]fdSymType, fdMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	fdstate := 0
	fdchar := -1
	fdp := -1
	goto fdstack

ret0:
	return 0

ret1:
	return 1

fdstack:
	/* put a state and value onto the stack */
	if fdDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", fdTokname(fdchar), fdStatname(fdstate))
	}

	fdp++
	if fdp >= len(fdS) {
		nyys := make([]fdSymType, len(fdS)*2)
		copy(nyys, fdS)
		fdS = nyys
	}
	fdS[fdp] = fdVAL
	fdS[fdp].yys = fdstate

fdnewstate:
	fdn = fdPact[fdstate]
	if fdn <= fdFlag {
		goto fddefault /* simple state */
	}
	if fdchar < 0 {
		fdchar = fdlex1(fdlex, &fdlval)
	}
	fdn += fdchar
	if fdn < 0 || fdn >= fdLast {
		goto fddefault
	}
	fdn = fdAct[fdn]
	if fdChk[fdn] == fdchar { /* valid shift */
		fdchar = -1
		fdVAL = fdlval
		fdstate = fdn
		if Errflag > 0 {
			Errflag--
		}
		goto fdstack
	}

fddefault:
	/* default state action */
	fdn = fdDef[fdstate]
	if fdn == -2 {
		if fdchar < 0 {
			fdchar = fdlex1(fdlex, &fdlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if fdExca[xi+0] == -1 && fdExca[xi+1] == fdstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			fdn = fdExca[xi+0]
			if fdn < 0 || fdn == fdchar {
				break
			}
		}
		fdn = fdExca[xi+1]
		if fdn < 0 {
			goto ret0
		}
	}
	if fdn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			fdlex.Error("syntax error")
			Nerrs++
			if fdDebug >= 1 {
				__yyfmt__.Printf("%s", fdStatname(fdstate))
				__yyfmt__.Printf(" saw %s\n", fdTokname(fdchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for fdp >= 0 {
				fdn = fdPact[fdS[fdp].yys] + fdErrCode
				if fdn >= 0 && fdn < fdLast {
					fdstate = fdAct[fdn] /* simulate a shift of "error" */
					if fdChk[fdstate] == fdErrCode {
						goto fdstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if fdDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", fdS[fdp].yys)
				}
				fdp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if fdDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", fdTokname(fdchar))
			}
			if fdchar == fdEofCode {
				goto ret1
			}
			fdchar = -1
			goto fdnewstate /* try again in the same state */
		}
	}

	/* reduction by production fdn */
	if fdDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", fdn, fdStatname(fdstate))
	}

	fdnt := fdn
	fdpt := fdp
	_ = fdpt // guard against "declared and not used"

	fdp -= fdR2[fdn]
	fdVAL = fdS[fdp+1]

	/* consult goto table to find next state */
	fdn = fdR1[fdn]
	fdg := fdPgo[fdn]
	fdj := fdg + fdS[fdp].yys + 1

	if fdj >= fdLast {
		fdstate = fdAct[fdg]
	} else {
		fdstate = fdAct[fdj]
		if fdChk[fdstate] != -fdn {
			fdstate = fdAct[fdg]
		}
	}
	// dummy call; replaced with literal code
	switch fdnt {

	case 1:
		//line ./parser/fd/parser.go.y:33
		{
			fdVAL.res = fdS[fdpt-0].res
			if lexer, ok := fdlex.(*FDLexer); ok {
				lexer.Result = fdVAL.res
			}
		}
	case 2:
		fdVAL.res = fdS[fdpt-0].res
	case 3:
		fdVAL.res = fdS[fdpt-0].res
	case 4:
		fdVAL.res = fdS[fdpt-0].res
	case 5:
		//line ./parser/fd/parser.go.y:49
		{
			fdVAL.res = NewPrimitiveType("byte")
		}
	case 6:
		//line ./parser/fd/parser.go.y:53
		{
			fdVAL.res = NewPrimitiveType("char")
		}
	case 7:
		//line ./parser/fd/parser.go.y:57
		{
			fdVAL.res = NewPrimitiveType("double")
		}
	case 8:
		//line ./parser/fd/parser.go.y:61
		{
			fdVAL.res = NewPrimitiveType("float")
		}
	case 9:
		//line ./parser/fd/parser.go.y:65
		{
			fdVAL.res = NewPrimitiveType("int")
		}
	case 10:
		//line ./parser/fd/parser.go.y:69
		{
			fdVAL.res = NewPrimitiveType("long")
		}
	case 11:
		//line ./parser/fd/parser.go.y:73
		{
			fdVAL.res = NewPrimitiveType("short")
		}
	case 12:
		//line ./parser/fd/parser.go.y:77
		{
			fdVAL.res = NewPrimitiveType("boolean")
		}
	case 13:
		//line ./parser/fd/parser.go.y:84
		{
			fdVAL.res = NewReferenceType(fdS[fdpt-1].token.Text)
		}
	case 14:
		//line ./parser/fd/parser.go.y:91
		{
			if fdS[fdpt-0].res.PrimitiveType || fdS[fdpt-0].res.ReferenceType {
				fdVAL.res = NewArrayType(fdS[fdpt-0].res, 1)
			} else if fdS[fdpt-0].res.ArrayType {
				fdVAL.res = NewArrayType(fdS[fdpt-0].res.ComponentType, fdS[fdpt-0].res.Dims+1)
			} else {
				panic("??? Siranai Kata da")
			}
		}
	case 15:
		fdVAL.res = fdS[fdpt-0].res
	}
	goto fdstack /* stack new state and value */
}
