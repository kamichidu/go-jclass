package jclass

type JavaMethod struct {
	*MethodInfo
	AccessFlags

	constantPool []ConstantPoolInfo
}

func NewJavaMethod(constantPool []ConstantPoolInfo, methodInfo *MethodInfo) *JavaMethod {
	return &JavaMethod{methodInfo, AccessFlag(methodInfo.AccessFlags), constantPool}
}

func (self *JavaMethod) Name() string {
	utf8Info := self.constantPool[self.NameIndex].(*ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaMethod) ReturnType() string {
	utf8Info := self.constantPool[self.DescriptorIndex].(*ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaMethod) parseDescriptor() {
}
