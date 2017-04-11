package jclass

import (
	"github.com/kamichidu/go-jclass/jvms"
)

type AccessFlags interface {
	IsPublic() bool
	IsPrivate() bool
	IsProtected() bool
	IsStatic() bool
	IsFinal() bool
	IsSuper() bool
	IsSynchronized() bool
	IsVolatile() bool
	IsBridge() bool
	IsTransient() bool
	IsVarargs() bool
	IsNative() bool
	IsInterface() bool
	IsAbstract() bool
	IsStrict() bool
	IsSynthetic() bool
	IsAnnotation() bool
	IsEnum() bool
}

type AccessFlag uint16

func (self AccessFlag) hasBit(bit uint16) bool {
	return uint16(self)&bit == bit
}

func (self AccessFlag) IsPublic() bool       { return self.hasBit(jvms.ACC_PUBLIC) }
func (self AccessFlag) IsPrivate() bool      { return self.hasBit(jvms.ACC_PRIVATE) }
func (self AccessFlag) IsProtected() bool    { return self.hasBit(jvms.ACC_PROTECTED) }
func (self AccessFlag) IsStatic() bool       { return self.hasBit(jvms.ACC_STATIC) }
func (self AccessFlag) IsFinal() bool        { return self.hasBit(jvms.ACC_FINAL) }
func (self AccessFlag) IsSuper() bool        { return self.hasBit(jvms.ACC_SUPER) }
func (self AccessFlag) IsSynchronized() bool { return self.hasBit(jvms.ACC_SYNCHRONIZED) }
func (self AccessFlag) IsVolatile() bool     { return self.hasBit(jvms.ACC_VOLATILE) }
func (self AccessFlag) IsBridge() bool       { return self.hasBit(jvms.ACC_BRIDGE) }
func (self AccessFlag) IsTransient() bool    { return self.hasBit(jvms.ACC_TRANSIENT) }
func (self AccessFlag) IsVarargs() bool      { return self.hasBit(jvms.ACC_VARARGS) }
func (self AccessFlag) IsNative() bool       { return self.hasBit(jvms.ACC_NATIVE) }
func (self AccessFlag) IsInterface() bool    { return self.hasBit(jvms.ACC_INTERFACE) }
func (self AccessFlag) IsAbstract() bool     { return self.hasBit(jvms.ACC_ABSTRACT) }
func (self AccessFlag) IsStrict() bool       { return self.hasBit(jvms.ACC_STRICT) }
func (self AccessFlag) IsSynthetic() bool    { return self.hasBit(jvms.ACC_SYNTHETIC) }
func (self AccessFlag) IsAnnotation() bool   { return self.hasBit(jvms.ACC_ANNOTATION) }
func (self AccessFlag) IsEnum() bool         { return self.hasBit(jvms.ACC_ENUM) }

var _ AccessFlags = AccessFlag(uint16(0))
