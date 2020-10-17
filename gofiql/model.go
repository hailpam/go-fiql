package gofiql

import "fmt"

// Node defines an Abstract Syntax Tree (AST) node, in turn
// composeed by left and right children. The information is
// reported by the internal expression.
type Node struct {
	lChild     *Node
	rChild     *Node
	expression *Expression
}

// Expression defines an expression in terms of left
// and right operands along with an operator.
type Expression struct {
	lOperand *string
	rOperand *string
	operator *string
}

// NewNode creates a new node and returns a pointer to it.
func NewNode() *Node {
	return &Node{}
}

// ToString stringified a node, composing its left
// and right children with an expression using the
// infix notation.
func (n *Node) ToString() string {
	var l, r string
	if n.lChild != nil {
		l = n.lChild.ToString()
	}
	if n.rChild != nil {
		r = n.rChild.ToString()
	}
	s := fmt.Sprintf("{ %s %s %s }", l, n.expression.ToString(), r)
	return s
}

// NewExpression creates a new expression and returns a
// pointer to it.
func NewExpression() *Expression {
	return &Expression{}
}

// ToString stringifies an expression, composing its
// operator with left and right operands using the
// infix notation.
func (e *Expression) ToString() string {
	var l, r, o string
	if e.lOperand != nil {
		l = *e.lOperand
	}
	if e.rOperand != nil {
		r = *e.rOperand
	}
	if e.operator != nil {
		o = *e.operator
	}
	s := fmt.Sprintf("[ %s %s %s ]", l, o, r)
	return s
}
