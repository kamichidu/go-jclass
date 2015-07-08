package md

import (
	c "github.com/kamichidu/go-jclass/parser/common"
	parser "github.com/kamichidu/go-jclass/parser/md"
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
            s += format(ast.Index(i))
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

func TestParseMethodDescriptor(t *testing.T) {
	ok := func(md string, expect interface{}) {
		t.Logf("Try to parse '%s'", md)
		if ret, err := parser.Parse(md); err == nil {
			if !reflect.DeepEqual(ret, expect) {
                t.Errorf("Expected %s", format(reflect.ValueOf(expect)))
                t.Errorf("Got      %s", format(reflect.ValueOf(ret)))
			}
			t.Log(" -> OK")
		} else {
			t.Errorf("  ERR: %v", err)
		}
	}

	ok("()V", &c.MethodDescriptor{
		ParameterDescriptor: make([]*c.ParameterDescriptor, 0),
		ReturnDescriptor: &c.ReturnDescriptor{
			VoidDescriptor: &c.VoidDescriptor{"void"},
		},
	})
    ok("(IDLjava/lang/Thread;)Ljava/lang/Object;", &c.MethodDescriptor{
        ParameterDescriptor: []*c.ParameterDescriptor{
            &c.ParameterDescriptor{
                FieldType: &c.FieldType{
                    BaseType: &c.BaseType{"int"},
                },
            },
            &c.ParameterDescriptor{
                FieldType: &c.FieldType{
                    BaseType: &c.BaseType{"double"},
                },
            },
            &c.ParameterDescriptor{
                FieldType: &c.FieldType{
                    ObjectType: &c.ObjectType{
                        ClassName: &c.ClassName{[]string{"java", "lang", "Thread"}},
                    },
                },
            },
        },
        ReturnDescriptor: &c.ReturnDescriptor{
            FieldType: &c.FieldType{
                ObjectType: &c.ObjectType{
                    ClassName: &c.ClassName{[]string{"java", "lang", "Object"}},
                },
            },
        },
    })
}
