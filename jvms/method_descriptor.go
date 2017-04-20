package jvms

import (
	"bufio"
	"fmt"
	"io"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.3.3

type MethodDescriptorInfo struct {
	ReturnTypeInfo    *FieldDescriptorInfo
	ParameterTypeInfo []*FieldDescriptorInfo
}

func ParseMethodDescriptor(r io.Reader) (*MethodDescriptorInfo, error) {
	ast := new(astMethodDescriptor)
	if err := methodDescriptor(ast, bufio.NewReader(r)); err != nil {
		return nil, err
	}
	info := new(MethodDescriptorInfo)
	info.ParameterTypeInfo = make([]*FieldDescriptorInfo, 0)
	if ast.ParameterDescriptor != nil {
		for _, ast := range ast.ParameterDescriptor {
			info.ParameterTypeInfo = append(info.ParameterTypeInfo, toFieldDescriptorInfo(ast.FieldType))
		}
	}
	if ast.ReturnDescriptor.FieldType != nil {
		info.ReturnTypeInfo = toFieldDescriptorInfo(ast.ReturnDescriptor.FieldType)
	} else {
		info.ReturnTypeInfo = &FieldDescriptorInfo{
			PrimitiveType: true,
			TypeName:      "void",
		}
	}
	return info, nil
}

type astMethodDescriptor struct {
	ParameterDescriptor []*astParameterDescriptor
	ReturnDescriptor    *astReturnDescriptor
}

// "(" ParameterDescriptor* ")" ReturnDescriptor
func methodDescriptor(out *astMethodDescriptor, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '(' {
		return fmt.Errorf("MethodDescriptor must starts with `(', but starts with `%c'", err)
	}

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
		child := new(astParameterDescriptor)
		if err = parameterDescriptor(child, r); err != nil {
			return err
		}
		out.ParameterDescriptor = append(out.ParameterDescriptor, child)
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ')' {
		return fmt.Errorf("MethodDescriptor must indicates with `)', but with `%c'", err)
	}

	out.ReturnDescriptor = new(astReturnDescriptor)
	return returnDescriptor(out.ReturnDescriptor, r)
}

type astParameterDescriptor struct {
	FieldType *astFieldType
}

// FieldType
func parameterDescriptor(out *astParameterDescriptor, r runeReader) error {
	out.FieldType = new(astFieldType)
	return fieldType(out.FieldType, r)
}

type astReturnDescriptor struct {
	FieldType      *astFieldType
	VoidDescriptor *astVoidDescriptor
}

// FieldType
// VoidDescriptor
func returnDescriptor(out *astReturnDescriptor, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if err = r.UnreadRune(); err != nil {
		return err
	}
	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'L', '[':
		out.FieldType = new(astFieldType)
		return fieldType(out.FieldType, r)
	case 'V':
		out.VoidDescriptor = new(astVoidDescriptor)
		return voidDescriptor(out.VoidDescriptor, r)
	default:
		return fmt.Errorf("ReturnDescriptor must be FieldType or VoidDescriptor, but unknown prefix `%c' found", c)
	}
}

type astVoidDescriptor struct {
	Token string
}

// "V"
func voidDescriptor(out *astVoidDescriptor, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'V' {
		return fmt.Errorf("VoidDescriptor must be `V', but is `%c'", c)
	}
	out.Token = "V"
	return nil
}
