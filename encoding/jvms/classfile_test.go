package jvms

import (
	"os"
	"reflect"
	"testing"

	"github.com/kamichidu/go-jclass/jvms"
	"github.com/kr/pretty"
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
			// ConstantPool:      make([]jvms.ConstantPoolInfo, 0x20),
			ConstantPool: []jvms.ConstantPoolInfo{
				nil,
				&jvms.ConstantMethodrefInfo{
					ClassIndex:       0x0008,
					NameAndTypeIndex: 0x0015,
				},
				&jvms.ConstantClassInfo{
					NameIndex: 0x0016,
				},
				&jvms.ConstantMethodrefInfo{
					ClassIndex:       0x0002,
					NameAndTypeIndex: 0x0015,
				},
				&jvms.ConstantMethodrefInfo{
					ClassIndex:       0x0002,
					NameAndTypeIndex: 0x0017,
				},
				&jvms.ConstantStringInfo{
					StringIndex: 0x0018,
				},
				&jvms.ConstantMethodrefInfo{
					ClassIndex:       0x0002,
					NameAndTypeIndex: 0x0019,
				},
				&jvms.ConstantClassInfo{
					NameIndex: 0x001a,
				},
				&jvms.ConstantClassInfo{
					NameIndex: 0x001b,
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0003,
					Bytes:  []uint8{0x6d, 0x73, 0x67},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0012,
					Bytes:  []uint8{0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0001,
					Bytes:  []uint8{0x6e},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0001,
					Bytes:  []uint8{0x49},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0006,
					Bytes:  []uint8{0x3c, 0x69, 0x6e, 0x69, 0x74, 0x3e},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0003,
					Bytes:  []uint8{0x28, 0x29, 0x56},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0004,
					Bytes:  []uint8{0x43, 0x6f, 0x64, 0x65},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x000f,
					Bytes:  []uint8{0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0003,
					Bytes:  []uint8{0x73, 0x61, 0x79},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0026,
					Bytes:  []uint8{0x28, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x000a,
					Bytes:  []uint8{0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0009,
					Bytes:  []uint8{0x54, 0x65, 0x73, 0x74, 0x2e, 0x6a, 0x61, 0x76, 0x61},
				},
				&jvms.ConstantNameAndTypeInfo{
					NameIndex:       0x000d,
					DescriptorIndex: 0x000e,
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0017,
					Bytes:  []uint8{0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72},
				},
				&jvms.ConstantNameAndTypeInfo{
					NameIndex:       0x001c,
					DescriptorIndex: 0x001d,
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0002,
					Bytes:  []uint8{0x68, 0x69},
				},
				&jvms.ConstantNameAndTypeInfo{
					NameIndex:       0x001e,
					DescriptorIndex: 0x001f,
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0004,
					Bytes:  []uint8{0x54, 0x65, 0x73, 0x74},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0010,
					Bytes:  []uint8{0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0006,
					Bytes:  []uint8{0x61, 0x70, 0x70, 0x65, 0x6e, 0x64},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x002d,
					Bytes:  []uint8{0x28, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x3b},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0008,
					Bytes:  []uint8{0x74, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67},
				},
				&jvms.ConstantUtf8Info{
					Length: 0x0014,
					Bytes:  []uint8{0x28, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b},
				},
			},
			AccessFlags:     uint16(0x0021),
			ThisClass:       uint16(0x0007),
			SuperClass:      uint16(0x0008),
			InterfacesCount: uint16(0x0000),
			Interfaces:      make([]uint16, 0x0),
			FieldsCount:     uint16(0x0002),
			Fields: []*jvms.FieldInfo{
				&jvms.FieldInfo{
					AccessFlags:     0x0009,
					NameIndex:       0x0009,
					DescriptorIndex: 0x000a,
					AttributesCount: 0x0000,
					Attributes:      make([]jvms.AttributeInfo, 0),
				},
				&jvms.FieldInfo{
					AccessFlags:     0x0002,
					NameIndex:       0x000b,
					DescriptorIndex: 0x000c,
					AttributesCount: 0x0000,
					Attributes:      make([]jvms.AttributeInfo, 0),
				},
			},
			MethodsCount: uint16(0x0002),
			Methods: []*jvms.MethodInfo{
				&jvms.MethodInfo{
					AccessFlags:     0x0001,
					NameIndex:       0x000d,
					DescriptorIndex: 0x000e,
					AttributesCount: 0x0001,
					Attributes: []jvms.AttributeInfo{
						&jvms.GenericAttributeInfo{
							AttributeNameIndex_: 0x000f,
							AttributeLength_:    0x001d,
							Info_:               []uint8{0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x0, 0x5, 0x2a, 0xb7, 0x0, 0x1, 0xb1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x10, 0x0, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x0, 0x0, 0x5},
						},
					},
				},
				&jvms.MethodInfo{
					AccessFlags:     0x0001,
					NameIndex:       0x0011,
					DescriptorIndex: 0x0012,
					AttributesCount: 0x0001,
					Attributes: []jvms.AttributeInfo{
						&jvms.GenericAttributeInfo{
							AttributeNameIndex_: 0x000f,
							AttributeLength_:    0x002c,
							Info_:               []uint8{0x0, 0x2, 0x0, 0x2, 0x0, 0x0, 0x0, 0x14, 0xbb, 0x0, 0x2, 0x59, 0xb7, 0x0, 0x3, 0x2b, 0xb6, 0x0, 0x4, 0x12, 0x5, 0xb6, 0x0, 0x4, 0xb6, 0x0, 0x6, 0xb0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x10, 0x0, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x0, 0x0, 0x6},
						},
					},
				},
			},
			AttributesCount: uint16(0x0001),
			Attributes: []jvms.AttributeInfo{
				&jvms.SourceFileAttribute{
					AttributeNameIndex_: 0x0013,
					AttributeLength_:    0x0002,
					SourceFileIndex:     0x0014,
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
			t.Fail()
			diffs := pretty.Diff(cf, c.Expected)
			for _, diff := range diffs {
				t.Log(diff)
			}
		}
	}
}
