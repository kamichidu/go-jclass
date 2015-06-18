package data

const (
	MethodInfo_ACC_PUBLIC       = 0x0001 // Declared public; may be accessed from outside its package.
	MethodInfo_ACC_PRIVATE      = 0x0002 // Declared private; accessible only within the defining class.
	MethodInfo_ACC_PROTECTED    = 0x0004 // Declared protected; may be accessed within subclasses.
	MethodInfo_ACC_STATIC       = 0x0008 // Declared static.
	MethodInfo_ACC_FINAL        = 0x0010 // Declared final; must not be overridden (§5.4.5).
	MethodInfo_ACC_SYNCHRONIZED = 0x0020 // Declared synchronized; invocation is wrapped by a monitor use.
	MethodInfo_ACC_BRIDGE       = 0x0040 // A bridge method, generated by the compiler.
	MethodInfo_ACC_VARARGS      = 0x0080 // Declared with variable number of arguments.
	MethodInfo_ACC_NATIVE       = 0x0100 // Declared native; implemented in a language other than Java.
	MethodInfo_ACC_ABSTRACT     = 0x0400 // Declared abstract; no implementation is provided.
	MethodInfo_ACC_STRICT       = 0x0800 // Declared strictfp; floating-point mode is FP-strict.
	MethodInfo_ACC_SYNTHETIC    = 0x1000 // Declared synthetic; not present in the source code.
)

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}
