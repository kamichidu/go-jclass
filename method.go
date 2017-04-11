package jclass

import (
	ejvms "github.com/kamichidu/go-jclass/encoding/jvms"
	"github.com/kamichidu/go-jclass/jvms"
	"strings"
)

type JavaMethod struct {
	*jvms.MethodInfo
	AccessFlags

	constantPool      []jvms.ConstantPoolInfo
	returnTypeInfo    *ejvms.FieldDescriptorInfo
	parameterTypeInfo []*ejvms.FieldDescriptorInfo
}

func NewJavaMethod(constantPool []jvms.ConstantPoolInfo, methodInfo *jvms.MethodInfo) *JavaMethod {
	return &JavaMethod{
		MethodInfo:        methodInfo,
		AccessFlags:       AccessFlag(methodInfo.AccessFlags),
		constantPool:      constantPool,
		returnTypeInfo:    nil,
		parameterTypeInfo: nil,
	}
}

func (self *JavaMethod) Name() string {
	utf8Info := self.constantPool[self.NameIndex].(*jvms.ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaMethod) ReturnType() string {
	self.parseDescriptor()

	return self.returnTypeInfo.String()
}

func (self *JavaMethod) ParameterTypes() []string {
	self.parseDescriptor()

	typs := make([]string, len(self.parameterTypeInfo))
	for i := 0; i < len(typs); i++ {
		typs[i] = self.parameterTypeInfo[i].String()
	}
	return typs
}

func (self *JavaMethod) parseDescriptor() {
	utf8Info := self.constantPool[self.DescriptorIndex].(*jvms.ConstantUtf8Info)
	descriptor := utf8Info.JavaString()

	info, err := ejvms.ParseMethodDescriptor(strings.NewReader(descriptor))
	if err != nil {
		// TODO: Error handling
		self.returnTypeInfo = new(ejvms.FieldDescriptorInfo)
		self.parameterTypeInfo = make([]*ejvms.FieldDescriptorInfo, 0)
		return
	}
	self.returnTypeInfo = info.ReturnTypeInfo
	self.parameterTypeInfo = info.ParameterTypeInfo
}