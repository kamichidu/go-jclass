package jvms

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.3.2
// FieldDescriptor:
//   FieldType
//
// FieldType:
//   BaseType
//   ObjectType
//   ArrayType
//
// BaseType:
//   "B"
//   "C"
//   "D"
//   "F"
//   "I"
//   "J"
//   "S"
//   "Z"
//
// ObjectType:
//   "L" ClassName ";"
//
// ArrayType:
//   "[" ComponentType
//
// ComponentType:
//   FieldType

type FieldDescriptorInfo struct {
	TypeName      string
	PrimitiveType bool
	ArrayType     bool
	ArrayDepth    int
}

func (self *FieldDescriptorInfo) String() string {
	return self.TypeName + strings.Repeat("[]", self.ArrayDepth)
}

func ParseFieldDescriptor(r io.Reader) (*FieldDescriptorInfo, error) {
	info := new(FieldDescriptorInfo)
	err := fieldDescriptor(info, bufio.NewReader(r))
	return info, err
}

func fieldDescriptor(out *FieldDescriptorInfo, r *bufio.Reader) error {
	return fieldType(out, r)
}

func fieldType(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if err = r.UnreadRune(); err != nil {
		return err
	}

	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		return baseType(out, r)
	case 'L':
		return objectType(out, r)
	case '[':
		return arrayType(out, r)
	default:
		return fmt.Errorf("FieldType must be BaseType, ObjectType or ArrayType, but unknown prefix `%c' found", c)
	}
}

func baseType(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	switch c {
	case 'B':
		out.TypeName = "byte"
		out.PrimitiveType = true
		return nil
	case 'C':
		out.TypeName = "char"
		out.PrimitiveType = true
		return nil
	case 'D':
		out.TypeName = "double"
		out.PrimitiveType = true
		return nil
	case 'F':
		out.TypeName = "float"
		out.PrimitiveType = true
		return nil
	case 'I':
		out.TypeName = "int"
		out.PrimitiveType = true
		return nil
	case 'J':
		out.TypeName = "long"
		out.PrimitiveType = true
		return nil
	case 'S':
		out.TypeName = "short"
		out.PrimitiveType = true
		return nil
	case 'Z':
		out.TypeName = "boolean"
		out.PrimitiveType = true
		return nil
	default:
		return fmt.Errorf("BaseType must be `B', `C', `D', `F', `I', `J', `S' or `Z', but is `%c'", c)
	}
}

func objectType(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'L' {
		return fmt.Errorf("ObjectType must starts with `L', but starts with `%c'", c)
	}

	chars := make([]rune, 0)
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			return err
		}
		if c == ';' {
			if err = r.UnreadRune(); err != nil {
				return err
			}
			break
		}
		if c == '/' {
			c = '.'
		}
		chars = append(chars, c)
	}
	out.TypeName = string(chars)

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ';' {
		return fmt.Errorf("ObjectType must ends with `;', but ends with `%c'", c)
	}
	return nil
}

func arrayType(out *FieldDescriptorInfo, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '[' {
		return fmt.Errorf("ArrayType must start with `[', but starts with `%c'", c)
	}
	out.ArrayType = true
	out.ArrayDepth++
	return componentType(out, r)
}

func componentType(out *FieldDescriptorInfo, r *bufio.Reader) error {
	return fieldType(out, r)
}
