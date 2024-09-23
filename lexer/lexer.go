package lexer

import "github.com/johneliud/Kisumu-Programming-Language/token"

// Lexer holds the input string and manages the current reading position.
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// Initializes a new Lexer instance with the input string and reads the first character.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Reads the next character from the input and advances positions.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Analyzes the current character and returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

	// Analyze the current character to determine the token type
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.NOT, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LESS_THAN, l.ch)
	case '>':
		tok = newToken(token.GREATER_THAN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMI_COLON, l.ch)
	case '(':
		tok = newToken(token.LEFT_PARENTHESIS, l.ch)
	case ')':
		tok = newToken(token.RIGHT_PARENTHESIS, l.ch)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LEFT_BRACKET, l.ch)
	case ']':
		tok = newToken(token.RIGHT_BRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Reads a string literal until the closing quote or end of input.
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()

		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

// Creates a new token with the specified type and literal value.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Reads a sequence of letters from the input.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Checks if a character is a valid letter or underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Reads a sequence of digits from the input.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Checks if a character is a valid digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Skips over spaces, tabs, and newlines.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Returns the next character in the input without advancing the position.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
