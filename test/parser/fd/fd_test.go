package fd

import (
	fdparser "github.com/kamichidu/go-jclass/parser/fd"
	"testing"
)

func TestParse(t *testing.T) {
	ok := func(fd string, expect string) {
		t.Logf("Try to parse '%s'", fd)
		if ret, err := fdparser.Parse(fd); err == nil {
			if ret != expect {
				t.Errorf("Expected %s", expect)
				t.Errorf("Got      %s", ret)
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("B", "byte")
	ok("C", "char")
	ok("D", "double")
	ok("F", "float")
	ok("I", "int")
	ok("J", "long")
	ok("S", "short")
	ok("Z", "boolean")
	ok("Ljava/lang/Object;", "java.lang.Object")
	ok("[B", "byte[]")
	ok("[C", "char[]")
	ok("[D", "double[]")
	ok("[F", "float[]")
	ok("[I", "int[]")
	ok("[J", "long[]")
	ok("[S", "short[]")
	ok("[Z", "boolean[]")
	ok("[Ljava/lang/Object;", "java.lang.Object[]")
	ok("[[B", "byte[][]")
	ok("[[C", "char[][]")
	ok("[[D", "double[][]")
	ok("[[F", "float[][]")
	ok("[[I", "int[][]")
	ok("[[J", "long[][]")
	ok("[[S", "short[][]")
	ok("[[Z", "boolean[][]")
	ok("[[Ljava/lang/Object;", "java.lang.Object[][]")
}

func BenchmarkParseFileDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fdparser.Parse("[[[[[[Ljava/util/List;")
	}
}
