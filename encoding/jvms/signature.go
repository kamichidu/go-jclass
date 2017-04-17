package jvms

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ClassSignatureInfo struct {
}

type FieldTypeSignatureInfo struct {
}

type MethodTypeSignatureInfo struct {
}

func (self *ClassSignatureInfo) Emit(grammar string, token string) {
	fmt.Printf("ClassSignatureInfo -- %s -- %#v\n", grammar, token)
}

func (self *FieldTypeSignatureInfo) Emit(grammar string, token string) {
	fmt.Printf("FieldTypeSignatureInfo -- %s -- %s\n", grammar, token)
}

func (self *MethodTypeSignatureInfo) Emit(grammar string, token string) {
	fmt.Printf("MethodTypeSignatureInfo -- %s -- %s\n", grammar, token)
}

type Emitter interface {
	Emit(grammar string, token string)
}

func ParseClassSignature(r io.Reader) error {
	info := new(ClassSignatureInfo)
	return classSignature(info, bufio.NewReader(r))
}

func ParseFieldTypeSignature(r io.Reader) error {
	info := new(FieldTypeSignatureInfo)
	return fieldTypeSignature(info, bufio.NewReader(r))
}

func ParseMethodTypeSignature(r io.Reader) error {
	info := new(MethodTypeSignatureInfo)
	return methodTypeSignature(info, bufio.NewReader(r))
}

// FormalTypeParameters? SuperclassSignature SuperinterfaceSignature*
func classSignature(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}

	if c == '<' {
		if err = formalTypeParameters(out, r); err != nil {
			return err
		}
	}
	if err = superclassSignature(out, r); err != nil {
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
		if err = superinterfaceSignature(out, r); err != nil {
			return err
		}
	}
	return nil
}

// '<' FormalTypeParameter+ '>'
func formalTypeParameters(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '<' {
		return errorPrefixUnmatch("FormalTypeParameters", []rune{'<'}, c)
	}
	out.Emit("FormalTypeParameters", "<")

	for {
		if err = formalTypeParameter(out, r); err != nil {
			return err
		}

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
		out.Emit("FormalTypeParameters", ">")
		return nil
	}
}

// Identifier ClassBound InterfaceBound*
func formalTypeParameter(out Emitter, r *bufio.Reader) error {
	var err error
	if err = identifier(out, r); err != nil {
		return err
	}
	if err = classBound(out, r); err != nil {
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

		if err = interfaceBound(out, r); err != nil {
			return err
		}
	}
	return nil
}

func identifier(out Emitter, r *bufio.Reader) error {
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
	out.Emit("Identifier", string(ident))
	return nil
}

// ':' FieldTypeSignature?
func classBound(out Emitter, r *bufio.Reader) error {
	if c, _, err := r.ReadRune(); err != nil {
		return err
	} else if c != ':' {
		return errorPrefixUnmatch("ClassBound", []rune{':'}, c)
	}
	out.Emit("ClassBound", ":")

	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T':
		return fieldTypeSignature(out, r)
	default:
		return nil
	}
}

// ':' FieldTypeSignature
func interfaceBound(out Emitter, r *bufio.Reader) error {
	if c, _, err := r.ReadRune(); err != nil {
		return err
	} else if c != ':' {
		return errorPrefixUnmatch("InterfaceBound", []rune{':'}, c)
	}
	out.Emit("InterfaceBound", ":")

	return fieldTypeSignature(out, r)
}

// ClassTypeSignature
func superinterfaceSignature(out Emitter, r *bufio.Reader) error {
	return classTypeSignature(out, r)
}

// ClassTypeSignature
func superclassSignature(out Emitter, r *bufio.Reader) error {
	return classTypeSignature(out, r)
}

// ClassTypeSignature
// ArrayTypeSignature
// TypeVariableSignature
func fieldTypeSignature(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L':
		return classTypeSignature(out, r)
	case '[':
		return arrayTypeSignature(out, r)
	case 'T':
		return typeVariableSignature(out, r)
	default:
		return errorPrefixUnmatch("FieldTypeSignature", []rune{'L', '[', 'T'}, c)
	}
}

// 'L' PackageSpecifier? SimpleClassTypeSignature ClassTypeSignatureSuffix* ';'
func classTypeSignature(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'L' {
		return errorPrefixUnmatch("ClassTypeSignature", []rune{'L'}, c)
	}
	out.Emit("ClassTypeSignature", "L")

	if found, err := lookahead(r, '/'); err != nil {
		return err
	} else if found {
		if err = packageSpecifier(out, r); err != nil {
			return err
		}
	}

	if err = simpleClassTypeSignature(out, r); err != nil {
		return err
	}

loop:
	for {
		c, err = peek(r)
		if err != nil {
		}
		switch c {
		case '.':
			if err = classTypeSignatureSuffix(out, r); err != nil {
				return err
			}
		default:
			break loop
		}
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return err
	} else if c != ';' {
		return errorSuffixUnmatch("ClassTypeSignature", []rune{';'}, c)
	} else {
		out.Emit("ClassTypeSignature", ";")
		return nil
	}
}

// Identifier '/' PackageSpecifier*
func packageSpecifier(out Emitter, r *bufio.Reader) error {
	if err := identifier(out, r); err != nil {
		return err
	}
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '/' {
		return fmt.Errorf("%s expected `%c', but `%c'", "PackageSpecifier", '/', c)
	}
	out.Emit("PackageSpecifier", "/")

	if found, err := lookahead(r, '/'); err != nil {
		return err
	} else if found {
		return packageSpecifier(out, r)
	} else {
		return nil
	}
}

