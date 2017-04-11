package jvms

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/kamichidu/go-jclass/jvms"
)

func ParseClassFile(r io.Reader) (*jvms.ClassFile, error) {
	var err error
	cf := new(jvms.ClassFile)
	for _, v := range []interface{}{&cf.Magic, &cf.MinorVersion, &cf.MajorVersion, &cf.ConstantPoolCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	if cf.Magic != 0xcafebabe {
		return nil, fmt.Errorf("Illegal magic binary: %v", cf.Magic)
	}
	// Constant pool starts with index 1
	cf.ConstantPool = make([]jvms.ConstantPoolInfo, cf.ConstantPoolCount)
	for i := uint16(1); i < cf.ConstantPoolCount; i++ {
		cpInfo, err := ParseConstantPool(r)
		if err != nil {
			return nil, err
		}
		cf.ConstantPool[i] = cpInfo

		// Some cp_info consumes double indices
		switch cpInfo.Tag() {
		case jvms.CONSTANT_Long, jvms.CONSTANT_Double:
			i++
			cf.ConstantPool[i] = cpInfo
		}
	}
	for _, v := range []interface{}{&cf.AccessFlags, &cf.ThisClass, &cf.SuperClass, &cf.InterfacesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	cf.Interfaces = make([]uint16, cf.InterfacesCount)
	if err = binary.Read(r, binary.BigEndian, &cf.Interfaces); err != nil {
		return nil, err
	}
	if err = binary.Read(r, binary.BigEndian, &cf.FieldsCount); err != nil {
		return nil, err
	}
	cf.Fields = make([]*jvms.FieldInfo, cf.FieldsCount)
	for i := uint16(0); i < cf.FieldsCount; i++ {
		if cf.Fields[i], err = ParseFieldInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.MethodsCount); err != nil {
		return nil, err
	}
	cf.Methods = make([]*jvms.MethodInfo, cf.MethodsCount)
	for i := uint16(0); i < cf.MethodsCount; i++ {
		if cf.Methods[i], err = ParseMethodInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.AttributesCount); err != nil {
		return nil, err
	}
	cf.Attributes = make([]jvms.AttributeInfo, cf.AttributesCount)
	for i := uint16(0); i < cf.AttributesCount; i++ {
		if cf.Attributes[i], err = ParseAttributeInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	return cf, nil
}

func ParseConstantPool(r io.Reader) (jvms.ConstantPoolInfo, error) {
	var (
		tag uint8
		err error
	)
	if tag, err = u1(r); err != nil {
		return nil, err
	}

	var (
		cpInfo jvms.ConstantPoolInfo
		data   []interface{}
	)
	switch tag {
	case jvms.CONSTANT_Class:
		info := new(jvms.ConstantClassInfo)
		data = []interface{}{
			&info.NameIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_Fieldref:
		info := new(jvms.ConstantFieldrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_Methodref:
		info := new(jvms.ConstantMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_InterfaceMethodref:
		info := new(jvms.ConstantInterfaceMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_String:
		info := new(jvms.ConstantStringInfo)
		data = []interface{}{
			&info.StringIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_Integer:
		info := new(jvms.ConstantIntegerInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case jvms.CONSTANT_Float:
		info := new(jvms.ConstantFloatInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case jvms.CONSTANT_Long:
		info := new(jvms.ConstantLongInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case jvms.CONSTANT_Double:
		info := new(jvms.ConstantDoubleInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case jvms.CONSTANT_NameAndType:
		info := new(jvms.ConstantNameAndTypeInfo)
		data = []interface{}{
			&info.NameIndex,
			&info.DescriptorIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_MethodHandle:
		info := new(jvms.ConstantMethodHandleInfo)
		data = []interface{}{
			&info.ReferenceKind,
			&info.ReferenceIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_MethodType:
		info := new(jvms.ConstantMethodTypeInfo)
		data = []interface{}{
			&info.DescriptorIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_InvokeDynamic:
		info := new(jvms.ConstantInvokeDynamicInfo)
		data = []interface{}{
			&info.BootstrapMethodAttrIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jvms.CONSTANT_Utf8:
		info := new(jvms.ConstantUtf8Info)
		if info.Length, err = u2(r); err != nil {
			return nil, err
		}
		info.Bytes = make([]uint8, info.Length)
		if err = binary.Read(r, binary.BigEndian, &info.Bytes); err != nil {
			return nil, err
		}
		return info, nil
	default:
		return nil, fmt.Errorf("Read unknown constant pool tag: %d", tag)
	}
	for _, v := range data {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	return cpInfo, err
}

func ParseFieldInfo(constantPool []jvms.ConstantPoolInfo, r io.Reader) (*jvms.FieldInfo, error) {
	var err error
	fi := new(jvms.FieldInfo)
	for _, v := range []interface{}{&fi.AccessFlags, &fi.NameIndex, &fi.DescriptorIndex, &fi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	fi.Attributes = make([]jvms.AttributeInfo, fi.AttributesCount)
	for i := uint16(0); i < fi.AttributesCount; i++ {
		if fi.Attributes[i], err = ParseAttributeInfo(constantPool, r); err != nil {
			return nil, err
		}
	}
	return fi, nil
}

func ParseMethodInfo(constantPool []jvms.ConstantPoolInfo, r io.Reader) (*jvms.MethodInfo, error) {
	var err error
	mi := new(jvms.MethodInfo)
	for _, v := range []interface{}{&mi.AccessFlags, &mi.NameIndex, &mi.DescriptorIndex, &mi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	mi.Attributes = make([]jvms.AttributeInfo, mi.AttributesCount)
	for i := uint16(0); i < mi.AttributesCount; i++ {
		if mi.Attributes[i], err = ParseAttributeInfo(constantPool, r); err != nil {
			return nil, err
		}
	}
	return mi, nil
}

func ParseAttributeInfo(constantPool []jvms.ConstantPoolInfo, r io.Reader) (jvms.AttributeInfo, error) {
	var err error
	var (
		attributeNameIndex uint16
		attributeLength    uint32
	)
	for _, v := range []interface{}{&attributeNameIndex, &attributeLength} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}

	utf8Info := constantPool[attributeNameIndex].(*jvms.ConstantUtf8Info)
	var ai jvms.AttributeInfo
	switch utf8Info.JavaString() {
	default:
		gai := &jvms.GenericAttributeInfo{
			AttributeNameIndex_: attributeNameIndex,
			AttributeLength_:    attributeLength,
			Info_:               make([]uint8, attributeLength),
		}
		ai = gai
		err = binary.Read(r, binary.BigEndian, &gai.Info_)
	}
	return ai, err
}

func u1(r io.Reader) (uint8, error) {
	var data uint8
	err := binary.Read(r, binary.BigEndian, &data)
	return data, err
}

func u2(r io.Reader) (uint16, error) {
	var data uint16
	err := binary.Read(r, binary.BigEndian, &data)
	return data, err
}

func u4(r io.Reader) (uint32, error) {
	var data uint32
	err := binary.Read(r, binary.BigEndian, &data)
	return data, err
}
