package token

import (
	"testing"
)

// Tests to make sure that all the different token kinds have different string representations
func TestTokenKindStrings(test *testing.T) {
	if len(tokens) != int(KindEnd)+1 {
		test.Logf("Token kind test found the number of Kind and the length of 'tokens' which holds the name of kind do not match")
		test.Fail()
		return
	}

	kinds := make([]Kind, 0, KindEnd)
	for i := Kind(0); i < KindEnd; i++ {
		kinds = append(kinds, i)
	}

	kindStrings := make(map[string]struct{})
	for _, kind := range kinds {
		kindString := kind.String()
		if _, exist := kindStrings[kindString]; exist {
			test.Logf("Token kind test found duplicate string for token kind %v ('%v')\n", kind, kindString)
			test.Fail()
			return
		}
		kindStrings[kindString] = struct{}{}
	}
}
