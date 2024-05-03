/*
 BSD License
 
 Copyright (c) 2013, Tom Everett All rights reserved.
 
 Redistribution and use in source and binary forms, with or without modification, are permitted
 provided that the following conditions are met:
 
 1. Redistributions of source code must retain the above copyright notice, this list of conditions
 and the following disclaimer. 2. Redistributions in binary form must reproduce the above copyright
 notice, this list of conditions and the following disclaimer in the documentation and/or other
 materials provided with the distribution. 3. Neither the name of Tom Everett nor the names of its
 contributors may be used to endorse or promote products derived from this software without specific
 prior written permission.
 
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
 FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
 CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
 IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
/*
 Adapted from pascal.g by Hakki Dogusan, Piet Schoutteten and Marton Papp
 */

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

grammar Pascal;

options {
    caseInsensitive = true;
}

program
    : PROGRAM identifier SEMI block DOT EOF
    ;

identifier
    : IDENT
    ;

block
    : (
        constantDefinitionPart
        | typeDefinitionPart
        | variableDeclarationPart
        | procedureAndFunctionDeclarationPart
    )* compoundStatement
    ;

constantDefinitionPart
    : CONST (constantDefinition SEMI)+
    ;

constantDefinition
    : identifier EQUAL constant
    ;

constantChr
    : CHR LPAREN unsignedInteger RPAREN
    ;

constant
    : unsignedNumber
    | sign unsignedNumber
    | identifier
    | sign identifier
    | string
    | constantChr
    ;

unsignedNumber
    : unsignedInteger
    | unsignedReal
    ;

unsignedInteger
    : NUM_INT
    ;

unsignedReal
    : NUM_REAL
    ;

sign
    : PLUS
    | MINUS
    ;

bool_
    : TRUE
    | FALSE
    ;

string
    : STRING_LITERAL {$STRING_LITERAL.SetText($STRING_LITERAL.GetText()[1:len($STRING_LITERAL.GetText())-1])}
    ;

typeDefinitionPart
    : TYPE (typeDefinition SEMI)+
    ;

typeDefinition
    : identifier EQUAL (type_ | functionType | procedureType)
    ;

functionType
    : FUNCTION (formalParameterList)? COLON resultType
    ;

procedureType
    : PROCEDURE (formalParameterList)?
    ;

type_
    : simpleType
    | structuredType
    | pointerType
    ;

simpleType
    : scalarType
    | subrangeType
    | typeIdentifier
    | stringtype
    ;

scalarType
    : LPAREN identifierList RPAREN
    ;

subrangeType
    : constant DOTDOT constant
    ;

typeIdentifier
    : identifier
    | (CHAR | BOOLEAN | INTEGER | REAL | STRING)
    ;

structuredType
    : arrayType
    | recordType
    | setType
    | fileType
    ;

stringtype
    : STRING LBRACK (identifier | unsignedNumber) RBRACK
    ;

arrayType
    : ARRAY LBRACK typeList RBRACK OF componentType
    | ARRAY LBRACK2 typeList RBRACK2 OF componentType
    ;

typeList
    : indexType (COMMA indexType)*
    ;

indexType
    : simpleType
    ;

componentType
    : type_
    ;

recordType
    : RECORD fieldList? END
    ;

fieldList
    : fixedPart (SEMI variantPart)?
    | variantPart
    ;

fixedPart
    : recordSection (SEMI recordSection)*
    ;

recordSection
    : identifierList COLON type_
    ;

variantPart
    : CASE tag OF variant (SEMI variant)*
    ;

tag
    : identifier COLON typeIdentifier
    | typeIdentifier
    ;

variant
    : constList COLON LPAREN fieldList RPAREN
    ;

setType
    : SET OF baseType
    ;

baseType
    : simpleType
    ;

fileType
    : FILE OF type_
    | FILE
    ;

pointerType
    : POINTER typeIdentifier
    ;

variableDeclarationPart
    : VAR variableDeclaration (SEMI variableDeclaration)* SEMI
    ;

