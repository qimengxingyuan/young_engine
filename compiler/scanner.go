package compiler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/qimengxingyuan/young_engine/token"
)

const eofRune = -1

type Scanner struct {
	source   []rune
	position int
	length   int
	ch       rune
}

func NewScanner(source string) *Scanner {
	runes := []rune(source)

	if len(runes) == 0 {
		runes = append(runes, rune(eofRune))
	}

	return &Scanner{
		source: runes,
		length: len(runes),
		ch:     runes[0],
	}
}

// read returns the character at the pos of position and advancing
// the scanner. If the scanner is at Eof, read returns -1.
func (scanner *Scanner) read() rune {
	var char rune

	if !scanner.canRead() {
		return eofRune
	}

	char = scanner.source[scanner.position]
	scanner.ch = scanner.peek()

	scanner.position += 1

	return char
}

// cur returns the character at the pos of position
func (scanner *Scanner) cur() rune {
	return scanner.ch
}

// peek returns the rune following the most recently read character without
// advancing the scanner. If the scanner is at Eof, peek returns -1.
func (scanner *Scanner) peek() rune {
	if scanner.position < scanner.length-1 {
		return scanner.source[scanner.position+1]
	}
	return eofRune
}

func (scanner *Scanner) canRead() bool {
	return scanner.position < scanner.length
}

func (scanner *Scanner) skipWhitespace() {
	for scanner.canRead() && unicode.IsSpace(scanner.cur()) {
		scanner.read()
	}
}

func (scanner *Scanner) scanIdentifier() string {
	startPos := scanner.position
	for isLetter(scanner.peek()) || isDigit(scanner.peek()) {
		scanner.read()
	}
	scanner.read()
	return string(scanner.source[startPos:scanner.position])
}

func (scanner *Scanner) scanNumber() string {
	startPos := scanner.position
	for isDigit(scanner.peek()) || isDot(scanner.peek()) {
		scanner.read()
	}
	scanner.read()

	return string(scanner.source[startPos:scanner.position])
}

func (scanner *Scanner) scanString() (string, error) {
	var err error
	quote := scanner.read() // consume " or \'
	startPos := scanner.position
	endPos := scanner.position
	for {
		ch := scanner.read()
		if isEof(ch) { // the scanner ends, but the terminator of string literal is not read
			err = errors.New("string literal not terminated")
			break
		}
		if ch == quote { // read the terminator of string literal
			endPos = scanner.position - 1 // give up  " or \'
			break
		}

		if ch == '\\' { // escape characters
			if err = scanner.scanEscape(quote); err != nil {
				break
			}
		}
	}

	return string(scanner.source[startPos:endPos]), err
}

func (scanner *Scanner) scanRawString() (string, error) {
	var err error
	quote := scanner.read() // consume `
	startPos := scanner.position
	endPos := scanner.position

	for {
		ch := scanner.read()
		if isEof(ch) {
			err = errors.New("raw string literal not terminated")
			break
		}

		if ch == quote {
			endPos = scanner.position - 1 // consume `
			break
		}
	}

	lit := scanner.source[startPos:endPos]

	return string(lit), err
}

// scanEscape parses an escape sequence where rune is the accepted
// escaped quote. In case of a syntax error, it stops at the offending
// character (without consuming it) and returns error message. Otherwise
// it returns nil.
func (scanner *Scanner) scanEscape(quote rune) error {

	var n int
	var err error
	var base, max uint32
	switch scanner.peek() {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'', '"', quote:
		scanner.read()
		return nil
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		scanner.read()
		n, base, max = 2, 16, 255
	case 'u':
		scanner.read()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		scanner.read()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if scanner.peek() < 0 {
			msg = "escape sequence not terminated"
		}
		err = errors.New(msg)
		return err
	}

	var x uint32
	for n > 0 {
		d := uint32(digitVal(scanner.peek()))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", scanner.peek())
			if scanner.peek() < 0 {
				msg = "escape sequence not terminated"
			}
			err = errors.New(msg)
			return err
		}
		x = x*base + d
		scanner.read()
		n--
	}

	if x > max || 0xD800 <= x && x < 0xE000 {
		err = errors.New("escape sequence is invalid Unicode code point")
		return err
	}

	return nil
}

