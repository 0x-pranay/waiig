package parser

import (
  "monkey/ast"
  "monkey/lexer"
  "monkey/token"
  "fmt"
)

type Parser struct {
  l *lexer.Lexer

  curToken token.Token
  peekToken token.Token

  errors []string
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{
    l :     l,
    errors: []string{},
  }
  
  // Read two tokens, so curToken and peekToken are both set
  p.nextToken()
  p.nextToken()
  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

// The first thing parseProgram does is 
// construct the root node of the AST, an *ast.Program

func (p *Parser) ParseProgram() *ast.Program {
  program := &ast.Program{}
  program.Statements = []ast.Statement{}

  for !p.curTokenIs(token.EOF) {
    stmt := p.parseStatement()
    if stmt != nil {
      program.Statements = append(program.Statements, stmt)
    }
    p.nextToken()
  }

  return program
}

func (p *Parser) parseStatement() ast.Statement {
  switch p.curToken.Type {
  case token.LET:
    return p.parseLetStatement()
  case token.RETURN:
    return p.parseReturnStatement()
  default:
    return nil
  }
}

// parseLetStatement
// 1. construct *ast.LetStatement node
// 2. moves to next token making assertions with expectPeek method. If yes moves to nextToken else return false. 
// 2a. Expect token.IDENT, then construct *ast.Identifier node
// 2b. Expect =, and finally jumps over the expression until it encounters a semicolon.
// 2c. TODO: handle expression

func (p *Parser) parseLetStatement() *ast.LetStatement {
  stmt := &ast.LetStatement{Token: p.curToken}
  // if nextToken is token.IDENT it moves to the nextToken
  if !p.expectPeek(token.IDENT) {
    return nil
  }
  stmt.Name = &ast.Identifier{Token: p.curToken, Value : p.curToken.Literal }
  if !p.expectPeek(token.ASSIGN) {
    return nil
  }
  // TODO: we're skipping the expression until we
  // encounter a semicolon
  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }
  return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
  stmt := &ast.ReturnStatement{ Token: p.curToken }

  p.nextToken()

  // TODO: We're skipping the expression until we encounter a semicolon
  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }
  return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
  return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
  return p.peekToken.Type == t
}

// an assertion function.
// Primary purpose is to enforce the correctness of the order of tokens by checking the type of the next token using `peekToken`.
// If type if correct then advance to next token using `nextToken`
// Else return nil. (rightnow) returned nil gets ignored in ParseProgram, which results in entire statement being ignored because of error in input. 
// ~TODO~ DONE: add error handling to our parser here
func (p *Parser) expectPeek(t token.TokenType) bool {
  if p.peekTokenIs(t) {
    p.nextToken()
    return true
  } else {
    p.peekError(t)
    return false
  }
}

func (p *Parser) Errors() []string {
  return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
  msg := fmt.Sprintf("expected next token to be %s, got %s instead",
      t, p.peekToken.Type)
  p.errors = append(p.errors, msg)
}
