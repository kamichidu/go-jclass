package fd

import (
	c "github.com/kamichidu/go-jclass/parser/common"
	fdparser "github.com/kamichidu/go-jclass/parser/fd"
	"reflect"
	"testing"
)

func format(ast reflect.Value) string {
	if ast.Kind() == reflect.Ptr {
		if ast.IsNil() {
			return "nil"
		} else {
			return format(ast.Elem())
		}
	}
	if ast.Kind() == reflect.Slice || ast.Kind() == reflect.Array {
		s := "["
		for i := 0; i < ast.Len(); i++ {
			s += ast.Index(i).String()
			s += ","
		}
		s += "]"
		return s
	} else if ast.Kind() != reflect.Struct {
		return ast.String()
	}

	s := ast.Type().String()
	s += "("
	for i := 0; i < ast.NumField(); i++ {
		s += format(ast.Field(i))
		s += ","
	}
	s += ")"
	return s
}

func TestParse(t *testing.T) {
	ok := func(fd string, expect interface{}) {
		t.Logf("Try to parse '%s'", fd)
		if ret, err := fdparser.Parse(fd); err == nil {
			if !reflect.DeepEqual(ret, expect) {
				t.Errorf("Expected %s", format(reflect.ValueOf(expect)))
				t.Errorf("Got      %s", format(reflect.ValueOf(ret)))
			}
		} else {
			t.Errorf("error: %v", err)
		}
	}

	ok("B", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"byte"}}})
	ok("C", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"char"}}})
	ok("D", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"double"}}})
	ok("F", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"float"}}})
	ok("I", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"int"}}})
	ok("J", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"long"}}})
	ok("S", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"short"}}})
	ok("Z", &c.FieldDescriptor{&c.FieldType{BaseType: &c.BaseType{"boolean"}}})
	ok("Ljava/lang/Object;", &c.FieldDescriptor{&c.FieldType{ObjectType: &c.ObjectType{&c.ClassName{[]string{"java", "lang", "Object"}}}}})
	// ok("[B", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"byte"}}}})
	// ok("[C", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"char"}}}})
	// ok("[D", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"double"}}}})
	// ok("[F", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"float"}}}})
	// ok("[I", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"int"}}}})
	// ok("[J", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"long"}}}})
	// ok("[S", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"short"}}}})
	// ok("[Z", &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"boolean"}}}})
	// ok("[Ljava/lang/Object;", &c.ArrayType{&c.ComponentType{&c.FieldType{ObjectType: &c.ObjectType{&c.ClassName{[]string{"java", "lang", "Object"}}}}}})
	// ok("[[B", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"byte"}}}}}}})
	// ok("[[C", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"char"}}}}}}})
	// ok("[[D", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"double"}}}}}}})
	// ok("[[F", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"float"}}}}}}})
	// ok("[[I", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"int"}}}}}}})
	// ok("[[J", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"long"}}}}}}})
	// ok("[[S", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"short"}}}}}}})
	// ok("[[Z", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{BaseType: &c.BaseType{"boolean"}}}}}}})
	// ok("[[Ljava/lang/Object;", &c.ArrayType{&c.ComponentType{&c.FieldType{ArrayType: &c.ArrayType{&c.ComponentType{&c.FieldType{ObjectType: &c.ObjectType{&c.ClassName{[]string{"java", "lang", "Object"}}}}}}}}})
}

func BenchmarkParseFileDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fdparser.Parse("[[[[[[Ljava/util/List;")
	}
}
