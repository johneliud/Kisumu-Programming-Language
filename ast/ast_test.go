package ast

import (
	"testing"

	"github.com/johneliud/Kisumu-Programming-Language/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"}, Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "var myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
