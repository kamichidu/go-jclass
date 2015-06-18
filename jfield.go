package jclass

import (
	"github.com/kamichidu/go-jclass/data"
)

type JField struct {
	enclosing *JClass
	data      *data.FieldInfo
}

func newJField(enclosing *JClass, data *data.FieldInfo) *JField {
	return &JField{
		enclosing: enclosing,
		data:      data,
	}
}

func (self *JField) GetAccessFlags() uint16 {
	return self.data.AccessFlags
}

func (self *JField) GetName() string {
	return self.enclosing.getUtf8String(self.data.NameIndex)
}

func (self *JField) GetDescriptor() string {
	return self.enclosing.getUtf8String(self.data.DescriptorIndex)
}

func (self *JField) GetAttributes() []*JAttribute {
	attributes := make([]*JAttribute, self.data.AttributesCount)
	for i := uint16(0); i < self.data.AttributesCount; i++ {
		attributes[i] = newJAttributeWithJField(self.enclosing, self, &self.data.Attributes[i])
	}
	return attributes
}
