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
### Parsing process
#### Preprocessing
#### The parsing operations
#### Postprocessing
### Edge cases
#### Operator precedence
#### Omitted semicolons
### Error handling
#### Error detection
#### Information about the error
#### Error recovery
#### Error reporting
### Performance
#### Time complexity
#### Space complexity
#### Memory usage
#### Performance tests
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