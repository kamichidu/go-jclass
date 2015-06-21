%{
package md

import (
    "github.com/kamichidu/go-jclass/parser/fd"
)

type MDToken struct {
    Id int
    Text string
    Pos int
}
%}

%type <params> ParameterDescriptors
%type <jtype> ParameterDescriptor
%type <jtype> ReturnDescriptor
%type <jtype> FieldType
%type <jtype> BaseType
%type <jtype> ObjectType
%type <jtype> ArrayType
%type <jtype> ComponentType

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '[' '(' ')'
%token <token> CLASS_NAME

%union {
    jtype  *fd.FDInfo
    params []*fd.FDInfo
    token  MDToken
}

%%

MethodDescriptor
    : '(' ParameterDescriptors ')' ReturnDescriptor
        {
            if l, ok := mdlex.(*MDLexer); ok {
                l.Result = &MDInfo{
                    ReturnType:     $4,
                    ParameterTypes: $2,
                }
            }
        }
    ;

ParameterDescriptors
    :
        {
            $$ = make([]*fd.FDInfo, 0)
        }
    | ParameterDescriptors ParameterDescriptor
        {
            $$ = append($1, $2)
        }
    ;

ParameterDescriptor
    : FieldType
    ;

ReturnDescriptor
    : FieldType
    | 'V'
        {
            $$ = fd.NewPrimitiveType("void")
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
            $$ = fd.NewPrimitiveType("byte")
        }
    | 'C'
        {
            $$ = fd.NewPrimitiveType("char")
        }
    | 'D'
        {
            $$ = fd.NewPrimitiveType("double")
        }
    | 'F'
        {
            $$ = fd.NewPrimitiveType("float")
        }
    | 'I'
        {
            $$ = fd.NewPrimitiveType("int")
        }
    | 'J'
        {
            $$ = fd.NewPrimitiveType("long")
        }
    | 'S'
        {
            $$ = fd.NewPrimitiveType("short")
        }
    | 'Z'
        {
            $$ = fd.NewPrimitiveType("boolean")
        }
    ;

ObjectType
    : 'L' CLASS_NAME ';'
        {
            $$ = fd.NewReferenceType($2.Text)
        }
    ;

ArrayType
    : '[' ComponentType
        {
            if $2.PrimitiveType || $2.ReferenceType {
                $$ = fd.NewArrayType($2, 1)
            } else if $2.ArrayType {
                $$ = fd.NewArrayType($2.ComponentType, $2.Dims + 1)
            } else {
                panic("??? Siranai Kata da")
            }
        }
    ;

ComponentType
    : FieldType
    ;

%%
