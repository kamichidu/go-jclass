package md

import (
	mdparser "github.com/kamichidu/go-jclass/parser/md"
	"testing"
)

func TestParseMethodDescriptor(t *testing.T) {
	ok := func(md string, expectReturnType string, expectParameterTypes []string) {
		t.Logf("Try to parse '%s'", md)
		if ret, err := mdparser.Parse(md); err == nil {
			if ret.GetReturnType().GetTypeName() != expectReturnType {
				t.Errorf("Expected '%s', but got '%s'", expectReturnType, ret.GetReturnType().GetTypeName())
			}
			t.Logf(" -> %s", ret.GetReturnType().GetTypeName())

			if len(ret.GetParameterTypes()) != len(expectParameterTypes) {
				t.Errorf("Expected %d, but got %d", len(expectParameterTypes), len(ret.GetParameterTypes()))
			}
			t.Logf(" -> It has %d parameters", len(ret.GetParameterTypes()))

			for i := 0; i < len(ret.GetParameterTypes()); i++ {
				expect := expectParameterTypes[i]
				actual := ret.GetParameterTypes()[i]
				if actual.GetTypeName() != expect {
					t.Errorf("Expected '%s', buto got '%s'", expect, actual.GetTypeName())
				}
				t.Logf("  %d: %s", i, actual.GetTypeName())
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", "java/lang/Object", []string{"int", "double", "java/lang/Thread"})
}
