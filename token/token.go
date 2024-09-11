package token

// TokenType defines the type for representing different token types in the language.
type TokenType string

// Token struct represents a lexical token with its type and literal value.
type Token struct {
	Type    TokenType
	Literal string
}

// Constants representing the types of tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	NOT      = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	EQUAL     = "=="
	NOT_EQUAL = "!="

	COMMA      = ","
	SEMI_COLON = ";"
	COLON      = ":"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	STRING        = "STRING"
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"
)

// maps keywords in the language to their respective token types.
var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Checks if the given identifier is a keyword, and returns the corresponding token type. If it's not a keyword, it returns the IDENTIFIER token type.
func LookupIdentifier(identifier string) TokenType {
	tok, ok := keywords[identifier]
	if ok {
		return tok
	}
	return IDENTIFIER
}
