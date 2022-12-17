package executor

import (
	"errors"
	"fmt"
	"math"
)

// Parameters is a collection of named parameters that can be used by an EvaluableExpression to retrieve parameters
// when an expression tries to use them.
type Parameters interface {

	// Get gets the parameter of the given name, or an error if the parameter is unavailable.
	// Failure to find the given parameter should be indicated by returning an error.
	Get(name string) (interface{}, error)
}

type MapParameters map[string]interface{}

var DummyParameters = MapParameters(map[string]interface{}{})

func (p MapParameters) Get(name string) (interface{}, error) {
	value, found := p[name]
	if !found {
		errorMessage := "No parameter '" + name + "' found."
		return nil, errors.New(errorMessage)
	}

	return value, nil
}

var (
	divideZeroErr = errors.New("engine: number divide by zero")
)

type operator func(root, left *Node, right *Node, parameters Parameters) (interface{}, TypeFlags, error)

func noopOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return right.value, right.tp, nil
}

// +
func addOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsString() && right.tp.IsString() {
		return left.value.(string) + right.value.(string), TypeString, nil
	} else {
		return execNumberBinOp(left, right, PLUS)
	}
}

// -
func subtractOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return execNumberBinOp(left, right, MINUS)
}

// -
func negateOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if right.tp == TypeFloat {
		return -right.value.(float64), right.tp, nil
	} else {
		return -right.value.(int64), right.tp, nil
	}
}

// +
func positOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return right.value.(float64), right.tp, nil
}

// *
func multiplyOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return execNumberBinOp(left, right, MULTIPLY)
}

// /
func divideOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return execNumberBinOp(left, right, DIVIDE)
}

// %
func modulusOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return execNumberBinOp(left, right, MODULUS)
}

// >=
func gteOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, GTE)
	} else {
		return left.value.(string) >= right.value.(string), TypeBool, nil
	}
}

// >
func gtOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, GT)
	} else {
		return left.value.(string) > right.value.(string), TypeBool, nil
	}
}

// <=
func lteOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, LTE)
	} else {
		return left.value.(string) <= right.value.(string), TypeBool, nil
	}
}

// <
func ltOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, LT)
	} else {
		return left.value.(string) < right.value.(string), TypeBool, nil
	}
}

// ==
func equalOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, EQ)
	} else if left.tp.IsString() && right.tp.IsString() {
		return left.value.(string) == right.value.(string), TypeBool, nil
	} else {
		return left.value.(bool) == right.value.(bool), TypeBool, nil
	}
}

// !=
func notEqualOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	if left.tp.IsNumber() && right.tp.IsNumber() {
		return execNumberBinOp(left, right, NEQ)
	} else if left.tp.IsString() && right.tp.IsString() {
		return left.value.(string) != right.value.(string), TypeBool, nil
	} else {
		return left.value.(bool) != right.value.(bool), TypeBool, nil
	}
}

// &&
func andOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return left.value.(bool) && right.value.(bool), TypeBool, nil
}

// ||
func orOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return left.value.(bool) || right.value.(bool), TypeBool, nil
}

// !
func invertOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return !right.value.(bool), TypeBool, nil
}

// value
func parameterOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	value, err := parameters.Get(root.value.(string))
	if err != nil {
		return nil, TypeNull, err
	}

	val, tp := getType(value)
	if tp.IsNull() {
		return val, tp, errors.New("unsupported type")
	}
	return value, tp, nil
}

// literal
func literalOperator(root, left, right *Node, parameters Parameters) (interface{}, TypeFlags, error) {
	return root.value, root.tp, nil
}

func int2float(n interface{}) float64 {
	switch v := n.(type) {
	case int64:
		return float64(v)
	case float64:
		return v
	default:
		msg := fmt.Sprintf("unreachable code:%v", v)
		panic(msg)
	}
}

func execNumberBinOp(l, r *Node, op Symbol) (interface{}, TypeFlags, error) {
	v1, t1 := l.value.(int64)
	v2, t2 := r.value.(int64)

	var v3, v4 float64
	isInt := t1 && t2
	if !isInt || op == DIVIDE {
		v3, v4 = int2float(l.value), int2float(r.value)
	}
	switch op {
	case PLUS:
		if isInt {
			return v1 + v2, TypeInteger, nil
		}
		return v3 + v4, TypeFloat, nil
	case MINUS:
		if isInt {
			return v1 - v2, TypeInteger, nil
		}
		return v3 - v4, TypeFloat, nil
	case MULTIPLY:
		if isInt {
			return v1 * v2, TypeInteger, nil
		}
		return v3 * v4, TypeFloat, nil
	case DIVIDE:
		if isInt {
			if v2 == 0 {
				return nil, TypeNull, divideZeroErr
			}
			if v1%v2 == 0 {
				return v1 / v2, TypeInteger, nil
			}
		}
		if v4 == 0.0 {
			return nil, TypeNull, divideZeroErr
		}
		return v3 / v4, TypeFloat, nil
	case MODULUS:
		if isInt {
			return v1 % v2, TypeInteger, nil
		}
		return math.Mod(v3, v4), TypeInteger, nil

	case GTE:
		if isInt {
			return v1 >= v2, TypeBool, nil
		}
		return v3 >= v4, TypeBool, nil
	case GT:
		if isInt {
			return v1 > v2, TypeBool, nil
		}
		return v3 > v4, TypeBool, nil
	case LTE:
		if isInt {
			return v1 <= v2, TypeBool, nil
		}
		return v3 <= v4, TypeBool, nil
	case LT:
		if isInt {
			return v1 < v2, TypeBool, nil
		}
		return v3 < v4, TypeBool, nil
	case EQ:
		if isInt {
			return v1 == v2, TypeBool, nil
		}
		return v3 == v4, TypeBool, nil
	case NEQ:
		if isInt {
			return v1 != v2, TypeBool, nil
		}
		return v3 != v4, TypeBool, nil
	default:
		return nil, TypeNull, errors.New("engine: unreachable code")
	}
}
