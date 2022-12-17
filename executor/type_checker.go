package executor

type typeChecker func(left *Node, right *Node) bool

// + > = < >= <=
func numberOrStringChecker(left *Node, right *Node) bool {
	return (left.tp.IsString() && right.tp.IsString()) || (left.tp.IsNumber() && right.tp.IsNumber())
}

// - * / %
func doubleNumberChecker(left *Node, right *Node) bool {
	return left.tp.IsNumber() && right.tp.IsNumber()
}

func matchChecker(left *Node, right *Node) bool {
	return left.tp == right.tp
}

func doubleBoolChecker(left *Node, right *Node) bool {
	return left.tp.IsBool() && right.tp.IsBool()
}

func singleBoolChecker(left *Node, right *Node) bool {
	return right.tp.IsBool()
}

func singleNumberChecker(left *Node, right *Node) bool {
	return right.tp.IsNumber()
}
