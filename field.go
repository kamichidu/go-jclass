package jclass

type JavaField struct {
	*FieldInfo
	AccessFlags

	constantPool []ConstantPoolInfo
}

func NewJavaField(constantPool []ConstantPoolInfo, fieldInfo *FieldInfo) *JavaField {
	return &JavaField{fieldInfo, AccessFlag(fieldInfo.AccessFlags), constantPool}
}

func (self *JavaField) Name() string {
	utf8Info := self.constantPool[self.NameIndex].(*ConstantUtf8Info)
	return utf8Info.JavaString()
}