variableDeclaration
    : varNames = identifierList COLON pascalType = type_
    ;

procedureAndFunctionDeclarationPart
    : procedureOrFunctionDeclaration SEMI
    ;

procedureOrFunctionDeclaration
    : procedureDeclaration
    | functionDeclaration
    ;

procedureDeclaration
    : PROCEDURE name = identifier (paramList = formalParameterList)? SEMI block
    ;

formalParameterList
    : LPAREN params += formalParameterSection (SEMI params += formalParameterSection)* RPAREN
    ;

formalParameterSection
    : paramNames = identifierList COLON paramType = typeIdentifier
    ;

identifierList
    : ids += identifier (COMMA ids += identifier)*
    ;

constList
    : constant (COMMA constant)*
    ;

functionDeclaration
    : FUNCTION name = identifier (paramList = formalParameterList)? COLON returnType = resultType SEMI block
    ;

resultType
    : typeIdentifier
    ;

statement
    : simpleStatement
    | structuredStatement
    ;

simpleStatement
    : assignmentStatement
    | procedureStatement
    | emptyStatement_
    ;

assignmentStatement
    : varName = variable ASSIGN expression
    ;

variable
    : (AT identifier | identifier) (
        LBRACK expression (COMMA expression)* RBRACK
        | LBRACK2 expression (COMMA expression)* RBRACK2
        | DOT identifier
        | POINTER
    )*
    ;

expression
    : op = NOT expression                           # NotOp
    | expression op = (AND | OR) expression         # BoolOp
    | expression op = relationaloperator expression # RelOp
    | expression op = (STAR | SLASH) expression     # MulDivOp
    | expression op = addsuboperator expression     # AddSubOp
    | signedFactor                                  # ExpSignedFactor
    ;

relationaloperator
    : EQUAL
    | NOT_EQUAL
    | LT
    | LE
    | GE
    | GT
    | IN
    ;

addsuboperator
    : PLUS
    | MINUS
    | OR
    ;

term
    : signedFactor (multiplicativeoperator term)?
    ;

multiplicativeoperator
    : STAR
    | SLASH
    | DIV
    | MOD
    | AND
    ;

signedFactor
    : (PLUS | MINUS)? factor
    ;

factor
    : id = variable            # factorVariable
    | LPAREN expression RPAREN # factorExpression
    | functionDesignator       # factorFunctionDesignator
    | unsignedConstant         # factorUnsignedConstant
    | set_                     # factorSet
    // | NOT factor
    | bool_ # factorBool
    ;

unsignedConstant
    : unsignedNumber
    | constantChr
    | string
    | NIL
    ;

functionDesignator
    : functionID = identifier LPAREN (parameterList)? RPAREN
    ;

parameterList
    : actualParameter (COMMA actualParameter)*
    ;

set_
    : LBRACK elementList RBRACK
    | LBRACK2 elementList RBRACK2
    ;

elementList
    : element (COMMA element)*
    |
    ;

element
    : expression (DOTDOT expression)?
    ;

procedureStatement
    : procedureID = identifier (LPAREN parameterList RPAREN)?
    ;

actualParameter
    : expression parameterwidth*
    ;

parameterwidth
    : COLON expression
    ;

emptyStatement_
    :
    ;

empty_
    :
    /* empty */
    ;

structuredStatement
    : compoundStatement
    | conditionalStatement
    | repetetiveStatement
    | withStatement
    ;

compoundStatement
    : BEGIN statements END
    ;

statements
    : statement (SEMI statement)*
    ;

conditionalStatement
    : ifStatement
    | caseStatement
    ;

ifStatement
    : IF expression thenStatement (: elseStatement)?
    ;

thenStatement
    : THEN statement
    ;

elseStatement
    : ELSE statement
    ;

caseStatement
    : CASE expression OF caseListElement (SEMI caseListElement)* (SEMI ELSE statements)? END
    ;

caseListElement
    : constList COLON statement
    ;

