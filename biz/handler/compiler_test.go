package handler

import (
	"os/exec"
	"testing"
	"time"
)

func TestCompiler(t *testing.T) {
	//rule := `!true || !!(7 > 9)`
	//rule := `-8 + 2 - 4`
	rule := `7.7.7  * -9 + 8 * 9`
	//rule := `--7  * -9 + -8 * 9`
	//rule := "s1 != 'abc123' && s2 != 'abc\n123'"
	//rule := "\"abc\n1234\"== 'abc\n123'"
	node, err := Compiler(rule)
	if err != nil {
		t.Error(err)
		return
	}

	// print and open svg
	node.PrintSvg("node")
	exec.Command("cmd", "/c", "start", "node.svg").Start()
	exec.Command("open", "node.svg").Start()
	time.Sleep(1 * time.Second)

	// eval
	params := map[string]interface{}{}
	err = node.Eval(params)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(node.GetVal())
	}
}
