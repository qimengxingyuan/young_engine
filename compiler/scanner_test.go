package compiler

import (
	"fmt"
	"testing"

	"github.com/qimengxingyuan/young_engine/token"
)

func TestLexer(t *testing.T) {
	//rule := "a==true"
	//rule := "1 + 1 + 1"
	rule := "s1 != 'abc\n123' \n && s2 != 'abc\\n123'"
	//rule := "100+200a==300"
	//rule := ""

	scanner := NewScanner(rule)

	for {
		tok, err := scanner.Scan()
		if err != nil {
			t.Log(err)
			break
		}
		if tok.Kind == token.Eof {
			break
		}
		fmt.Printf("kind=[%v], val=[%v], pos=[%d]\n", tok.Kind, tok.Value, tok.Position)
	}

	fmt.Printf("scanner done\n")
}
