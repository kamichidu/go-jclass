package jclass

import (
	"github.com/k0kubun/pp"
	"testing"
)

func init() {
	pp.ColoringEnabled = false
}

func TestJMethod(t *testing.T) {
	jc, err := NewJavaClassFromFilename("./testdata/String.class")
	if err != nil {
		t.Fatal(err)
	}

	methods := jc.Method("toUpperCase")
	if len(methods) != 2 {
		t.Errorf("toUpperCase() has 2 overloads, but got %s", pp.Sprint(methods))
	}
	for _, method := range methods {
		if len(method.ParameterTypes()) == 0 {
			if method.ReturnType() != "java.lang.String" {
				t.Errorf("toUpperCase(void) returns java.lang.String, but as %s", method.ReturnType())
			}
		} else {
			if method.ReturnType() != "java.lang.String" {
				t.Errorf("toUpperCase(java.util.Locale) returns java.lang.String, but as %s", method.ReturnType())
			}
			if len(method.ParameterTypes()) != 1 {
				t.Errorf("toUpperCase(java.util.Locale) takes 1 argument, but as %s", pp.Sprint(method.ParameterTypes()))
			}
			if method.ParameterTypes()[0] != "java.util.Locale" {
				t.Errorf("toUpperCase(java.util.Locale) takes java.util.Locale, but as %s", pp.Sprint(method.ParameterTypes()))
			}
		}
	}
}
