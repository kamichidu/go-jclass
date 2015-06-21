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
	return self.enclosing.getUtf8String(self.data.NameIndex)
}

func (self *JMethod) GetDescriptor() string {
	return self.enclosing.getUtf8String(self.data.DescriptorIndex)
}

func (self *JMethod) GetParameterTypes() []jclass.JType {
	ret, err := md.Parse(self.GetDescriptor())
	if err != nil {
		panic(err)
	}
	return ret.GetParameterTypes()
}

func (self *JMethod) GetReturnType() jclass.JType {
	ret, err := md.Parse(self.GetDescriptor())
	if err != nil {
		panic(err)
	}
	return ret.GetReturnType()
}

func (self *JMethod) GetAttributes() []*JAttribute {
	attributes := make([]*JAttribute, self.data.AttributesCount)
	for i := uint16(0); i < self.data.AttributesCount; i++ {
		attributes[i] = newJAttributeWithJMethod(self.enclosing, self, &self.data.Attributes[i])
	}
	return attributes
}
