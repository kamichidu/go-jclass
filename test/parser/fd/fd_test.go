package fd

import (
	fdparser "github.com/kamichidu/go-jclass/parser/fd"
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

func TestParse(t *testing.T) {
	ok := func(fd string, expect *fdparser.FDInfo) {
		t.Logf("Try to parse '%s'", fd)
		if ret, err := fdparser.Parse(fd); err == nil {
			if !reflect.DeepEqual(ret, expect) {
				t.Errorf("Expected %v", expect)
				t.Errorf("Got      %v", ret)
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("B", p("byte"))
	ok("C", p("char"))
	ok("D", p("double"))
	ok("F", p("float"))
	ok("I", p("int"))
	ok("J", p("long"))
	ok("S", p("short"))
	ok("Z", p("boolean"))
	ok("Ljava/lang/Object;", r("java/lang/Object"))
	ok("[B", a(p("byte"), 1))
	ok("[C", a(p("char"), 1))
	ok("[D", a(p("double"), 1))
	ok("[F", a(p("float"), 1))
	ok("[I", a(p("int"), 1))
	ok("[J", a(p("long"), 1))
	ok("[S", a(p("short"), 1))
	ok("[Z", a(p("boolean"), 1))
	ok("[Ljava/lang/Object;", a(r("java/lang/Object"), 1))
	ok("[[B", a(p("byte"), 2))
	ok("[[C", a(p("char"), 2))
	ok("[[D", a(p("double"), 2))
	ok("[[F", a(p("float"), 2))
	ok("[[I", a(p("int"), 2))
	ok("[[J", a(p("long"), 2))
	ok("[[S", a(p("short"), 2))
	ok("[[Z", a(p("boolean"), 2))
	ok("[[Ljava/lang/Object;", a(r("java/lang/Object"), 2))
}

func BenchmarkParseFileDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fdparser.Parse("[[[[[[Ljava/util/List;")
	}
}
