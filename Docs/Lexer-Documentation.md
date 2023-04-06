# Ecla Lexer #

## Table of content ##

- Technical specifications
    - context
    - Inner structure
        - Token
        - identifier
- Inner workings
    - Lexer process
        - composite and situation token
        - both direction lexing
        - behavior
    - Performance
        - Time complexity
        - Space complexity
        - Memory usage
        - Performance tests
- Grammar
    - Table of token

## Technical specification ##

### Context ###

In computer science, lexical analysis, lexing or tokenization is the process of converting a
sequence of characters (such as in a computer program or web page) into a sequence of lexical tokens (strings with an
assigned and thus identified meaning).

In our context, Ecla Lexer use the input string to create a list of token usable for the Ecla
Parser. the input string is the actual source code of the Ecla file that we want to execute as a string.

### Inner Structure ###

for the tokenization of the Ecla lexer, we need to declare our token, and identifier.

* #### Token ####

```go
type Token struct {
    TokenType string
    Value     string
    Position  int
    Line      int
}
```

- Token
    - TokenType refers to one of our token, listed in the Table of token describe below in this doc.
    - Value contain all the string considered as the specified Token.
    - Position contain the position in the line were our token begin.
    - Line contain the line number of our token

* #### Identifier ####

an identifier is here to link a token type with his known syntax. You can find the known syntax of each token in the
Table of token.

the Identifier variable declared as below is used everywhere to compare each substring of our input string with all
the known syntax.

```go
type identifier struct {
    Identifier string
    Syntaxe    []string
}
var Identifier []identifier = []identifier{
...
}
```

- identifier
    - Identifier refers to the tokenType.
    - Syntax refers to all the known syntax.

## Inner workings ##

### Lexer process ###

* #### Composite and Involved token ####

We have defined 2 attributes for each token. they may be **_Composite_** and/or **_Involved_**.

- A **_Composite_** token is a token that may need a specific token's sequence to exist, and is the result of this 
    sequence by modifying the first token of the composite operation rather than creating a new token thanks to an 
    identifier. Thereby it may not have an identifier.
    - example: `+=` => `token ADD + token ASSIGN = token ADDASSIGN`
        ADDASSIGN is so a Composite token here. we don't know the actual ``+=`` syntax, but we know it by composition of
        the``+`` which is a syntax for ADD, and ``=`` which is a syntax for EQUAL.
- An **_Involved_** token is a token that can induce a token just after. An Involved token can induce more than one other
    token.
  - example: for the same example as the previous one for Composite, `ADD` is Involved to the `ADDASSAGN`. 

if a token is modified after its instantiation, it means that the token can induce the new  tokenType, and that the new 
TokenType is a composite token of the previous tokenType.

we are using those 2 attributes because its easier to understand the behavior of our lexer with that in mind.

Sometimes, a token can be Composite and Situational, like the INT token. in this case, the token have a specific
identifier to initiate the behavior of instantiation and compose with itself for a limited duration.

* #### both direction lexing ####

Our lexer work with a 1 character increment each iteration on the input string. we compare all the substring that we
didn't assign to a token with all the know syntax. but with this simple behavior, a basic situation can occur and brake
everything:

input string : `Murloc()`

if we use this input string for references, and do the Lexer work step by step, we can see the problem:

- `M` is not a know syntax
- `Mu` is not a know syntax
- `Mur` is not a know syntax
- `Murl` is not a know syntax
- `Murlo` is not a know syntax
- `Murloc` is not a know syntax
- `Murloc(` is not a know syntax

For this last step, we see the mistake: even if a known syntax occur later in the substring , nothing happen, that's why
for each iteration, we check if our substring backward is a known syntax to avoid this mistake.

* #### Behavior ####

Like said in the last section, Our lexer work with a 1 character increment each iteration on the input string.

We define each iteration the substring that we will be compared to all the known syntax, to identify first the 
token that should be placed with just the substring information. We browse our `Identifier` array, and for each 
`identifier` we compare the substring with the known syntax of this `identifier`.

