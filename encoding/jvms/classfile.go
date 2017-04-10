package jvms

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/kamichidu/go-jclass"
)

func ParseClassFile(r io.Reader) (*jclass.ClassFile, error) {
	var err error
	cf := new(jclass.ClassFile)
	for _, v := range []interface{}{&cf.Magic, &cf.MinorVersion, &cf.MajorVersion, &cf.ConstantPoolCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	if cf.Magic != 0xcafebabe {
		return nil, fmt.Errorf("Illegal magic binary: %v", cf.Magic)
	}
	// Constant pool starts with index 1
	cf.ConstantPool = make([]jclass.ConstantPoolInfo, cf.ConstantPoolCount)
	for i := uint16(1); i < cf.ConstantPoolCount; i++ {
		cpInfo, err := ParseConstantPool(r)
		if err != nil {
			return nil, err
		}
		cf.ConstantPool[i] = cpInfo

		// Some cp_info consumes double indices
		switch cpInfo.Tag() {
		case jclass.CONSTANT_Long, jclass.CONSTANT_Double:
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
	cf.Fields = make([]*jclass.FieldInfo, cf.FieldsCount)
	for i := uint16(0); i < cf.FieldsCount; i++ {
		if cf.Fields[i], err = ParseFieldInfo(r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.MethodsCount); err != nil {
		return nil, err
	}
	cf.Methods = make([]*jclass.MethodInfo, cf.MethodsCount)
	for i := uint16(0); i < cf.MethodsCount; i++ {
		if cf.Methods[i], err = ParseMethodInfo(r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.AttributesCount); err != nil {
		return nil, err
	}
	cf.Attributes = make([]*jclass.AttributeInfo, cf.AttributesCount)
	for i := uint16(0); i < cf.AttributesCount; i++ {
		if cf.Attributes[i], err = ParseAttributeInfo(r); err != nil {
			return nil, err
		}
	}
	return cf, nil
}

func ParseConstantPool(r io.Reader) (jclass.ConstantPoolInfo, error) {
	var (
		tag uint8
		err error
	)
	if tag, err = u1(r); err != nil {
		return nil, err
	}

	var (
		cpInfo jclass.ConstantPoolInfo
		data   []interface{}
	)
	switch tag {
	case jclass.CONSTANT_Class:
		info := new(jclass.ConstantClassInfo)
		data = []interface{}{
			&info.NameIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_Fieldref:
		info := new(jclass.ConstantFieldrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_Methodref:
		info := new(jclass.ConstantMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_InterfaceMethodref:
		info := new(jclass.ConstantInterfaceMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_String:
		info := new(jclass.ConstantStringInfo)
		data = []interface{}{
			&info.StringIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_Integer:
		info := new(jclass.ConstantIntegerInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case jclass.CONSTANT_Float:
		info := new(jclass.ConstantFloatInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case jclass.CONSTANT_Long:
		info := new(jclass.ConstantLongInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case jclass.CONSTANT_Double:
		info := new(jclass.ConstantDoubleInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case jclass.CONSTANT_NameAndType:
		info := new(jclass.ConstantNameAndTypeInfo)
		data = []interface{}{
			&info.NameIndex,
			&info.DescriptorIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_MethodHandle:
		info := new(jclass.ConstantMethodHandleInfo)
		data = []interface{}{
			&info.ReferenceKind,
			&info.ReferenceIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_MethodType:
		info := new(jclass.ConstantMethodTypeInfo)
		data = []interface{}{
			&info.DescriptorIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_InvokeDynamic:
		info := new(jclass.ConstantInvokeDynamicInfo)
		data = []interface{}{
			&info.BootstrapMethodAttrIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case jclass.CONSTANT_Utf8:
		info := new(jclass.ConstantUtf8Info)
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

func ParseFieldInfo(r io.Reader) (*jclass.FieldInfo, error) {
	var err error
	fi := new(jclass.FieldInfo)
	for _, v := range []interface{}{&fi.AccessFlags, &fi.NameIndex, &fi.DescriptorIndex, &fi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	fi.Attributes = make([]*jclass.AttributeInfo, fi.AttributesCount)
	for i := uint16(0); i < fi.AttributesCount; i++ {
		if fi.Attributes[i], err = ParseAttributeInfo(r); err != nil {
			return nil, err
		}
	}
	return fi, nil
}

func ParseMethodInfo(r io.Reader) (*jclass.MethodInfo, error) {
	var err error
	mi := new(jclass.MethodInfo)
	for _, v := range []interface{}{&mi.AccessFlags, &mi.NameIndex, &mi.DescriptorIndex, &mi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	mi.Attributes = make([]*jclass.AttributeInfo, mi.AttributesCount)
	for i := uint16(0); i < mi.AttributesCount; i++ {
		if mi.Attributes[i], err = ParseAttributeInfo(r); err != nil {
			return nil, err
		}
	}
	return mi, nil
}

func ParseAttributeInfo(r io.Reader) (*jclass.AttributeInfo, error) {
	var err error
	ai := new(jclass.AttributeInfo)
	for _, v := range []interface{}{&ai.AttributeNameIndex, &ai.AttributeLength} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	ai.Info = make([]uint8, ai.AttributeLength)
	err = binary.Read(r, binary.BigEndian, &ai.Info)
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
