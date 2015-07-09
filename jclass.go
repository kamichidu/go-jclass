package jclass

import (
	"bufio"
	"fmt"
	"github.com/kamichidu/go-jclass/data"
	"github.com/kamichidu/go-jclass/parser/md"
	"io"
	"os"
	"reflect"
	"strings"
)

const (
	ACC_PUBLIC       = 0x0001
	ACC_PRIVATE      = 0x0002
	ACC_PROTECTED    = 0x0004
	ACC_STATIC       = 0x0008
	ACC_FINAL        = 0x0010
	ACC_SUPER        = 0x0020
	ACC_SYNCHRONIZED = 0x0020
	ACC_BRIDGE       = 0x0040
	ACC_VOLATILE     = 0x0040
	ACC_TRANSIENT    = 0x0080
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_INTERFACE    = 0x0200
	ACC_ABSTRACT     = 0x0400
	ACC_STRICT       = 0x0800
	ACC_SYNTHETIC    = 0x1000
	ACC_ANNOTATION   = 0x2000
	ACC_ENUM         = 0x4000
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

func (self *JClass) IsPublic() bool {
	return (self.GetAccessFlags() & ACC_PUBLIC) == ACC_PUBLIC
}

func (self *JClass) IsFinal() bool {
	return (self.GetAccessFlags() & ACC_FINAL) == ACC_FINAL
}

func (self *JClass) IsSuper() bool {
	return (self.GetAccessFlags() & ACC_SUPER) == ACC_SUPER
}

func (self *JClass) IsInterface() bool {
	return (self.GetAccessFlags() & ACC_INTERFACE) == ACC_INTERFACE
}

func (self *JClass) IsAbstract() bool {
	return (self.GetAccessFlags() & ACC_ABSTRACT) == ACC_ABSTRACT
}

func (self *JClass) IsSynthetic() bool {
	return (self.GetAccessFlags() & ACC_SYNTHETIC) == ACC_SYNTHETIC
}

func (self *JClass) IsAnnotation() bool {
	return (self.GetAccessFlags() & ACC_ANNOTATION) == ACC_ANNOTATION
}

func (self *JClass) IsEnum() bool {
	return (self.GetAccessFlags() & ACC_ENUM) == ACC_ENUM
}

func (self *JClass) GetPackageName() string {
	name := self.GetName()
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return ""
	}

	return name[0:idx]
}

func (self *JClass) GetName() string {
	classInfo := getClassInfo(self.data.ConstantPool, self.data.ThisClass)
	name := getUtf8String(self.data.ConstantPool, classInfo.NameIndex)
	name = strings.Replace(name, "/", ".", -1)
	return name
}

func (self *JClass) GetCanonicalName() string {
	name := self.GetName()
	name = strings.Replace(name, "$", ".", -1)
	return name
}

func (self *JClass) GetSimpleName() string {
	name := self.GetName()
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return name
	}
	return name[idx+1:]
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

func (self *JClass) GetMethod(name string) []*JMethod {
	methods := make([]*JMethod, 0)
	for _, m := range self.GetMethods() {
		if m.GetName() == name {
			methods = append(methods, m)
		}
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

func (self *JClass) GetEnclosingClass() string {
	var attr *data.EnclosingMethodAttribute
	if found := self.GetAttribute(reflect.TypeOf(attr)); found != nil {
		attr, _ = found.(*data.EnclosingMethodAttribute)
	} else {
		return ""
	}
	classInfo := getClassInfo(self.data.ConstantPool, attr.ClassIndex)
	return getUtf8String(self.data.ConstantPool, classInfo.NameIndex)
}

func (self *JClass) GetEnclosingMethod() (string, []string, string) {
	var attr *data.EnclosingMethodAttribute
	if found := self.GetAttribute(reflect.TypeOf(attr)); found != nil {
		attr, _ = found.(*data.EnclosingMethodAttribute)
	} else {
		return "", []string{}, ""
	}
	if attr.MethodIndex == uint16(0) {
		return "", []string{}, ""
	}
	if nameAndTypeInfo, ok := self.data.ConstantPool[attr.MethodIndex].(*data.NameAndTypeInfo); ok {
		name := getUtf8String(self.data.ConstantPool, nameAndTypeInfo.NameIndex)
		descriptor := getUtf8String(self.data.ConstantPool, nameAndTypeInfo.DescriptorIndex)

		paramTypes, retType, _, err := md.Parse(descriptor)
		if err != nil {
			panic("Internal error!")
		}
		return name, paramTypes, retType
	} else {
		panic("Internal error!")
	}
}

func (self *JClass) GetSignature() string {
	var attr *data.SignatureAttribute
	if found := self.GetAttribute(reflect.TypeOf(attr)); found != nil {
		attr, _ = found.(*data.SignatureAttribute)
	} else {
		return ""
	}
	return getUtf8String(self.data.ConstantPool, attr.SignatureIndex)
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
