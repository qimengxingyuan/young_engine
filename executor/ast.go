package executor

import (
	"io"
	"os"
)

type Node struct {
	symbol Symbol
	value  interface{}
	tp     TypeFlags

	leftNode, rightNode *Node

	// the operator that will be used to evaluate this node (such as adding [left] to [right] and return the result)
	operator operator

	// ensures that both left and right values are appropriate for this node. Returns an error if they aren't operable.
	typeChecker typeChecker
}

// NewNodeWithPrefixFix The symbol `+` - can represent both unary and binary operators,
// and the priority of the operator is unique and needs to be adjusted on a case-by-case basis
// - In abstract syntax tree, if the left subtree of the node is empty and the right subtree is not empty,
//   it can be judged to be a negative sign. The symbol needs to be corrected
// - If the right subtree of the right subtree is a symbol other than the prefix symbol [+、-],
//	  The node order needs to be corrected
func NewNodeWithPrefixFix(right *Node, symbol Symbol, value interface{}) *Node {
	needFixed := needFixedSymbol[symbol]
	if !needFixed {
		panic("should not use this new node function for current symbol")
	}
	if right != nil && right.rightNode != nil && right.symbol != NEGATIVE && right.symbol != POSITIVE {
		right.leftNode = NewNode(nil, right.leftNode, symbol, value)
		return right
	} else {
		return NewNode(nil, right, symbol, value)
	}
}

func NewNode(left, right *Node, symbol Symbol, value interface{}) *Node {
	return NewNodeWithType(left, right, symbol, value, TypeNull)
}

func NewNodeWithType(left, right *Node, symbol Symbol, value interface{}, tp TypeFlags) *Node {
	return &Node{
		symbol:      symbol,
		value:       value,
		tp:          tp,
		leftNode:    left,
		rightNode:   right,
		operator:    symbol.getOperator(),
		typeChecker: symbol.getTypeChecker(),
	}
}

func (n *Node) PrintSvg(name string) {
	svgFile := name + ".svg"
	file, err := os.Create(svgFile)
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(file, n.xmlTree())
	if err != nil {
		panic(err)
	}
	file.Close()
}

func (n *Node) Eval(parameters map[string]interface{}) error {
	if n == nil {
		return nil
	}

	if n.leftNode != nil {
		err := n.leftNode.Eval(parameters)
		if err != nil {
			return err
		}
	}

	if n.rightNode != nil {
		err := n.rightNode.Eval(parameters)
		if err != nil {
			return err
		}
	}

	if n.typeChecker != nil {
		if !n.typeChecker(n.leftNode, n.rightNode) {
			return n.symbol.formatTypeError(n.leftNode, n.rightNode)
		}
	}

	ret, tp, err := n.operator(n, n.leftNode, n.rightNode, MapParameters(parameters))
	if err != nil {
		return err
	}
	n.value = ret
	n.tp = tp

	return nil
}

func (n *Node) GetVal() (interface{}, TypeFlags) {
	return n.value, n.tp
}
