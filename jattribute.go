package jclass

import (
	"github.com/kamichidu/go-jclass/data"
	"reflect"
)

type JAttribute struct {
	jclass *JClass
	rel    interface{}
	data   data.AttributeInfo
}

func newJAttributeWithJClass(jclass *JClass, data data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    jclass,
		data:   data,
	}
}

func newJAttributeWithJField(jclass *JClass, rel *JField, data data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    rel,
		data:   data,
	}
}

func newJAttributeWithJMethod(jclass *JClass, rel *JMethod, data data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    rel,
		data:   data,
	}
}

func (self *JAttribute) GetName() string {
	val := reflect.ValueOf(self.data).Elem().Elem().Elem()
	nameIndex := uint16(val.FieldByName("AttributeNameIndex").Uint())
	return getUtf8String(self.jclass.data.ConstantPool, nameIndex)
}
