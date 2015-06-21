package test

import (
    "github.com/kamichidu/go-jclass"
    "testing"
)

func Test(t *testing.T) {
    jc, err := jclass.NewJClassWithFilename("String.class")
    if err != nil {
        t.Error(err)
    }

    if jc.GetName() != "java/lang/String" {
        t.Errorf("Name %s", jc.GetName())
    }
    if jc.GetSimpleName() != "String" {
        t.Errorf("Simple name %s", jc.GetSimpleName())
    }
    if jc.GetCanonicalName() != "java.lang.String" {
        t.Errorf("Canonical name %s", jc.GetCanonicalName())
    }
}
