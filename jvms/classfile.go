package jvms

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1
// ClassFile {
//     u4             magic;
//     u2             minor_version;
//     u2             major_version;
//     u2             constant_pool_count;
//     cp_info        constant_pool[constant_pool_count-1];
//     u2             access_flags;
//     u2             this_class;
//     u2             super_class;
//     u2             interfaces_count;
//     u2             interfaces[interfaces_count];
//     u2             fields_count;
//     field_info     fields[fields_count];
//     u2             methods_count;
//     method_info    methods[methods_count];
//     u2             attributes_count;
//     attribute_info attributes[attributes_count];
// }
type ClassFile struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      []ConstantPoolInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []uint16
	FieldsCount       uint16
	Fields            []*FieldInfo
	MethodsCount      uint16
	Methods           []*MethodInfo
	AttributesCount   uint16
	Attributes        []AttributeInfo
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
// cp_info {
//     u1 tag;
//     u1 info[];
// }
type ConstantPoolInfo interface {
	Tag() uint8
	Info() []uint8
}

// CONSTANT_Class_info {
//     u1 tag;
//     u2 name_index;
// }
type ConstantClassInfo struct {
	NameIndex uint16
}

func (self *ConstantClassInfo) Tag() uint8 {
	return CONSTANT_Class
}

func (self *ConstantClassInfo) Info() []uint8 {
	return u16_u8slice(self.NameIndex)
}

// CONSTANT_Fieldref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantFieldrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (self *ConstantFieldrefInfo) Tag() uint8 {
	return CONSTANT_Fieldref
}

func (self *ConstantFieldrefInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u16_u8slice(self.ClassIndex)...)
	info = append(info, u16_u8slice(self.NameAndTypeIndex)...)
	return info
}

// CONSTANT_Methodref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantMethodrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (self *ConstantMethodrefInfo) Tag() uint8 {
	return CONSTANT_Methodref
}

func (self *ConstantMethodrefInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u16_u8slice(self.ClassIndex)...)
	info = append(info, u16_u8slice(self.NameAndTypeIndex)...)
	return info
}

// CONSTANT_InterfaceMethodref_info {
//     u1 tag;
//     u2 class_index;
//     u2 name_and_type_index;
// }
type ConstantInterfaceMethodrefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (self *ConstantInterfaceMethodrefInfo) Tag() uint8 {
	return CONSTANT_InterfaceMethodref
}

func (self *ConstantInterfaceMethodrefInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u16_u8slice(self.ClassIndex)...)
	info = append(info, u16_u8slice(self.NameAndTypeIndex)...)
	return info
}

// CONSTANT_String_info {
//     u1 tag;
//     u2 string_index;
// }
type ConstantStringInfo struct {
	StringIndex uint16
}

func (self *ConstantStringInfo) Tag() uint8 {
	return CONSTANT_String
}

func (self *ConstantStringInfo) Info() []uint8 {
	return u16_u8slice(self.StringIndex)
}

// CONSTANT_Integer_info {
//     u1 tag;
//     u4 bytes;
// }
type ConstantIntegerInfo struct {
	Bytes uint32
}

func (self *ConstantIntegerInfo) Tag() uint8 {
	return CONSTANT_Integer
}

func (self *ConstantIntegerInfo) Info() []uint8 {
	return u32_u8slice(self.Bytes)
}

func (self *ConstantIntegerInfo) Boolean() bool {
	return self.Bytes != 0
}

func (self *ConstantIntegerInfo) Byte() int8 {
	return int8(self.Bytes)
}

func (self *ConstantIntegerInfo) Char() rune {
	return rune(self.Bytes)
}

func (self *ConstantIntegerInfo) Short() int16 {
	return int16(self.Bytes)
}

func (self *ConstantIntegerInfo) Int() int32 {
	return int32(self.Bytes)
}

// CONSTANT_Float_info {
//     u1 tag;
//     u4 bytes;
// }
type ConstantFloatInfo struct {
	Bytes uint32
}

func (self *ConstantFloatInfo) Tag() uint8 {
	return CONSTANT_Float
}

func (self *ConstantFloatInfo) Info() []uint8 {
	return u32_u8slice(self.Bytes)
}

func (self *ConstantFloatInfo) Float() float32 {
	return math.Float32frombits(self.Bytes)
}

// CONSTANT_Long_info {
//     u1 tag;
//     u4 high_bytes;
//     u4 low_bytes;
// }
type ConstantLongInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

func (self *ConstantLongInfo) Tag() uint8 {
	return CONSTANT_Long
}

func (self *ConstantLongInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u32_u8slice(self.HighBytes)...)
	info = append(info, u32_u8slice(self.LowBytes)...)
	return info
}

func (self *ConstantLongInfo) Long() int64 {
	return int64(uint64(self.HighBytes)<<32 | uint64(self.LowBytes))
}

