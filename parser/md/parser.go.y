%{
package md

import (
    c "github.com/kamichidu/go-jclass/parser/common"
)
%}

%type <methodDescriptor> MethodDescriptor
%type <parameterDescriptors> ParameterDescriptors
%type <parameterDescriptor> ParameterDescriptor
%type <returnDescriptor> ReturnDescriptor
%type <fieldType> FieldType
%type <baseType> BaseType
%type <objectType> ObjectType
%type <arrayType> ArrayType
%type <componentType> ComponentType
%type <className> ClassName

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '[' '(' ')'
%token <token> IDENTIFIER

%union {
    methodDescriptor     *c.MethodDescriptor
    parameterDescriptors []*c.ParameterDescriptor
    parameterDescriptor  *c.ParameterDescriptor
    returnDescriptor     *c.ReturnDescriptor
    fieldType            *c.FieldType
    baseType             *c.BaseType
    objectType           *c.ObjectType
    arrayType            *c.ArrayType
    componentType        *c.ComponentType
    className            *c.ClassName
    token                *c.Token
}

%%

MethodDescriptor
    : '(' ParameterDescriptors ')' ReturnDescriptor
        {
            $$ = &c.MethodDescriptor{
                ParameterDescriptor: $2,
                ReturnDescriptor: $4,
            }
            if l, ok := mdlex.(*MDLexer); ok {
                l.Result = $$
            }
        }
    ;

ParameterDescriptors
    :
        {
            $$ = make([]*c.ParameterDescriptor, 0)
        }
    | ParameterDescriptors ParameterDescriptor
        {
            $$ = append($1, $2)
        }
    ;

ParameterDescriptor
    : FieldType
        {
            $$ = &c.ParameterDescriptor{
                FieldType: $1,
            }
        }
    ;

ReturnDescriptor
    : FieldType
        {
            $$ = &c.ReturnDescriptor{
                FieldType: $1,
            }
        }
    | 'V'
        {
            $$ = &c.ReturnDescriptor{
                VoidDescriptor: &c.VoidDescriptor{"void"},
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
