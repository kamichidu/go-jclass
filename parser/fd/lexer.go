package fd

import (
	"errors"
	c "github.com/kamichidu/go-jclass/parser/common"
	"strings"
)

type FDLexer struct {
	*c.Lexer
	Result *c.FieldDescriptor
}

func (self *FDLexer) Lex(lval *fdSymType) int {
	token := self.Lexer.Lex()
    if token == nil {
        return 0
    }
	lval.token = token
	return token.Id
}

func Parse(descriptor string) (*c.FieldDescriptor, error) {
	lexer := &FDLexer{c.NewLexer(descriptor), nil}
	ret := fdParse(lexer)
	if ret != 0 {
		return nil, errors.New(strings.Join(lexer.Errors, "\n"))
	}

	return lexer.Result, nil
}
