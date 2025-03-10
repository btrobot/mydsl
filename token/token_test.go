package token

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	tests := []struct {
		tokenType TokenType
		literal   string
		line      int
		column    int
	}{
		{IDENT, "x", 1, 1},
		{INT, "5", 1, 3},
		{PLUS, "+", 1, 5},
		{SEMICOLON, ";", 1, 6},
	}

	for i, tt := range tests {
		token := NewToken(tt.tokenType, tt.literal, tt.line, tt.column)

		if token.Type != tt.tokenType {
			t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.tokenType, token.Type)
		}

		if token.Literal != tt.literal {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.literal, token.Literal)
		}

		if token.Line != tt.line {
			t.Errorf("tests[%d] - line wrong. expected=%d, got=%d",
				i, tt.line, token.Line)
		}

		if token.Column != tt.column {
			t.Errorf("tests[%d] - column wrong. expected=%d, got=%d",
				i, tt.column, token.Column)
		}
	}
}

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		ident    string
		expected TokenType
	}{
		{"x", IDENT},
		{"let", LET},
		{"function", FUNCTION},
		{"if", IF},
		{"else", ELSE},
		{"true", TRUE},
		{"false", FALSE},
		{"return", RETURN},
		{"unknown", IDENT},
	}

	for i, tt := range tests {
		tokenType := LookupIdent(tt.ident)

		if tokenType != tt.expected {
			t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expected, tokenType)
		}
	}
}
