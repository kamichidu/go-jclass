//line ./parser/md/parser.go.y:2
package md

import __yyfmt__ "fmt"

//line ./parser/md/parser.go.y:2
import (
	"github.com/kamichidu/go-jclass/parser/fd"
)

type MDToken struct {
	Id   int
	Text string
	Pos  int
}

//line ./parser/md/parser.go.y:27
type mdSymType struct {
	yys    int
	jtype  *fd.FDInfo
	params []*fd.FDInfo
	token  MDToken
}

const CLASS_NAME = 57346

var mdToknames = []string{
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
	"'('",
	"')'",
	"CLASS_NAME",
}
var mdStatenames = []string{}

const mdEofCode = 1
const mdErrCode = 2
const mdMaxDepth = 200

//line ./parser/md/parser.go.y:144

//line yacctab:1
var mdExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const mdNprod = 22
const mdPrivate = 57344

var mdTokenNames []string
var mdStates []string

const mdLast = 59

var mdAct = []int{

	11, 12, 13, 14, 15, 16, 17, 18, 19, 24,
	20, 3, 2, 28, 7, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 5, 20, 1, 21, 11, 12,
	13, 14, 15, 16, 17, 18, 19, 6, 20, 25,
	10, 9, 23, 8, 22, 4, 27, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 26,
}
var mdPact = []int{

	-3, -1000, -5, -4, 11, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -8,
	24, -4, -1000, -1000, 0, -1000, -1000, -1000, -1000,
}
var mdPgo = []int{

	0, 45, 44, 24, 37, 43, 41, 40, 39, 26,
}
var mdR1 = []int{

	0, 9, 9, 1, 1, 2, 3, 3, 4, 4,
	4, 5, 5, 5, 5, 5, 5, 5, 5, 6,
	7, 8,
}
var mdR2 = []int{

	0, 3, 4, 0, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	2, 1,
}
var mdChk = []int{

	-1000, -9, 15, 16, -1, -3, -4, 18, -5, -6,
	-7, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	14, 16, -2, -4, 17, -8, -4, -3, 13,
}
var mdDef = []int{

	0, -2, 3, 0, 0, 1, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 0,
	0, 0, 4, 5, 0, 20, 21, 2, 19,
}
var mdTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	3, 3, 3, 3, 3, 3, 4, 5, 6, 3,
	7, 3, 3, 8, 9, 3, 12, 3, 3, 3,
	3, 3, 3, 10, 3, 3, 18, 3, 3, 3,
	11, 14,
}
var mdTok2 = []int{

	2, 3, 17,
}
var mdTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var mdDebug = 0

type mdLexer interface {
	Lex(lval *mdSymType) int
	Error(s string)
}

const mdFlag = -1000

func mdTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(mdToknames) {
		if mdToknames[c-4] != "" {
			return mdToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func mdStatname(s int) string {
	if s >= 0 && s < len(mdStatenames) {
		if mdStatenames[s] != "" {
			return mdStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func mdlex1(lex mdLexer, lval *mdSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = mdTok1[0]
		goto out
	}
	if char < len(mdTok1) {
		c = mdTok1[char]
		goto out
	}
	if char >= mdPrivate {
		if char < mdPrivate+len(mdTok2) {
			c = mdTok2[char-mdPrivate]
			goto out
		}
	}
	for i := 0; i < len(mdTok3); i += 2 {
		c = mdTok3[i+0]
		if c == char {
			c = mdTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = mdTok2[1] /* unknown char */
	}
	if mdDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", mdTokname(c), uint(char))
	}
	return c
}

func mdParse(mdlex mdLexer) int {
	var mdn int
	var mdlval mdSymType
	var mdVAL mdSymType
	mdS := make([]mdSymType, mdMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	mdstate := 0
	mdchar := -1
	mdp := -1
	goto mdstack

ret0:
	return 0

ret1:
	return 1

mdstack:
	/* put a state and value onto the stack */
	if mdDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", mdTokname(mdchar), mdStatname(mdstate))
	}

	mdp++
	if mdp >= len(mdS) {
		nyys := make([]mdSymType, len(mdS)*2)
		copy(nyys, mdS)
		mdS = nyys
	}
	mdS[mdp] = mdVAL
	mdS[mdp].yys = mdstate

mdnewstate:
	mdn = mdPact[mdstate]
	if mdn <= mdFlag {
		goto mddefault /* simple state */
	}
	if mdchar < 0 {
		mdchar = mdlex1(mdlex, &mdlval)
	}
	mdn += mdchar
	if mdn < 0 || mdn >= mdLast {
		goto mddefault
	}
	mdn = mdAct[mdn]
	if mdChk[mdn] == mdchar { /* valid shift */
		mdchar = -1
		mdVAL = mdlval
		mdstate = mdn
		if Errflag > 0 {
			Errflag--
		}
		goto mdstack
	}

mddefault:
	/* default state action */
	mdn = mdDef[mdstate]
	if mdn == -2 {
		if mdchar < 0 {
			mdchar = mdlex1(mdlex, &mdlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if mdExca[xi+0] == -1 && mdExca[xi+1] == mdstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			mdn = mdExca[xi+0]
			if mdn < 0 || mdn == mdchar {
				break
			}
		}
		mdn = mdExca[xi+1]
		if mdn < 0 {
			goto ret0
		}
	}
	if mdn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			mdlex.Error("syntax error")
			Nerrs++
			if mdDebug >= 1 {
				__yyfmt__.Printf("%s", mdStatname(mdstate))
				__yyfmt__.Printf(" saw %s\n", mdTokname(mdchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for mdp >= 0 {
				mdn = mdPact[mdS[mdp].yys] + mdErrCode
				if mdn >= 0 && mdn < mdLast {
					mdstate = mdAct[mdn] /* simulate a shift of "error" */
					if mdChk[mdstate] == mdErrCode {
						goto mdstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if mdDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", mdS[mdp].yys)
				}
				mdp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if mdDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", mdTokname(mdchar))
			}
			if mdchar == mdEofCode {
				goto ret1
			}
			mdchar = -1
			goto mdnewstate /* try again in the same state */
		}
	}

	/* reduction by production mdn */
	if mdDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", mdn, mdStatname(mdstate))
	}

	mdnt := mdn
	mdpt := mdp
	_ = mdpt // guard against "declared and not used"

	mdp -= mdR2[mdn]
	mdVAL = mdS[mdp+1]

	/* consult goto table to find next state */
	mdn = mdR1[mdn]
	mdg := mdPgo[mdn]
	mdj := mdg + mdS[mdp].yys + 1

	if mdj >= mdLast {
		mdstate = mdAct[mdg]
	} else {
		mdstate = mdAct[mdj]
		if mdChk[mdstate] != -mdn {
			mdstate = mdAct[mdg]
		}
	}
	// dummy call; replaced with literal code
	switch mdnt {

	case 1:
		//line ./parser/md/parser.go.y:37
		{
			if l, ok := mdlex.(*MDLexer); ok {
				l.Result = &MDInfo{
					ReturnType:     mdS[mdpt-0].jtype,
					ParameterTypes: make([]*fd.FDInfo, 0),
				}
			}
		}
	case 2:
		//line ./parser/md/parser.go.y:46
		{
			if l, ok := mdlex.(*MDLexer); ok {
				l.Result = &MDInfo{
					ReturnType:     mdS[mdpt-0].jtype,
					ParameterTypes: mdS[mdpt-2].params,
				}
			}
		}
	case 3:
		//line ./parser/md/parser.go.y:58
		{
			mdVAL.params = make([]*fd.FDInfo, 0)
		}
	case 4:
		//line ./parser/md/parser.go.y:62
		{
			mdVAL.params = append(mdS[mdpt-1].params, mdS[mdpt-0].jtype)
		}
	case 5:
		mdVAL.jtype = mdS[mdpt-0].jtype
	case 6:
		mdVAL.jtype = mdS[mdpt-0].jtype
	case 7:
		//line ./parser/md/parser.go.y:74
		{
			mdVAL.jtype = fd.NewPrimitiveType("void")
		}
	case 8:
		mdVAL.jtype = mdS[mdpt-0].jtype
	case 9:
		mdVAL.jtype = mdS[mdpt-0].jtype
	case 10:
		mdVAL.jtype = mdS[mdpt-0].jtype
	case 11:
		//line ./parser/md/parser.go.y:87
		{
			mdVAL.jtype = fd.NewPrimitiveType("byte")
		}
	case 12:
		//line ./parser/md/parser.go.y:91
		{
			mdVAL.jtype = fd.NewPrimitiveType("char")
		}
	case 13:
		//line ./parser/md/parser.go.y:95
		{
			mdVAL.jtype = fd.NewPrimitiveType("double")
		}
	case 14:
		//line ./parser/md/parser.go.y:99
		{
			mdVAL.jtype = fd.NewPrimitiveType("float")
		}
	case 15:
		//line ./parser/md/parser.go.y:103
		{
			mdVAL.jtype = fd.NewPrimitiveType("int")
		}
	case 16:
		//line ./parser/md/parser.go.y:107
		{
			mdVAL.jtype = fd.NewPrimitiveType("long")
		}
	case 17:
		//line ./parser/md/parser.go.y:111
		{
			mdVAL.jtype = fd.NewPrimitiveType("short")
		}
	case 18:
		//line ./parser/md/parser.go.y:115
		{
			mdVAL.jtype = fd.NewPrimitiveType("boolean")
		}
	case 19:
		//line ./parser/md/parser.go.y:122
		{
			mdVAL.jtype = fd.NewReferenceType(mdS[mdpt-1].token.Text)
		}
	case 20:
		//line ./parser/md/parser.go.y:129
		{
			if mdS[mdpt-0].jtype.PrimitiveType || mdS[mdpt-0].jtype.ReferenceType {
				mdVAL.jtype = fd.NewArrayType(mdS[mdpt-0].jtype, 1)
			} else if mdS[mdpt-0].jtype.ArrayType {
				mdVAL.jtype = fd.NewArrayType(mdS[mdpt-0].jtype.ComponentType, mdS[mdpt-0].jtype.Dims+1)
			} else {
				panic("??? Siranai Kata da")
			}
		}
	case 21:
		mdVAL.jtype = mdS[mdpt-0].jtype
	}
	goto mdstack /* stack new state and value */
}
