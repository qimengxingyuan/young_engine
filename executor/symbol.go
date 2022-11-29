package executor

import (
	"fmt"
	"github.com/qimengxingyuan/young_engine/token"
)

type Symbol int

const (
	VALUE    Symbol = iota
	LITERAL         // var string int float bool
	NOOP            // noop
	EQ              // ==
	NEQ             // !=
	GT              // >
	LT              // <
	GTE             // >=
	LTE             // <=
	AND             // &&
	OR              // ||
	PLUS            // +
	MINUS           // -
	MULTIPLY        // *
	DIVIDE          // /s
	MODULUS         // %
	INVERT          // ！
	POSITIVE        // +
	NEGATIVE        // -
)

const (
	binaryErrFmt = "type mismatch for operator [%s]: left='%s', right='%s'"
	unaryErrFmt  = "type mismatch for operator [%s]: right='%s'"
)

var (
	MultiKindsToSymbol = map[token.Kind]Symbol{
		token.Multiply: MULTIPLY,
		token.Divide:   DIVIDE,
		token.Modulus:  MODULUS,
	}

	AddKindsToSymbol = map[token.Kind]Symbol{
		token.Addition:    PLUS,
		token.Subtraction: MINUS,
	}

	CompareKindsToSymbol = map[token.Kind]Symbol{
		token.GreaterThan:  GT,
		token.GreaterEqual: GTE,
		token.LessThan:     LT,
		token.LessEqual:    LTE,
		token.Equal:        EQ,
		token.NotEqual:     NEQ,
	}

	OrKindsToSymbol = map[token.Kind]Symbol{
		token.Or: OR,
	}

	AndKindsToSymbol = map[token.Kind]Symbol{
		token.And: AND,
	}

	NotKindsToSymbol = map[token.Kind]Symbol{
		token.Not: INVERT,
	}

	PrefixKindsToSymbol = map[token.Kind]Symbol{
		token.Addition:    POSITIVE,
		token.Subtraction: NEGATIVE,
	}

	fixedSymbolMap = map[Symbol]Symbol{
		PLUS:  POSITIVE,
		MINUS: NEGATIVE,
	}

	symbolToOperator = map[Symbol]operator{
		VALUE:    parameterOperator,
		LITERAL:  literalOperator,
		NOOP:     noopOperator,
		EQ:       equalOperator,
		NEQ:      notEqualOperator,
		GT:       gtOperator,
		LT:       ltOperator,
		GTE:      gteOperator,
		LTE:      lteOperator,
		AND:      andOperator,
		OR:       orOperator,
		PLUS:     addOperator,
		MINUS:    subtractOperator,
		MULTIPLY: multiplyOperator,
		DIVIDE:   divideOperator,
		MODULUS:  modulusOperator,
		INVERT:   invertOperator,
		NEGATIVE: negateOperator,
		POSITIVE: noopOperator,
	}

	symbolToTypeChecker = map[Symbol]typeChecker{
		VALUE:    nil,
		LITERAL:  nil,
		NOOP:     nil,
		EQ:       matchChecker,
		NEQ:      matchChecker,
		GT:       numberOrStringChecker,
		LT:       numberOrStringChecker,
		GTE:      numberOrStringChecker,
		LTE:      numberOrStringChecker,
		AND:      doubleBoolChecker,
		OR:       doubleBoolChecker,
		PLUS:     numberOrStringChecker,
		MINUS:    doubleNumberChecker,
		MULTIPLY: doubleNumberChecker,
		DIVIDE:   doubleNumberChecker,
		MODULUS:  doubleNumberChecker,
		INVERT:   singleBoolChecker,
		NEGATIVE: singleNumberChecker,
		POSITIVE: singleNumberChecker,
	}
)

func (s Symbol) String() string {
	switch s {
	case NOOP:
		return "NOOP"
	case VALUE:
		return "VALUE"
	case EQ:
		return "="
	case NEQ:
		return "!="
	case GT:
		return ">"
	case LT:
		return "<"
	case GTE:
		return ">="
	case LTE:
		return "<="
	case AND:
		return "&&"
	case OR:
		return "||"
	case PLUS, POSITIVE:
		return "+"
	case MINUS, NEGATIVE:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case MODULUS:
		return "%"
	case INVERT:
		return "!"
	}
	return ""
}

func (s Symbol) getOperator() operator {
	if op, exist := symbolToOperator[s]; exist {
		return op
	}
	return nil
}

func (s Symbol) getTypeChecker() typeChecker {
	if checker, exist := symbolToTypeChecker[s]; exist {
		return checker
	}
	return nil
}

func (s Symbol) formatTypeError(left, right *Node) error {
	switch s {
	case PLUS, MINUS, MULTIPLY, DIVIDE, MODULUS:
		return fmt.Errorf(binaryErrFmt, s.String(), left.tp.String(), right.tp.String())
	case GT, GTE, LT, LTE, EQ, NEQ, AND, OR:
		return fmt.Errorf(binaryErrFmt, s.String(), left.tp.String(), right.tp.String())
	case NEGATIVE, POSITIVE, INVERT:
		return fmt.Errorf(unaryErrFmt, s.String(), right.tp.String())
	default:
		return fmt.Errorf("type error for %v", s.String())
	}
}
