package jclass

import (
	ejvms "github.com/kamichidu/go-jclass/encoding/jvms"
	"github.com/kamichidu/go-jclass/jvms"
	"strings"
)

type JavaField struct {
	*jvms.FieldInfo
	AccessFlags

	constantPool []jvms.ConstantPoolInfo
}

func NewJavaField(constantPool []jvms.ConstantPoolInfo, fieldInfo *jvms.FieldInfo) *JavaField {
	return &JavaField{fieldInfo, AccessFlag(fieldInfo.AccessFlags), constantPool}
}

func (self *JavaField) Name() string {
	utf8Info := self.constantPool[self.NameIndex].(*jvms.ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaField) Type() string {
	utf8Info := self.constantPool[self.DescriptorIndex].(*jvms.ConstantUtf8Info)
	descriptor := utf8Info.JavaString()

	info, err := ejvms.ParseFieldDescriptor(strings.NewReader(descriptor))
	if err != nil {
		// TODO: Error handling
		return ""
	}
	return info.String()
}
