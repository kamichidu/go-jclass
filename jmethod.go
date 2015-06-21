package jclass

import (
	"github.com/kamichidu/go-jclass/data"
	"github.com/kamichidu/go-jclass/parser/md"
)

type JMethod struct {
	enclosing *JClass
	data      *data.MethodInfo
}

func newJMethod(enclosing *JClass, data *data.MethodInfo) *JMethod {
	return &JMethod{
		enclosing: enclosing,
		data:      data,
	}
}

func (self *JMethod) GetAccessFlags() uint16 {
	return self.data.AccessFlags
}

func (self *JMethod) GetName() string {
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.NameIndex)
}

func (self *JMethod) getDescriptor() string {
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.DescriptorIndex)
}

func (self *JMethod) GetParameterTypes() []JType {
	ret, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	types := make([]JType, len(ret.ParameterTypes))
	for i := 0; i < len(ret.ParameterTypes); i++ {
		types[i] = newJType(ret.ParameterTypes[i])
	}
	return types
}

func (self *JMethod) GetReturnType() JType {
	ret, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	return newJType(ret.ReturnType)
}

func (self *JMethod) GetAttributes() []*JAttribute {
	attributes := make([]*JAttribute, self.data.AttributesCount)
	for i := uint16(0); i < self.data.AttributesCount; i++ {
		attributes[i] = newJAttribute(self.enclosing.data.ConstantPool, &self.data.Attributes[i])
	}
	return attributes
}
