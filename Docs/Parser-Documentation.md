# The Ecla parser documentation

## Table of content

- Technical specifications
  - Type of parser
  - Inner structure
    - Parser
    - AST
    - File
    - Node
- Inner workings
  - Parsing process
    - Preprocessing
    - The parsing operations
    - Postprocessing
  - Edge cases
    - Operator precedence
    - Omitted semicolons
  - Error handling
    - Error detection
    - Information about the error
    - Error recovery
    - Error reporting
  - Performance
    - Time complexity
    - Space complexity
    - Memory usage
    - Performance tests
- Grammar
  - Table of symbols
  - Grammar rules
    - Node rule 1
    - Node rule 2
    - Node rule 3


## Technical specifications
Here are the technical specifications of the Ecla parser.

### Type of parser
The type of parser used by the Ecla interpreter is an LL(1) recursive descent parser based on context-free grammar or CFG.  
LL(1) means that the parser is a left-to-right parser that reads one token at a time and that it can only look one token ahead.  
Recursive descent means that the parser is a top-down parser that uses a recursive function to parse the input.  
And context-free grammar means that the parser uses a grammar which as the following properties :

- V a finite set of symbols such that V is a set of non-terminals symbols.
- Σ is a finite set of symbols such that V∩Σ=∅ and Σ is a set of terminal symbols.
- P is a set of production rules in the form A→w where A∈V and w∈(V∪Σ)*.
- And S is the start symbol such that S∈V.

### Inner structure
The parser is composed of 4 main components that are the `Parser`, the `AST`, the `Files` and the `Nodes`.
#### Parser
The `Parser` structure is one of the most important structure of the Ecla parser.  
It is composed but not limited to :

- The Tokens yielded by the Ecla lexer.
- The Current `File`.
- And the error handler.

It also contains all the parsing methods that are used to parse the input.
#### AST
The `AST` structure or abstract syntax tree is the structure that contains an abstract representation of the input as parsed by the `Parser`.
It contains only one field which is an array of Nodes.
#### Files
The `File` structure is the structure that contains the information about an Ecla source file.
It contains but not limited to :

- The parse tree of the `File` source code.
- The imported packages.
- His dependencies.
- The variables declared in the `File`.
- The functions declared in the `File`.
- And the Consumed comments.

It also contains all the preprocessing and postprocessing methods that are used to parse the input.
#### Nodes
The `Node` interface is the building block of the AST.
There exists 3 flavors of nodes :

- The Expression nodes.
- The Statement nodes.
- And the Declaration nodes.

Those 3 types of nodes are derived from the main Node interface containing :

- The start position of the `Node`.
- And the end position of the `Node`.

All Nodes that represent `Expressions`, `Statements` or `Declarations` implement one of the 3 flavors of nodes.
for example, the `BinaryExpr` node implements the `Expression` interface.
## Inner workings
Here is a detailed description of the Ecla parser inner workings.
### Parsing process
The parsing process is composed of 3 main steps :
#### Preprocessing
During the preprocessing step, the Ecla parser will remove all the comments from the input.  
This is done by the `File` structure which contains the `ConsumeComments` method.  
All the comments are stored in the `File` structure in the `Comments` field.
#### The parsing operations
During the parsing operations step, the `Parser` will parse the input and build the AST.  
Add details about the parsing operations here.  
After the parsing operations step, the `Parser` will have built the AST in the `File` structure.  

#### Postprocessing
During the postprocessing step, the Ecla parser will check for dependencies resolution.  
This is done by the `File` structure which contains the `DepChecker` method.  
It is run at the end of the parsing operations and throws an error if a dependency is not resolved.  
### Edge cases
Here are the edge cases that are handled by the Ecla parser.  
#### Operator precedence
The Ecla parser handles the operator precedence by using a precedence table.  

Here it is :

|          **Operator**           | **Precedence** |
|:-------------------------------:|:--------------:|
|         `&#124;&#124;`          |       1        |
|              `&&`               |       2        |
| `=`, `!=`, `<`, `<=`, `>`, `>=` |       3        |
|            `+`, `-`             |       4        |
|       `*`, `/`, `%`, `//`       |       5        |

All the operators are not yet implemented so the table is subject to change.  

#### Omitted semicolons
The Ecla parser handles the omitting of semicolons after conditional and loop statements.  
By tracking the state of the `IsEndOfBrace` field of the `Parser` structure, the parser skip the `EOL` check.  
If a semicolon is still present, the parser will throw a `Warning` to notify the user that the semicolon is not needed.
### Error handling
Here are the error handling methods used by the Ecla parser.
#### Error detection
The detection of errors is done by the `Parser`, it is done in a failsafe way.  
By making the error the first priority, if the `Parser` encounters an error, it will throw an error and stop the parsing process instead of continuing and throwing a panic.
Furthermore, the `Parser` uses the centralised error handling system of the Ecla interpreter.
#### Information about the error
The `Parser` will throw an error with the following information :

- The time of the error.
- The type of the error.
- The error message.
- The line of the error.
- The column of the error.
- A little snippet of the source code around the error.

some information may be missing like the file name because it is not implemented yet.
#### Error recovery
Recovering from an error is not implemented yet but will be implemented in the future by using the `try` and `catch` keywords.
#### Error reporting
The error reporting is not implemented yet but will be implemented in the future by implementing logging.
### Performance
Here are the performance metrics of the Ecla parser.
#### Time complexity
Need to be measured.
#### Space complexity
Need to be measured.
#### Memory usage
Need to be measured.
#### Performance tests
Need te create performance tests.
## Grammar
### Table of symbols
|    **Usage**     | **Notation** |
|:----------------:|:------------:|
|    definition    |      =       |
|  concatenation   |      ,       |
|   termination    |      ;       |
|   alternation    |      \       |
|     optional     |   [ ... ]    |
|    repetition    |   { ... }    |
|     grouping     |   ( ... )    |
| terminal string  |   " ... "    |
| terminal string  |   ' ... '    |
|     comment      |  (* ... *)   |
| special sequence |   ? ... ?    |
|    exception     |      -       |

>as proposed by the ISO/IEC 14977 standard, page 7, table 1
>
### Grammar rules
#### Node rule 1
#### Node rule 2
#### Node rule 3