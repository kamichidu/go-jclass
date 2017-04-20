package jclass

import (
	"fmt"
	ejvms "github.com/kamichidu/go-jclass/encoding/jvms"
	"github.com/kamichidu/go-jclass/jvms"
	"strings"
)

type JavaField struct {
	*jvms.FieldInfo
	AccessFlags

	constantPool []jvms.ConstantPoolInfo
}

func NewJavaField(constantPool []jvms.ConstantPoolInfo, fieldInfo *jvms.FieldInfo) *JavaField {
	return &JavaField{fieldInfo, AccessFlag(fieldInfo.AccessFlags), constantPool}
}

func (self *JavaField) Name() string {
	utf8Info := self.constantPool[self.NameIndex].(*jvms.ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaField) Type() string {
	utf8Info := self.constantPool[self.DescriptorIndex].(*jvms.ConstantUtf8Info)
	descriptor := utf8Info.JavaString()

	info, err := ejvms.ParseFieldDescriptor(strings.NewReader(descriptor))
	if err != nil {
		// TODO: Error handling
		return ""
	}
	return info.String()
}

func (self *JavaField) ConstantValue() (interface{}, bool) {
	var cv *jvms.ConstantValueAttribute
	for _, attr := range self.FieldInfo.Attributes {
		if v, ok := attr.(*jvms.ConstantValueAttribute); ok {
			cv = v
			break
		}
	}
	if cv == nil {
		return nil, false
	}

	switch v := self.constantPool[cv.ConstantvalueIndex].(type) {
	case *jvms.ConstantLongInfo:
		return v.Long(), true
	case *jvms.ConstantFloatInfo:
		return v.Float(), true
	case *jvms.ConstantDoubleInfo:
		return v.Double(), true
	case *jvms.ConstantIntegerInfo:
		switch self.Type() {
		case "boolean":
			return v.Boolean(), true
		case "char":
			return v.Char(), true
		case "byte":
			return v.Byte(), true
		case "short":
			return v.Short(), true
		case "int":
			return v.Int(), true
		default:
			panic(fmt.Sprintf("Unsupported field type: %s", self.Type()))
		}
	case *jvms.ConstantStringInfo:
		return self.constantPool[v.StringIndex].(*jvms.ConstantUtf8Info).JavaString(), true
	default:
		panic(fmt.Sprintf("Unsupported constant value type: %#v", v))
	}
}
