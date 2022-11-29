package compiler

import (
	"unicode"
	"unicode/utf8"
)

func lower(ch rune) rune { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter

// isLetter reports whether a given 'rune' is classified as a Letter.
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isDigit(ch rune) bool {
	return isDecimal(ch) || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func isDot(ch rune) bool {
	return ch == '.'
}

func isEof(ch rune) bool {
	return ch == eofRune
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= lower(ch) && lower(ch) <= 'f':
		return int(lower(ch) - 'a' + 10)
	}
	return 16 // larger than any legal digit val
}

func parseBool(b string) bool {
	return b == "true"
}
