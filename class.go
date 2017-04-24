package jclass

import (
	"github.com/kamichidu/go-jclass/jvms"
	"io"
	"os"
	"strings"
)

type JavaClass struct {
	AccessFlags

	classFile *jvms.ClassFile
	fields    map[string]*JavaField
}

func NewJavaClass(classFile *jvms.ClassFile) *JavaClass {
	return &JavaClass{
		AccessFlags: AccessFlag(classFile.AccessFlags),
		classFile:   classFile,
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
	classInfo := self.classFile.ConstantPool[self.classFile.ThisClass].(*jvms.ConstantClassInfo)
	utf8Info := self.classFile.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
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
	for _, interfaceIndex := range self.classFile.Interfaces {
		classInfo := self.classFile.ConstantPool[interfaceIndex].(*jvms.ConstantClassInfo)
		utf8Info := self.classFile.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
		interfaceNames = append(interfaceNames, strings.Replace(utf8Info.JavaString(), "/", ".", -1))
	}
	return interfaceNames
}

func (self *JavaClass) SuperClass() string {
	classInfo, ok := self.classFile.ConstantPool[self.classFile.SuperClass].(*jvms.ConstantClassInfo)
	if !ok {
		return ""
	}
	utf8Info := self.classFile.ConstantPool[classInfo.NameIndex].(*jvms.ConstantUtf8Info)
	return strings.Replace(utf8Info.JavaString(), "/", ".", -1)
}

func (self *JavaClass) Fields() []*JavaField {
	fields := make([]*JavaField, 0)
	for _, fieldInfo := range self.classFile.Fields {
		fields = append(fields, newJavaField(self.classFile.ConstantPool, fieldInfo))
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

func (self *JavaClass) Constructors() []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		switch method.Name() {
		case "<init>":
			return method.IsPublic()
		default:
			return false
		}
	})
}

func (self *JavaClass) DeclaredConstructors() []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		switch method.Name() {
		case "<init>":
			return true
		default:
			return false
		}
	})
}

func (self *JavaClass) Methods() []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		switch method.Name() {
		case "<init>", "<clinit>":
			return false
		default:
			return method.IsPublic()
		}
	})
}

func (self *JavaClass) DeclaredMethods() []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		switch method.Name() {
		case "<init>", "<clinit>":
			return false
		default:
			return true
		}
	})
}

func (self *JavaClass) Method(name string) []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		return method.IsPublic() && method.Name() == name
	})
}

func (self *JavaClass) DeclaredMethod(name string) []*JavaMethod {
	return self.filterMethods(func(method *JavaMethod) bool {
		return method.Name() == name
	})
}

func (self *JavaClass) filterMethods(filter func(method *JavaMethod) bool) []*JavaMethod {
	methods := make([]*JavaMethod, 0)
	for _, methodInfo := range self.classFile.Methods {
		method := newJavaMethod(self.classFile.ConstantPool, methodInfo)
		if filter(method) {
			methods = append(methods, method)
		}
	}
	return methods
}

func (self *JavaClass) SourceFile() string {
	for _, attr := range self.classFile.Attributes {
		sourceFile, ok := attr.(*jvms.SourceFileAttribute)
		if !ok {
			continue
		}
		utf8Info := self.classFile.ConstantPool[sourceFile.SourceFileIndex].(*jvms.ConstantUtf8Info)
		return utf8Info.JavaString()
	}
	return ""
}

func (self *JavaClass) ClassFile() *jvms.ClassFile {
	return self.classFile
}
