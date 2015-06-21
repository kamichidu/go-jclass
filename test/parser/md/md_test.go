package md

import (
	fdparser "github.com/kamichidu/go-jclass/parser/fd"
	mdparser "github.com/kamichidu/go-jclass/parser/md"
	"reflect"
	"testing"
)

func p(typeName string) *fdparser.FDInfo {
	return &fdparser.FDInfo{
		TypeName:      typeName,
		PrimitiveType: true,
	}
}

func r(typeName string) *fdparser.FDInfo {
	return &fdparser.FDInfo{
		TypeName:      typeName,
		ReferenceType: true,
	}
}

func a(ct *fdparser.FDInfo, dims int) *fdparser.FDInfo {
	return &fdparser.FDInfo{
		ComponentType: ct,
		Dims:          dims,
		ArrayType:     true,
	}
}

func m(retType *fdparser.FDInfo, params []*fdparser.FDInfo) *mdparser.MDInfo {
	return &mdparser.MDInfo{
		ReturnType:     retType,
		ParameterTypes: params,
	}
}

func TestParseMethodDescriptor(t *testing.T) {
	ok := func(md string, expect *mdparser.MDInfo) {
		t.Logf("Try to parse '%s'", md)
		if ret, err := mdparser.Parse(md); err == nil {
			if !reflect.DeepEqual(ret, expect) {
				t.Errorf("Expected %#v", expect)
				t.Errorf("Got      %#v", ret)
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", m(r("java/lang/Object"), []*fdparser.FDInfo{p("int"), p("double"), r("java/lang/Thread")}))
}