// CONSTANT_Double_info {
//     u1 tag;
//     u4 high_bytes;
//     u4 low_bytes;
// }
type ConstantDoubleInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

func (self *ConstantDoubleInfo) Tag() uint8 {
	return CONSTANT_Double
}

func (self *ConstantDoubleInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u32_u8slice(self.HighBytes)...)
	info = append(info, u32_u8slice(self.LowBytes)...)
	return info
}

func (self *ConstantDoubleInfo) Double() float64 {
	return math.Float64frombits(uint64(self.HighBytes)<<32 | uint64(self.LowBytes))
}

// CONSTANT_NameAndType_info {
//     u1 tag;
//     u2 name_index;
//     u2 descriptor_index;
// }
type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) Tag() uint8 {
	return CONSTANT_NameAndType
}

func (self *ConstantNameAndTypeInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u16_u8slice(self.NameIndex)...)
	info = append(info, u16_u8slice(self.DescriptorIndex)...)
	return info
}

// CONSTANT_Utf8_info {
//     u1 tag;
//     u2 length;
//     u1 bytes[length];
// }
type ConstantUtf8Info struct {
	Length uint16
	Bytes  []uint8
}

func (self *ConstantUtf8Info) Tag() uint8 {
	return CONSTANT_Utf8
}

func (self *ConstantUtf8Info) Info() []uint8 {
	info := u16_u8slice(self.Length)
	info = append(info, self.Bytes...)
	return info
}

func (self *ConstantUtf8Info) JavaString() string {
	// TODO: Support java's modified utf8
	return string(self.Bytes)
}

// CONSTANT_MethodHandle_info {
//     u1 tag;
//     u1 reference_kind;
//     u2 reference_index;
// }
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (self *ConstantMethodHandleInfo) Tag() uint8 {
	return CONSTANT_MethodHandle
}

func (self *ConstantMethodHandleInfo) Info() []uint8 {
	info := []uint8{self.ReferenceKind}
	info = append(info, u16_u8slice(self.ReferenceIndex)...)
	return info
}

// CONSTANT_MethodType_info {
//     u1 tag;
//     u2 descriptor_index;
// }
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func (self *ConstantMethodTypeInfo) Tag() uint8 {
	return CONSTANT_MethodType
}

func (self *ConstantMethodTypeInfo) Info() []uint8 {
	return u16_u8slice(self.DescriptorIndex)
}

// CONSTANT_InvokeDynamic_info {
//     u1 tag;
//     u2 bootstrap_method_attr_index;
//     u2 name_and_type_index;
// }
type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) Tag() uint8 {
	return CONSTANT_InvokeDynamic
}

