package jclass

import (
	"testing"
)

func TestJavaClassNames(t *testing.T) {
	cases := []struct {
		Filename      string
		PackageName   string
		SimpleName    string
		CanonicalName string
		Name          string
	}{
		{"./testdata/Constants.class", "", "Constants", "Constants", "Constants"},
		{"./testdata/String.class", "java.lang", "String", "java.lang.String", "java.lang.String"},
		{"./testdata/Map$Entry.class", "java.util", "Entry", "java.util.Map.Entry", "java.util.Map$Entry"},
	}
	for _, c := range cases {
		class, err := NewJavaClassFromFilename(c.Filename)
		if err != nil {
			t.Fatalf("NewJavaClassFromFilename: %s", err)
		}

		if class.PackageName() != c.PackageName {
			t.Errorf("PackageName %v, wants %v", class.PackageName(), c.PackageName)
		}
		if class.SimpleName() != c.SimpleName {
			t.Errorf("SimpleName %v, wants %v", class.SimpleName(), c.SimpleName)
		}
		if class.CanonicalName() != c.CanonicalName {
			t.Errorf("CanonicalName %v, wants %v", class.CanonicalName(), c.CanonicalName)
		}
		if class.Name() != c.Name {
			t.Errorf("Name %v, wants %v", class.Name(), c.Name)
		}
	}
}
