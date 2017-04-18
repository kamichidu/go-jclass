package jvms

import (
	"fmt"
	"github.com/kr/pretty"
	"io"
)

type ClassSignatureInfo struct {
}

type FieldTypeSignatureInfo struct {
}

type MethodTypeSignatureInfo struct {
}

func ParseClassSignature(r io.Reader) error {
	// info := new(ClassSignatureInfo)
	ast := new(astClassSignature)
	if err := classSignature(ast, newReader(r)); err != nil {
		return err
	}
	pretty.Printf("%# v", ast)
	return nil
}

func ParseFieldTypeSignature(r io.Reader) error {
	// info := new(FieldTypeSignatureInfo)
	ast := new(astFieldTypeSignature)
	return fieldTypeSignature(ast, newReader(r))
}

func ParseMethodTypeSignature(r io.Reader) error {
	// info := new(MethodTypeSignatureInfo)
	ast := new(astMethodTypeSignature)
	return methodTypeSignature(ast, newReader(r))
}

type astClassSignature struct {
	FormalTypeParameters    *astFormalTypeParameters
	SuperclassSignature     *astSuperclassSignature
	SuperinterfaceSignature []*astSuperinterfaceSignature
}

// FormalTypeParameters? SuperclassSignature SuperinterfaceSignature*
func classSignature(out *astClassSignature, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}

	if c == '<' {
		out.FormalTypeParameters = new(astFormalTypeParameters)
		if err = formalTypeParameters(out.FormalTypeParameters, r); err != nil {
			return err
		}
	}

	out.SuperclassSignature = new(astSuperclassSignature)
	if err = superclassSignature(out.SuperclassSignature, r); err != nil {
		return err
	}

	for {
		c, err := peek(r)
		if err != nil {
			return err
		}
		if c != 'L' {
			break
		}
		child := new(astSuperinterfaceSignature)
		if err = superinterfaceSignature(child, r); err != nil {
			return err
		}
		out.SuperinterfaceSignature = append(out.SuperinterfaceSignature, child)
	}
	return nil
}

type astFormalTypeParameters struct {
	FormalTypeParameter []*astFormalTypeParameter
}

// '<' FormalTypeParameter+ '>'
func formalTypeParameters(out *astFormalTypeParameters, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '<' {
		return errorPrefixUnmatch("FormalTypeParameters", []rune{'<'}, c)
	}

	for {
		child := new(astFormalTypeParameter)
		if err = formalTypeParameter(child, r); err != nil {
			return err
		}
		out.FormalTypeParameter = append(out.FormalTypeParameter, child)

		if c, err = peek(r); err != nil {
			return err
		} else if c == '>' {
			break
		}
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != '>' {
		return errorSuffixUnmatch("FormalTypeParameters", []rune{'>'}, c)
	} else {
		return nil
	}
}

type astFormalTypeParameter struct {
	Identifier     *astIdentifier
	ClassBound     *astClassBound
	InterfaceBound []*astInterfaceBound
}

// Identifier ClassBound InterfaceBound*
func formalTypeParameter(out *astFormalTypeParameter, r runeReader) error {
	var err error

	out.Identifier = new(astIdentifier)
	if err = identifier(out.Identifier, r); err != nil {
		return err
	}

	out.ClassBound = new(astClassBound)
	if err = classBound(out.ClassBound, r); err != nil {
		return err
	}

	for {
		c, err := peek(r)
		if err != nil {
			return err
		}
		if c != ':' {
			break
		}

		child := new(astInterfaceBound)
		if err = interfaceBound(child, r); err != nil {
			return err
		}
		out.InterfaceBound = append(out.InterfaceBound, child)
	}
	return nil
}

type astIdentifier struct {
	Text string
}

func identifier(out *astIdentifier, r runeReader) error {
	ident := make([]rune, 0)
loop:
	for {
		c, err := peek(r)
		if err != nil {
			return err
		}
		switch {
		case c >= 'a' && c <= 'z':
			fallthrough
		case c >= 'A' && c <= 'Z':
			fallthrough
		case c >= '0' && c <= '9':
			fallthrough
		case c == '_':
			c, _, err = r.ReadRune()
			if err != nil {
				return err
			}
			ident = append(ident, c)
		default:
			break loop
		}
	}
	out.Text = string(ident)
	return nil
}

type astClassBound struct {
	FieldTypeSignature *astFieldTypeSignature
}

// ':' FieldTypeSignature?
func classBound(out *astClassBound, r runeReader) error {
	if c, _, err := r.ReadRune(); err != nil {
		return err
	} else if c != ':' {
		return errorPrefixUnmatch("ClassBound", []rune{':'}, c)
	}

	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T':
		out.FieldTypeSignature = new(astFieldTypeSignature)
		return fieldTypeSignature(out.FieldTypeSignature, r)
	default:
		return nil
	}
}

