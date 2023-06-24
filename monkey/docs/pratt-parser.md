Parsing let and return statements are straightforward but parsing expressions are hard.

Expressions come in different form
- involving prefix operator
  `-5`
  `!true`
  `!false`
- infix operator (or "binary operators")
  `5+5`
  `5-5`
  `5/5`
  `5*5`
- Operator precedence
  `5 + 5*5`
  `(5+5)*5`
- comparison operators
  `foo==bar`
  `foo != bar`
  `foo < bar`
  `foo > bar`
- grouped expressions
  `5 * (5 + 5)`
  `((5+5)*5)*5`
- call expressions
  `add(2,3)`
  `add(add(2,3), add(5,10))`
  `max(5, add(5, (5*5)))`
- using Identifiers
  `foo * bar/ foobar`
  `add(foo, bar)`
- first class functions
  `let add = fn(x,y){ return x+y };`
  `fn(x,y){ return x+y }(5,5)`
  `(fn(x){ return x*2 }(5) + 10) * 10
- If statements as expressions
  `let result = if (10>5) { true } else {false};` // true

# TOP DOWN OPERATOR PRECEDENCE (OR PRATT PARSING)

The main idea of Pratt parser is to associate parsing functions (aka "semantic code") with token type. Whenever this token type is encountered, the parsing functions are called to parse the appropriate expressions and return an AST node that represents it. 
- Each token can type can have up to two parsing functions depending on whethers its prefix or an infix position. 
















### Further reading
- http://javascript.crockford.com/tdop/tdop.html
- http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy
