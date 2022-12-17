package test

import (
	"testing"
)

func TestRune(t *testing.T) {
	t.Log(rune(0))
	t.Log(string(rune(40)))

	s := "100+200a==300"
	t.Log(s[0:13])
}
