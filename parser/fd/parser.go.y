%{
package fd

import (
)

type FDToken struct {
    Id int
    Text string
    Pos int
}
%}

%type <res> FieldDescriptor
%type <res> FieldType
%type <res> BaseType
%type <res> ObjectType
%type <res> ArrayType
%type <res> ComponentType

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '['
%token <token> CLASS_NAME

%union {
    res   *FDInfo
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
            $$ = NewPrimitiveType("byte")
        }
    | 'C'
        {
            $$ = NewPrimitiveType("char")
        }
    | 'D'
        {
            $$ = NewPrimitiveType("double")
        }
    | 'F'
        {
            $$ = NewPrimitiveType("float")
        }
    | 'I'
        {
            $$ = NewPrimitiveType("int")
        }
    | 'J'
        {
            $$ = NewPrimitiveType("long")
        }
    | 'S'
        {
            $$ = NewPrimitiveType("short")
        }
    | 'Z'
        {
            $$ = NewPrimitiveType("boolean")
        }
    ;

ObjectType
    : 'L' CLASS_NAME ';'
        {
            $$ = NewReferenceType($2.Text)
        }
    ;

ArrayType
    : '[' ComponentType
        {
            if $2.PrimitiveType || $2.ReferenceType {
                $$ = NewArrayType($2, 1)
            } else if $2.ArrayType {
                $$ = NewArrayType($2.ComponentType, $2.Dims + 1)
            } else {
                panic("??? Siranai Kata da")
            }
        }
    ;

ComponentType
    : FieldType
    ;

%%
