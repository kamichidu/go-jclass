package common

import (
	"fmt"
)

type LexerInput struct {
	Text   []rune
	Length int
	Pos    int
}

const (
	EOF        = '\x00'
	IDENTIFIER = 57346
)

func (self *LexerInput) Next() rune {
	if self.Eof() {
		return EOF
	}

	c := self.Text[self.Pos]
	self.Pos++
	return c
}

func (self *LexerInput) Peek() rune {
	if self.Eof() {
		return EOF
	}

	return self.Text[self.Pos]
}

func (self *LexerInput) Eof() bool {
	return self.Pos >= self.Length
}

type Lexer struct {
	Input   *LexerInput
	Errors  []string
	inState bool
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input: &LexerInput{
			Text:   []rune(input),
			Length: len(input),
			Pos:    0,
		},
		Errors: make([]string, 0),
	}
}

func (self *Lexer) Lex() *Token {
	if self.Input.Eof() {
		return nil
	}

	pos := self.Input.Pos
	ch := self.Input.Next()
	if ch == ';' {
		self.inState = false
	}
	var charsAsId []rune
	if self.inState {
		charsAsId = []rune{'/'}
	} else {
		charsAsId = []rune{'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', '[', '(', ')', 'V', '^', ':', '<', '>', 'T', '.', '+', '-', '*', ';', '/'}
	}
	if ch == 'L' || ch == 'T' {
		self.inState = true
	}
	for _, charAsId := range charsAsId {
		if ch == charAsId {
			return &Token{
				Id:   int(ch),
				Text: string(ch),
				Pos:  pos,
			}
		}
	}

	// Identifier
	startPos := pos
	for isIdentifierChar(self.Input.Peek()) {
		self.Input.Next()
	}
	return &Token{
		Id:   IDENTIFIER,
		Text: string(self.Input.Text[startPos:self.Input.Pos]),
		Pos:  startPos,
	}
}

func isIdentifierChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '$'
}

func (self *Lexer) Error(s string) {
	if !self.Input.Eof() {
		desc := fmt.Sprintf("%v at %c (column %d)", s, self.Input.Peek(), self.Input.Pos)
		self.Errors = append(self.Errors, desc)
	} else {
		self.Errors = append(self.Errors, "Unexpected EOF")
	}
}
