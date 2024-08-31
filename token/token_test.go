package token

import (
	"testing"
)
	
	func TestLookupIdentReturnsCorrectTokenTypeForKnownKeywords(t *testing.T) {
		keywords := map[string]TokenType{
			"if":     IF,
			"else":   ELSE,
			"return": RETURN,
		}
	
		for ident, expectedToken := range keywords {
			result := LookupIdent(ident)
			if result != expectedToken {
				t.Errorf("LookupIdent(%s) = %v; want %v", ident, result, expectedToken)
			}
		}
	}

	func TestLookupIdentHandlesEmptyStringInput(t *testing.T) {
		result := LookupIdent("")
		if result != IDENT {
			t.Errorf("LookupIdent(\"\") = %v; want %v", result, IDENT)
		}
	}
	
	