type astInterfaceBound struct {
	FieldTypeSignature *astFieldTypeSignature
}

// ':' FieldTypeSignature
func interfaceBound(out *astInterfaceBound, r runeReader) error {
	if c, _, err := r.ReadRune(); err != nil {
		return err
	} else if c != ':' {
		return errorPrefixUnmatch("InterfaceBound", []rune{':'}, c)
	}

	out.FieldTypeSignature = new(astFieldTypeSignature)
	return fieldTypeSignature(out.FieldTypeSignature, r)
}

type astSuperinterfaceSignature struct {
	ClassTypeSignature *astClassTypeSignature
}

// ClassTypeSignature
func superinterfaceSignature(out *astSuperinterfaceSignature, r runeReader) error {
	out.ClassTypeSignature = new(astClassTypeSignature)
	return classTypeSignature(out.ClassTypeSignature, r)
}

type astSuperclassSignature struct {
	ClassTypeSignature *astClassTypeSignature
}

// ClassTypeSignature
func superclassSignature(out *astSuperclassSignature, r runeReader) error {
	out.ClassTypeSignature = new(astClassTypeSignature)
	return classTypeSignature(out.ClassTypeSignature, r)
}

type astFieldTypeSignature struct {
	ClassTypeSignature    *astClassTypeSignature
	ArrayTypeSignature    *astArrayTypeSignature
	TypeVariableSignature *astTypeVariableSignature
}

// ClassTypeSignature
// ArrayTypeSignature
// TypeVariableSignature
func fieldTypeSignature(out *astFieldTypeSignature, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L':
		out.ClassTypeSignature = new(astClassTypeSignature)
		return classTypeSignature(out.ClassTypeSignature, r)
	case '[':
		out.ArrayTypeSignature = new(astArrayTypeSignature)
		return arrayTypeSignature(out.ArrayTypeSignature, r)
	case 'T':
		out.TypeVariableSignature = new(astTypeVariableSignature)
		return typeVariableSignature(out.TypeVariableSignature, r)
	default:
		return errorPrefixUnmatch("FieldTypeSignature", []rune{'L', '[', 'T'}, c)
	}
}

type astClassTypeSignature struct {
	PackageSpecifier         *astPackageSpecifier
	SimpleClassTypeSignature *astSimpleClassTypeSignature
	ClassTypeSignatureSuffix []*astClassTypeSignatureSuffix
}

type astPackageSpecifier struct {
	Token string
}

// 'L' PackageSpecifier? SimpleClassTypeSignature ClassTypeSignatureSuffix* ';'
func classTypeSignature(out *astClassTypeSignature, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'L' {
		return errorPrefixUnmatch("ClassTypeSignature", []rune{'L'}, c)
	}

	// extract chars until meet '/'
	buffer := make([]rune, 0)
	packageSpecifier := make([]rune, 0)
	unreadRunes := 0
	for c != ';' {
		c, _, err = r.ReadRune()
		if err != nil {
			break
		}
		unreadRunes++

		buffer = append(buffer, c)
		if c == '/' {
			packageSpecifier = append(packageSpecifier, buffer...)
			buffer = make([]rune, 0)
			unreadRunes = 0
		}
	}
	for unreadRunes > 0 {
		if err = r.UnreadRune(); err != nil {
			return err
		}
		unreadRunes--
	}
	if len(packageSpecifier) > 0 {
		out.PackageSpecifier = new(astPackageSpecifier)
		out.PackageSpecifier.Token = string(packageSpecifier)
	}

	out.SimpleClassTypeSignature = new(astSimpleClassTypeSignature)
	if err = simpleClassTypeSignature(out.SimpleClassTypeSignature, r); err != nil {
		return err
	}

	for {
		if c, err = peek(r); err != nil {
			return err
		} else if c == '.' {
			child := new(astClassTypeSignatureSuffix)
			if err = classTypeSignatureSuffix(child, r); err != nil {
				return err
			}
			out.ClassTypeSignatureSuffix = append(out.ClassTypeSignatureSuffix, child)
		} else {
			break
		}
	}

	if c, _, err = r.ReadRune(); err != nil {
		return err
	} else if c != ';' {
		return errorSuffixUnmatch("ClassTypeSignature", []rune{';'}, c)
	} else {
		return nil
	}
}

