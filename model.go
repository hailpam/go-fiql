package gofiql

type Node struct {
	lChild     *Node
	rChild     *Node
	constraint *Constraint
}

type Constraint struct {
	lOperand *string
	rOperand *string
	operator *string
}
