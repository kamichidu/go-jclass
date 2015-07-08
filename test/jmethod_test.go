package test

import (
	"github.com/k0kubun/pp"
	"github.com/kamichidu/go-jclass"
	"testing"
)

func init() {
	pp.ColoringEnabled = false
}

func TestJMethod(t *testing.T) {
	jc, err := jclass.NewJClassWithFilename("./String.class")
	if err != nil {
		t.Fatal(err)
	}

	methods := jc.GetMethod("toUpperCase")
	if len(methods) != 2 {
		t.Errorf("toUpperCase() has 2 overloads, but got %s", pp.Sprint(methods))
	}
	for _, method := range methods {
		if len(method.GetParameterTypes()) == 0 {
			if method.GetReturnType() != "java.lang.String" {
				t.Errorf("toUpperCase(void) returns java.lang.String, but as %s", method.GetReturnType())
			}
		} else {
			if method.GetReturnType() != "java.lang.String" {
				t.Errorf("toUpperCase(java.util.Locale) returns java.lang.String, but as %s", method.GetReturnType())
			}
			if len(method.GetParameterTypes()) != 1 {
				t.Errorf("toUpperCase(java.util.Locale) takes 1 argument, but as %s", pp.Sprint(method.GetParameterTypes()))
			}
			if method.GetParameterTypes()[0] != "java.util.Locale" {
				t.Errorf("toUpperCase(java.util.Locale) takes java.util.Locale, but as %s", pp.Sprint(method.GetParameterTypes()))
			}
		}
	}
}
