package md

import (
	parser "github.com/kamichidu/go-jclass/parser/md"
	"reflect"
	"strings"
	"testing"
)

func TestParseMethodDescriptor(t *testing.T) {
	type Expect struct {
		P []string
		R string
	}

	ok := func(md string, expect *Expect) {
		t.Logf("Try to parse '%s'", md)
		if params, ret, n, err := parser.Parse(md); err == nil {
			if !reflect.DeepEqual(params, expect.P) {
				t.Errorf("Expected %s", strings.Join(expect.P, ", "))
				t.Errorf("Got      %s", strings.Join(params, ", "))
			}
			if ret != expect.R {
				t.Errorf("Expected %s", expect.R)
				t.Errorf("Got      %s", ret)
			}
			if n != len(md) {
				t.Errorf("Consumed %d characters, but expected %d", n, len(md))
			}
			t.Log(" -> OK")
		} else {
			t.Errorf("  ERR: %v", err)
		}
	}

	ok("()V", &Expect{
		P: []string{},
		R: "void",
	})
	ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", &Expect{
		P: []string{"int", "double", "java.lang.Thread"},
		R: "java.lang.Object",
	})
}
