%{
package fd

import (
    "github.com/kamichidu/go-jclass"
)

type FDToken struct {
    Id int
    Text string
    Pos int
}
%}

%type <jtype> FieldDescriptor
%type <jtype> FieldType
%type <jtype> BaseType
%type <jtype> ObjectType
%type <jtype> ArrayType
%type <jtype> ComponentType

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '['
%token <token> CLASS_NAME

%union {
    jtype jclass.JType
    token FDToken
}

%%

FieldDescriptor
    : FieldType
        {
            $$ = $1
            if lexer, ok := fdlex.(*FDLexer); ok {
                lexer.Result = $$
            }
        }
    ;

FieldType
    : BaseType
    | ObjectType
    | ArrayType
    ;

BaseType
    : 'B'
        {
            $$ = jclass.NewJPrimitiveType("byte")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "byte"
            // }
        }
    | 'C'
        {
            $$ = jclass.NewJPrimitiveType("char")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "char"
            // }
        }
    | 'D'
        {
            $$ = jclass.NewJPrimitiveType("double")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "double"
            // }
        }
    | 'F'
        {
            $$ = jclass.NewJPrimitiveType("float")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "float"
            // }
        }
    | 'I'
        {
            $$ = jclass.NewJPrimitiveType("int")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "int"
            // }
        }
    | 'J'
        {
            $$ = jclass.NewJPrimitiveType("long")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "long"
            // }
        }
    | 'S'
        {
            $$ = jclass.NewJPrimitiveType("short")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "short"
            // }
        }
    | 'Z'
        {
            $$ = jclass.NewJPrimitiveType("boolean")
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = "boolean"
            // }
        }
    ;

ObjectType
    : 'L' CLASS_NAME ';'
        {
            $$ = jclass.NewJReferenceType($2.Text)
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.TypeName = $2.Text
            // }
        }
    ;

ArrayType
    : '[' ComponentType
        {
            switch jtype := $2.(type) {
            case *jclass.JPrimitiveType:
                $$ = jclass.NewJArrayType(jtype, 1)
            case *jclass.JReferenceType:
                $$ = jclass.NewJArrayType(jtype, 1)
            case *jclass.JArrayType:
                $$ = jclass.NewJArrayType(jtype.GetComponentType(), jtype.GetDims() + 1)
            default:
                panic("??? Siranai Kata da")
            }
            // $$ = jclass.NewJArrayType($2, 1)
            // if lexer, ok := yylex.(*DescriptorLexer); ok {
            //     lexer.result.Dims++
            // }
        }
    ;

ComponentType
    : FieldType
    ;

%%
