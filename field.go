package jclass

import (
	"fmt"
	"github.com/kamichidu/go-jclass/jvms"
	"strings"
)

type JavaField struct {
	AccessFlags

	constantPool []jvms.ConstantPoolInfo
	fieldInfo    *jvms.FieldInfo
}

func newJavaField(constantPool []jvms.ConstantPoolInfo, fieldInfo *jvms.FieldInfo) *JavaField {
	return &JavaField{AccessFlag(fieldInfo.AccessFlags), constantPool, fieldInfo}
}

func (self *JavaField) Name() string {
	utf8Info := self.constantPool[self.fieldInfo.NameIndex].(*jvms.ConstantUtf8Info)
	return utf8Info.JavaString()
}

func (self *JavaField) Type() string {
	utf8Info := self.constantPool[self.fieldInfo.DescriptorIndex].(*jvms.ConstantUtf8Info)
	descriptor := utf8Info.JavaString()

	info, err := jvms.ParseFieldDescriptor(strings.NewReader(descriptor))
	if err != nil {
		// TODO: Error handling
		return ""
	}
	return info.String()
}

func (self *JavaField) ConstantValue() (interface{}, bool) {
	var cv *jvms.ConstantValueAttribute
	for _, attr := range self.fieldInfo.Attributes {
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
