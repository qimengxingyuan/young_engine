package token

import (
	"errors"
	"fmt"
	"strconv"
)

// Kind Represents all valid types of tokens that a token can be.
type Kind int

const (
	Illegal Kind = iota
	Eof
	KindBegin

	//literal_beg

	/*
	* literals of bool, int, string and variables
	* */
	Identifier     // variables
	BoolLiteral    // true, false
	IntegerLiteral // 12345
	FloatLiteral   // 123.45
	StringLiteral  // "abc"
	//literal_end

	//operator_beg

	/*
	* single character operator
	* */
	OpenParen  // (
	CloseParen // )

	/*
	* arithmetic operator
	* */
	Addition    // +
	Subtraction // -
	Multiply    // *
	Divide      // /
	Modulus     // %

	/*
	* cmp operator
	* */
	GreaterThan  // >
	LessThan     // <
	GreaterEqual // >=
	LessEqual    // <=
	Equal        // ==
	NotEqual     // !=

	/*
	* logic operator
	* */
	And // &&
	Or  // ||
	Not // !

	//operator_end

	KindEnd
)

var tokens = [...]string{
	/*
	* special tokens
	* */
	Illegal: "Illegal",
	Eof:     "Eof",

	/*
	* literals of nil, bool, number, string
	* */
	Identifier:     "Identifier",
	BoolLiteral:    "BoolLiteral",
	IntegerLiteral: "IntegerLiteral",
	FloatLiteral:   "FloatLiteral",
	StringLiteral:  "StringLiteral",

	/*
	* single character operator
	* */
	OpenParen:  "(",
	CloseParen: ")",

	/*
	* arithmetic operator
	* */
	Addition:    "+",
	Subtraction: "-",
	Multiply:    "*",
	Divide:      "/",
	Modulus:     "%",

	/*
	* cmp operator
	* */
	GreaterThan:  ">",
	LessThan:     "<",
	GreaterEqual: ">=",
	LessEqual:    "<=",
	Equal:        "==",
	NotEqual:     "!=",

	/*
	* logic operator
	* */
	And: "&&",
	Or:  "||",
	Not: "!",

	KindEnd: "KindEnd",
}

var operatorToKind = map[string]Kind{
	"(": OpenParen,
	")": CloseParen,

	"+": Addition,
	"-": Subtraction,
	"*": Multiply,
	"/": Divide,
	"%": Modulus,

	">":  GreaterThan,
	"<":  LessThan,
	">=": GreaterEqual,
	"<=": LessEqual,
	"==": Equal,
	"!=": NotEqual,

	"&&": And,
	"||": Or,
	"!":  Not,
}

// String returns the string corresponding to the token tok.
// For operators the string is the actual token character sequence (e.g., for the token Addition, the string is "+").
// For all other tokens the string corresponds to the token constant name (e.g. for the token Identifier, the string
// is "Identifier").
func (k Kind) String() string {
	s := ""
	if 0 <= k && k < Kind(len(tokens)) {
		s = tokens[k]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(k)) + ")"
	}
	return s
}

func (k Kind) IsIllegal() bool {
	return k == Illegal
}

func (k Kind) IsEof() bool {
	return k == Eof
}

func (k Kind) GetLexerState() (LexerState, error) {
	state, exist := validLexerStates[k]
	if exist {
		return state, nil
	}
	errorMsg := fmt.Sprintf("No lexer state found for token kind '%v'\n", k.String())
	return validLexerStates[0], errors.New(errorMsg)
}
