package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	NOT     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESS = "<"
	GREATER = ">"

	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Delimiters
	COMMA     = ","
	SEMI_COLON = ";"

	LEFT_PARENTHESIS = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE = "{"
	RIGHT_BRACE = "}"

	// Keywords
	FUNC 	 = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":   FUNC,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}