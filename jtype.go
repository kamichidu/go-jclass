package jclass

import (
	"strings"
)

type JType interface {
	GetTypeName() string
}

type JPrimitiveType struct {
	typeName string
}

func NewJPrimitiveType(typeName string) *JPrimitiveType {
	return &JPrimitiveType{
		typeName: typeName,
	}
}

func (self *JPrimitiveType) GetTypeName() string {
	return self.typeName
}

type JReferenceType struct {
	typeName string
}

func NewJReferenceType(typeName string) *JReferenceType {
	return &JReferenceType{
		typeName: typeName,
	}
}

func (self *JReferenceType) GetTypeName() string {
	return self.typeName
}

type JArrayType struct {
	jtype JType
	dims  int
}

func NewJArrayType(jtype JType, dims int) *JArrayType {
	return &JArrayType{
		jtype: jtype,
		dims:  dims,
	}
}

func (self *JArrayType) GetTypeName() string {
	return self.jtype.GetTypeName() + strings.Repeat("[]", self.dims)
}

func (self *JArrayType) GetComponentType() JType {
	return self.jtype
}

func (self *JArrayType) GetDims() int {
	return self.dims
}