repetetiveStatement
    : whileStatement
    | repeatStatement
    | forStatement
    ;

whileStatement
    : WHILE expression whileBlock
    ;

whileBlock
    : DO statement
    ;

repeatStatement
    : REPEAT statements UNTIL expression
    ;

forStatement
    : FOR forInit forUntil DO statement
    ;

forInit
    : varName = variable ASSIGN expression
    ;

forUntil
    : step = (TO | DOWNTO) expression
    ;

withStatement
    : WITH recordVariableList DO statement
    ;

recordVariableList
    : variable (COMMA variable)*
    ;

AND
    : 'AND'
    ;

ARRAY
    : 'ARRAY'
    ;

BEGIN
    : 'BEGIN'
    ;

BOOLEAN
    : 'BOOLEAN'
    ;

CASE
    : 'CASE'
    ;

CHAR
    : 'CHAR'
    ;

CHR
    : 'CHR'
    ;

CONST
    : 'CONST'
    ;

DIV
    : 'DIV'
    ;

DO
    : 'DO'
    ;

DOWNTO
    : 'DOWNTO'
    ;

ELSE
    : 'ELSE'
    ;

END
    : 'END'
    ;

FILE
    : 'FILE'
    ;

FOR
    : 'FOR'
    ;

FUNCTION
    : 'FUNCTION'
    ;

IF
    : 'IF'
    ;

IN
    : 'IN'
    ;

INTEGER
    : 'INTEGER'
    ;

MOD
    : 'MOD'
    ;

NIL
    : 'NIL'
    ;

NOT
    : 'NOT'
    ;

OF
    : 'OF'
    ;

OR
    : 'OR'
    ;

PROCEDURE
    : 'PROCEDURE'
    ;

PROGRAM
    : 'PROGRAM'
    ;

REAL
    : 'REAL'
    ;

RECORD
    : 'RECORD'
    ;

REPEAT
    : 'REPEAT'
    ;

SET
    : 'SET'
    ;

THEN
    : 'THEN'
    ;

TO
    : 'TO'
    ;

TYPE
    : 'TYPE'
    ;

UNTIL
    : 'UNTIL'
    ;

VAR
    : 'VAR'
    ;

WHILE
    : 'WHILE'
    ;

WITH
    : 'WITH'
    ;

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

STAR
    : '*'
    ;

SLASH
    : '/'
    ;

ASSIGN
    : ':='
    ;

COMMA
    : ','
    ;

SEMI
    : ';'
    ;

COLON
    : ':'
    ;

EQUAL
    : '='
    ;

NOT_EQUAL
    : '<>'
    ;

LT
    : '<'
    ;

LE
    : '<='
    ;

GE
    : '>='
    ;

GT
    : '>'
    ;

LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

LBRACK
    : '['
    ;

LBRACK2
    : '(.'
    ;

RBRACK
    : ']'
    ;

RBRACK2
    : '.)'
    ;

POINTER
    : '^'
    ;

AT
    : '@'
    ;

DOT
    : '.'
    ;

DOTDOT
    : '..'
    ;

LCURLY
    : '{'
    ;

RCURLY
    : '}'
    ;

STRING
    : 'STRING'
    ;

TRUE
    : 'TRUE'
    ;

FALSE
    : 'FALSE'
    ;

WS
    : [ \t\r\n] -> skip
    ;

COMMENT_1
    : '(*' .*? '*)' -> skip
    ;

COMMENT_2
    : '{' .*? '}' -> skip
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

IDENT
    : ('A' .. 'Z') ('A' .. 'Z' | '0' .. '9' | '_')*
    ;

STRING_LITERAL
    : '\'' ('\'\'' | ~ ('\''))* '\''
    ;

NUM_INT
    : ('0' .. '9')+
    ;

NUM_REAL
    : ('0' .. '9')+ (('.' ('0' .. '9')+ (EXPONENT)?)? | EXPONENT)
    ;

fragment EXPONENT
    : ('E') ('+' | '-')? ('0' .. '9')+
    ;