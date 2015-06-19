package data

const (
	FieldInfo_ACC_PUBLIC    = 0x0001 // Declared public; may be accessed from outside its package.
	FieldInfo_ACC_PRIVATE   = 0x0002 // Declared private; usable only within the defining class.
	FieldInfo_ACC_PROTECTED = 0x0004 // Declared protected; may be accessed within subclasses.
	FieldInfo_ACC_STATIC    = 0x0008 // Declared static.
	FieldInfo_ACC_FINAL     = 0x0010 // Declared final; never directly assigned to after object construction (JLS §17.5).
	FieldInfo_ACC_VOLATILE  = 0x0040 // Declared volatile; cannot be cached.
	FieldInfo_ACC_TRANSIENT = 0x0080 // Declared transient; not written or read by a persistent object manager.
	FieldInfo_ACC_SYNTHETIC = 0x1000 // Declared synthetic; not present in the source code.
	FieldInfo_ACC_ENUM      = 0x4000 // Declared as an element of an enum.
)

type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}
