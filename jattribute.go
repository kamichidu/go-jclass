package jclass

import (
	"github.com/kamichidu/go-jclass/data"
)

type JAttribute struct {
	jclass *JClass
	rel    interface{}
	data   *data.AttributeInfo
}

func newJAttributeWithJClass(jclass *JClass, data *data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    jclass,
		data:   data,
	}
}

func newJAttributeWithJField(jclass *JClass, rel *JField, data *data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    rel,
		data:   data,
	}
}

func newJAttributeWithJMethod(jclass *JClass, rel *JMethod, data *data.AttributeInfo) *JAttribute {
	return &JAttribute{
		jclass: jclass,
		rel:    rel,
		data:   data,
	}
}

func (self *JAttribute) GetName() string {
	return self.jclass.getUtf8String(self.data.AttributeNameIndex)
}

func (self *JAttribute) GetInfo() []uint8 {
	return self.data.Info
}
