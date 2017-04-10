package jvms

import (
	"io"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.3.3
// MethodDescriptor:
//   "(" ParameterDescriptor* ")" ReturnDescriptor
//
// ParameterDescriptor:
//   FieldType
//
// ReturnDescriptor:
//   FieldType
// 	   VoidDescriptor
//
// VoidDescriptor:
//   "V"

func ParseMethodDescriptor(r io.Reader) {
}
