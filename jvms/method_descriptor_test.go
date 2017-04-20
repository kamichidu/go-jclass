package jvms

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseMethodDescriptor(t *testing.T) {
	ok := func(md string, expect *MethodDescriptorInfo) {
		t.Logf("Try to parse '%s'", md)
		if actual, err := ParseMethodDescriptor(strings.NewReader(md)); err == nil {
			if !reflect.DeepEqual(actual, expect) {
				t.Errorf("Expected %#v", expect)
				t.Errorf("Got      %#v", actual)
			}
			t.Log(" -> OK")
		} else {
			t.Errorf("  ERR: %v", err)
		}
	}

	ok("()V", &MethodDescriptorInfo{
		ParameterTypeInfo: []*FieldDescriptorInfo{},
		ReturnTypeInfo:    &FieldDescriptorInfo{"void", true, false, 0},
	})
	ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", &MethodDescriptorInfo{
		ParameterTypeInfo: []*FieldDescriptorInfo{
			&FieldDescriptorInfo{"int", true, false, 0},
			&FieldDescriptorInfo{"double", true, false, 0},
			&FieldDescriptorInfo{"java.lang.Thread", false, false, 0},
		},
		ReturnTypeInfo: &FieldDescriptorInfo{"java.lang.Object", false, false, 0},
	})
}
