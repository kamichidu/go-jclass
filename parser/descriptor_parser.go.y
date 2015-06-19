%{
package parser

import (
)

type Token struct {
    Id int
    Text string
    Pos int
}
%}

%type <typeName> FieldDescriptor
%type <typeName> FieldType
%type <typeName> BaseType
%type <typeName> ObjectType
%type <typeName> ArrayType
%type <typeName> ComponentType

%token 'B' 'C' 'D' 'F' 'I' 'J' 'S' 'Z' 'L' ';' '['
%token <token> CLASS_NAME

%union {
    typeName string
    token Token
}

%%

FieldDescriptor
    : FieldType
    ;

FieldType
    : BaseType
    | ObjectType
    | ArrayType
    ;

BaseType
    : 'B'
        {
            $$ = "byte"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'C'
        {
            $$ = "char"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'D'
        {
            $$ = "double"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'F'
        {
            $$ = "float"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'I'
        {
            $$ = "int"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'J'
        {
            $$ = "long"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'S'
        {
            $$ = "short"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    | 'Z'
        {
            $$ = "boolean"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    ;

ObjectType
    : 'L' CLASS_NAME ';'
        {
            $$ += "L" + $2.Text + ";"
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    ;

ArrayType
    : '[' ComponentType
        {
            $$ += "[" + $2
            if lexer, ok := yylex.(*DescriptorLexer); ok {
                lexer.Result += $$
            }
        }
    ;

ComponentType
    : FieldType
    ;

%%