// Identifier TypeArguments?
func simpleClassTypeSignature(out Emitter, r *bufio.Reader) error {
	if err := identifier(out, r); err != nil {
		return err
	}
	c, err := peek(r)
	if err != nil {
		return err
	}
	if c == '<' {
		return typeArguments(out, r)
	} else {
		return nil
	}
}

// '.' SimpleClassTypeSignature
func classTypeSignatureSuffix(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '.' {
		return errorPrefixUnmatch("ClassTypeSignatureSuffix", []rune{'.'}, c)
	}
	return simpleClassTypeSignature(out, r)
}

// 'T' Identifier ';'
func typeVariableSignature(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != 'T' {
		return errorPrefixUnmatch("TypeVariableSignature", []rune{'T'}, c)
	}

	if err = identifier(out, r); err != nil {
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

// '<' TypeArgument+ '>'
func typeArguments(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	} else if c != '<' {
		return errorPrefixUnmatch("TypeArguments", []rune{'<'}, c)
	}

loop:
	for {
		if err = typeArgument(out, r); err != nil {
			return err
		}

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

// WildcardIndicator? FieldTypeSignature
// '*'
func typeArgument(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case '+', '-':
		if err = wildcardIndicator(out, r); err != nil {
			return err
		}
		fallthrough
	case 'L', '[', 'T':
		return fieldTypeSignature(out, r)
	case '*':
		r.ReadRune()
		return nil
	default:
		return errorPrefixUnmatch("TypeArgument", []rune{'+', '-', 'L', '[', 'T', '*'}, c)
	}
}

// '+'
// '-'
func wildcardIndicator(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	switch c {
	case '+':
		return nil
	case '-':
		return nil
	default:
		return errorPrefixUnmatch("WildcardIndicator", []rune{'+', '-'}, c)
	}
}

// '[' TypeSignature
func arrayTypeSignature(out Emitter, r *bufio.Reader) error {
	c, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if c != '[' {
		return errorPrefixUnmatch("ArrayTypeSignature", []rune{'['}, c)
	}

	return typeSignature(out, r)
}

// FieldTypeSignature
// BaseType
func typeSignature(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T':
		return fieldTypeSignature(out, r)
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		// TODO
		out := new(FieldDescriptorInfo)
		return baseType(out, r)
	default:
		return errorPrefixUnmatch("TypeSignature", []rune{'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z'}, c)
	}
}

// FormalTypeParameters? '(' TypeSignature* ')' ReturnType ThrowsSignature*
func methodTypeSignature(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	if c == '<' {
		if err = formalTypeParameters(out, r); err != nil {
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
			return fieldTypeSignature(out, r)
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

	if err = returnType(out, r); err != nil {
		return err
	}

	for {
		c, err = peek(r)
		if err != nil {
			return err
		}
		switch c {
		case '^':
			if err = throwsSignature(out, r); err != nil {
				return err
			}
		default:
			return nil
		}
	}
}

// TypeSignature
// VoidDescriptor
func returnType(out Emitter, r *bufio.Reader) error {
	c, err := peek(r)
	if err != nil {
		return err
	}
	switch c {
	case 'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		return typeSignature(out, r)
	case 'V':
		// TODO
		out := new(FieldDescriptorInfo)
		return voidDescriptor(out, r)
	default:
		return errorPrefixUnmatch("ReturnType", []rune{'L', '[', 'T', 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 'V'}, c)
	}
}

// '^' ClassTypeSignature
// '^' TypeVariableSignature
func throwsSignature(out Emitter, r *bufio.Reader) error {
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
		return classTypeSignature(out, r)
	case 'T':
		return typeVariableSignature(out, r)
	default:
		return fmt.Errorf("", c)
	}
}

// utilities
func peek(r *bufio.Reader) (rune, error) {
	c, _, err := r.ReadRune()
	if err != nil {
		return c, err
	}
	return c, r.UnreadRune()
}

func lookahead(r *bufio.Reader, ch rune) (bool, error) {
	var err error

	found := false
	unreadRunes := 0
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		unreadRunes++
		fmt.Printf("LookAhead ReadRune %c\n", c)

		if c == ch {
			found = true
			break
		}
	}
	for unreadRunes > 0 {
		if err = r.UnreadRune(); err != nil {
			return false, err
		}
		fmt.Println("LookAhead UnreadRune")
		unreadRunes--
	}
	return found, nil
}

func errorPrefixUnmatch(syntax string, expects []rune, actual rune) error {
	chars := make([]string, 0)
	if len(expects) == 1 {
		chars = append(chars, fmt.Sprintf("`%c'", expects[0]))
	} else {
		for _, c := range expects {
			chars = append(chars, fmt.Sprintf("`%c'", c))
		}
	}
	return fmt.Errorf("%s must starts with %s, but with `%c'", syntax, strings.Join(chars, ", "), actual)
}

func errorSuffixUnmatch(syntax string, expects []rune, actual rune) error {
	chars := make([]string, 0)
	if len(expects) == 1 {
		chars = append(chars, fmt.Sprintf("`%c'", expects[0]))
	} else {
		for _, c := range expects {
			chars = append(chars, fmt.Sprintf("`%c'", c))
		}
	}
	return fmt.Errorf("%s must ends with %s, but with `%c'", syntax, strings.Join(chars, ", "), actual)
}
