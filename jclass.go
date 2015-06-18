package jclass

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kamichidu/go-jclass/data"
	"io"
)

func parseU1(in *bufio.Reader) (uint8, error) {
	value, err := in.ReadByte()
	if err != nil {
		panic(err)
		//return value, err
	}
	return value, nil
}

func parseU2(in *bufio.Reader) (uint16, error) {
	high, err := parseU1(in)
	if err != nil {
		return 0, err
	}
	low, err := parseU1(in)
	if err != nil {
		return 0, err
	}
	value := uint16(high)<<8 | uint16(low)
	return value, nil
}

func parseU4(in *bufio.Reader) (uint32, error) {
	high, err := parseU2(in)
	if err != nil {
		return 0, err
	}
	low, err := parseU2(in)
	if err != nil {
		return 0, err
	}
	value := uint32(high)<<16 | uint32(low)
	return value, nil
}

func parseCpInfo(in *bufio.Reader) (data.CpInfo, error) {
	tag, err := parseU1(in)
	if err != nil {
		return nil, err
	}

	switch tag {
	case data.CpInfo_Tag_Class:
		cpInfo := data.ClassInfo{Tag: tag}
		cpInfo.NameIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Fieldref:
		cpInfo := data.FieldrefInfo{Tag: tag}
		cpInfo.ClassIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.NameAndTypeIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Methodref:
		cpInfo := data.MethodrefInfo{Tag: tag}
		cpInfo.ClassIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.NameAndTypeIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_InterfaceMethodref:
		cpInfo := data.InterfaceMethodrefInfo{Tag: tag}
		cpInfo.ClassIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.NameAndTypeIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_String:
		cpInfo := data.StringInfo{Tag: tag}
		cpInfo.StringIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Integer:
		cpInfo := data.IntegerInfo{Tag: tag}
		cpInfo.Bytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Float:
		cpInfo := data.FloatInfo{Tag: tag}
		cpInfo.Bytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Long:
		cpInfo := data.LongInfo{Tag: tag}
		cpInfo.HighBytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		cpInfo.LowBytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Double:
		cpInfo := data.DoubleInfo{Tag: tag}
		cpInfo.HighBytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		cpInfo.LowBytes, err = parseU4(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_NameAndType:
		cpInfo := data.NameAndTypeInfo{Tag: tag}
		cpInfo.NameIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.DescriptorIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_Utf8:
		cpInfo := data.Utf8Info{Tag: tag}
		cpInfo.Length, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.Bytes = make([]uint8, cpInfo.Length)
		for i := uint16(0); i < cpInfo.Length; i++ {
			b, err := in.ReadByte()
			if err != nil {
				return nil, err
			}
			cpInfo.Bytes[i] = b
		}
		return cpInfo, nil
	case data.CpInfo_Tag_MethodHandle:
		cpInfo := data.MethodHandleInfo{Tag: tag}
		cpInfo.ReferenceKind, err = parseU1(in)
		if err != nil {
			return nil, err
		}
		cpInfo.ReferenceIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_MethodType:
		cpInfo := data.MethodTypeInfo{Tag: tag}
		cpInfo.DescriptorIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	case data.CpInfo_Tag_InvokeDynamic:
		cpInfo := data.InvokeDynamicInfo{Tag: tag}
		cpInfo.BootstrapMethodAttrIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		cpInfo.NameAndTypeIndex, err = parseU2(in)
		if err != nil {
			return nil, err
		}
		return cpInfo, nil
	}
	return nil, errors.New(fmt.Sprintf("Illegal tag value detected: %d", tag))
}

func parseFieldInfo(in *bufio.Reader) (*data.FieldInfo, error) {
	var err error

	fieldInfo := &data.FieldInfo{}
	fieldInfo.AccessFlags, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	fieldInfo.NameIndex, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	fieldInfo.DescriptorIndex, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	fieldInfo.AttributesCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	fieldInfo.Attributes = make([]data.AttributeInfo, fieldInfo.AttributesCount)
	for i := uint16(0); i < fieldInfo.AttributesCount; i++ {
		attributeInfo, err := parseAttributeInfo(in)
		if err != nil {
			return nil, err
		}
		fieldInfo.Attributes[i] = *attributeInfo
	}
	return fieldInfo, nil
}

func parseMethodInfo(in *bufio.Reader) (*data.MethodInfo, error) {
	var err error

	methodInfo := &data.MethodInfo{}
	methodInfo.AccessFlags, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	methodInfo.NameIndex, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	methodInfo.DescriptorIndex, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	methodInfo.AttributesCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	methodInfo.Attributes = make([]data.AttributeInfo, methodInfo.AttributesCount)
	for i := uint16(0); i < methodInfo.AttributesCount; i++ {
		attributeInfo, err := parseAttributeInfo(in)
		if err != nil {
			return nil, err
		}
		methodInfo.Attributes[i] = *attributeInfo
	}
	return methodInfo, nil
}

func parseAttributeInfo(in *bufio.Reader) (*data.AttributeInfo, error) {
	var err error

	attributeInfo := &data.AttributeInfo{}
	if err != nil {
		return nil, err
	}
	attributeInfo.AttributeNameIndex, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	attributeInfo.AttributeLength, err = parseU4(in)
	if err != nil {
		return nil, err
	}
	attributeInfo.Info = make([]uint8, attributeInfo.AttributeLength)
	for i := uint32(0); i < attributeInfo.AttributeLength; i++ {
		b, err := parseU1(in)
		if err != nil {
			return nil, err
		}
		attributeInfo.Info[i] = b
	}
	return attributeInfo, nil
}

func parseClassFile(in *bufio.Reader) (*data.ClassFile, error) {
	var err error

	classFile := &data.ClassFile{}
	classFile.Magic, err = parseU4(in)
	if err != nil {
		return nil, err
	} else if classFile.Magic != 0xcafebabe {
		return nil, errors.New("It's not a java class file")
	}
	classFile.MinorVersion, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.MajorVersion, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.ConstantPoolCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	// ConstantPool index starts with 1, make padding
	classFile.ConstantPool = make([]data.CpInfo, classFile.ConstantPoolCount)
	for i := uint16(1); i < classFile.ConstantPoolCount; i++ {
		cpInfo, err := parseCpInfo(in)
		if err != nil {
			return nil, err
		}
		classFile.ConstantPool[i] = cpInfo

		// LongInfo and DoubleInfo use 2 spaces
		switch cpInfo.(type) {
		case data.LongInfo, data.DoubleInfo:
			i++
			classFile.ConstantPool[i] = cpInfo
		}
	}
	classFile.AccessFlags, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.ThisClass, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.SuperClass, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.InterfacesCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.Interfaces = make([]uint16, classFile.InterfacesCount)
	for i := uint16(0); i < classFile.InterfacesCount; i++ {
		value, err := parseU2(in)
		if err != nil {
			return nil, err
		}
		classFile.Interfaces[i] = value
	}
	classFile.FieldsCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.Fields = make([]data.FieldInfo, classFile.FieldsCount)
	for i := uint16(0); i < classFile.FieldsCount; i++ {
		fieldInfo, err := parseFieldInfo(in)
		if err != nil {
			return nil, err
		}
		classFile.Fields[i] = *fieldInfo
	}
	classFile.MethodsCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.Methods = make([]data.MethodInfo, classFile.MethodsCount)
	for i := uint16(0); i < classFile.MethodsCount; i++ {
		methodInfo, err := parseMethodInfo(in)
		if err != nil {
			return nil, err
		}
		classFile.Methods[i] = *methodInfo
	}
	classFile.AttributesCount, err = parseU2(in)
	if err != nil {
		return nil, err
	}
	classFile.Attributes = make([]data.AttributeInfo, classFile.AttributesCount)
	for i := uint16(0); i < classFile.AttributesCount; i++ {
		attributeInfo, err := parseAttributeInfo(in)
		if err != nil {
			return nil, err
		}
		classFile.Attributes[i] = *attributeInfo
	}

	return classFile, nil
}

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

func (self *JClass) getClassInfo(index uint16) data.ClassInfo {
	classInfo, ok := self.data.ConstantPool[index].(data.ClassInfo)
	if !ok {
		panic(fmt.Sprintf("given index (%d) indicates a invalid cp_info", index))
	}
	return classInfo
}

func (self *JClass) getUtf8String(index uint16) string {
	utf8Info, ok := self.data.ConstantPool[index].(data.Utf8Info)
	if !ok {
		panic(fmt.Sprintf("given index (%d) indicates a invalid cp_info", index))
	}
	return string(utf8Info.Bytes)
}

func (self *JClass) GetAccessFlags() uint16 {
	return self.data.AccessFlags
}

func (self *JClass) GetThisClass() string {
	classInfo := self.getClassInfo(self.data.ThisClass)
	return self.getUtf8String(classInfo.NameIndex)
}

func (self *JClass) GetSuperclass() string {
	classInfo := self.getClassInfo(self.data.SuperClass)
	return self.getUtf8String(classInfo.NameIndex)
}

func (self *JClass) GetInterfaces() []string {
	interfaces := make([]string, self.data.InterfacesCount)
	for i := uint16(0); i < self.data.InterfacesCount; i++ {
		classInfo := self.getClassInfo(self.data.Interfaces[i])
		interfaces[i] = self.getUtf8String(classInfo.NameIndex)
	}
	return interfaces
}

func (self *JClass) GetFields() []*JField {
	fields := make([]*JField, self.data.FieldsCount)
	for i := uint16(0); i < self.data.FieldsCount; i++ {
		fields[i] = newJField(self, &self.data.Fields[i])
	}
	return fields
}

func (self *JClass) GetMethods() []*JMethod {
	methods := make([]*JMethod, self.data.MethodsCount)
	for i := uint16(0); i < self.data.MethodsCount; i++ {
		methods[i] = newJMethod(self, &self.data.Methods[i])
	}
	return methods
}

func (self *JClass) GetAttributes() []*JAttribute {
	attributes := make([]*JAttribute, self.data.AttributesCount)
	for i := uint16(0); i < self.data.AttributesCount; i++ {
		attributes[i] = newJAttributeWithJClass(self, &self.data.Attributes[i])
	}
	return attributes
}
