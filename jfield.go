package jclass

import (
	"fmt"
	"github.com/kamichidu/go-jclass/data"
	c "github.com/kamichidu/go-jclass/parser/common"
	"github.com/kamichidu/go-jclass/parser/fd"
	"math"
	"reflect"
	"strings"
)

type JField struct {
	cp   []data.CpInfo
	data *data.FieldInfo
}

func newJField(cp []data.CpInfo, data *data.FieldInfo) *JField {
	return &JField{
		cp:   cp,
		data: data,
	}
}

func (self *JField) GetAccessFlags() uint16 {
	return self.data.AccessFlags
}

func (self *JField) GetName() string {
	return getUtf8String(self.cp, self.data.NameIndex)
}

func (self *JField) getDescriptor() string {
	return getUtf8String(self.cp, self.data.DescriptorIndex)
}

func (self *JField) GetType() string {
	ret, err := fd.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}

	return ret
}

func (self *JField) GetAttribute(typ reflect.Type) data.AttributeInfo {
	for _, attr := range self.data.Attributes {
		if reflect.TypeOf(attr).AssignableTo(typ) {
			return attr
		}
	}
	return nil
}

func (self *JField) GetConstantValue() interface{} {
	var cv *data.ConstantValueAttribute
	if found := self.GetAttribute(reflect.TypeOf(cv)); found != nil {
		cv, _ = found.(*data.ConstantValueAttribute)
	} else {
		return nil
	}
	cpInfo := self.cp[cv.ConstantValueIndex]
	switch value := cpInfo.(type) {
	case data.LongInfo:
		// long
		return int64(uint64(value.HighBytes)<<32 | uint64(value.LowBytes))
	case data.FloatInfo:
		// float
		return math.Float32frombits(value.Bytes)
	case data.DoubleInfo:
		// double
		return math.Float64frombits(uint64(value.HighBytes)<<32 | uint64(value.LowBytes))
	case data.IntegerInfo:
		// int, short, char, byte, or boolean
		switch self.GetType() {
		case "boolean":
			return value.Bytes != 0
		case "byte":
			return int8(value.Bytes)
		case "char":
			return rune(value.Bytes)
		case "short":
			return int16(value.Bytes)
		case "int":
			fallthrough
		default:
			return int32(value.Bytes)
		}
	case data.StringInfo:
		// String
		return getUtf8String(self.cp, value.StringIndex)
	default:
		panic("???")
	}
}

func toString(ast *c.FieldDescriptor) string {
	return toStringImpl(ast)
}

func toStringImpl(val interface{}) string {
	switch v := val.(type) {
	case *c.FieldDescriptor:
		return toStringImpl(v.FieldType)
	case *c.FieldType:
		if v.BaseType != nil {
			return toStringImpl(v.BaseType)
		} else if v.ObjectType != nil {
			return toStringImpl(v.ObjectType)
		} else {
			return toStringImpl(v.ArrayType)
		}
	case *c.BaseType:
		return v.Text
	case *c.ObjectType:
		return toStringImpl(v.ClassName)
	case *c.ClassName:
		return strings.Join(v.Identifier, ".")
	case *c.ArrayType:
		return toStringImpl(v.ComponentType) + "[]"
	case *c.ComponentType:
		return toStringImpl(v.FieldType)
	case *c.MethodDescriptor:
		params := make([]string, len(v.ParameterDescriptor))
		for i := 0; i < len(params); i++ {
			params[i] = toStringImpl(v.ParameterDescriptor[i])
		}
		ret := toStringImpl(v.ReturnDescriptor)

		return strings.Join(params, ", ") + " : " + ret
	case *c.ReturnDescriptor:
		if v.FieldType != nil {
			return toStringImpl(v.FieldType)
		} else {
			return toStringImpl(v.VoidDescriptor)
		}
	case *c.VoidDescriptor:
		return v.Text
	case *c.ParameterDescriptor:
		return toStringImpl(v.FieldType)
	default:
		panic(fmt.Errorf("Unknown type (%s) detected.", reflect.TypeOf(v).String()))
	}
}