If this browse of all the know syntax doesn't find an associate token, we check backward like previously explained with
the same behavior of the onward check.

when we find a known syntax, we also find the related token. with that in mind, we can now check the context around 
our new token. The first thing we check is if a COMMENT or COMMENTGROUP is being composed, and if so, we compose our 
current new token with the previous one. if not, we then check if a STRING is being composed, and if so, we compose our
current new token with the previous one just like before. We know if a COMMENT or COMMENTGROUP is being composes by using
some boolean, and some index to locate the last COMMENT created. for the STRING, it's just a boolean that switch on and 
off at each DQUOTE token, nothing more. after those 2 check , we can work on some composite token:

each involved token have his little function to create his composite one.

and if after all of this no composite token are created, it means that the token is not involved or include in any 
composite token, so we can just create the new token as it is.

At the very end, we must place an EOF token, so we created, after our loop, a new EOF token.

You should have seen that normally, our token contain 4 value: his type, his related substrings, his position, and his
line. we talked a lot of the first 2, but not a lot with the last 2:Each time a token is created, we use this function 
to detect the position and the line of the new token. Because the line correspond with the IN FILE line, we must check 
the `\n` character to know the line, and position.

### Performance ###

* #### Time complexity ####

* #### Space complexity ####

* #### Memory usage ####

* #### Performance tests ####

## Grammar ##

### Table of Token ###

