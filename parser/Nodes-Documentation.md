# Nodes

## Table of content

- Nodes inner workings
  - What is a node
  - Node types
- Nodes documentation
  - Expression nodes
    - Literal node
    - BinaryExpr node
    - UnaryExpr node
    - ParenExpr node
    - ArrayLiteral node
    - MapLiteral node
    - IndexableAccessExpr node
    - FunctionCallExpr node
    - MethodCallExpr node
  - Statement nodes
    - PrintStmt node
    - TypeStmt node
    - VariableAssignStmt node
    - IfStmt node
    - ElseStmt node
    - WhileStmt node
    - ForStmt node
    - ImportStmt node
    - ReturnStmt node
  - Declaration nodes
    - VariableDecl node
    - FunctionDecl node

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

#### PrintStmt node
The `PrintStmt` node represents a print statement in the Ecla language.

##### Fields
The `PrintStmt` node is defined as follows :

```go
    type PrintStmt struct {
        PrintToken lexer.Token
        Lparen     lexer.Token
        Rparen     lexer.Token
        Expression Expr
    }
```

The `PrintToken` field is the token that represents the print statement.
The `Lparen` field is the left parenthesis of the print statement.
The `Rparen` field is the right parenthesis of the print statement.
The `Expression` field is the expression of the print statement.

##### Code Example
a print statement is a statement that contains an expression surrounded by parenthesis.

for example :
```ecla   
    print(1);
    print("hello world");
```

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
    var a = 1;
    var a += 1;
    var a, b += 1, 2;
```


