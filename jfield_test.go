// +build ignore

package jclass

import (
	"testing"
)

func TestJField_GetConstantValue(t *testing.T) {
	jc, _ := NewJClassWithFilename("./Constants.class")
	if f := jc.GetField("stringValue"); f.GetConstantValue() != "hoge" {
		t.Errorf("stringValue constant has \"hoge\", but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("booleanValue"); f.GetConstantValue() != true {
		t.Errorf("booleanValue constant has true, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("byteValue"); f.GetConstantValue() != int8(127) {
		t.Errorf("byteValue constant has 127, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("charValue"); f.GetConstantValue() != 'H' {
		t.Errorf("charValue constant has 'H', but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("shortValue"); f.GetConstantValue() != int16(32767) {
		t.Errorf("shortValue constant has 32767, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("intValue"); f.GetConstantValue() != int32(2147483647) {
		t.Errorf("intValue constant has 2147483647, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("longValue"); f.GetConstantValue() != int64(9223372036854775807) {
		t.Errorf("longValue constant has 9223372036854775807, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("floatValue"); f.GetConstantValue() != float32(3.4028235E38) {
		t.Errorf("floatValue constant has 3.4028235E38, but got %v", f.GetConstantValue())
	}
	if f := jc.GetField("doubleValue"); f.GetConstantValue() != float64(1.7976931348623157E308) {
		t.Errorf("doubleValue constant has 1.7976931348623157E308, but got %v", f.GetConstantValue())
	}
}

func TestJField_GetType(t *testing.T) {
	jc, _ := NewJClassWithFilename("./Constants.class")
	if f := jc.GetField("stringValue"); f.GetType() != "java.lang.String" {
		t.Errorf("stringValue constant has \"java.lang.String\", but got %v", f.GetType())
	}
	if f := jc.GetField("booleanValue"); f.GetType() != "boolean" {
		t.Errorf("booleanValue constant has boolean, but got %v", f.GetType())
	}
	if f := jc.GetField("byteValue"); f.GetType() != "byte" {
		t.Errorf("byteValue constant has byte, but got %v", f.GetType())
	}
	if f := jc.GetField("charValue"); f.GetType() != "char" {
		t.Errorf("charValue constant has char, but got %v", f.GetType())
	}
	if f := jc.GetField("shortValue"); f.GetType() != "short" {
		t.Errorf("shortValue constant has short, but got %v", f.GetType())
	}
	if f := jc.GetField("intValue"); f.GetType() != "int" {
		t.Errorf("intValue constant has int, but got %v", f.GetType())
	}
	if f := jc.GetField("longValue"); f.GetType() != "long" {
		t.Errorf("longValue constant has long, but got %v", f.GetType())
	}
	if f := jc.GetField("floatValue"); f.GetType() != "float" {
		t.Errorf("floatValue constant has float, but got %v", f.GetType())
	}
	if f := jc.GetField("doubleValue"); f.GetType() != "double" {
		t.Errorf("doubleValue constant has double, but got %v", f.GetType())
	}
}
