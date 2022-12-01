package compiler

import "testing"

func TestParser_CheckBalance(t *testing.T) {
	exp := `1 + 1 + 1`
	// 词法分析
	tokenScanner := NewScanner(exp)
	tokens, err := tokenScanner.Lexer()
	if err != nil {
		t.Error(err)
	}

	// 语法分析
	parser := NewParser(tokens)
	parser.Print()
	err = parser.checkBalance()
	if err != nil {
		t.Error(err)
	}

	err = parser.ParseSyntax()
	if err != nil {
		t.Error(err)
	}
}
