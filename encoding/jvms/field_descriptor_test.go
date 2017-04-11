package jvms

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseFieldDescriptor(t *testing.T) {
	ok := func(fd string, expect *FieldDescriptorInfo) {
		t.Logf("Try to parse '%s'", fd)
		if actual, err := ParseFieldDescriptor(strings.NewReader(fd)); err == nil {
			if !reflect.DeepEqual(actual, expect) {
				t.Errorf("Failed:\nExpected %#v\nActual   %#v", expect, actual)
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("B", &FieldDescriptorInfo{"byte", true, false, 0})
	ok("C", &FieldDescriptorInfo{"char", true, false, 0})
	ok("D", &FieldDescriptorInfo{"double", true, false, 0})
	ok("F", &FieldDescriptorInfo{"float", true, false, 0})
	ok("I", &FieldDescriptorInfo{"int", true, false, 0})
	ok("J", &FieldDescriptorInfo{"long", true, false, 0})
	ok("S", &FieldDescriptorInfo{"short", true, false, 0})
	ok("Z", &FieldDescriptorInfo{"boolean", true, false, 0})
	ok("Ljava/lang/Object;", &FieldDescriptorInfo{"java.lang.Object", false, false, 0})
	ok("[B", &FieldDescriptorInfo{"byte", true, true, 1})
	ok("[C", &FieldDescriptorInfo{"char", true, true, 1})
	ok("[D", &FieldDescriptorInfo{"double", true, true, 1})
	ok("[F", &FieldDescriptorInfo{"float", true, true, 1})
	ok("[I", &FieldDescriptorInfo{"int", true, true, 1})
	ok("[J", &FieldDescriptorInfo{"long", true, true, 1})
	ok("[S", &FieldDescriptorInfo{"short", true, true, 1})
	ok("[Z", &FieldDescriptorInfo{"boolean", true, true, 1})
	ok("[Ljava/lang/Object;", &FieldDescriptorInfo{"java.lang.Object", false, true, 1})
	ok("[[B", &FieldDescriptorInfo{"byte", true, true, 2})
	ok("[[C", &FieldDescriptorInfo{"char", true, true, 2})
	ok("[[D", &FieldDescriptorInfo{"double", true, true, 2})
	ok("[[F", &FieldDescriptorInfo{"float", true, true, 2})
	ok("[[I", &FieldDescriptorInfo{"int", true, true, 2})
	ok("[[J", &FieldDescriptorInfo{"long", true, true, 2})
	ok("[[S", &FieldDescriptorInfo{"short", true, true, 2})
	ok("[[Z", &FieldDescriptorInfo{"boolean", true, true, 2})
	ok("[[Ljava/lang/Object;", &FieldDescriptorInfo{"java.lang.Object", false, true, 2})
}

func BenchmarkParseFieldDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFieldDescriptor(strings.NewReader("[[[[[[Ljava/util/List;"))
	}
}
