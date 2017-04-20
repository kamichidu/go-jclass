package jclass

import (
	"github.com/kamichidu/go-jclass/jvms"
	"io"
	"os"
	"strings"
)

type JavaClass struct {
	*jvms.ClassFile
	AccessFlags

	fields  map[string]*JavaField
	methods map[string][]*JavaMethod
}

func NewJavaClass(classFile *jvms.ClassFile) *JavaClass {
	return &JavaClass{
		ClassFile:   classFile,
		AccessFlags: AccessFlag(classFile.AccessFlags),
	}
}

func NewJavaClassFromReader(r io.Reader) (*JavaClass, error) {
	cf, err := jvms.ParseClassFile(r)
	if err != nil {
		return nil, err
	}
	return NewJavaClass(cf), nil
}

func NewJavaClassFromFilename(filename string) (*JavaClass, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return NewJavaClassFromReader(file)
}

func (self *JavaClass) PackageName() string {
	items := strings.Split(self.Name(), ".")
	return strings.Join(items[:len(items)-1], ".")
}

func (self *JavaClass) CanonicalName() string {
	return strings.Replace(self.Name(), "$", ".", -1)
}

func (self *JavaClass) Name() string {
	classInfo := self.ConstantPool[self.ThisClass].(*jvms.ConstantClassInfo)
	utf8Info := self.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) SimpleName() string {
	items := strings.Split(self.CanonicalName(), ".")
	return items[len(items)-1]
}

func (self *JavaClass) IsClass() bool {
	return !(self.IsInterface() && self.IsEnum() && self.IsAnnotation())
}

func (self *JavaClass) Interfaces() []string {
	interfaceNames := make([]string, 0)
	for _, interfaceIndex := range self.ClassFile.Interfaces {
		classInfo := self.ConstantPool[interfaceIndex].(*jvms.ConstantClassInfo)
		utf8Info := self.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
		interfaceNames = append(interfaceNames, strings.Replace(utf8Info.JavaString(), "/", ".", -1))
	}
	return interfaceNames
}

func (self *JavaClass) SuperClass() string {
	classInfo, ok := self.ConstantPool[self.ClassFile.SuperClass].(*jvms.ConstantClassInfo)
	if !ok {
		return ""
	}
	utf8Info := self.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) Fields() []*JavaField {
	fields := make([]*JavaField, 0)
	for _, fieldInfo := range self.ClassFile.Fields {
		fields = append(fields, NewJavaField(self.ConstantPool, fieldInfo))
	}
	return fields
}

func (self *JavaClass) Field(name string) *JavaField {
	if self.fields == nil {
		self.fields = make(map[string]*JavaField)
		for _, field := range self.Fields() {
			self.fields[field.Name()] = field
		}
	}

	if field, found := self.fields[name]; found {
		return field
	} else {
		return nil
	}
}

func (self *JavaClass) Methods() []*JavaMethod {
	methods := make([]*JavaMethod, 0)
	for _, methodInfo := range self.ClassFile.Methods {
		methods = append(methods, NewJavaMethod(self.ConstantPool, methodInfo))
	}
	return methods
}

func (self *JavaClass) Method(name string) []*JavaMethod {
	if self.methods == nil {
		self.methods = make(map[string][]*JavaMethod)
		for _, method := range self.Methods() {
			self.methods[method.Name()] = append(self.methods[method.Name()], method)
		}
	}

	if methods, found := self.methods[name]; found {
		return methods
	} else {
		return make([]*JavaMethod, 0)
	}
}

func (self *JavaClass) SourceFile() string {
	for _, attr := range self.ClassFile.Attributes {
		sourceFile, ok := attr.(*jvms.SourceFileAttribute)
		if !ok {
			continue
		}
		utf8Info := self.ConstantPool[sourceFile.SourceFileIndex].(*jvms.ConstantUtf8Info)
		return utf8Info.JavaString()
	}
	return ""
}
