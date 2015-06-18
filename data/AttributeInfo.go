package data

const (
	AttributeInfo_ClassFile_SourceFile                            = "SourceFile"
	AttributeInfo_ClassFile_InnerClasses                          = "InnerClasses"
	AttributeInfo_ClassFile_EnclosingMethod                       = "EnclosingMethod"
	AttributeInfo_ClassFile_SourceDebugExtension                  = "SourceDebugExtension"
	AttributeInfo_ClassFile_BootstrapMethods                      = "BootstrapMethods"
	AttributeInfo_FieldInfo_ConstantValue                         = "ConstantValue"
	AttributeInfo_MethodInfo_Code                                 = "Code"
	AttributeInfo_MethodInfo_Exceptions                           = "Exceptions"
	AttributeInfo_MethodInfo_RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	AttributeInfo_MethodInfo_RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	AttributeInfo_MethodInfo_AnnotationDefault                    = "AnnotationDefault"
	AttributeInfo_MethodInfo_MethodParameters                     = "MethodParameters"
	AttributeInfo_ClassFile_Synthetic                             = "Synthetic"
	AttributeInfo_FieldInfo_Synthetic                             = "Synthetic"
	AttributeInfo_MethodInfo_Synthetic                            = "Synthetic"
	AttributeInfo_ClassFile_Deprecated                            = "Deprecated"
	AttributeInfo_FieldInfo_Deprecated                            = "Deprecated"
	AttributeInfo_MethodInfo_Deprecated                           = "Deprecated"
	AttributeInfo_ClassFile_Signature                             = "Signature"
	AttributeInfo_FieldInfo_Signature                             = "Signature"
	AttributeInfo_MethodInfo_Signature                            = "Signature"
	AttributeInfo_ClassFile_RuntimeVisibleAnnotations             = "RuntimeVisibleAnnotations"
	AttributeInfo_FieldInfo_RuntimeVisibleAnnotations             = "RuntimeVisibleAnnotations"
	AttributeInfo_MethodInfo_RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	AttributeInfo_ClassFile_RuntimeInvisibleAnnotations           = "RuntimeInvisibleAnnotations"
	AttributeInfo_FieldInfo_RuntimeInvisibleAnnotations           = "RuntimeInvisibleAnnotations"
	AttributeInfo_MethodInfo_RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	AttributeInfo_Code_LineNumberTable                            = "LineNumberTable"
	AttributeInfo_Code_LocalVariableTable                         = "LocalVariableTable"
	AttributeInfo_Code_LocalVariableTypeTable                     = "LocalVariableTypeTable"
	AttributeInfo_Code_StackMapTable                              = "StackMapTable"
	AttributeInfo_ClassFile_RuntimeVisibleTypeAnnotations         = "RuntimeVisibleTypeAnnotations"
	AttributeInfo_FieldInfo_RuntimeVisibleTypeAnnotations         = "RuntimeVisibleTypeAnnotations"
	AttributeInfo_MethodInfo_RuntimeVisibleTypeAnnotations        = "RuntimeVisibleTypeAnnotations"
	AttributeInfo_Code_RuntimeVisibleTypeAnnotations              = "RuntimeVisibleTypeAnnotations"
	AttributeInfo_ClassFile_RuntimeInvisibleTypeAnnotations       = "RuntimeInvisibleTypeAnnotations"
	AttributeInfo_FieldInfo_RuntimeInvisibleTypeAnnotations       = "RuntimeInvisibleTypeAnnotations"
	AttributeInfo_MethodInfo_RuntimeInvisibleTypeAnnotations      = "RuntimeInvisibleTypeAnnotations"
	AttributeInfo_Code_RuntimeInvisibleTypeAnnotations            = "RuntimeInvisibleTypeAnnotations"
)

type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []uint8
}
