# Writing An Interpreter in Go

## Introduction

We are going to parse and evaluate our own language called Monkey.

Monkey has these features: 
- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

Major parts 
- lexer
- parser
- AST
- the internal object system
- the evaluator


## Chapter 1: Lexing


┌─────────────────┐     ┌───────────┐      ┌────────────────────────┐
│                 │     │           │      │                        │
│   Source Code   ├─────►  Tokens   ├──────► Abstract Syntax Tree   │
│                 │     │           │      │                        │
└─────────────────┘     └───────────┘      └────────────────────────┘

The first transformation, from source code to tokens is called "lexical analysis" or "lexing".

- Current lexer only supports ASCII character set since we use `byte` to store current character. If we need to support full Unicode range or UTF-8 better use rune (aka int32) type which uses 4 bytes(32bits) 

### extending token set and lexer
- To support ==, !, !=, - , /, *, <, > and keywords `true`, `false`, `if`, `else, and `return`
- We classify new tokens into one-character token (`-`, `+` ), two-character token (e.g. ==) and keyword token (e.g. return)

## Chapter 2: Parsing
- A parser turns its input into a data structure that represents the input.
- The data structure used for the internal representation of tghe source code is called a "syntax tree" or an "abstract syntax tree" (AST for short). The "abstract" is based on the fact that certain details visible in source code are omitted in AST. e.g. Semicolons, newlines, whitespaces etc depending on the language.

- What parsers do ?
  They take source code as input (either as text or tokens) and produce
a data structure which represents this source code. While building up the data structure, they
unavoidably analyse the input, checking that it conforms to the expected structure. Thus the
process of parsing is also called syntactic analysis.

### parser generators
- yacc, bison or ANTLR. 
- These are tools that takes formal description of a language as input and produce parsers as their output. This output is code that can be compiled/interpreted and itself fed with source code as input to produce a syntax tree.
- Context free grammer (CFG) is widely used input format for may parser generators.
- CFG is a set of rules that describe how to form  correct sentences in a language.
- Backus-Naur Form (BNF) or the Extended Backus-Naur Form (EBNF) are the most common CFG notational formats.
- Here is a full description of EcmaScript syntax in BNF - https://tomcopeland.blogs.com/EcmaScript.html


### Parser for monkey programming language
- Two main strategies when parsing a programming language
  1. Top-down parsing
    - other variations
      - recursive descent parsing
      - early parsing or predictive parsing
  2. Botton-up parsing

- We are going to write recursive descent parser. Its a top down operator precedence parser, somtimes called "Pratt parser", after its inventor Vaughan Pratt.
