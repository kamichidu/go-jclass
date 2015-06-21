package md

import (
	fdparser "github.com/kamichidu/go-jclass/parser/fd"
	mdparser "github.com/kamichidu/go-jclass/parser/md"
	"reflect"
	"testing"
)

func pp(val interface{}) string {
    v := reflect.ValueOf(val)
    return v.String()
}

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
                t.Errorf("Expected %s", pp(expect))
                t.Errorf("Got      %s", pp(ret))
			}
            t.Log(" -> OK")
		} else {
			t.Errorf("  ERR: %v", err)
		}
	}

	ok("()V", m(p("void"), []*fdparser.FDInfo{}))
	ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", m(r("java/lang/Object"), []*fdparser.FDInfo{p("int"), p("double"), r("java/lang/Thread")}))
}
