# Nodes

## Table of content

- [Nodes inner workings](#nodes-inner-workings)
  - [What is a node](#what-is-a-node)
  - [Node types](#node-types)
- [Nodes documentation](#nodes-documentation)
  - [Expression nodes](#expression-nodes)
    - [Literal node](#literal-node)
    - [BinaryExpr node](#binaryexpr-node)
    - [UnaryExpr node](#unaryexpr-node)
    - [ParenExpr node](#parenexpr-node)
    - [ArrayLiteral node](#arrayliteral-node)
    - [MapLiteral node](#mapliteral-node)
    - [IndexableAccessExpr node](#indexableaccessexpr-node)
    - [FunctionCallExpr node](#functioncallexpr-node)
    - [MethodCallExpr node](#methodcallexpr-node)
  - [Statement nodes](#statement-nodes)
    - [TypeStmt node](#typestmt-node)
    - [VariableAssignStmt node](#variableassignstmt-node)
    - [IfStmt node](#ifstmt-node)
    - [ElseStmt node](#elsestmt-node)
    - [WhileStmt node](#whilestmt-node)
    - [ForStmt node](#forstmt-node)
    - [ImportStmt node](#importstmt-node)
    - [ReturnStmt node](#returnstmt-node)
    - [MurlocStmt node](#murlocstmt-node)
  - [Declaration nodes](#declaration-nodes)
    - [VariableDecl node](#variabledecl-node)
    - [FunctionDecl node](#functiondecl-node)

## Nodes inner workings

### What is a node
The `Node` interface is the building block of the Ecla AST.
It contains the start and end position methods and needs to be implemented by all the nodes.

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

***

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

***

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

***

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

***

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

***

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

***

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

***

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

***

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

***

#### MethodCallExpr node
The `MethodCallExpr` node represents a method call expression in the Ecla language.

##### Fields
The `MethodCallExpr` node is defined as follows :

```go
    type MethodCallExpr struct {
        MethodCallToken lexer.Token
        ObjectName      string
        FunctionCall    FunctionCallExpr
    }
```

The `MethodCallToken` field is the token that represents the method call.
The `ObjectName` field is the name of the object.
The `FunctionCall` field is the function call of the method call.

##### Code Example
a method call expression is an expression that contains an object name and a function call.

for example :
```ecla   
    console.print("hello world")
    math.pi()
```

***

### Statement nodes
This part of the documentation will cover all the statement nodes.
All the statement nodes implement the `Stmt` interface via the `stmtNode` method.

***

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

***

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

***

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

***

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

***

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

***

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
        print("hello world");
    }
```

***

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

***

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

***

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

***


### Declaration nodes
This part of the documentation will cover all the declaration nodes.

***

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

***

#### FunctionDecl node
The `FunctionDecl` node represents a function declaration in the Ecla language.

##### Fields
The `FunctionDecl` node is defined as follows :

```go
    type FunctionDecl struct {
        FunctionToken   lexer.Token
        Name            string
        LeftParamParen  lexer.Token
        RightParamParen lexer.Token
        Parameters      []FunctionParams
        LeftRetsParen   lexer.Token
        RightRetsParen  lexer.Token
        ReturnTypes     []string
        LeftBrace       lexer.Token
        RightBrace      lexer.Token
        Body            []Node
    }
```

The `FunctionToken` field is the token that represents the function declaration.
The `Name` field is the name of the function.
The `LeftParamParen` field is the left parenthesis of the function declaration.
The `RightParamParen` field is the right parenthesis of the function declaration.
The `Parameters` field is the parameters of the function declaration.
The `LeftRetsParen` field is the left parenthesis of the function declaration.
The `RightRetsParen` field is the right parenthesis of the function declaration.
The `ReturnTypes` field is the return types of the function declaration.
The `LeftBrace` field is the left brace of the function declaration.
The `RightBrace` field is the right brace of the function declaration.
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

***

