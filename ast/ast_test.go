package ast

import (
	"testing"
	
	"github.com/btrobot/mydsl/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestLetStatement(t *testing.T) {
	letStmt := &LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let", Line: 1, Column: 1},
		Name: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "x", Line: 1, Column: 5},
			Value: "x",
		},
		Value: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "5", Line: 1, Column: 9},
			Value: 5,
		},
	}
	
	if letStmt.TokenLiteral() != "let" {
		t.Errorf("letStmt.TokenLiteral() wrong. got=%q", letStmt.TokenLiteral())
	}
	
	if letStmt.String() != "let x = 5;" {
		t.Errorf("letStmt.String() wrong. got=%q", letStmt.String())
	}
	
	line, col := letStmt.Position()
	if line != 1 || col != 1 {
		t.Errorf("letStmt.Position() wrong. got=(%d, %d)", line, col)
	}
}

func TestIdentifier(t *testing.T) {
	ident := &Identifier{
		Token: token.Token{Type: token.IDENT, Literal: "x", Line: 1, Column: 5},
		Value: "x",
	}
	
	if ident.TokenLiteral() != "x" {
		t.Errorf("ident.TokenLiteral() wrong. got=%q", ident.TokenLiteral())
	}
	
	if ident.String() != "x" {
		t.Errorf("ident.String() wrong. got=%q", ident.String())
	}
	
	line, col := ident.Position()
	if line != 1 || col != 5 {
		t.Errorf("ident.Position() wrong. got=(%d, %d)", line, col)
	}
}
