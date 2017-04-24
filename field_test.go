package jclass

import (
	"testing"
)

func TestJavaFieldConstantValue(t *testing.T) {
	cases := []struct {
		N string
		T string
		V interface{}
	}{
		{"stringValue", "java.lang.String", "hoge"},
		{"booleanValue", "boolean", true},
		{"byteValue", "byte", int8(127)},
		{"charValue", "char", 'H'},
		{"shortValue", "short", int16(32767)},
		{"intValue", "int", int32(2147483647)},
		{"longValue", "long", int64(9223372036854775807)},
		{"floatValue", "float", float32(3.4028235E38)},
		{"doubleValue", "double", float64(1.7976931348623157E308)},
	}
	class, _ := NewJavaClassFromFilename("./testdata/Constants.class")
	for _, c := range cases {
		f := class.DeclaredField(c.N)
		if f == nil {
			t.Errorf("Field not found for %s", c.N)
			continue
		}
		if f.Type() != c.T {
			t.Errorf("Field type mismatch %v, wants %v", f.Type(), c.T)
			continue
		}
		v, ok := f.ConstantValue()
		if !ok {
			t.Errorf("ConstantValue not found for %s", c.N)
			continue
		}
		if v != c.V {
			t.Errorf("ConstantValue mismatch %v, wants %v", c.V, v)
		}
	}
}
