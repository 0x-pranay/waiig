package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifier and literals
  IDENT = "IDENT"
  INT = "INT"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  LT = "<"
  GT = ">"

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE= "TRUE"
  FALSE= "FALSE"
  IF= "IF"
  ELSE= "ELSE"
  RETURN= "RETURN"
)

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

// checks the keyword table to see if given identifier is a keyword.
// if keyword it returns that keyword's TokenType 
// else returns token.IDENT, which is the TokenType for all user defined identifiers
func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
