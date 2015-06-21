package jclass

import (
	"github.com/kamichidu/go-jclass/data"
	"reflect"
)

type JAttribute struct {
	cp   []data.CpInfo
	data data.AttributeInfo
}

func newJAttribute(cp []data.CpInfo, data data.AttributeInfo) *JAttribute {
	return &JAttribute{
		cp:   cp,
		data: data,
	}
}

func (self *JAttribute) GetName() string {
	val := reflect.ValueOf(self.data).Elem().Elem().Elem()
	nameIndex := uint16(val.FieldByName("AttributeNameIndex").Uint())
	return getUtf8String(self.cp, nameIndex)
}
