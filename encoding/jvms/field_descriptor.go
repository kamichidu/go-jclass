package jvms

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.3.2

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
	ast := new(astFieldDescriptor)
	if err := fieldDescriptor(ast, bufio.NewReader(r)); err != nil {
		return nil, err
	}
	return toFieldDescriptorInfo(ast.FieldType), nil
}

func toFieldDescriptorInfo(ast *astFieldType) *FieldDescriptorInfo {
	info := new(FieldDescriptorInfo)
	switch {
	case ast.BaseType != nil:
		info.PrimitiveType = true
		info.TypeName = baseType2TypeName(ast.BaseType.Token)
	case ast.ObjectType != nil:
		info.TypeName = ast.ObjectType.ClassName.Token
	case ast.ArrayType != nil:
		info.ArrayType = true
		info.ArrayDepth = 1
		curr := ast.ArrayType.ComponentType
		for curr != nil {
			switch {
			case curr.FieldType.ArrayType != nil:
				info.ArrayDepth++
				curr = curr.FieldType.ArrayType.ComponentType
			case curr.FieldType.BaseType != nil:
				info.PrimitiveType = true
				info.TypeName = baseType2TypeName(curr.FieldType.BaseType.Token)
				curr = nil
			case curr.FieldType.ObjectType != nil:
				info.TypeName = curr.FieldType.ObjectType.ClassName.Token
				curr = nil
			default:
				curr = nil
			}
		}
	}
	return info
}

type astFieldDescriptor struct {
	FieldType *astFieldType
}

// FieldType
func fieldDescriptor(out *astFieldDescriptor, r runeReader) error {
	out.FieldType = new(astFieldType)
	return fieldType(out.FieldType, r)
}

type astFieldType struct {
	BaseType   *astBaseType
	ObjectType *astObjectType
	ArrayType  *astArrayType
}

// BaseType
// ObjectType
// ArrayType
func fieldType(out *astFieldType, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if err = r.UnreadRune(); err != nil {
		return err
	}

	switch c {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		out.BaseType = new(astBaseType)
		return baseType(out.BaseType, r)
	case 'L':
		out.ObjectType = new(astObjectType)
		return objectType(out.ObjectType, r)
	case '[':
		out.ArrayType = new(astArrayType)
		return arrayType(out.ArrayType, r)
	default:
		return fmt.Errorf("FieldType must be BaseType, ObjectType or ArrayType, but unknown prefix `%c' found", c)
	}
}

type astBaseType struct {
	Token string
}

// "B"
// "C"
// "D"
// "F"
// "I"
// "J"
// "S"
// "Z"
func baseType(out *astBaseType, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	switch c {
	case 'B':
		out.Token = "B"
		return nil
	case 'C':
		out.Token = "C"
		return nil
	case 'D':
		out.Token = "D"
		return nil
	case 'F':
		out.Token = "F"
		return nil
	case 'I':
		out.Token = "I"
		return nil
	case 'J':
		out.Token = "J"
		return nil
	case 'S':
		out.Token = "S"
		return nil
	case 'Z':
		out.Token = "Z"
		return nil
	default:
		return fmt.Errorf("BaseType must be `B', `C', `D', `F', `I', `J', `S' or `Z', but is `%c'", c)
	}
}

type astObjectType struct {
	ClassName *astClassName
}

type astClassName struct {
	Token string
}

// "L" ClassName ";"
func objectType(out *astObjectType, r runeReader) error {
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
	out.ClassName = new(astClassName)
	out.ClassName.Token = string(chars)

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ';' {
		return fmt.Errorf("ObjectType must ends with `;', but ends with `%c'", c)
	}
	return nil
}

type astArrayType struct {
	ComponentType *astComponentType
}

// "[" ComponentType
func arrayType(out *astArrayType, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '[' {
		return fmt.Errorf("ArrayType must start with `[', but starts with `%c'", c)
	}

	out.ComponentType = new(astComponentType)
	return componentType(out.ComponentType, r)
}

type astComponentType struct {
	FieldType *astFieldType
}

// FieldType
func componentType(out *astComponentType, r runeReader) error {
	out.FieldType = new(astFieldType)
	return fieldType(out.FieldType, r)
}
