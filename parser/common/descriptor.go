package common

type FieldDescriptor struct {
	FieldType *FieldType
}

type FieldType struct {
	BaseType   *BaseType
	ObjectType *ObjectType
	ArrayType  *ArrayType
}

type BaseType struct {
	Text string
}

type ObjectType struct {
	ClassName *ClassName
}

type ClassName struct {
    Identifier []string
}

type ArrayType struct {
	ComponentType *ComponentType
}

type ComponentType struct {
	FieldType *FieldType
}

type MethodDescriptor struct {
	ParameterDescriptor []*ParameterDescriptor
	ReturnDescriptor     *ReturnDescriptor
}

type ParameterDescriptor struct {
	FieldType *FieldType
}

type ReturnDescriptor struct {
	FieldType      *FieldType
	VoidDescriptor *VoidDescriptor
}

type VoidDescriptor struct {
	Text string
}
