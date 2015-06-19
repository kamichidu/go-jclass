package fd

import (
	"errors"
	"github.com/kamichidu/go-jclass"
	"strings"
)

type FDLexer struct {
	Text   []rune
	Length int
	Pos    int
	Result jclass.JType
	Errors []string
}

func (self *FDLexer) Lex(lval *fdSymType) int {
	if self.Pos >= self.Length {
		return 0
	}

	c := self.Text[self.Pos]
	pos := self.Pos
	self.Pos++
	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', ';', '[':
		lval.token = FDToken{
			Id:   int(c),
			Text: string(c),
			Pos:  pos,
		}
		return lval.token.Id
	default:
		var className []rune
		className = append(className, c)
		for self.Pos < self.Length {
			c = self.Text[self.Pos]
			if c == ';' {
				break
			}
			className = append(className, c)
			self.Pos++
		}
		lval.token = FDToken{
			Id:   CLASS_NAME,
			Text: string(className),
			Pos:  pos,
		}
		return lval.token.Id
	}
}

func (self *FDLexer) Error(s string) {
	self.Errors = append(self.Errors, s)
}

func Parse(descriptor string) (jclass.JType, error) {
	lexer := &FDLexer{
		Text:   []rune(descriptor),
		Length: len(descriptor),
		Pos:    0,
	}
	ret := fdParse(lexer)
	if ret != 0 {
		return nil, errors.New(strings.Join(lexer.Errors, "\n"))
	}

	return lexer.Result, nil
}
