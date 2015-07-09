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

func (self *JMethod) IsPublic() bool {
	return (self.GetAccessFlags() & ACC_PUBLIC) == ACC_PUBLIC
}

func (self *JMethod) IsPrivate() bool {
	return (self.GetAccessFlags() & ACC_PRIVATE) == ACC_PRIVATE
}

func (self *JMethod) IsProtected() bool {
	return (self.GetAccessFlags() & ACC_PROTECTED) == ACC_PROTECTED
}

func (self *JMethod) IsStatic() bool {
	return (self.GetAccessFlags() & ACC_STATIC) == ACC_STATIC
}

func (self *JMethod) IsFinal() bool {
	return (self.GetAccessFlags() & ACC_FINAL) == ACC_FINAL
}

func (self *JMethod) IsSynchronized() bool {
	return (self.GetAccessFlags() & ACC_SYNCHRONIZED) == ACC_SYNCHRONIZED
}

func (self *JMethod) IsBridge() bool {
	return (self.GetAccessFlags() & ACC_BRIDGE) == ACC_BRIDGE
}

func (self *JMethod) IsVarargs() bool {
	return (self.GetAccessFlags() & ACC_VARARGS) == ACC_VARARGS
}

func (self *JMethod) IsNative() bool {
	return (self.GetAccessFlags() & ACC_NATIVE) == ACC_NATIVE
}

func (self *JMethod) IsAbstract() bool {
	return (self.GetAccessFlags() & ACC_ABSTRACT) == ACC_ABSTRACT
}

func (self *JMethod) IsStrict() bool {
	return (self.GetAccessFlags() & ACC_STRICT) == ACC_STRICT
}

func (self *JMethod) IsSynthetic() bool {
	return (self.GetAccessFlags() & ACC_SYNTHETIC) == ACC_SYNTHETIC
}

func (self *JMethod) GetName() string {
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.NameIndex)
}

func (self *JMethod) getDescriptor() string {
	return getUtf8String(self.enclosing.data.ConstantPool, self.data.DescriptorIndex)
}

func (self *JMethod) GetParameterTypes() []string {
	params, _, _, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	return params
}

func (self *JMethod) GetReturnType() string {
	_, ret, _, err := md.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}
	return ret
}

func (self *JMethod) GetAttributes() []data.AttributeInfo {
	return self.data.Attributes
}