// Identifier '/' PackageSpecifier*
// func packageSpecifier(out Emitter, r *bufio.Reader) error {
// 	if err := identifier(out, r); err != nil {
// 		return err
// 	}
// 	c, _, err := r.ReadRune()
// 	if err != nil {
// 		return err
// 	} else if c != '/' {
// 		return fmt.Errorf("%s expected `%c', but `%c'", "PackageSpecifier", '/', c)
// 	}
// 	out.Emit("PackageSpecifier", "/")
//
// 	if found, err := lookahead(r, '/'); err != nil {
// 		return err
// 	} else if found {
// 		return packageSpecifier(out, r)
// 	} else {
// 		return nil
// 	}
// }

type astSimpleClassTypeSignature struct {
	Identifier    *astIdentifier
	TypeArguments *astTypeArguments
}

// Identifier TypeArguments?
func simpleClassTypeSignature(out *astSimpleClassTypeSignature, r runeReader) error {
	out.Identifier = new(astIdentifier)
	if err := identifier(out.Identifier, r); err != nil {
		return err
	}

	c, err := peek(r)
	if err != nil {
		return err
	}
	if c == '<' {
		out.TypeArguments = new(astTypeArguments)
		return typeArguments(out.TypeArguments, r)
	} else {
		return nil
	}
}

type astClassTypeSignatureSuffix struct {
	SimpleClassTypeSignature *astSimpleClassTypeSignature
}

// '.' SimpleClassTypeSignature
func classTypeSignatureSuffix(out *astClassTypeSignatureSuffix, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '.' {
		return errorPrefixUnmatch("ClassTypeSignatureSuffix", []rune{'.'}, c)
	}

	out.SimpleClassTypeSignature = new(astSimpleClassTypeSignature)
	return simpleClassTypeSignature(out.SimpleClassTypeSignature, r)
}

type astTypeVariableSignature struct {
	Identifier *astIdentifier
}

// 'T' Identifier ';'
func typeVariableSignature(out *astTypeVariableSignature, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'T' {
		return errorPrefixUnmatch("TypeVariableSignature", []rune{'T'}, c)
	}

	out.Identifier = new(astIdentifier)
	if err = identifier(out.Identifier, r); err != nil {
		return err
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ';' {
		return errorSuffixUnmatch("TypeVariableSignature", []rune{';'}, c)
	} else {
		return nil
	}
}

type astTypeArguments struct {
	TypeArgument []*astTypeArgument
}

// '<' TypeArgument+ '>'
func typeArguments(out *astTypeArguments, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '<' {
		return errorPrefixUnmatch("TypeArguments", []rune{'<'}, c)
	}

loop:
	for {
		child := new(astTypeArgument)
		if err = typeArgument(child, r); err != nil {
			return err
		}
		out.TypeArgument = append(out.TypeArgument, child)

		c, err = peek(r)
		if err != nil {
			return err
		}
		switch c {
		case '+', '-', 'L', '[', 'T', '*':
			// go next
		default:
			break loop
		}
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != '>' {
		return errorSuffixUnmatch("TypeArguments", []rune{'>'}, c)
	} else {
		return nil
	}
}

type astTypeArgument struct {
	WildcardIndicator  *astWildcardIndicator
	FieldTypeSignature *astFieldTypeSignature
	Token              string
}

// WildcardIndicator? FieldTypeSignature
// '*'
func typeArgument(out *astTypeArgument, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case '+', '-':
		out.WildcardIndicator = new(astWildcardIndicator)
		if err = wildcardIndicator(out.WildcardIndicator, r); err != nil {
			return err
		}
		fallthrough
	case 'L', '[', 'T':
		out.FieldTypeSignature = new(astFieldTypeSignature)
		return fieldTypeSignature(out.FieldTypeSignature, r)
	case '*':
		r.ReadRune()
		out.Token = "*"
		return nil
	default:
		return errorPrefixUnmatch("TypeArgument", []rune{'+', '-', 'L', '[', 'T', '*'}, c)
	}
}

type astWildcardIndicator struct {
	Token string
}

// '+'
// '-'
func wildcardIndicator(out *astWildcardIndicator, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	switch c {
	case '+':
		out.Token = "+"
		return nil
	case '-':
		out.Token = "-"
		return nil
	default:
		return errorPrefixUnmatch("WildcardIndicator", []rune{'+', '-'}, c)
	}
}

type astArrayTypeSignature struct {
	TypeSignature *astTypeSignature
}

// '[' TypeSignature
func arrayTypeSignature(out *astArrayTypeSignature, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if c != '[' {
		return errorPrefixUnmatch("ArrayTypeSignature", []rune{'['}, c)
	}

	out.TypeSignature = new(astTypeSignature)
	return typeSignature(out.TypeSignature, r)
}

type astTypeSignature struct {
	FieldTypeSignature *astFieldTypeSignature
	BaseType           *astBaseType
}

// FieldTypeSignature
// BaseType
func typeSignature(out *astTypeSignature, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T':
		out.FieldTypeSignature = new(astFieldTypeSignature)
		return fieldTypeSignature(out.FieldTypeSignature, r)
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		out.BaseType = new(astBaseType)
		return baseType(out.BaseType, r)
	default:
		return errorPrefixUnmatch("TypeSignature", []rune{'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z'}, c)
	}
}

