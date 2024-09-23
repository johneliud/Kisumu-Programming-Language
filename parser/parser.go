package parser


import (
   "fmt"
   "strconv"


   "github.com/johneliud/Kisumu-Programming-Language/ast"
   "github.com/johneliud/Kisumu-Programming-Language/lexer"
   "github.com/johneliud/Kisumu-Programming-Language/token"
)


type Parser struct {
   l            *lexer.Lexer
   currentToken token.Token
   peekToken    token.Token
   errors       []string


   prefixParseFns map[token.TokenType]prefixParseFn
   infixParseFns  map[token.TokenType]infixParseFn
}


type (
   prefixParseFn func() ast.Expression
   infixParseFn  func(ast.Expression) ast.Expression
)


const (
   _ int = iota
   LOWEST
   EQUALS
   LESSGREATER
   SUM
   PRODUCT
   PREFIX
   CALL
   INDEX
)


var precedences = map[token.TokenType]int{
   token.EQUAL:            EQUALS,
   token.NOT_EQUAL:        EQUALS,
   token.LESS_THAN:        LESSGREATER,
   token.GREATER_THAN:     LESSGREATER,
   token.PLUS:             SUM,
   token.MINUS:            SUM,
   token.SLASH:            PRODUCT,
   token.ASTERISK:         PRODUCT,
   token.LEFT_PARENTHESIS: CALL,
   token.LEFT_BRACKET:     INDEX,
}


func New(l *lexer.Lexer) *Parser {
   p := &Parser{
       l:      l,
       errors: []string{},
   }
   p.nextToken()
   p.nextToken()
   p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
   p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
   p.registerPrefix(token.INT, p.parseIntegerLiteral)
   p.registerPrefix(token.NOT, p.parsePrefixExpression)
   p.registerPrefix(token.MINUS, p.parsePrefixExpression)
   p.registerPrefix(token.TRUE, p.parseBoolean)
   p.registerPrefix(token.FALSE, p.parseBoolean)
   p.registerPrefix(token.LEFT_PARENTHESIS, p.parseGroupedExpression)
   p.registerPrefix(token.IF, p.parseIfExpression)
   p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
   p.registerPrefix(token.STRING, p.parseStringLiteral)
   p.registerPrefix(token.LEFT_BRACKET, p.parseArrayLiteral)
   p.registerPrefix(token.LEFT_BRACE, p.parseHashLiteral)


   p.infixParseFns = make(map[token.TokenType]infixParseFn)
   p.registerInfix(token.PLUS, p.parseInfixExpression)
   p.registerInfix(token.MINUS, p.parseInfixExpression)
   p.registerInfix(token.SLASH, p.parseInfixExpression)
   p.registerInfix(token.ASTERISK, p.parseInfixExpression)
   p.registerInfix(token.EQUAL, p.parseInfixExpression)
   p.registerInfix(token.NOT_EQUAL, p.parseInfixExpression)
   p.registerInfix(token.LESS_THAN, p.parseInfixExpression)
   p.registerInfix(token.GREATER_THAN, p.parseInfixExpression)
   p.registerInfix(token.LEFT_PARENTHESIS, p.parseCallExpression)
   p.registerInfix(token.LEFT_BRACKET, p.parseIndexExpression)


   return p
}


func (p *Parser) peekPrecedence() int {
   if p, ok := precedences[p.peekToken.Type]; ok {
       return p
   }
   return LOWEST
}


func (p *Parser) currentPrecedence() int {
   if p, ok := precedences[p.currentToken.Type]; ok {
       return p
   }
   return LOWEST
}


func (p *Parser) parseIdentifier() ast.Expression {
   return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}


func (p *Parser) parseBoolean() ast.Expression {
   return &ast.Boolean{Token: p.currentToken, Value: p.currentTokenIs(token.TRUE)}
}


func (p *Parser) parseIntegerLiteral() ast.Expression {
   lit := &ast.IntegerLiteral{Token: p.currentToken}


   value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
   if err != nil {
       msg := fmt.Sprintf("could not pass %q as integer", p.currentToken.Literal)
       p.errors = append(p.errors, msg)
       return nil
   }
   lit.Value = value
   return lit
}


