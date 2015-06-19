package parser

import (
	"errors"
	"fmt"
	"strings"
)

type DescriptorLexer struct {
	Text   []rune
	Length int
	Pos    int
	Result string
	Errors []string
}

func (self *DescriptorLexer) Lex(lval *yySymType) int {
	fmt.Printf("lval: %v\n", lval)

	if self.Pos >= self.Length {
		return 0
	}

	c := self.Text[self.Pos]
	pos := self.Pos
	self.Pos++
	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', ';', '[':
		lval.token = Token{
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
		lval.token = Token{
			Id:   CLASS_NAME,
			Text: string(className),
			Pos:  pos,
		}
		return lval.token.Id
	}
}

func (self *DescriptorLexer) Error(s string) {
	self.Errors = append(self.Errors, s)
}

type FieldDescriptor struct {
	TypeName string
}

func ParseFieldDescriptor(descriptor string) (*FieldDescriptor, error) {
	lexer := &DescriptorLexer{
		Text:   []rune(descriptor),
		Length: len(descriptor),
		Pos:    0,
	}
	ret := yyParse(lexer)
	if ret != 0 {
		return nil, errors.New(strings.Join(lexer.Errors, "\n"))
	}

	return &FieldDescriptor{
        TypeName: lexer.Result,
	}, nil
}
