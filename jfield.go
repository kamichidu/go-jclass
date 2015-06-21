package jclass

import (
	"github.com/kamichidu/go-jclass/data"
	"github.com/kamichidu/go-jclass/parser/fd"
	"math"
	"reflect"
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

func (self *JField) GetType() JType {
	fdinfo, err := fd.Parse(self.getDescriptor())
	if err != nil {
		panic(err)
	}

	return newJType(fdinfo)
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
		switch self.GetType().GetTypeName() {
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

func newJType(fdinfo *fd.FDInfo) JType {
	if fdinfo.PrimitiveType {
		return NewJPrimitiveType(fdinfo.TypeName)
	} else if fdinfo.ReferenceType {
		return NewJReferenceType(fdinfo.TypeName)
	} else if fdinfo.ArrayType {
		ct := fdinfo.ComponentType
		if ct.PrimitiveType {
			return NewJArrayType(NewJPrimitiveType(ct.TypeName), fdinfo.Dims)
		} else if ct.ReferenceType {
			return NewJArrayType(NewJReferenceType(ct.TypeName), fdinfo.Dims)
		} else {
			panic("???")
		}
	} else {
		panic("???")
	}
}