func (p *Parser) parsePrefixExpression() ast.Expression {
   expression := &ast.PrefixExpression{
       Token:    p.currentToken,
       Operator: p.currentToken.Literal,
   }
   p.nextToken()
   expression.Right = p.parseExpression(PREFIX)
   return expression
}


func (p *Parser) peekError(t token.TokenType) {
   msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
   p.errors = append(p.errors, msg)
}


func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
   expression := &ast.InfixExpression{
       Token:    p.currentToken,
       Operator: p.currentToken.Literal,
       Left:     left,
   }
   precedence := p.currentPrecedence()
   p.nextToken()
   expression.Right = p.parseExpression(precedence)
   return expression
}


func (p *Parser) parseGroupedExpression() ast.Expression {
   p.nextToken()
   exp := p.parseExpression(LOWEST)
   if !p.expectPeek(token.RIGHT_PARENTHESIS) {
       return nil
   }
   return exp
}


func (p *Parser) parseIfExpression() ast.Expression {
   expression := &ast.IfExpression{Token: p.currentToken}


   if !p.expectPeek(token.LEFT_PARENTHESIS) {
       return nil
   }


   p.nextToken()
   expression.Condition = p.parseExpression(LOWEST)


   if !p.expectPeek(token.RIGHT_PARENTHESIS) {
       return nil
   }


   if !p.expectPeek(token.LEFT_BRACE) {
       return nil
   }


   expression.Consequence = p.parseBlockStatement()


   if p.peekTokenIs(token.ELSE) {
       p.nextToken()


       if !p.expectPeek(token.LEFT_BRACE) {
           return nil
       }
       expression.Alternative = p.parseBlockStatement()
   }
   return expression
}


func (p *Parser) parseBlockStatement() *ast.BlockStatement {
   block := &ast.BlockStatement{Token: p.currentToken}
   block.Statements = []ast.Statement{}
   p.nextToken()


   for !p.currentTokenIs(token.RIGHT_BRACE) && !p.currentTokenIs(token.EOF) {
       stmt := p.parseStatement()
       if stmt != nil {
           block.Statements = append(block.Statements, stmt)
       }
       p.nextToken()
   }
   return block
}


func (p *Parser) parseFunctionLiteral() ast.Expression {
   lit := &ast.FunctionLiteral{Token: p.currentToken}


   if !p.expectPeek(token.LEFT_PARENTHESIS) {
       return nil
   }


   lit.Parameters = p.parseFunctionParameters()


   if !p.expectPeek(token.LEFT_BRACE) {
       return nil
   }
   lit.Body = p.parseBlockStatement()
   return lit
}


func (p *Parser) parseFunctionParameters() []*ast.Identifier {
   identifiers := []*ast.Identifier{}


   if p.peekTokenIs(token.RIGHT_PARENTHESIS) {
       p.nextToken()
       return identifiers
   }
   p.nextToken()
   ident := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
   identifiers = append(identifiers, ident)


   for p.peekTokenIs(token.COMMA) {
       p.nextToken()
       p.nextToken()
       ident := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
       identifiers = append(identifiers, ident)
   }


   if !p.expectPeek(token.RIGHT_PARENTHESIS) {
       return nil
   }
   return identifiers
}


func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
   exp := &ast.CallExpression{Token: p.currentToken, Function: function}
   exp.Arguments = p.parseExpressionList(token.RIGHT_PARENTHESIS)
   return exp
}


func (p *Parser) parseArrayLiteral() ast.Expression {
   array := &ast.ArrayLiteral{Token: p.currentToken}
   array.Elements = p.parseExpressionList(token.RIGHT_BRACKET)
   return array
}


func (p *Parser) parseHashLiteral() ast.Expression {
   hash := &ast.HashLiteral{Token: p.currentToken}
   hash.Pairs = make(map[ast.Expression]ast.Expression)


   for !p.peekTokenIs(token.RIGHT_BRACE) {
       p.nextToken()
       key := p.parseExpression(LOWEST)


       if !p.expectPeek(token.COLON) {
           return nil
       }
       p.nextToken()
       value := p.parseExpression(LOWEST)
       hash.Pairs[key] = value


       if !p.peekTokenIs(token.RIGHT_BRACE) && !p.expectPeek(token.COMMA) {
           return nil
       }
   }


   if !p.expectPeek(token.RIGHT_BRACE) {
       return nil
   }
   return hash
}


