package jvms

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
// TODO: ConstantValueAttribute
// TODO: DeprecatedAttribute
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
// TODO: SignatureAttribute
// TODO: SourceDebugExtensionAttribute
// TODO: SourceFileAttribute
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
