package young_engine

import (
	"github.com/qimengxingyuan/young_engine/executor"
	"testing"
	"time"
)

func TestCompiler(t *testing.T) {
	//rule := `true || 7 > 9`
	rule := `29 + "hhh"`
	//rule := `7 * 9 + 8 * 4`
	node, err := Compiler(rule)
	if err != nil {
		t.Log(err)
	}

	node.Print()
	time.Sleep(1 * time.Second)
	err = node.Eval(executor.DummyParameters)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(node.GetVal())
	}
}