type astMethodTypeSignature struct {
	FormalTypeParameters *astFormalTypeParameters
	TypeSignature        []*astTypeSignature
	ReturnType           *astReturnType
	ThrowsSignature      []*astThrowsSignature
}

// FormalTypeParameters? '(' TypeSignature* ')' ReturnType ThrowsSignature*
func methodTypeSignature(out *astMethodTypeSignature, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	if c == '<' {
		out.FormalTypeParameters = new(astFormalTypeParameters)
		if err = formalTypeParameters(out.FormalTypeParameters, r); err != nil {
			return err
		}
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != '(' {
		return fmt.Errorf("MethodTypeSignature expects `(', but is `%c'", c)
	}

loop:
	for {
		c, err = peek(r)
		if err != nil {
			return err
		}
		switch c {
		case 'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
			child := new(astTypeSignature)
			if err = typeSignature(child, r); err != nil {
				return err
			}
			out.TypeSignature = append(out.TypeSignature, child)
		default:
			break loop
		}
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ')' {
		return fmt.Errorf("MethodTypeSignature expects `)', but is `%c'", c)
	}

	out.ReturnType = new(astReturnType)
	if err = returnType(out.ReturnType, r); err != nil {
		return err
	}

	for {
		c, err = peek(r)
		if err != nil {
			return err
		}
		switch c {
		case '^':
			child := new(astThrowsSignature)
			if err = throwsSignature(child, r); err != nil {
				return err
			}
			out.ThrowsSignature = append(out.ThrowsSignature, child)
		default:
			return nil
		}
	}
}

type astReturnType struct {
	TypeSignature  *astTypeSignature
	VoidDescriptor *astVoidDescriptor
}

// TypeSignature
// VoidDescriptor
func returnType(out *astReturnType, r runeReader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		out.TypeSignature = new(astTypeSignature)
		return typeSignature(out.TypeSignature, r)
	case 'V':
		out.VoidDescriptor = new(astVoidDescriptor)
		return voidDescriptor(out.VoidDescriptor, r)
	default:
		return errorPrefixUnmatch("ReturnType", []rune{'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'V'}, c)
	}
}

type astThrowsSignature struct {
	ClassTypeSignature    *astClassTypeSignature
	TypeVariableSignature *astTypeVariableSignature
}

// '^' ClassTypeSignature
// '^' TypeVariableSignature
func throwsSignature(out *astThrowsSignature, r runeReader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '^' {
		return errorPrefixUnmatch("ThrowsSignature", []rune{'^'}, c)
	}

	c, err = peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L':
		out.ClassTypeSignature = new(astClassTypeSignature)
		return classTypeSignature(out.ClassTypeSignature, r)
	case 'T':
		out.TypeVariableSignature = new(astTypeVariableSignature)
		return typeVariableSignature(out.TypeVariableSignature, r)
	default:
		return fmt.Errorf("", c)
	}
}