func (self *ConstantInvokeDynamicInfo) Info() []uint8 {
	info := make([]uint8, 0)
	info = append(info, u16_u8slice(self.BootstrapMethodAttrIndex)...)
	info = append(info, u16_u8slice(self.NameAndTypeIndex)...)
	return info
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5
// field_info {
//     u2             access_flags;
//     u2             name_index;
//     u2             descriptor_index;
//     u2             attributes_count;
//     attribute_info attributes[attributes_count];
// }
type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6
// method_info {
//     u2             access_flags;
//     u2             name_index;
//     u2             descriptor_index;
//     u2             attributes_count;
//     attribute_info attributes[attributes_count];
// }
type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7
// attribute_info {
//     u2 attribute_name_index;
//     u4 attribute_length;
//     u1 info[attribute_length];
// }
type AttributeInfo interface {
	AttributeNameIndex() uint16
	AttributeLength() uint32
	Info() []uint8
}

// TODO: AnnotationDefaultAttribute
// TODO: BootstrapMethodsAttribute
// TODO: CodeAttribute

type ConstantValueAttribute struct {
	AttributeNameIndex_ uint16
	AttributeLength_    uint32
	ConstantvalueIndex  uint16
}

func (self *ConstantValueAttribute) AttributeNameIndex() uint16 {
	return self.AttributeNameIndex_
}

func (self *ConstantValueAttribute) AttributeLength() uint32 {
	return self.AttributeLength_
}

func (self *ConstantValueAttribute) Info() []uint8 {
	return u16_u8slice(self.ConstantvalueIndex)
}

type DeprecatedAttribute struct {
	AttributeNameIndex_ uint16
	AttributeLength_    uint32
}

func (self *DeprecatedAttribute) AttributeNameIndex() uint16 {
	return self.AttributeNameIndex_
}

func (self *DeprecatedAttribute) AttributeLength() uint32 {
	return self.AttributeLength_
}

func (self *DeprecatedAttribute) Info() []uint8 {
	return []uint8{}
}

// TODO: EnclosingMethodAttribute
// TODO: ExceptionsAttribute
// TODO: InnerClassesAttribute
// TODO: LineNumberTableAttribute
// TODO: LocalVariableTableAttribute
// TODO: LocalVariableTypeTableAttribute
// TODO: MethodParametersAttribute
// TODO: RuntimeInvisibleAnnotationsAttribute
// TODO: RuntimeInvisibleParameterAnnotationsAttribute
// TODO: RuntimeInvisibleTypeAnnotationsAttribute
// TODO: RuntimeVisibleAnnotationsAttribute
// TODO: RuntimeVisibleParameterAnnotationsAttribute
// TODO: RuntimeVisibleTypeAnnotationsAttribute

type SignatureAttribute struct {
	AttributeNameIndex_ uint16
	AttributeLength_    uint32
	SignatureIndex      uint16
}

func (self *SignatureAttribute) AttributeNameIndex() uint16 {
	return self.AttributeNameIndex_
}

func (self *SignatureAttribute) AttributeLength() uint32 {
	return self.AttributeLength_
}

func (self *SignatureAttribute) Info() []uint8 {
	return u16_u8slice(self.SignatureIndex)
}

// TODO: SourceDebugExtensionAttribute

type SourceFileAttribute struct {
	AttributeNameIndex_ uint16
	AttributeLength_    uint32
	SourceFileIndex     uint16
}

func (self *SourceFileAttribute) AttributeNameIndex() uint16 {
	return self.AttributeNameIndex_
}

func (self *SourceFileAttribute) AttributeLength() uint32 {
	return self.AttributeLength_
}

func (self *SourceFileAttribute) Info() []uint8 {
	return u16_u8slice(self.SourceFileIndex)
}

// TODO: StackMapTableAttribute
// TODO: SyntheticAttribute

type GenericAttributeInfo struct {
	AttributeNameIndex_ uint16
	AttributeLength_    uint32
	Info_               []uint8
}

func (self *GenericAttributeInfo) AttributeNameIndex() uint16 {
	return self.AttributeNameIndex_
}

func (self *GenericAttributeInfo) AttributeLength() uint32 {
	return self.AttributeLength_
}

func (self *GenericAttributeInfo) Info() []uint8 {
	return self.Info_
}

// Utilities
func u16_u8slice(n uint16) []uint8 {
	b := make([]uint8, 2)
	b[0] = uint8((n & 0xff00) >> 8)
	b[1] = uint8((n & 0x00ff) >> 0)
	return b
}

func u32_u8slice(n uint32) []uint8 {
	b := make([]uint8, 4)
	b[0] = uint8((n & 0xff000000) >> 24)
	b[1] = uint8((n & 0x00ff0000) >> 16)
	b[2] = uint8((n & 0x0000ff00) >> 8)
	b[3] = uint8((n & 0x000000ff) >> 0)
	return b
}

// Parse functions

func ParseClassFile(r io.Reader) (*ClassFile, error) {
	var err error
	cf := new(ClassFile)
	for _, v := range []interface{}{&cf.Magic, &cf.MinorVersion, &cf.MajorVersion, &cf.ConstantPoolCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	if cf.Magic != 0xcafebabe {
		return nil, fmt.Errorf("Illegal magic binary: %v", cf.Magic)
	}
	// Constant pool starts with index 1
	cf.ConstantPool = make([]ConstantPoolInfo, cf.ConstantPoolCount)
	for i := uint16(1); i < cf.ConstantPoolCount; i++ {
		cpInfo, err := parseConstantPool(r)
		if err != nil {
			return nil, err
		}
		cf.ConstantPool[i] = cpInfo

		// Some cp_info consumes double indices
		switch cpInfo.Tag() {
		case CONSTANT_Long, CONSTANT_Double:
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
	cf.Fields = make([]*FieldInfo, cf.FieldsCount)
	for i := uint16(0); i < cf.FieldsCount; i++ {
		if cf.Fields[i], err = parseFieldInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.MethodsCount); err != nil {
		return nil, err
	}
	cf.Methods = make([]*MethodInfo, cf.MethodsCount)
	for i := uint16(0); i < cf.MethodsCount; i++ {
		if cf.Methods[i], err = parseMethodInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	if err = binary.Read(r, binary.BigEndian, &cf.AttributesCount); err != nil {
		return nil, err
	}
	cf.Attributes = make([]AttributeInfo, cf.AttributesCount)
	for i := uint16(0); i < cf.AttributesCount; i++ {
		if cf.Attributes[i], err = parseAttributeInfo(cf.ConstantPool, r); err != nil {
			return nil, err
		}
	}
	return cf, nil
}

func parseConstantPool(r io.Reader) (ConstantPoolInfo, error) {
	var (
		tag uint8
		err error
	)
	if tag, err = u1(r); err != nil {
		return nil, err
	}

	var (
		cpInfo ConstantPoolInfo
		data   []interface{}
	)
	switch tag {
	case CONSTANT_Class:
		info := new(ConstantClassInfo)
		data = []interface{}{
			&info.NameIndex,
		}
		cpInfo = info
	case CONSTANT_Fieldref:
		info := new(ConstantFieldrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case CONSTANT_Methodref:
		info := new(ConstantMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case CONSTANT_InterfaceMethodref:
		info := new(ConstantInterfaceMethodrefInfo)
		data = []interface{}{
			&info.ClassIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case CONSTANT_String:
		info := new(ConstantStringInfo)
		data = []interface{}{
			&info.StringIndex,
		}
		cpInfo = info
	case CONSTANT_Integer:
		info := new(ConstantIntegerInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case CONSTANT_Float:
		info := new(ConstantFloatInfo)
		data = []interface{}{
			&info.Bytes,
		}
		cpInfo = info
	case CONSTANT_Long:
		info := new(ConstantLongInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case CONSTANT_Double:
		info := new(ConstantDoubleInfo)
		data = []interface{}{
			&info.HighBytes,
			&info.LowBytes,
		}
		cpInfo = info
	case CONSTANT_NameAndType:
		info := new(ConstantNameAndTypeInfo)
		data = []interface{}{
			&info.NameIndex,
			&info.DescriptorIndex,
		}
		cpInfo = info
	case CONSTANT_MethodHandle:
		info := new(ConstantMethodHandleInfo)
		data = []interface{}{
			&info.ReferenceKind,
			&info.ReferenceIndex,
		}
		cpInfo = info
	case CONSTANT_MethodType:
		info := new(ConstantMethodTypeInfo)
		data = []interface{}{
			&info.DescriptorIndex,
		}
		cpInfo = info
	case CONSTANT_InvokeDynamic:
		info := new(ConstantInvokeDynamicInfo)
		data = []interface{}{
			&info.BootstrapMethodAttrIndex,
			&info.NameAndTypeIndex,
		}
		cpInfo = info
	case CONSTANT_Utf8:
		info := new(ConstantUtf8Info)
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

func parseFieldInfo(constantPool []ConstantPoolInfo, r io.Reader) (*FieldInfo, error) {
	var err error
	fi := new(FieldInfo)
	for _, v := range []interface{}{&fi.AccessFlags, &fi.NameIndex, &fi.DescriptorIndex, &fi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	fi.Attributes = make([]AttributeInfo, fi.AttributesCount)
	for i := uint16(0); i < fi.AttributesCount; i++ {
		if fi.Attributes[i], err = parseAttributeInfo(constantPool, r); err != nil {
			return nil, err
		}
	}
	return fi, nil
}

func parseMethodInfo(constantPool []ConstantPoolInfo, r io.Reader) (*MethodInfo, error) {
	var err error
	mi := new(MethodInfo)
	for _, v := range []interface{}{&mi.AccessFlags, &mi.NameIndex, &mi.DescriptorIndex, &mi.AttributesCount} {
		if err = binary.Read(r, binary.BigEndian, v); err != nil {
			return nil, err
		}
	}
	mi.Attributes = make([]AttributeInfo, mi.AttributesCount)
	for i := uint16(0); i < mi.AttributesCount; i++ {
		if mi.Attributes[i], err = parseAttributeInfo(constantPool, r); err != nil {
			return nil, err
		}
	}
	return mi, nil
}

func parseAttributeInfo(constantPool []ConstantPoolInfo, r io.Reader) (AttributeInfo, error) {
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

	utf8Info := constantPool[attributeNameIndex].(*ConstantUtf8Info)
	var ai AttributeInfo
	switch utf8Info.JavaString() {
	case Attribute_ConstantValue:
		cv := &ConstantValueAttribute{
			AttributeNameIndex_: attributeNameIndex,
			AttributeLength_:    attributeLength,
		}
		ai = cv
		err = binary.Read(r, binary.BigEndian, &cv.ConstantvalueIndex)
	case Attribute_Deprecated:
		ai = &DeprecatedAttribute{
			AttributeNameIndex_: attributeNameIndex,
			AttributeLength_:    attributeLength,
		}
		err = nil
	case Attribute_Signature:
		sig := &SignatureAttribute{
			AttributeNameIndex_: attributeNameIndex,
			AttributeLength_:    attributeLength,
		}
		ai = sig
		err = binary.Read(r, binary.BigEndian, &sig.SignatureIndex)
	case Attribute_SourceFile:
		sf := &SourceFileAttribute{
			AttributeNameIndex_: attributeNameIndex,
			AttributeLength_:    attributeLength,
		}
		ai = sf
		err = binary.Read(r, binary.BigEndian, &sf.SourceFileIndex)
	default:
		gai := &GenericAttributeInfo{
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
