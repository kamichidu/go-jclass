//line parser/md/parser.go.y:2
package md

import __yyfmt__ "fmt"

//line parser/md/parser.go.y:2
import (
	c "github.com/kamichidu/go-jclass/parser/common"
)

//line parser/md/parser.go.y:23
type mdSymType struct {
	yys                  int
	methodDescriptor     *c.MethodDescriptor
	parameterDescriptors []*c.ParameterDescriptor
	parameterDescriptor  *c.ParameterDescriptor
	returnDescriptor     *c.ReturnDescriptor
	fieldType            *c.FieldType
	baseType             *c.BaseType
	objectType           *c.ObjectType
	arrayType            *c.ArrayType
	componentType        *c.ComponentType
	className            *c.ClassName
	token                *c.Token
}

const IDENTIFIER = 57346

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
	"IDENTIFIER",
}
var mdStatenames = []string{}

const mdEofCode = 1
const mdErrCode = 2
const mdMaxDepth = 200

//line parser/md/parser.go.y:185

//line yacctab:1
var mdExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const mdNprod = 23
const mdPrivate = 57344

var mdTokenNames []string
var mdStates []string

const mdLast = 58

var mdAct = []int{

	10, 11, 12, 13, 14, 15, 16, 17, 18, 29,
	19, 24, 2, 23, 22, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 25, 19, 9, 4, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 6, 19, 27,
	8, 7, 21, 20, 5, 28, 3, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 26,
}
var mdPact = []int{

	-3, -1000, -1000, 11, -4, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -6, 24,
	-1000, -1000, -1000, 26, -1000, -1000, -1000, -1000, -8, -1000,
}
var mdPgo = []int{

	0, 47, 46, 44, 43, 37, 41, 40, 26, 24,
	13,
}
var mdR1 = []int{

	0, 1, 2, 2, 3, 4, 4, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 6, 7, 8,
	9, 10, 10,
}
var mdR2 = []int{

	0, 4, 0, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 2,
	1, 3, 1,
}
var mdChk = []int{

	-1000, -1, 15, -2, 16, -3, -5, -6, -7, -8,
	4, 5, 6, 7, 8, 9, 10, 11, 12, 14,
	-4, -5, 18, -10, 17, -9, -5, 13, 19, 17,
}
var mdDef = []int{

	0, -2, 2, 0, 0, 3, 4, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 0, 0,
	1, 5, 6, 0, 22, 19, 20, 18, 0, 21,
}
var mdTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 3, 3, 3, 19, 3, 3,
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
		//line parser/md/parser.go.y:41
		{
			mdVAL.methodDescriptor = &c.MethodDescriptor{
				ParameterDescriptor: mdS[mdpt-2].parameterDescriptors,
				ReturnDescriptor:    mdS[mdpt-0].returnDescriptor,
			}
			if l, ok := mdlex.(*MDLexer); ok {
				l.Result = mdVAL.methodDescriptor
			}
		}
	case 2:
		//line parser/md/parser.go.y:54
		{
			mdVAL.parameterDescriptors = make([]*c.ParameterDescriptor, 0)
		}
	case 3:
		//line parser/md/parser.go.y:58
		{
			mdVAL.parameterDescriptors = append(mdS[mdpt-1].parameterDescriptors, mdS[mdpt-0].parameterDescriptor)
		}
	case 4:
		//line parser/md/parser.go.y:65
		{
			mdVAL.parameterDescriptor = &c.ParameterDescriptor{
				FieldType: mdS[mdpt-0].fieldType,
			}
		}
	case 5:
		//line parser/md/parser.go.y:74
		{
			mdVAL.returnDescriptor = &c.ReturnDescriptor{
				FieldType: mdS[mdpt-0].fieldType,
			}
		}
	case 6:
		//line parser/md/parser.go.y:80
		{
			mdVAL.returnDescriptor = &c.ReturnDescriptor{
				VoidDescriptor: &c.VoidDescriptor{"void"},
			}
		}
	case 7:
		//line parser/md/parser.go.y:89
		{
			mdVAL.fieldType = &c.FieldType{
				BaseType: mdS[mdpt-0].baseType,
			}
		}
	case 8:
		//line parser/md/parser.go.y:95
		{
			mdVAL.fieldType = &c.FieldType{
				ObjectType: mdS[mdpt-0].objectType,
			}
		}
	case 9:
		//line parser/md/parser.go.y:101
		{
			mdVAL.fieldType = &c.FieldType{
				ArrayType: mdS[mdpt-0].arrayType,
			}
		}
	case 10:
		//line parser/md/parser.go.y:110
		{
			mdVAL.baseType = &c.BaseType{"byte"}
		}
	case 11:
		//line parser/md/parser.go.y:114
		{
			mdVAL.baseType = &c.BaseType{"char"}
		}
	case 12:
		//line parser/md/parser.go.y:118
		{
			mdVAL.baseType = &c.BaseType{"double"}
		}
	case 13:
		//line parser/md/parser.go.y:122
		{
			mdVAL.baseType = &c.BaseType{"float"}
		}
	case 14:
		//line parser/md/parser.go.y:126
		{
			mdVAL.baseType = &c.BaseType{"int"}
		}
	case 15:
		//line parser/md/parser.go.y:130
		{
			mdVAL.baseType = &c.BaseType{"long"}
		}
	case 16:
		//line parser/md/parser.go.y:134
		{
			mdVAL.baseType = &c.BaseType{"short"}
		}
	case 17:
		//line parser/md/parser.go.y:138
		{
			mdVAL.baseType = &c.BaseType{"boolean"}
		}
	case 18:
		//line parser/md/parser.go.y:145
		{
			mdVAL.objectType = &c.ObjectType{
				ClassName: mdS[mdpt-1].className,
			}
		}
	case 19:
		//line parser/md/parser.go.y:154
		{
			mdVAL.arrayType = &c.ArrayType{
				ComponentType: mdS[mdpt-0].componentType,
			}
		}
	case 20:
		//line parser/md/parser.go.y:163
		{
			mdVAL.componentType = &c.ComponentType{
				FieldType: mdS[mdpt-0].fieldType,
			}
		}
	case 21:
		//line parser/md/parser.go.y:172
		{
			mdVAL.className = &c.ClassName{
				Identifier: append(mdS[mdpt-2].className.Identifier, mdS[mdpt-0].token.Text),
			}
		}
	case 22:
		//line parser/md/parser.go.y:178
		{
			mdVAL.className = &c.ClassName{
				Identifier: []string{mdS[mdpt-0].token.Text},
			}
		}
	}
	goto mdstack /* stack new state and value */
}
