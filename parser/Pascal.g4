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
    : (variableDeclarationPart | procedureAndFunctionDeclarationPart)* compoundStatement
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

bool_
    : TRUE
    | FALSE
    ;

string
    : STRING_LITERAL {$STRING_LITERAL.SetText($STRING_LITERAL.GetText()[1:len($STRING_LITERAL.GetText())-1])}
    ;

typeIdentifier
    : identifier
    | (BOOLEAN | INTEGER | REAL | STRING)
    ;

variableDeclarationPart
    : VAR variableDeclaration (SEMI variableDeclaration)* SEMI
    ;

variableDeclaration
    : varNames = identifierList COLON pascalType = typeIdentifier
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

functionDeclaration
    : FUNCTION name = identifier (paramList = formalParameterList)? COLON returnType = typeIdentifier SEMI block
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
    : identifier (LBRACK expression (COMMA expression)* RBRACK | DOT identifier)*
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
    // | NOT factor
    | bool_ # factorBool
    ;

unsignedConstant
    : unsignedNumber
    | string
    ;

functionDesignator
    : functionID = identifier LPAREN (parameterList)? RPAREN
    ;

parameterList
    : actualParameter (COMMA actualParameter)*
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

structuredStatement
    : compoundStatement
    | conditionalStatement
    | repetetiveStatement
    ;

compoundStatement
    : BEGIN statements END
    ;

statements
    : statement (SEMI statement)*
    ;

conditionalStatement
    : ifStatement
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

AND
    : 'AND'
    ;

BEGIN
    : 'BEGIN'
    ;

BOOLEAN
    : 'BOOLEAN'
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

FOR
    : 'FOR'
    ;

FUNCTION
    : 'FUNCTION'
    ;

IF
    : 'IF'
    ;

INTEGER
    : 'INTEGER'
    ;

MOD
    : 'MOD'
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

REPEAT
    : 'REPEAT'
    ;

THEN
    : 'THEN'
    ;

TO
    : 'TO'
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