| token name     | Context                                                                                                                                                                                        | identifier                              | Composite                                    | Involved                 |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|----------------------------------------------|--------------------------|
| `TEXT`         | The default token                                                                                                                                                                              |                                         |                                              |                          |
| `STRING`       | Everything place after a `DQUOTE` token.<br/>The STRING token's completion end when a new `DQUOTE` token appear.                                                                               |                                         | `STRING`+`. . .` until `DQUOTE`              | `STRING`                 |
| `PRINT`        | **_Deprecated_** <br/>the "print" element.                                                                                                                                                     | `print`                                 |                                              |                          |
| `INT`          | A series of numeric characters that compose an integer.                                                                                                                                        | `0` `1` `2` `3` `4` `5` `6` `7` `8` `9` | `INT`+`INT`                                  | `INT` `FLOAT`            |
| `FLOAT`        | A series of numeric characters that compose a float.                                                                                                                                           |                                         | `INT`+`PERIOD` <br/> `FLOAT`+`INT`           | `FLOAT`                  |
| `ADD`          | Refers to the addition operator.                                                                                                                                                               | `+`                                     |                                              | `ADDASSIGN`              |
| `SUB`          | Refers to the subtraction operator.                                                                                                                                                            | `-`                                     |                                              | `SUBASSIGN`              |
| `MULT`         | Refers to the multiplication operator.                                                                                                                                                         | `*`                                     |                                              | `MULTASSIGN`             |
| `DIV`          | Refers to the division operator.                                                                                                                                                               | `/`                                     |                                              | `DIVASSIGN` `QOT`        |
| `MOD`          | Refers to modulo operator.                                                                                                                                                                     | `%`                                     |                                              | `MODASSIGN`              |
| `QOT`          | Refers to mathematical quotient operator.                                                                                                                                                      |                                         | `DIV`+`DIV`                                  |                          |
| `INC`          | Refers to the increment operator.                                                                                                                                                              |                                         | `ADD`+`ADD`                                  |                          |
| `DEC`          | Refers to the decrement operator.                                                                                                                                                              |                                         | `SUB`+`SUB`                                  |                          |
| `NOT`          | Refers to the "not" statement.                                                                                                                                                                 | `!`                                     |                                              | `NOTASSIGN`              |
| `ASSIGN`       | Refers to the assigment.                                                                                                                                                                       | `=`                                     |                                              | `EQUAL`                  |
| `ADDASSIGN`    | Refers to the additive assigment.                                                                                                                                                              |                                         | `ADD`+`ASSIGN`                               |                          |
| `SUBASSIGN`    | Refers to the subtractive assigment.                                                                                                                                                           |                                         | `SUB`+`ASSIGN`                               |                          |
| `MULTASSIGN`   | Refers to the multiplicative assigment.                                                                                                                                                        |                                         | `MULT`+`ASSIGN`                              |                          |
| `DIVASSIGN`    | Refers to the divide assigment.                                                                                                                                                                |                                         | `DIV`+`ASSIGN`                               |                          |
| `LSS`          | Refers to the "lesser than" statement.                                                                                                                                                         | `<`                                     |                                              | `LEQ`                    |
| `GTR`          | Refers to the "greater than" statement.                                                                                                                                                        | `>`                                     |                                              | `GEQ`                    |
| `NEQ`          | Refers to the "not equal" statement.                                                                                                                                                           |                                         | `NOT`+`ASSIGN`                               |                          |
| `EQUAL`        | Refers to the "is equal" statement.                                                                                                                                                            |                                         | `ASSIGN`+`ASSIGN`                            |                          |
| `LEQ`          | Refers to the "lesser or equal than" statement.                                                                                                                                                |                                         | `LSS`+`ASSIGN`                               |                          |
| `GEQ`          | Refers to the "greater or equal than" statement.                                                                                                                                               |                                         | `GTR`+`ASSIGN`                               |                          |
| `XOR`          | Refers to the "exclusive or" statement.                                                                                                                                                        | `^`                                     |                                              |                          |
| `OR`           | Refers to the "or" statement.                                                                                                                                                                  | `&#124;&#124;`                          |                                              |                          |
| `AND`          | Refers to the "and" statement.                                                                                                                                                                 | `&&`                                    |                                              |                          |
| `LPARENT`      | Refers to the left parenthesis.                                                                                                                                                                | `(`                                     |                                              |                          |
| `RPARENT`      | Refers to the right parenthesis.                                                                                                                                                               | `)`                                     |                                              |                          |
| `EOL`          | Refers to the end of line character.                                                                                                                                                           | `;`                                     |                                              |                          |
| `DQUOTE`       | Refers to the double quote character.                                                                                                                                                          | `"`                                     |                                              | `STRING` without compose |
| `PERIOD`       | Refers to the period character.                                                                                                                                                                | `.`                                     |                                              |                          |
| `COLON`        | Refers to the colon character.                                                                                                                                                                 | `:`                                     |                                              |                          |
| `LBRACE`       | Refers to the left brace character.                                                                                                                                                            | `{`                                     |                                              |                          |
| `RBRACE`       | Refers to the right brace character.                                                                                                                                                           | `}`                                     |                                              |                          |
| `LBRACKET`     | Refers to the left bracket character.                                                                                                                                                          | `[`                                     |                                              |                          |
| `RBRACKET`     | Refers to the right bracket character.                                                                                                                                                         | `]`                                     |                                              |                          |
| `COMMA`        | Refers to the comma character.                                                                                                                                                                 | `,`                                     |                                              |                          |
| `BOOL`         | Refers to a "true" or "false" statement.                                                                                                                                                       | `true` `false`                          |                                              |                          |
| `EOF`          | Always empty. Mark the end of the file. Always appear at the end of the token list.                                                                                                            |                                         |                                              |                          |
| `COMMENT`      | Refers to everything place after the specify identifier with the specify identifier include.<br/>The `COMMENT` token's completion end when a break line appear.<br/>Can be composite of itself | `#`                                     | `COMMENT`+`. . .` until `EOL`                | `COMMENT`                |
| `COMMENTGROUP` | Refers to everything place after the first specify identifier with identifier include.<br/>The `COMMENTGROUP` token's completion end when the second specify identifier appear.                | `#/` `/#`                               | `COMMENTGROUP`+`. . . ` until `COMMENTGROUP` | `COMMENTGROUP`           |
| `MURLOC`       | Mrgle, Mmmm Uuua !                                                                                                                                                                             | `mgrlgrl`                               |                                              |                          |