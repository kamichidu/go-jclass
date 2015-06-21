package jclass

import (
	"github.com/kamichidu/go-jclass/data"
	"github.com/kamichidu/go-jclass/parser/fd"
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
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.NameIndex)
}

func (self *JField) GetDescriptor() string {
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.DescriptorIndex)
}

func (self *JField) GetType() JType {
	fdinfo, err := fd.Parse(self.GetDescriptor())
	if err != nil {
		panic(err)
	}

	return newJType(fdinfo)
}

func (self *JField) GetAttributes() []*JAttribute {
	attributes := make([]*JAttribute, self.data.AttributesCount)
	for i := uint16(0); i < self.data.AttributesCount; i++ {
		attributes[i] = newJAttribute(self.enclosing.data.ConstantPool, &self.data.Attributes[i])
	}
	return attributes
}

func newJType(fdinfo *fd.FDInfo) JType {
	if fdinfo.PrimitiveType {
		return NewJPrimitiveType(fdinfo.TypeName)
	} else if fdinfo.ReferenceType {
		return NewJReferenceType(fdinfo.TypeName)
	} else if fdinfo.ArrayType {
		ct := fdinfo.ComponentType
		if ct.PrimitiveType {
			return NewJArrayType(NewJPrimitiveType(ct.TypeName), fdinfo.Dims)
		} else if ct.ReferenceType {
			return NewJArrayType(NewJReferenceType(ct.TypeName), fdinfo.Dims)
		} else {
			panic("???")
		}
	} else {
		panic("???")
	}
}
