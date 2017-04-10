package jvms

import (
	"github.com/kamichidu/go-jclass"
	"os"
	"reflect"
	"testing"
)

func BenchmarkParseClassFile(b *testing.B) {
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

func TestParseClassFile(t *testing.T) {
	cases := []struct {
		Filename string
		Expected *jclass.ClassFile
	}{
		{"./testdata/String.class", &jclass.ClassFile{
			Magic:             uint32(0xcafebabe),
			MinorVersion:      uint16(0x0000),
			MajorVersion:      uint16(0x0034),
			ConstantPoolCount: uint16(0x0222),
			ConstantPool: []jclass.ConstantPoolInfo{
				&jclass.ConstantMethodrefInfo{
					ClassIndex:       uint16(0x0092),
					NameAndTypeIndex: uint16(0x012e),
				},
			},
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
