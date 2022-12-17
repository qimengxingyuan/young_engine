package token

type LexerState struct {
	isEOF          bool
	validNextKinds []Kind
}

// lexer states. Define all the grammar rules here
var validLexerStates = map[Kind]LexerState{
	Eof: {
		isEOF:          true,
		validNextKinds: []Kind{},
	},
	/*
		The following token types are allowed at the beginning of an expression.
	*/
	Illegal: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			BoolLiteral,    // true, false
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			StringLiteral,  // "abc"
			OpenParen,      // (
			Addition,       // +
			Subtraction,    // -
			Not,            // !
		},
	},

	/*
	* literals of bool, int, string and variables
	* */
	Identifier: {
		isEOF: true,
		validNextKinds: []Kind{
			CloseParen, // )
			// arithmetic operator
			Addition,    // +
			Subtraction, // -
			Multiply,    // *
			Divide,      // /
			Modulus,     // %

			// cmp operator
			GreaterThan,  // >
			LessThan,     // <
			GreaterEqual, // >=
			LessEqual,    // <=
			Equal,        // ==
			NotEqual,     // !=

			// logic operator
			And, // &&
			Or,  // ||
			Eof,
		},
	},
	BoolLiteral: {
		isEOF: true,
		validNextKinds: []Kind{
			CloseParen, // )
			Equal,      // ==
			NotEqual,   // !=

			// logic operator
			And, // &&
			Or,  // ||
			Eof,
		},
	},
	IntegerLiteral: {
		isEOF: true,
		validNextKinds: []Kind{
			CloseParen, // )
			// arithmetic operator
			Addition,    // +
			Subtraction, // -
			Multiply,    // *
			Divide,      // /
			Modulus,     // %

			// cmp operator
			GreaterThan,  // >
			LessThan,     // <
			GreaterEqual, // >=
			LessEqual,    // <=
			Equal,        // ==
			NotEqual,     // !=
			Eof,

			// logic
			And, // &&
			Or,  // ||

		},
	},
	FloatLiteral: {
		isEOF: true,
		validNextKinds: []Kind{
			CloseParen, // )
			// arithmetic operator
			Addition,    // +
			Subtraction, // -
			Multiply,    // *
			Divide,      // /
			Modulus,     // %

			// cmp operator
			GreaterThan,  // >
			LessThan,     // <
			GreaterEqual, // >=
			LessEqual,    // <=
			Equal,        // ==
			NotEqual,     // !=
			Eof,

			// logic
			And, // &&
			Or,  // ||
		},
	},
	StringLiteral: {
		isEOF: true,
		validNextKinds: []Kind{
			CloseParen, // )
			Addition,   // +
			Equal,      // ==
			NotEqual,   // !=
			//GreaterThan,  // >
			//LessThan,     // <
			//GreaterEqual, // >=
			//LessEqual,    // <=
			Eof,

			// logic
			And, // &&
			Or,  // ||
		},
	},

	/*
	* single character operator
	* */
	OpenParen: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			BoolLiteral,    // true, false
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			StringLiteral,  // "abc"
			Addition,       // +
			Subtraction,    // -
			Not,            // !
		},
	},
	CloseParen: {
		isEOF: true,
		validNextKinds: []Kind{
			Addition,     // +
			Subtraction,  // -
			Multiply,     // *
			Divide,       // /
			Modulus,      // %
			GreaterThan,  // >
			LessThan,     // <
			GreaterEqual, // >=
			LessEqual,    // <=
			Equal,        // ==
			NotEqual,     // !=
			And,          // &&
			Or,           // ||
			Eof,
		},
	},
	/*
	* arithmetic operator
	* */
	Addition: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			StringLiteral,  // "abc"
			OpenParen,      // (
			Subtraction,
			Addition,
		},
	},
	Subtraction: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,
			Subtraction,
		},
	},
	Multiply: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,
			Subtraction,
		},
	},
	Divide: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,
			Subtraction,
		},
	},
	Modulus: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
		},
	},

	/*
	* cmp operator
	* */
	GreaterThan: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},
	LessThan: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},
	GreaterEqual: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},
	LessEqual: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},
	Equal: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			StringLiteral,  // "abc"
			BoolLiteral,    // true, false
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},
	NotEqual: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,     // variables
			IntegerLiteral, // 12345
			FloatLiteral,   // 123.45
			StringLiteral,  // "abc"
			BoolLiteral,    // true, false
			OpenParen,      // (
			Addition,       // + 145 > +146
			Subtraction,    // - 145 > -146
		},
	},

	/*
	* logic operator
	* */
	And: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,  // variables
			BoolLiteral, // true, false
			OpenParen,   // (
			Subtraction, // true || -7 > 9
			Addition,
			FloatLiteral,
			IntegerLiteral,
			Not,
		},
	},
	Or: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,  // variables
			BoolLiteral, // true, false
			OpenParen,   // (
			Subtraction, // true || -7 > 9
			Addition,
			FloatLiteral,
			IntegerLiteral,
			Not,
		},
	},
	Not: {
		isEOF: false,
		validNextKinds: []Kind{
			Identifier,  // variables
			BoolLiteral, // true, false
			OpenParen,   // (
			Not,
		},
	},
}

func (s *LexerState) CanTransitionTo(k Kind) bool {
	for _, validKind := range s.validNextKinds {
		if validKind == k {
			return true
		}
	}
	return false
}

func (s *LexerState) IsEOF() bool {
	return s.isEOF
}
