package data

const (
	CpInfo_Tag_Class              = 7
	CpInfo_Tag_Fieldref           = 9
	CpInfo_Tag_Methodref          = 10
	CpInfo_Tag_InterfaceMethodref = 11
	CpInfo_Tag_String             = 8
	CpInfo_Tag_Integer            = 3
	CpInfo_Tag_Float              = 4
	CpInfo_Tag_Long               = 5
	CpInfo_Tag_Double             = 6
	CpInfo_Tag_NameAndType        = 12
	CpInfo_Tag_Utf8               = 1
	CpInfo_Tag_MethodHandle       = 15
	CpInfo_Tag_MethodType         = 16
	CpInfo_Tag_InvokeDynamic      = 18
)

type CpInfo interface {
}

type ClassInfo struct {
	Tag  uint8
	NameIndex uint16
}

type FieldrefInfo struct {
	Tag  uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type MethodrefInfo struct {
	Tag  uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type InterfaceMethodrefInfo struct {
	Tag  uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type StringInfo struct {
	Tag  uint8
	StringIndex uint16
}

type IntegerInfo struct {
	Tag  uint8
	Bytes uint32
}

type FloatInfo struct {
	Tag  uint8
	Bytes uint32
}

type LongInfo struct {
	Tag  uint8
	HighBytes uint32
	LowBytes  uint32
}

type DoubleInfo struct {
	Tag  uint8
	HighBytes uint32
	LowBytes  uint32
}

type NameAndTypeInfo struct {
	Tag  uint8
	NameIndex       uint16
	DescriptorIndex uint16
}

type Utf8Info struct {
	Tag  uint8
	Length uint16
	Bytes  []uint8
}

type MethodHandleInfo struct {
	Tag  uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type MethodTypeInfo struct {
	Tag  uint8
	DescriptorIndex uint16
}

type InvokeDynamicInfo struct {
	Tag  uint8
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
