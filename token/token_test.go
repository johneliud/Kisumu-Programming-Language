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
			result := LookupIdentifier(ident)
			if result != expectedToken {
				t.Errorf("LookupIdent(%s) = %v; want %v", ident, result, expectedToken)
			}
		}
	}

	func TestLookupIdentHandlesEmptyStringInput(t *testing.T) {
		result := LookupIdentifier("")
		if result != IDENTIFIER {
			t.Errorf("LookupIdent(\"\") = %v; want %v", result, IDENTIFIER)
		}
	}
	
	