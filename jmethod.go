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

func (self *JMethod) GetParameterTypes() []string {
	ret, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	types := make([]string, len(ret.ParameterDescriptor))
	for i := 0; i < len(types); i++ {
		types[i] = toStringImpl(ret.ParameterDescriptor[i])
	}
	return types
}

func (self *JMethod) GetReturnType() string {
	ret, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	return toStringImpl(ret.ReturnDescriptor)
}

func (self *JMethod) GetAttributes() []data.AttributeInfo {
	return self.data.Attributes
}
