package jvms

import (
	"bufio"
	"fmt"
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

type MethodDescriptorInfo struct {
	ReturnTypeInfo    *FieldDescriptorInfo
	ParameterTypeInfo []*FieldDescriptorInfo
}

func ParseMethodDescriptor(r io.Reader) (*MethodDescriptorInfo, error) {
	info := new(MethodDescriptorInfo)
	err := methodDescriptor(info, bufio.NewReader(r))
	return info, err
}

func methodDescriptor(out *MethodDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '(' {
		return fmt.Errorf("MethodDescriptor must starts with `(', but starts with `%c'", err)
	}

	out.ParameterTypeInfo = make([]*FieldDescriptorInfo, 0)
	for {
		c, _, err = r.ReadRune()
		if err != nil {
			return err
		}
		if err = r.UnreadRune(); err != nil {
			return err
		}
		if c == ')' {
			break
		}
		paramInfo := new(FieldDescriptorInfo)
		if err = parameterDescriptor(paramInfo, r); err != nil {
			return err
		}
		out.ParameterTypeInfo = append(out.ParameterTypeInfo, paramInfo)
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ')' {
		return fmt.Errorf("MethodDescriptor must indicates with `)', but with `%c'", err)
	}

	out.ReturnTypeInfo = new(FieldDescriptorInfo)
	return returnDescriptor(out.ReturnTypeInfo, r)
}

func parameterDescriptor(out *FieldDescriptorInfo, r *bufio.Reader) error {
	return fieldType(out, r)
}

func returnDescriptor(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if err = r.UnreadRune(); err != nil {
		return err
	}
	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', '[':
		return fieldType(out, r)
	case 'V':
		return voidDescriptor(out, r)
	default:
		return fmt.Errorf("ReturnDescriptor must be FieldType or VoidDescriptor, but unknown prefix `%c' found", c)
	}
}

func voidDescriptor(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'V' {
		return fmt.Errorf("VoidDescriptor must be `V', but is `%c'", c)
	}
	out.TypeName = "void"
	out.PrimitiveType = true
	return nil
}
