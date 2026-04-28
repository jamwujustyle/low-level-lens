Step 1: What is a Lexer and What is a Token?
Imagine I give you the sentence: "I love coding". As a human, you instantly see three words. But to a computer, it's just a raw string of 13 characters: ['I', ' ', 'l', 'o', 'v', 'e', ' ', 'c', 'o', 'd', 'i', 'n', 'g'].

A Lexer (short for Lexical Analyzer) acts like a scanner. It reads those individual characters one by one and groups them into meaningful "words" called Tokens.

So, if we feed our Lexer the math problem 10 + 5, it shouldn't see ['1', '0', ' ', '+', ' ', '5']. It needs to group them into three Tokens:

Number(10)
Operator(+)
Number(5)

Lexer: Turns text into tokens.
Parser: Turns tokens into a tree.
Semantic Analyzer: Validates the math.
Code Generator: Translates the tree into your custom Assembly and Binary.
Virtual CPU: Executes the binary instructions and manages registers.
