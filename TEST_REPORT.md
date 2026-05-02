# Low-Level-Lens Compiler: Final Project Report

## 1. Language Overview
The Low-Level-Lens compiler is designed to process arithmetic expressions and generate machine code for a custom 32-bit Virtual CPU (VCPU). The language supports standard symbolic operators, word-based operators, Arabic numerals, and an extensible numeral system (Roman numerals and number words).

## 2. Grammar Definition (EBNF)
The language syntax is formally described using the following EBNF grammar:

```ebnf
<expression>       ::= <term> { ( "+" | "-" | "plus" | "minus" ) <term> }
<term>             ::= <factor> { ( "*" | "/" | "times" | "divided" ) <factor> }
<factor>           ::= <number> | "(" <expression> ")"
<number>           ::= <arabic_numeral> | <roman_numeral> | <number_word>
<arabic_numeral>   ::= [0-9]+
<roman_numeral>    ::= "i" | "v" | "x"
<number_word>      ::= "one" | "two" | "three" | "ten"
```

## 3. Test Cases & Results
The following 18 test cases were executed against the compiler and VCPU. All cases passed successfully.

### Valid Expressions (10/10)
| Case | Input | Expected R0 | Result |
| :--- | :--- | :--- | :--- |
| Basic Add | `2 + 3 * 4` | 14 | PASS |
| Parens | `(2 + 3) * 4` | 20 | PASS |
| Precedence | `18 - 6 / 2` | 15 | PASS |
| Division | `100 / 5` | 20 | PASS |
| Subtraction | `42 - 19` | 23 | PASS |
| Nested Parens | `10 + (2 * 3)` | 16 | PASS |
| Chain Ops | `50 * 2 / 10` | 10 | PASS |
| Sub with Parens | `(100 - 50) + 10` | 60 | PASS |
| Multiply Chain | `5 * 5 * 5` | 125 | PASS |
| Long Addition | `1 + 2 + 3 + 4 + 5` | 15 | PASS |

### Invalid Expressions (5/5)
| Case | Input | Expected Error | Result |
| :--- | :--- | :--- | :--- |
| Div by Zero | `10 / (5 - 5)` | Semantic error: division by zero | PASS |
| Mismatched Parens | `(2 + 3` | Syntax Error | PASS |
| Double Operator | `2 ++ 3` | Syntax Error | PASS |
| Invalid Char | `2 $ 3` | Unknown Token | PASS |
| Bad Word Structure| `10 plus plus 5` | Syntax Error | PASS |

### Language Extensions (3/3)
| Case | Input | Expected R0 | Result |
| :--- | :--- | :--- | :--- |
| Roman Numerals | `X + V` | 15 | PASS |
| Word Operators | `10 plus 5` | 15 | PASS |
| Mixed Extension | `ten times v` | 50 | PASS |

## 4. Design Decisions
- **Lexical Analysis**: A hand-written scanner categorizes tokens into Numbers, Operators, and Parentheses. It uses a keyword map to normalize word operators (`plus`) to their symbolic equivalents (`+`) at the token level, allowing for easy expansion.
- **Parsing**: A Recursive Descent Parser (Top-Down) was chosen for its readability and ease of implementation for arithmetic grammars. It constructs an Abstract Syntax Tree (AST) to represent the hierarchy of operations.
- **Code Generation**: The compiler generates bytecode for a stack-like register machine. It uses a recursive register allocation strategy to manage intermediate values.
- **Extensibility**: The system is highly modular. Adding a new numeral system (e.g., Binary) only requires updating the Lexer to recognize the format and the Parser to convert it to an integer.
