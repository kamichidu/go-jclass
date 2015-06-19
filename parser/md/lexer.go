package md

import (
	"errors"
	"fmt"
	"github.com/kamichidu/go-jclass"
	"strings"
)

type MethodInfo struct {
	returnType     jclass.JType
	parameterTypes []jclass.JType
}

func (self *MethodInfo) GetReturnType() jclass.JType {
	return self.returnType
}

func (self *MethodInfo) GetParameterTypes() []jclass.JType {
	return self.parameterTypes
}

type MDLexer struct {
	Text   []rune
	Length int
	Pos    int
	Result *MethodInfo
	Errors []string
}

func (self *MDLexer) Lex(lval *mdSymType) int {
	if self.Pos >= self.Length {
		return 0
	}

	c := self.Text[self.Pos]
	pos := self.Pos
	self.Pos++
	switch c {
    case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', ';', '[', '(', ')':
		lval.token = MDToken{
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
		lval.token = MDToken{
			Id:   CLASS_NAME,
			Text: string(className),
			Pos:  pos,
		}
		return lval.token.Id
	}
}

func (self *MDLexer) Error(s string) {
    desc := fmt.Sprintf("%v at %c (column %d)", s, self.Text[self.Pos], self.Pos)
	self.Errors = append(self.Errors, desc)
}

func Parse(descriptor string) (*MethodInfo, error) {
	lexer := &MDLexer{
		Text:   []rune(descriptor),
		Length: len(descriptor),
		Pos:    0,
        Result: &MethodInfo{},
	}
	ret := mdParse(lexer)
	if ret != 0 {
		return nil, errors.New(strings.Join(lexer.Errors, "\n"))
	}

	return lexer.Result, nil
}
