%{
package fd

import (
    c "github.com/kamichidu/go-jclass/parser/common"
)
%}

%type <fieldDescriptor> FieldDescriptor
%type <fieldType> FieldType
%type <baseType> BaseType
%type <objectType> ObjectType
%type <arrayType> ArrayType
%type <componentType> ComponentType
%type <className> ClassName

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '['
%token <token> IDENTIFIER

%union {
    fieldDescriptor *c.FieldDescriptor
    fieldType       *c.FieldType
    baseType        *c.BaseType
    objectType      *c.ObjectType
    arrayType       *c.ArrayType
    componentType   *c.ComponentType
    className       *c.ClassName
    token           *c.Token
}

%%

FieldDescriptor
    : FieldType
        {
            $$ = &c.FieldDescriptor{
                FieldType: $1,
            }
            if lexer, ok := fdlex.(*FDLexer); ok {
                lexer.Result = $$
            }
        }
    ;

FieldType
    : BaseType
        {
            $$ = &c.FieldType{
                BaseType: $1,
            }
        }
    | ObjectType
        {
            $$ = &c.FieldType{
                ObjectType: $1,
            }
        }
    | ArrayType
        {
            $$ = &c.FieldType{
                ArrayType: $1,
            }
        }
    ;

BaseType
    : 'B'
        {
            $$ = &c.BaseType{"byte"}
        }
    | 'C'
        {
            $$ = &c.BaseType{"char"}
        }
    | 'D'
        {
            $$ = &c.BaseType{"double"}
        }
    | 'F'
        {
            $$ = &c.BaseType{"float"}
        }
    | 'I'
        {
            $$ = &c.BaseType{"int"}
        }
    | 'J'
        {
            $$ = &c.BaseType{"long"}
        }
    | 'S'
        {
            $$ = &c.BaseType{"short"}
        }
    | 'Z'
        {
            $$ = &c.BaseType{"boolean"}
        }
    ;

ObjectType
    : 'L' ClassName ';'
        {
            $$ = &c.ObjectType{
                ClassName: $2,
            }
        }
    ;

ArrayType
    : '[' ComponentType
        {
            $$ = &c.ArrayType{
                ComponentType: $2,
            }
        }
    ;

ComponentType
    : FieldType
        {
            $$ = &c.ComponentType{
                FieldType: $1,
            }
        }
    ;

ClassName
    : ClassName '/' IDENTIFIER
        {
            $$ = &c.ClassName{
                Identifier: append($1.Identifier, $3.Text),
            }
        }
    | IDENTIFIER
        {
            $$ = &c.ClassName{
                Identifier: []string{$1.Text},
            }
        }
    ;

%%
