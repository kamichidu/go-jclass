package jvms

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1-200-E.1
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5-200-A.1
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6-200-A.1

const (
	ACC_PUBLIC       uint16 = 0x0001
	ACC_PRIVATE             = 0x0002
	ACC_PROTECTED           = 0x0004
	ACC_STATIC              = 0x0008
	ACC_FINAL               = 0x0010
	ACC_SUPER               = 0x0020
	ACC_SYNCHRONIZED        = 0x0020
	ACC_VOLATILE            = 0x0040
	ACC_BRIDGE              = 0x0040
	ACC_TRANSIENT           = 0x0080
	ACC_VARARGS             = 0x0080
	ACC_NATIVE              = 0x0100
	ACC_INTERFACE           = 0x0200
	ACC_ABSTRACT            = 0x0400
	ACC_STRICT              = 0x0800
	ACC_SYNTHETIC           = 0x1000
	ACC_ANNOTATION          = 0x2000
	ACC_ENUM                = 0x4000
)

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4-140
const (
	CONSTANT_Class              uint8 = 7
	CONSTANT_Fieldref                 = 9
	CONSTANT_Methodref                = 10
	CONSTANT_InterfaceMethodref       = 11
	CONSTANT_String                   = 8
	CONSTANT_Integer                  = 3
	CONSTANT_Float                    = 4
	CONSTANT_Long                     = 5
	CONSTANT_Double                   = 6
	CONSTANT_NameAndType              = 12
	CONSTANT_Utf8                     = 1
	CONSTANT_MethodHandle             = 15
	CONSTANT_MethodType               = 16
	CONSTANT_InvokeDynamic            = 18
)

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7
const (
	Attribute_AnnotationDefault                    = "AnnotationDefault"
	Attribute_BootstrapMethods                     = "BootstrapMethods"
	Attribute_Code                                 = "Code"
	Attribute_ConstantValue                        = "ConstantValue"
	Attribute_Deprecated                           = "Deprecated"
	Attribute_EnclosingMethod                      = "EnclosingMethod"
	Attribute_Exceptions                           = "Exceptions"
	Attribute_InnerClasses                         = "InnerClasses"
	Attribute_LineNumberTable                      = "LineNumberTable"
	Attribute_LocalVariableTable                   = "LocalVariableTable"
	Attribute_LocalVariableTypeTable               = "LocalVariableTypeTable"
	Attribute_MethodParameters                     = "MethodParameters"
	Attribute_RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	Attribute_RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	Attribute_RuntimeInvisibleTypeAnnotations      = "RuntimeInvisibleTypeAnnotations"
	Attribute_RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	Attribute_RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	Attribute_RuntimeVisibleTypeAnnotations        = "RuntimeVisibleTypeAnnotations"
	Attribute_Signature                            = "Signature"
	Attribute_SourceDebugExtension                 = "SourceDebugExtension"
	Attribute_SourceFile                           = "SourceFile"
	Attribute_StackMapTable                        = "StackMapTable"
	Attribute_Synthetic                            = "Synthetic"
)
