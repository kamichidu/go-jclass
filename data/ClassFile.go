package data

const (
	ClassFile_ACC_PUBLIC     = 0x0001
	ClassFile_ACC_FINAL      = 0x0010
	ClassFile_ACC_SUPER      = 0x0020
	ClassFile_ACC_INTERFACE  = 0x0200
	ClassFile_ACC_ABSTRACT   = 0x0400
	ClassFile_ACC_SYNTHETIC  = 0x1000
	ClassFile_ACC_ANNOTATION = 0x2000
	ClassFile_ACC_ENUM       = 0x4000
)

type ClassFile struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      []CpInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []uint16
	FieldsCount       uint16
	Fields            []FieldInfo
	MethodsCount      uint16
	Methods           []MethodInfo
	AttributesCount   uint16
	Attributes        []AttributeInfo
}
