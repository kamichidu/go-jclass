package jclass

import (
	"bufio"
	"fmt"
	"github.com/kamichidu/go-jclass/data"
	"io"
	"os"
	"reflect"
	"strings"
)

type JClass struct {
	data data.ClassFile
}

func NewJClass(in io.Reader) (*JClass, error) {
	var err error

	reader := bufio.NewReader(in)

	classFile, err := parseClassFile(reader)
	if err != nil {
		return nil, err
	}

	jclass := &JClass{
		data: *classFile,
	}

	return jclass, nil
}

func NewJClassWithFilename(filename string) (*JClass, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return NewJClass(file)
}

func (self *JClass) GetAccessFlags() uint16 {
	return self.data.AccessFlags
}

func (self *JClass) GetPackageName() string {
	bname := self.GetClassName()
	idx := strings.LastIndex(bname, "/")
	if idx == -1 {
		return ""
	}

	return bname[0:idx]
}

func (self *JClass) GetClassName() string {
	classInfo := getClassInfo(self.data.ConstantPool, self.data.ThisClass)
	return getUtf8String(self.data.ConstantPool, classInfo.NameIndex)
}

func (self *JClass) GetSuperclass() string {
	if self.data.SuperClass == 0 {
		return ""
	}
	classInfo := getClassInfo(self.data.ConstantPool, self.data.SuperClass)
	return getUtf8String(self.data.ConstantPool, classInfo.NameIndex)
}

func (self *JClass) GetInterfaces() []string {
	interfaces := make([]string, self.data.InterfacesCount)
	for i := uint16(0); i < self.data.InterfacesCount; i++ {
		classInfo := getClassInfo(self.data.ConstantPool, self.data.Interfaces[i])
		interfaces[i] = getUtf8String(self.data.ConstantPool, classInfo.NameIndex)
	}
	return interfaces
}

func (self *JClass) GetFields() []*JField {
	fields := make([]*JField, self.data.FieldsCount)
	for i := uint16(0); i < self.data.FieldsCount; i++ {
		fields[i] = newJField(self.data.ConstantPool, &self.data.Fields[i])
	}
	return fields
}

func (self *JClass) GetField(name string) *JField {
	for _, field := range self.GetFields() {
		if field.GetName() == name {
			return field
		}
	}
	return nil
}

func (self *JClass) GetMethods() []*JMethod {
	methods := make([]*JMethod, self.data.MethodsCount)
	for i := uint16(0); i < self.data.MethodsCount; i++ {
		methods[i] = newJMethod(self, &self.data.Methods[i])
	}
	return methods
}

func (self *JClass) GetAttributes() []data.AttributeInfo {
	return self.data.Attributes
}

func (self *JClass) GetAttribute(typ reflect.Type) data.AttributeInfo {
	for _, attr := range self.GetAttributes() {
		if reflect.TypeOf(attr).AssignableTo(typ) {
			return attr
		}
	}
	return nil
}

type JInnerClass struct {
	name           string
	simpleName     string
	accessFlags    uint16
	outerClassName string
}

func (self *JInnerClass) GetName() string {
	return self.name
}

func (self *JInnerClass) GetSimpleName() string {
	return self.simpleName
}

func (self *JInnerClass) GetAccessFlags() uint16 {
	return self.accessFlags
}

func (self *JInnerClass) GetOuterClassName() string {
	return self.outerClassName
}

func (self *JClass) GetInnerClasses() []*JInnerClass {
	var attr *data.InnerClassesAttribute
	if found := self.GetAttribute(reflect.TypeOf(attr)); found != nil {
		attr, _ = found.(*data.InnerClassesAttribute)
	} else {
		return make([]*JInnerClass, 0)
	}
	inners := make([]*JInnerClass, attr.NumberOfClasses)
	for i, class := range attr.Classes {
		ici := getClassInfo(self.data.ConstantPool, class.InnerClassInfoIndex)
		inner := &JInnerClass{}
		inner.name = getUtf8String(self.data.ConstantPool, ici.NameIndex)
		inner.accessFlags = class.InnerClassAccessFlags
		if class.InnerNameIndex != 0 {
			inner.simpleName = getUtf8String(self.data.ConstantPool, class.InnerNameIndex)
		}
		if class.OuterClassInfoIndex != 0 {
			oci := getClassInfo(self.data.ConstantPool, class.OuterClassInfoIndex)
			inner.outerClassName = getUtf8String(self.data.ConstantPool, oci.NameIndex)
		}
		inners[i] = inner
	}
	return inners
}

func (self *JClass) GetSourceFile() string {
	var attr *data.SourceFileAttribute
	if found := self.GetAttribute(reflect.TypeOf(attr)); found != nil {
		attr, _ = found.(*data.SourceFileAttribute)
	} else {
		return ""
	}
	return getUtf8String(self.data.ConstantPool, attr.SourcefileIndex)
}

func getClassInfo(cp []data.CpInfo, index uint16) data.ClassInfo {
	classInfo, ok := cp[index].(data.ClassInfo)
	if !ok {
		panic(fmt.Sprintf("given index (%d) indicates a invalid cp_info", index))
	}
	return classInfo
}

func getUtf8String(cp []data.CpInfo, index uint16) string {
	utf8Info, ok := cp[index].(data.Utf8Info)
	if !ok {
		panic(fmt.Sprintf("given index (%d) indicates a invalid cp_info", index))
	}
	return string(utf8Info.Bytes)
}
