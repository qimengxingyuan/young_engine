package compiler

import (
	"fmt"
	"testing"

	"github.com/qimengxingyuan/young_engine/token"
)

func TestLexer(t *testing.T) {
	//rule := "f_2128 == \"app\" && f_13288.Hit&& (f_799 == \"\"|| !Exist( f_799  ) || !Exist( f_2458  ) || f_785 == \"\"|| !Exist( f_785  ) || f_802 == \"\"|| !Exist( f_802  ) || f_800 == \"\"|| !Exist( f_800  ) || f_784 == \"\" || !Exist(f_784) || !Exist(f_1126)|| f_1126== \"\" || !Exist(f_1006)|| f_1006==\"\"|| !Exist(f_2956)|| f_2956==\"\" || (   f_784 == \"android\"   &&(!Exist(f_9342) || f_9342==\"\" || !Exist(f_13469) || f_13469==\"\" || !Exist(f_2957) || f_2957==\"\" || f_13118==\"\" || !Exist(f_13118) || !Exist(f_12461) || f_12461==\"\"   )) || ( f_784 in ('iphone', 'ipad') && (!Exist(f_2959)|| f_2959 == \"\")))"
	//rule := "a==true"
	rule := "1 + 1 + 1"
	//rule := "ConvNum( ip1_str  )& 1 == 1"
	//rule := "s1 != 'abc\n123' \n && s2 != 'abc\\n123'"
	//rule := "n1 in (11 22 33 44xb)"
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