func (scanner *Scanner) scanSwitch2(kid0 token.Kind, ch1 rune, kid1 token.Kind) (string, token.Kind) {
	var s string
	s += string(scanner.read())
	if scanner.cur() == ch1 {
		s += string(scanner.read())
		return s, kid1
	}
	return s, kid0
}

func (scanner *Scanner) scanExpect1(kid token.Kind, ch rune) (string, token.Kind) {
	var s string
	s += string(scanner.read())
	if scanner.cur() == ch {
		s += string(scanner.read())
		return s, kid
	} else {
		return s, token.Illegal
	}
}

func (scanner *Scanner) Scan() (token.Token, error) {
	var tok token.Token
	var err error

	scanner.skipWhitespace()
	tok.Position = scanner.position

	switch ch := scanner.cur(); {
	case isEof(ch):
		tok.Kind = token.Eof
	case isLetter(ch):
		// if the first character is letter, this token must be an Identifier or BoolLiteral or otherwise
		literal := scanner.scanIdentifier()
		tok.Kind = token.Lookup(literal)
		tok.Value = literal
		// boolean?
		if tok.Kind == token.BoolLiteral {
			tok.Value = parseBool(literal)
		}
	case isDecimal(ch) || isDot(ch):
		// Decimal,
		literal := scanner.scanNumber()
		if strings.Contains(literal, ".") { // float
			tok.Value, err = strconv.ParseFloat(literal, 64)
			tok.Kind = token.FloatLiteral
		} else { // int
			tok.Value, err = strconv.ParseInt(literal, 10, 64)
			tok.Kind = token.IntegerLiteral
		}

		if err != nil {
			errorMsg := fmt.Sprintf("Unable to compiler numeric value '%v'", literal)
			return tok, errors.New(errorMsg)
		}
	default:
		switch ch {
		case '+', '-', '*', '/', '%', '(', ')': // 确定的单一运算符
			tok.Kind = token.LookupOperator(string(ch))
			tok.Value = scanner.read()
		case '"', '\'':
			tok.Kind = token.StringLiteral
			tok.Value, err = scanner.scanString()
		case '`':
			tok.Kind = token.StringLiteral
			tok.Value, err = scanner.scanRawString()
		case '<':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.LessThan, '=', token.LessEqual)
		case '>':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.GreaterThan, '=', token.GreaterEqual)
		case '!':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.Not, '=', token.NotEqual)
		case '=':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.Illegal, '=', token.Equal)
			if tok.Kind.IsIllegal() {
				return tok, errors.New("expected to get '==', but only found '='")
			}
		case '&':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.Illegal, '&', token.And)
			if tok.Kind.IsIllegal() {
				return tok, errors.New("expected to get '&&', but only found '&'")
			}
		case '|':
			tok.Value, tok.Kind = scanner.scanSwitch2(token.Illegal, '|', token.Or)
			if tok.Kind.IsIllegal() {
				return tok, errors.New("expected to get '||', but only found '|'")
			}
		default:
			tok.Kind = token.Illegal
			tok.Value = string(ch)
			errMsg := fmt.Sprintf("the scan found an illegal character '%v'", ch)
			return tok, errors.New(errMsg)
		}
	}

	return tok, err
}

func (scanner *Scanner) Lexer() ([]token.Token, error) {
	tokens := make([]token.Token, 0)

	var err error
	var tok token.Token
	for {
		tok, err = scanner.Scan()
		tokens = append(tokens, tok)
		if err != nil || tok.Kind == token.Eof {
			break
		}
	}

	return tokens, err
}
