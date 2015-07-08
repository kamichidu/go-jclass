package md

import (
	"errors"
	c "github.com/kamichidu/go-jclass/parser/common"
	"strings"
)

type MDLexer struct {
	*c.Lexer
	Result *c.MethodDescriptor
}

func (self *MDLexer) Lex(lval *mdSymType) int {
	token := self.Lexer.Lex()
	if token == nil {
		return 0
	}
	lval.token = token
	return token.Id
}

func Parse(descriptor string) (*c.MethodDescriptor, error) {
	lexer := &MDLexer{c.NewLexer(descriptor), nil}
	ret := mdParse(lexer)
	if ret != 0 {
		return nil, errors.New(strings.Join(lexer.Errors, "\n"))
	}

	return lexer.Result, nil
}
