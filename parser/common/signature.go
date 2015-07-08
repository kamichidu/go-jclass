package common

type JavaTypeSignature struct {
	ReferenceTypeSignature *ReferenceTypeSignature
	BaseType               *BaseType
}

type ReferenceTypeSignature struct {
	ClassTypeSignature    *ClassTypeSignature
	TypeVariableSignature *TypeVariableSignature
	ArrayTypeSignature    *ArrayTypeSignature
}

type ClassTypeSignature struct {
	PackageSpecifier         *PackageSpecifier
	SimpleClassTypeSignature *SimpleClassTypeSignature
	ClassTypeSignatureSuffix []*ClassTypeSignatureSuffix
}

type PackageSpecifier struct {
	Identifier []string
}

type SimpleClassTypeSignature struct {
	Identifier    string
	TypeArguments *TypeArguments
}

type TypeArguments struct {
	TypeArgument []*TypeArgument
}

type TypeArgument struct {
	WildcardIndicator      *WildcardIndicator
	ReferenceTypeSignature *ReferenceTypeSignature
	Text                   string
}

type WildcardIndicator struct {
	Text string
}

type ClassTypeSignatureSuffix struct {
	SimpleClassTypeSignature *SimpleClassTypeSignature
}

type TypeVariableSignature struct {
	Identifier string
}

type ArrayTypeSignature struct {
	JavaTypeSignature *JavaTypeSignature
}

type ClassSignature struct {
	TypeParameters          *TypeParameters
	SuperclassSignature     *SuperclassSignature
	SuperinterfaceSignature *SuperinterfaceSignature
}

type TypeParameters struct {
	TypeParameter []*TypeParameter
}

type TypeParameter struct {
	Identifier     string
	ClassBound     *ClassBound
	InterfaceBound []*InterfaceBound
}

type ClassBound struct {
	ReferenceTypeSignature *ReferenceTypeSignature
}

type InterfaceBound struct {
	ReferenceTypeSignature *ReferenceTypeSignature
}

type SuperclassSignature struct {
	ClassTypeSignature *ClassTypeSignature
}

type SuperinterfaceSignature struct {
	ClassTypeSignature *ClassTypeSignature
}

type MethodSignature struct {
	TypeParameters    *TypeParameters
	JavaTypeSignature []*JavaTypeSignature
	Result            *Result
	ThrowsSignature   []*ThrowsSignature
}

type Result struct {
	JavaTypeSignature *JavaTypeSignature
	VoidDescriptor    *VoidDescriptor
}

type ThrowsSignature struct {
	ClassTypeSignature    *ClassTypeSignature
	TypeVariableSignature *TypeVariableSignature
}

type FieldSignature struct {
	ReferenceTypeSignature *ReferenceTypeSignature
}
