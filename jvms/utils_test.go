package jvms

import (
	"strings"
	"testing"
)

func TestRuneReaderNoUnread(t *testing.T) {
	r := newReader(strings.NewReader("abcdefg"))

	for i, expect := range []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if c, _, err := r.ReadRune(); err != nil {
			t.Fatal(err)
		} else if c != expect {
			t.Fatalf("Expected %dth '%c', but is '%c'", i, expect, c)
		}
	}
}

func TestRuneReaderWithUnread(t *testing.T) {
	r := newReader(strings.NewReader("abcdefg"))

	for i, expect := range []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if c, _, err := r.ReadRune(); err != nil {
			t.Fatal(err)
		} else if c != expect {
			t.Fatalf("Expected %dth '%c', but is '%c'", i, expect, c)
		}

		if err := r.UnreadRune(); err != nil {
			t.Fatal(err)
		}

		if c, _, err := r.ReadRune(); err != nil {
			t.Fatal(err)
		} else if c != expect {
			t.Fatalf("Expected (%d + 1 - 1)th '%c', but is '%c'", i, expect, c)
		}

	}
}
