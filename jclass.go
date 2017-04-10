package jclass

import (
	"strings"
)

type JavaClass struct {
	*ClassFile
}

func NewJavaClass(classFile *ClassFile) *JavaClass {
	return &JavaClass{classFile}
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

func (self *JavaClass) IsPublic() bool {
	return self.AccessFlags&ACC_PUBLIC == ACC_PUBLIC
}

func (self *JavaClass) IsProtected() bool {
	return self.AccessFlags&ACC_PROTECTED == ACC_PROTECTED
}

func (self *JavaClass) IsPrivate() bool {
	return self.AccessFlags&ACC_PRIVATE == ACC_PRIVATE
}

func (self *JavaClass) IsFinal() bool {
	return self.AccessFlags&ACC_FINAL == ACC_FINAL
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
	return utf8Info.JavaString()
}
