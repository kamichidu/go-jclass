package jvms

import (
	"github.com/kamichidu/go-jclass/jvms"
	"os"
	"reflect"
	"testing"
)

func BenchmarkParseClassFileWithMidiumString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./testdata/String.class")
		if err != nil {
			b.Fatal(err)
		}
		defer file.Close()

		cf, err := ParseClassFile(file)
		if err != nil {
			b.Fatal(err)
		} else if cf == nil {
			b.Fatalf("nil value returned???")
		}
	}
}

func BenchmarkParseClassFileWithLargeClass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./testdata/ORBUtilSystemException.class")
		if err != nil {
			b.Fatal(err)
		}
		defer file.Close()

		cf, err := ParseClassFile(file)
		if err != nil {
			b.Fatal(err)
		} else if cf == nil {
			b.Fatalf("nil value returned???")
		}
	}
}

func TestParseClassFile(t *testing.T) {
	cases := []struct {
		Filename string
		Expected *jvms.ClassFile
	}{
		{"./testdata/Test.class", &jvms.ClassFile{
			Magic:             uint32(0xcafebabe),
			MinorVersion:      uint16(0x0000),
			MajorVersion:      uint16(0x0034),
			ConstantPoolCount: uint16(0x0020),
			ConstantPool:      make([]jvms.ConstantPoolInfo, 0x20),
			AccessFlags:       uint16(0x0021),
			ThisClass:         uint16(0x0007),
			SuperClass:        uint16(0x0008),
			InterfacesCount:   uint16(0x0000),
			Interfaces:        make([]uint16, 0x0),
			FieldsCount:       uint16(0x0002),
			Fields:            make([]*jvms.FieldInfo, 0x2),
			MethodsCount:      uint16(0x0002),
			Methods:           make([]*jvms.MethodInfo, 0x2),
			AttributesCount:   uint16(0x0001),
			Attributes:        make([]*jvms.AttributeInfo, 0x1),
		}},
	}
	for _, c := range cases {
		file, err := os.Open(c.Filename)
		if err != nil {
			t.Errorf("Can't open test file %s: %s", c.Filename, err)
			continue
		}
		defer file.Close()

		cf, err := ParseClassFile(file)
		if err != nil {
			t.Errorf("Parse failed: %s", err)
			continue
		} else if !reflect.DeepEqual(cf, c.Expected) {
			t.Errorf("Oops,\nexpected %#v\nactual   %#v", c.Expected, cf)
		}
	}
}
