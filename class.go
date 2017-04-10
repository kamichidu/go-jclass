package jclass

import (
	"strings"
)

type JavaClass struct {
	*ClassFile
	AccessFlags
}

func NewJavaClass(classFile *ClassFile) *JavaClass {
	return &JavaClass{classFile, AccessFlag(classFile.AccessFlags)}
}

func (self *JavaClass) CanonicalName() string {
	classInfo := self.ConstantPool[self.ThisClass].(*ConstantClassInfo)
	utf8Info := self.ConstantPool[classInfo.NameIndex].(*ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) Name() string {
	classInfo := self.ConstantPool[self.ThisClass].(*ConstantClassInfo)
	utf8Info := self.ConstantPool[classInfo.NameIndex].(*ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) IsClass() bool {
	return !(self.IsInterface() && self.IsEnum() && self.IsAnnotation())
}

func (self *JavaClass) Interfaces() []string {
	interfaceNames := make([]string, 0)
	for _, interfaceIndex := range self.ClassFile.Interfaces {
		classInfo := self.ConstantPool[interfaceIndex].(*ConstantClassInfo)
		utf8Info := self.ConstantPool[classInfo.NameIndex].(*ConstantUtf8Info)
		interfaceNames = append(interfaceNames, strings.Replace(utf8Info.JavaString(), "/", ".", -1))
	}
	return interfaceNames
}

func (self *JavaClass) SuperClass() string {
	classInfo := self.ConstantPool[self.ClassFile.SuperClass].(*ConstantClassInfo)
	utf8Info := self.ConstantPool[classInfo.NameIndex].(*ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) Fields() []*JavaField {
	fields := make([]*JavaField, 0)
	for _, fieldInfo := range self.ClassFile.Fields {
		fields = append(fields, NewJavaField(self.ConstantPool, fieldInfo))
	}
	return fields
}

func (self *JavaClass) Methods() []*JavaMethod {
	methods := make([]*JavaMethod, 0)
	for _, methodInfo := range self.ClassFile.Methods {
		methods = append(methods, NewJavaMethod(self.ConstantPool, methodInfo))
	}
	return methods
}
