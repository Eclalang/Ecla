# Nodes

## Table of content

- [Nodes inner workings](#nodes-inner-workings)
  - [What is a node](#what-is-a-node)
  - [Node types](#node-types)
- [Nodes documentation](#nodes-documentation)
  - [Expression nodes](#expression-nodes)
    - [AnonymousFunctionCallExpr node](#anonymousfunctioncallexpr-node)
    - [AnonymousFunctionExpr node](#anonymousfunctionexpr-node)
    - [ArrayLiteral node](#arrayliteral-node)
    - [BinaryExpr node](#binaryexpr-node)
    - [FunctionCallExpr node](#functioncallexpr-node)
    - [IndexableAccessExpr node](#indexableaccessexpr-node)
    - [Literal node](#literal-node)
    - [MapLiteral node](#mapliteral-node)
    - [ParenExpr node](#parenexpr-node)
    - [SelectorExpr node](#selectorexpr-node)
    - [StructInstantiationExpr node](#structinstantiationexpr-node)
    - [UnaryExpr node](#unaryexpr-node)
  - [Statement nodes](#statement-nodes)
    - [BlockStmt node](#blockstmt-node)
    - [ElseStmt node](#elsestmt-node)
    - [ForStmt node](#forstmt-node)
    - [IfStmt node](#ifstmt-node)
    - [ImportStmt node](#importstmt-node)
    - [MurlocStmt node](#murlocstmt-node)
    - [ReturnStmt node](#returnstmt-node)
    - [TypeStmt node](#typestmt-node)
    - [VariableAssignStmt node](#variableassignstmt-node)
    - [WhileStmt node](#whilestmt-node)
  - [Declaration nodes](#declaration-nodes)
    - [FunctionDecl node](#functiondecl-node)
    - [StructDecl node](#structdecl-node)
    - [VariableDecl node](#variabledecl-node)

## Nodes inner workings

### What is a node

The `Node` interface is the building block of the Ecla AST.
It contains methods to get the node line and position that needs to be implemented by all the nodes.

### Node types

There exists 3 flavors of nodes :

- The Expression nodes.
- The Statement nodes.
- And the Declaration nodes.

Those 3 types of nodes implement the main `Node` interface.
All Nodes that represent `Expressions`, `Statements` or `Declarations` implement one of the 3 flavors of nodes.
for example, the `BinaryExpr` node implements the `Expr` interface via the `exprNode` method.

## Nodes documentation

### Expression nodes

This part of the documentation will cover all the expression nodes.
All the expression nodes implement the `Expr` interface via the `exprNode` method and the `precedence` method.
The `precedence` method is used to determine the precedence of the expression node in the AST.

---

#### AnonymousFunctionCallExpr node

The `AnonymousFunctionCallExpr` node represents an anonymous function call expression in the Ecla language.

##### Fields

The `AnonymousFunctionCallExpr` node is defined as follows :

```go
    type AnonymousFunctionCallExpr struct {
        AnonymousFunction AnonymousFunctionExpr
        LeftParen         lexer.Token
        RightParen        lexer.Token
        Args              []Expr
    }
```

The `AnonymousFunction` field is the anonymous function of the anonymous function call expression.
The `LeftParen` field is the left parenthesis of the anonymous function call expression.
The `RightParen` field is the right parenthesis of the anonymous function call expression.
The `Args` field is the arguments of the anonymous function call expression.

##### Code Example

an anonymous function call expression is an expression that contains an anonymous function and an array of arguments surrounded by parenthesis.

for example :

```ecla
    function (a : int, b : int) {
        console.println(a + b);
    }(1,1);
```

---

#### AnonymousFunctionExpr node

The `AnonymousFunctionExpr` node represents an anonymous function expression in the Ecla language.
It uses the `FunctionPrototype` struct to represent the prototype of the function.

```go
        type FunctionPrototype struct {
        LeftParamParen  lexer.Token
        RightParamParen lexer.Token
        Parameters      []FunctionParams
        LeftRetsParen   lexer.Token
        RightRetsParen  lexer.Token
        ReturnTypes     []string
        LeftBrace       lexer.Token
        RightBrace      lexer.Token
    }
```

##### Fields

The `AnonymousFunctionExpr` node is defined as follows :

```go
    type AnonymousFunctionExpr struct {
        FunctionToken lexer.Token
        Prototype     FunctionPrototype
        Body          []Node
    }
```

The `FunctionToken` field is the token that represents the anonymous function expression.
The `Prototype` field is the prototype of the anonymous function expression.
The `Body` field is the body of the anonymous function expression.

##### Code Example

an anonymous function expression is an expression that contains a prototype and a body.

for example :

```ecla
    function (a : int, b : int) {
        console.println(a + b);
    }
    var add = function (a : int, b : int) (int) {
        return a + b;
    }
```

---

#### ArrayLiteral node

The `ArrayLiteral` node represents an array literal in the Ecla language.

##### Fields

The `ArrayLiteral` node is defined as follows :

```go
    type ArrayLiteral struct {
        LBRACKET lexer.Token
        Values   []Expr
        RBRACKET lexer.Token
    }
```

The `LBRACKET` field is the left bracket of the array literal.
The `Values` field is the values of the array literal.
The `RBRACKET` field is the right bracket of the array literal.

##### Code Example

an array literal is an expression that contains an array of expressions surrounded by brackets.

for example :

```ecla
    [1, 2, 3]
    [true, false]
```

---

#### BinaryExpr node

The `BinaryExpr` node represents a binary expression in the Ecla language.

##### Fields

The `BinaryExpr` node is defined as follows :

```go
    type BinaryExpr struct {
        LeftExpr  Expr
        Operator  lexer.Token
        RightExpr Expr
    }
```

The `LeftExpr` field is the left expression of the binary expression.
The `Operator` field is the operator of the binary expression.
The `RightExpr` field is the right expression of the binary expression.

##### Code Example

a binary expression is an expression that contains 2 expressions and an operator.
for example :

```ecla
    1 + 1
    1 - 1
    1 * 1
    1 / 1
    1 // 1
    1 % 1
    1 == 1
    1 != 1
    1 > 1
    1 >= 1
    1 < 1
    1 <= 1
    1 && 1
    1 || 1
```

---

#### FunctionCallExpr node

The `FunctionCallExpr` node represents a function call expression in the Ecla language.

##### Fields

The `FunctionCallExpr` node is defined as follows :

```go
    type FunctionCallExpr struct {
        FunctionCallToken lexer.Token
        Name              string
        LeftParen         lexer.Token
        RightParen        lexer.Token
        Args              []Expr
    }
```

The `FunctionCallToken` field is the token that represents the function call.
The `Name` field is the name of the function.
The `LeftParen` field is the left parenthesis of the function call.
The `RightParen` field is the right parenthesis of the function call.
The `Args` field is the arguments of the function call.

##### Code Example

a function call expression is an expression that contains a function name and an array of arguments surrounded by parenthesis.

for example :

```ecla
    a()
    a(1)
    a(1, 2)
```

---

#### IndexableAccessExpr node

The `IndexableAccessExpr` node represents an indexable access expression in the Ecla language.

##### Fields

The `IndexableAccessExpr` node is defined as follows :

```go
    type IndexableAccessExpr struct {
        VariableToken lexer.Token
        VariableName  string
        Indexes       []Expr
        LastBracket   lexer.Token
    }
```

The `VariableToken` field is the token that represents the variable name.
The `VariableName` field is the name of the variable.
The `Indexes` field is the indexes of the indexable access expression.

##### Code Example

an indexable access expression is an expression that contains a variable name and an array of indexes surrounded by brackets.

for example :

```ecla
    a[1]
    a[1][2]
```

---

#### Literal node

The `Literal` node represents a literal value in the Ecla language.

##### Fields

The `Literal` node is defined as follows :

```go
    type Literal struct {
        Token lexer.Token
        Type  string
        Value string
    }
```

The `Token` field is the token that represents the start of the literal value.
The `Type` field is the type of the literal value.
The `Value` field is the value of the literal value.

##### Code Example

a literal value is a value that is not a variable or a function call.
for example :

```ecla
    1
    "hello world"
    true
```

---

#### MapLiteral node

The `MapLiteral` node represents a map literal in the Ecla language.

##### Fields

The `MapLiteral` node is defined as follows :

```go
    type MapLiteral struct {
        LBRACE lexer.Token
        Keys   []Expr
        Values []Expr
        RBRACE lexer.Token
    }
```

The `LBRACE` field is the left brace of the map literal.
The `Keys` field is the keys of the map literal.
The `Values` field is the values of the map literal.
The `RBRACE` field is the right brace of the map literal.

##### Code Example

a map literal is an expression that contains an array of keys and an array of values surrounded by braces.

for example :

```ecla
    {1:1, 2:2, 3:3}
    {true:true, false:false}
```

---

#### ParenExpr node

The `ParenExpr` node represents a parenthesized expression in the Ecla language.

##### Fields

The `ParenExpr` node is defined as follows :

```go
    type ParenExpr struct {
        Lparen     lexer.Token
        Expression Expr
        Rparen     lexer.Token
    }
```

The `Lparen` field is the left parenthesis of the parenthesized expression.
The `Expression` field is the expression inside the parenthesized expression.
The `Rparen` field is the right parenthesis of the parenthesized expression.

##### Code Example

a parenthesized expression is an expression that is surrounded by parenthesis.

for example :

```ecla
    (1 + 1)
    (true)
```

---

#### SelectorExpr node

The `SelectorExpr` node represents a selector expression in the Ecla language.

##### Fields

The `SelectorExpr` node is defined as follows :

```go
    type SelectorExpr struct {
        Field lexer.Token
        Expr  Expr
        Sel   Expr
    }
```

The `Field` field is the field of the selector expression.
The `Expr` field is the expression of the selector expression.
The `Sel` field is the selector of the selector expression.

##### Code Example

a selector expression is an expression that contains an expression and a selector separated by a dot.

for example :

```ecla
    a.b
    a.b.c
    console.println("hello world");
    a[1].b["hello"].c
    returnStruct().a
```

---

#### StructInstantiationExpr node

The `StructInstantiationExpr` node represents a struct instantiation expression in the Ecla language.

##### Fields

The `StructInstantiationExpr` node is defined as follows :

```go
    type StructInstantiationExpr struct {
        StructNameToken lexer.Token
        Name            string
        LeftBrace       lexer.Token
        RightBrace      lexer.Token
        Args            []Expr
    }
```

The `StructNameToken` field is the token that represents the struct instantiation expression.
The `Name` field is the name of the struct.
The `LeftBrace` field is the left brace of the struct instantiation expression.
The `RightBrace` field is the right brace of the struct instantiation expression.
The `Args` field is the arguments of the struct instantiation expression.

##### Code Example

a struct instantiation expression is an expression that contains a name and an array of arguments surrounded by braces.

for example :

```ecla
    Point{1, 2}
    Person{"John", 20}
    var p = Point{1, 2}
    var class = Class{"Math", Person{"John", 20}}
```

---

#### UnaryExpr node

The `UnaryExpr` node represents a unary expression in the Ecla language.

##### Fields

The `UnaryExpr` node is defined as follows :

```go
    type UnaryExpr struct {
        Operator lexer.Token
        RightExpr Expr
    }
```

The `Operator` field is the operator of the unary expression.
The `RightExpr` field is the right expression of the unary expression.

##### Code Example

a unary expression is an expression that contains an operator and an expression.

for example :

```ecla
    -1
    !true
```

---

### Statement nodes

This part of the documentation will cover all the statement nodes.
All the statement nodes implement the `Stmt` interface via the `stmtNode` method.

---

#### BlockStmt node

The `BlockStmt` node represents a block statement in the Ecla language.

##### Fields

The `BlockStmt` node is defined as follows :

```go
    type BlockScopeStmt struct {
        LeftBrace  lexer.Token
        RightBrace lexer.Token
        Body       []Node
    }
```

The `LeftBrace` field is the left brace of the block statement.
The `RightBrace` field is the right brace of the block statement.
The `Body` field is the body of the block statement.

##### Code Example

a block statement is a statement that contains a body surrounded by braces.

for example :

```ecla
    {
        console.prinln("hello world");
        var a = 1;
        var b = 2;
        var c = a + b;
        console.println(c);
    }
```

---

#### ElseStmt node

The `ElseStmt` node represents an else statement in the Ecla language.

##### Fields

The `ElseStmt` node is defined as follows :

```go
    type ElseStmt struct {
        ElseToken  lexer.Token
        LeftBrace  lexer.Token
        RightBrace lexer.Token
        Body       []Node
        IfStmt     *IfStmt
    }
```

The `ElseToken` field is the token that represents the else statement.
The `LeftBrace` field is the left brace of the else statement.
The `RightBrace` field is the right brace of the else statement.
The `Body` field is the body of the else statement.
The `IfStmt` field is the if statement of the else statement.

##### Code Example

an else statement is a statement that contains a body surrounded by braces.

for example :

```ecla
    if (true) {
        print("hello world");
    } else {
        print("hello world");
    }

    if (true) {
        print("hello world");
    } else if (false) {
        print("hello world");
    }else {
        print("hello world");
    }
```

---

#### ForStmt node

The `ForStmt` node represents a for statement in the Ecla language.

##### Fields

The `ForStmt` node is defined as follows :

```go
    type ForStmt struct {
        ForToken             lexer.Token
        LeftParen            lexer.Token
        RightParen           lexer.Token
        InitDecl             Decl
        CondExpr             Expr
        PostAssignStmt       Stmt
        KeyToken, ValueToken lexer.Token
        RangeToken           lexer.Token
        RangeExpr            Expr
        LeftBrace            lexer.Token
        RightBrace           lexer.Token
        Body                 []Node
    }
```

The `ForToken` field is the token that represents the for statement.
The `LeftParen` field is the left parenthesis of the for statement.
The `RightParen` field is the right parenthesis of the for statement.
The `InitDecl` field is the init declaration of the for statement.
The `CondExpr` field is the condition expression of the for statement.
The `PostAssignStmt` field is the post assign statement of the for statement.
The `KeyToken` field is the key token of the for statement.
The `ValueToken` field is the value token of the for statement.
The `RangeToken` field is the range token of the for statement.
The `RangeExpr` field is the range expression of the for statement.
The `LeftBrace` field is the left brace of the for statement.
The `RightBrace` field is the right brace of the for statement.
The `Body` field is the body of the for statement.

##### Code Example

a for statement is a statement that contains an init declaration, a condition expression, a post assign statement and a body surrounded by braces.

for example :

```ecla
    for (var i int = 0; i < 10; i += 1) {
        console.println("hello world");
    }
```

---

#### IfStmt node

The `IfStmt` node represents an if statement in the Ecla language.

##### Fields

The `IfStmt` node is defined as follows :

```go
    type IfStmt struct {
        IfToken    lexer.Token
        LeftParen  lexer.Token
        RightParen lexer.Token
        Cond       Expr
        LeftBrace  lexer.Token
        RightBrace lexer.Token
        Body       []Node
        ElseStmt   *ElseStmt
    }
```

The `IfToken` field is the token that represents the if statement.
The `LeftParen` field is the left parenthesis of the if statement.
The `RightParen` field is the right parenthesis of the if statement.
The `Cond` field is the condition of the if statement.
The `LeftBrace` field is the left brace of the if statement.
The `RightBrace` field is the right brace of the if statement.
The `Body` field is the body of the if statement.
The `ElseStmt` field is the else statement of the if statement.

##### Code Example

an if statement is a statement that contains a condition and a body surrounded by braces.

for example :

```ecla
    if (true) {
        print("hello world");
    }
```

---

#### ImportStmt node

The `ImportStmt` node represents an import statement in the Ecla language.

##### Fields

The `ImportStmt` node is defined as follows :

```go
    type ImportStmt struct {
        ImportToken lexer.Token
        ModulePath  string
    }
```

The `ImportToken` field is the token that represents the import statement.
The `ModulePath` field is the module path of the import statement.

##### Code Example

an import statement is a statement that contains a module path.

for example :

```ecla
    import "console"
```

---

#### MurlocStmt node

The `MurlocStmt` node represents a murloc statement in the Ecla language.

##### Fields

The `MurlocStmt` node is defined as follows :

```go
    type MurlocStmt struct {
        MurlocToken lexer.Token
    }
```

The `MurlocToken` field is the token that represents the murloc statement.

##### Code Example

a murloc statement is a statement that contains a murloc token.

for example :

```ecla
    mgrlgrl;
```

---

#### ReturnStmt node

The `ReturnStmt` node represents a return statement in the Ecla language.

##### Fields

The `ReturnStmt` node is defined as follows :

```go
    type ReturnStmt struct {
        ReturnToken  lexer.Token
        ReturnValues []Expr
    }
```

The `ReturnToken` field is the token that represents the return statement.
The `ReturnValues` field is the return values of the return statement.

##### Code Example

a return statement is a statement that contains an array of return values.

for example :

```ecla
    return 1;
    return 1, 2;
```

---

#### TypeStmt node

The `TypeStmt` node represents a type statement in the Ecla language.

##### Fields

The `TypeStmt` node is defined as follows :

```go
    type TypeStmt struct {
        TypeToken  lexer.Token
        Lparen     lexer.Token
        Rparen     lexer.Token
        Expression Expr
    }
```

The `TypeToken` field is the token that represents the type statement.
The `Lparen` field is the left parenthesis of the type statement.
The `Rparen` field is the right parenthesis of the type statement.
The `Expression` field is the expression of the type statement.

##### Code Example

a type statement is a statement that contains an expression surrounded by parenthesis.

for example :

```ecla
    type(1);
    type("hello world");
```

---

#### VariableAssignStmt node

The `VariableAssignStmt` node represents a variable assign statement in the Ecla language.

##### Fields

The `VariableAssignStmt` node is defined as follows :

```go
    type VariableAssignStmt struct {
        VarToken lexer.Token
        Names    []Expr
        Operator string
        Values   []Expr
    }
```

The `VarToken` field is the token that represents the variable assign statement.
The `Names` field is the names of the variable assign statement.
The `Operator` field is the operator of the variable assign statement.
The `Values` field is the values of the variable assign statement.

##### Code Example

a variable assign statement is a statement that contains an array of names, an operator and an array of values.

for example :

```ecla
    a = 1;
    a += 1;
```

---

#### WhileStmt node

The `WhileStmt` node represents a while statement in the Ecla language.

##### Fields

The `WhileStmt` node is defined as follows :

```go
    type WhileStmt struct {
        WhileToken lexer.Token
        LeftParen  lexer.Token
        RightParen lexer.Token
        Cond       Expr
        LeftBrace  lexer.Token
        RightBrace lexer.Token
        Body       []Node
    }
```

The `WhileToken` field is the token that represents the while statement.
The `LeftParen` field is the left parenthesis of the while statement.
The `RightParen` field is the right parenthesis of the while statement.
The `Cond` field is the condition of the while statement.
The `LeftBrace` field is the left brace of the while statement.
The `RightBrace` field is the right brace of the while statement.
The `Body` field is the body of the while statement.

##### Code Example

a while statement is a statement that contains a condition and a body surrounded by braces.

for example :

```ecla
    while (true) {
        print("hello world");
    }
```

---

### Declaration nodes

This part of the documentation will cover all the declaration nodes.

---

#### FunctionDecl node

The `FunctionDecl` node represents a function declaration in the Ecla language.
It uses the `FunctionPrototype` struct to represent the prototype of the function.

```go
        type FunctionPrototype struct {
        LeftParamParen  lexer.Token
        RightParamParen lexer.Token
        Parameters      []FunctionParams
        LeftRetsParen   lexer.Token
        RightRetsParen  lexer.Token
        ReturnTypes     []string
        LeftBrace       lexer.Token
        RightBrace      lexer.Token
    }
```

##### Fields

The `FunctionDecl` node is defined as follows :

```go
        type FunctionDecl struct {
        FunctionToken lexer.Token
        Name          string
        Prototype     FunctionPrototype
        Body          []Node
    }
```

The `FunctionToken` field is the token that represents the function declaration.
The `Name` field is the name of the function.
The `Prototype` field is the prototype of the function declaration.
The `Body` field is the body of the function declaration.

##### Code Example

a function declaration is a declaration that contains a name, an array of parameters, an array of return types and a body surrounded by braces.

for example :

```ecla
    function add(a int, b int) (int) {
        return a + b;
    }

    function concat(a string, b string) (string) {
        return a + b;
    }

    function doNothing() {
    }
```

---

#### StructDecl node

The `StructDecl` node represents a struct declaration in the Ecla language.
it uses the `StructField` struct to represent the fields of the struct.

```go
    type StructField struct {
        Name string
        Type string
    }
```

##### Fields

The `StructDecl` node is defined as follows :

```go
    type StructDecl struct {
        StructToken lexer.Token
        Name        string
        LeftBrace   lexer.Token
        Fields      []StructField
        RightBrace  lexer.Token
    }
```

The `StructToken` field is the token that represents the struct declaration.
The `Name` field is the name of the struct.
The `LeftBrace` field is the left brace of the struct declaration.
The `Fields` field is the fields of the struct declaration.
The `RightBrace` field is the right brace of the struct declaration.

##### Code Example

a struct declaration is a declaration that contains a name and a succession of fields surrounded by braces.

for example :

```ecla
    struct Point {
        x int
        y int
    }

    struct Person {
        name string
        age int
    }
```

---

#### VariableDecl node

The `VariableDecl` node represents a variable declaration in the Ecla language.

##### Fields

The `VariableDecl` node is defined as follows :

```go
    type VariableDecl struct {
        VarToken lexer.Token
        Name     string
        Type     string
        Value    Expr
    }
```

The `VarToken` field is the token that represents the variable declaration.
The `Name` field is the name of the variable.
The `Type` field is the type of the variable.
The `Value` field is the value of the variable.

##### Code Example

a variable declaration is a declaration that contains a name, a type and a value.

for example :

```ecla
    var a int;
    var a int = 1;
    a := 1;
```

---