func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
   list := []ast.Expression{}


   if p.peekTokenIs(end) {
       p.nextToken()
       return list
   }


   p.nextToken()
   list = append(list, p.parseExpression(LOWEST))


   for p.peekTokenIs(token.COMMA) {
       p.nextToken()
       p.nextToken()
       list = append(list, p.parseExpression(LOWEST))
   }


   if !p.expectPeek(end) {
       return nil
   }
   return list
}


func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
   exp := &ast.IndexExpression{Token: p.currentToken, Left: left}
   p.nextToken()
   exp.Index = p.parseExpression(LOWEST)


   if !p.expectPeek(token.RIGHT_BRACKET) {
       return nil
   }
   return exp
}


func (p *Parser) parseStringLiteral() ast.Expression {
   return &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}


func (p *Parser) Errors() []string {
   return p.errors
}


func (p *Parser) nextToken() {
   p.currentToken = p.peekToken
   p.peekToken = p.l.NextToken()
}


func (p *Parser) ParseProgram() *ast.Program {
   program := &ast.Program{}
   program.Statements = []ast.Statement{}


   for p.currentToken.Type != token.EOF {
       stmt := p.parseStatement()
       if stmt != nil {
           program.Statements = append(program.Statements, stmt)
       }
       p.nextToken()
   }
   return program
}


func (p *Parser) parseStatement() ast.Statement {
   switch p.currentToken.Type {
   case token.VAR:
       return p.parseVarStatement()
   case token.RETURN:
       return p.parseReturnStatement()
   default:
       return p.parseExpressionStatement()
   }
}


func (p *Parser) parseVarStatement() *ast.VarStatement {
   stmt := &ast.VarStatement{Token: p.currentToken}


   if !p.expectPeek(token.IDENTIFIER) {
       return nil
   }
   stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}


   if !p.expectPeek(token.ASSIGN) {
       return nil
   }
   p.nextToken()
   stmt.Value = p.parseExpression(LOWEST)


   for !p.currentTokenIs(token.SEMI_COLON) {
       p.nextToken()
   }
   return stmt
}


func (p *Parser) parseExpression(precedence int) ast.Expression {
   prefix := p.prefixParseFns[p.currentToken.Type]
   if prefix == nil {
       p.noPrefixParseFnError(p.currentToken.Type)
       return nil
   }


   leftExp := prefix()


   for !p.peekTokenIs(token.SEMI_COLON) && precedence < p.peekPrecedence() {
       infix := p.infixParseFns[p.peekToken.Type]
       if infix == nil {
           return leftExp
       }
       p.nextToken()
       leftExp = infix(leftExp)
   }
   return leftExp
}


func (p *Parser) currentTokenIs(t token.TokenType) bool {
   return p.currentToken.Type == t
}


func (p *Parser) peekTokenIs(t token.TokenType) bool {
   return p.peekToken.Type == t
}


func (p *Parser) expectPeek(t token.TokenType) bool {
   if p.peekTokenIs(t) {
       p.nextToken()
       return true
   } else {
       p.peekError(t)
       return false
   }
}


func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
   stmt := &ast.ReturnStatement{Token: p.currentToken}
   p.nextToken()


   stmt.ReturnValue = p.parseExpression(LOWEST)


   for !p.currentTokenIs(token.SEMI_COLON) {
       p.nextToken()
   }
   return stmt
}


func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
   stmt := &ast.ExpressionStatement{Token: p.currentToken}
   stmt.Expression = p.parseExpression(LOWEST)
   if p.peekTokenIs(token.SEMI_COLON) {
       p.nextToken()
   }
   return stmt
}


func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
   p.prefixParseFns[tokenType] = fn
}


func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
   p.infixParseFns[tokenType] = fn
}


func (p *Parser) noPrefixParseFnError(t token.TokenType) {
   msg := fmt.Sprintf("no prefix parse function for %s found", t)
   p.errors = append(p.errors, msg)
}
