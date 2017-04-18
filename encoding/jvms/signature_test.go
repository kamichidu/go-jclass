package jvms

import (
	_ "github.com/kr/pretty"
	_ "reflect"
	"strings"
	"testing"
)

func TestParseClassSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
		{S: "Ljava/lang/Object;"},
		{S: "<E:Ljava/lang/Object;>Ljava/util/AbstractList<TE;>;Ljava/util/List<TE;>;Ljava/util/RandomAccess;Ljava/lang/Cloneable;Ljava/io/Serializable;"},
	}
	for _, c := range cases {
		if err := ParseClassSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse ClassSignature %#v: %s", c.S, err)
		}
	}
}

func TestParseFieldTypeSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
	// {S: ""},
	}
	for _, c := range cases {
		if err := ParseFieldTypeSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse FieldTypeSignature %#v: %s", c.S, err)
		}
	}
}

func TestParseMethodTypeSignature(t *testing.T) {
	cases := []struct {
		S string
	}{
	// {S: ""},
	}
	for _, c := range cases {
		if err := ParseMethodTypeSignature(strings.NewReader(c.S)); err != nil {
			t.Errorf("Can't parse MethodTypeSignature %#v: %s", c.S, err)
		}
	}
}
