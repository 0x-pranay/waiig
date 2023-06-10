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
