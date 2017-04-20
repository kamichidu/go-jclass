package jvms

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var debug = false

func baseType2TypeName(token string) string {
	switch token {
	case "B":
		return "byte"
	case "C":
		return "char"
	case "D":
		return "double"
	case "F":
		return "float"
	case "I":
		return "int"
	case "J":
		return "long"
	case "S":
		return "short"
	case "Z":
		return "boolean"
	default:
		return token
	}
}

func peek(r runeReader) (rune, error) {
	c, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return '\x00', nil
		} else {
			return c, err
		}
	}
	return c, r.UnreadRune()
}

func lookahead(r *bufio.Reader, ch rune) (bool, error) {
	var err error

	found := false
	unreadRunes := 0
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		unreadRunes++
		fmt.Printf("LookAhead ReadRune %c\n", c)

		if c == ch {
			found = true
			break
		}
	}
	for unreadRunes > 0 {
		if err = r.UnreadRune(); err != nil {
			return false, err
		}
		fmt.Println("LookAhead UnreadRune")
		unreadRunes--
	}
	return found, nil
}

func errorPrefixUnmatch(syntax string, expects []rune, actual rune) error {
	chars := make([]string, 0)
	if len(expects) == 1 {
		chars = append(chars, fmt.Sprintf("`%c'", expects[0]))
	} else {
		for _, c := range expects {
			chars = append(chars, fmt.Sprintf("`%c'", c))
		}
	}
	return fmt.Errorf("%s must starts with %s, but with `%c'", syntax, strings.Join(chars, ", "), actual)
}

func errorSuffixUnmatch(syntax string, expects []rune, actual rune) error {
	chars := make([]string, 0)
	if len(expects) == 1 {
		chars = append(chars, fmt.Sprintf("`%c'", expects[0]))
	} else {
		for _, c := range expects {
			chars = append(chars, fmt.Sprintf("`%c'", c))
		}
	}
	return fmt.Errorf("%s must ends with %s, but with `%c'", syntax, strings.Join(chars, ", "), actual)
}

type runeReader interface {
	ReadRune() (rune, int, error)
	UnreadRune() error
}

type reader struct {
	r *bufio.Reader

	buffer []rune
	offset int
}

func newReader(r io.Reader) *reader {
	return &reader{
		r:      bufio.NewReader(r),
		buffer: make([]rune, 0),
		offset: 0,
	}
}

func (self *reader) ReadRune() (rune, int, error) {
	if self.offset < len(self.buffer) {
		c := self.buffer[self.offset]
		self.offset++
		if debug {
			fmt.Printf("***ReadRune***** % 2d -> % 2d: %c\n", self.offset-1, self.offset, c)
		}
		return c, 1, nil
	} else {
		c, n, err := self.r.ReadRune()
		if err != nil {
			return c, n, err
		}
		self.buffer = append(self.buffer, c)
		self.offset++
		if debug {
			fmt.Printf("***ReadRune***** % 2d -> % 2d: %c\n", self.offset-1, self.offset, c)
		}
		return c, n, err
	}
}

func (self *reader) UnreadRune() error {
	if self.offset <= 0 {
		return fmt.Errorf("Can't UnreadRune() with offset 0")
	}
	self.offset--
	if debug {
		fmt.Printf("***UnreadRune*** % 2d -> % 2d\n", self.offset+1, self.offset)
	}
	return nil
}
