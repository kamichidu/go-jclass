package jclass

import (
	"github.com/kamichidu/go-jclass/jvms"
	"strings"
)

type JavaMethod struct {
	AccessFlags

	constantPool      []jvms.ConstantPoolInfo
	methodInfo        *jvms.MethodInfo
	returnTypeInfo    *jvms.FieldDescriptorInfo
	parameterTypeInfo []*jvms.FieldDescriptorInfo
}

func newJavaMethod(constantPool []jvms.ConstantPoolInfo, methodInfo *jvms.MethodInfo) *JavaMethod {
	return &JavaMethod{
		AccessFlags:       AccessFlag(methodInfo.AccessFlags),
		constantPool:      constantPool,
		methodInfo:        methodInfo,
		returnTypeInfo:    nil,
		parameterTypeInfo: nil,
	}
}

func (self *JavaMethod) Name() string {
	utf8Info := self.constantPool[self.methodInfo.NameIndex].(*jvms.ConstantUtf8Info)
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
	utf8Info := self.constantPool[self.methodInfo.DescriptorIndex].(*jvms.ConstantUtf8Info)
	descriptor := utf8Info.JavaString()

	info, err := jvms.ParseMethodDescriptor(strings.NewReader(descriptor))
	if err != nil {
		// TODO: Error handling
		self.returnTypeInfo = new(jvms.FieldDescriptorInfo)
		self.parameterTypeInfo = make([]*jvms.FieldDescriptorInfo, 0)
		return
	}
	self.returnTypeInfo = info.ReturnTypeInfo
	self.parameterTypeInfo = info.ParameterTypeInfo
}
