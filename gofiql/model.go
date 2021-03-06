package gofiql

import "fmt"

// Node defines an Abstract Syntax Tree (AST) node, in turn
// composeed by left and right children. The information is
// reported by the internal expression.
//
// The AST is organized as a composition of nodes and
// expressions, looking like:
//
//    <logical_operator>
//        <lef_expression>
//        <logical_operator>
//            <left_expression>
//            <right_expression>
type Node struct {
	// lChild points to the left child of this node, if any.
	lChild *Node
	// rChild points to the right child of this node, if any.
	rChild *Node
	// expression points to the expression represented by this
	// node.
	expression *Expression
}

// Expression defines an expression in terms of left
// and right operands along with an operator.
//
// The expression for the AST is organized as binary operators
// and operands, looking like:
//
//    <operator>
//        <left_operand>
//        <right_operand>
type Expression struct {
	// lOperand defines the left operand for this operator.
	lOperand *string
	// rOperand defines the right operand of this operator.
	rOperand *string
	// operator defines the operator for this expression and
	// so this node.
	operator *string
}

// NewNode creates a new node and returns a pointer to it.
func NewNode() *Node {
	return &Node{}
}

// String stringified a node, composing its left
// and right children with an expression using the
// infix notation.
func (n *Node) String() string {
	var l, r string
	if n.lChild != nil {
		l = (*n.lChild).String()
	}
	if n.rChild != nil {
		r = (*n.rChild).String()
	}
	s := fmt.Sprintf("{ %s %s %s }", l, n.expression.String(), r)
	return s
}

// Accept takes in input a visitor and applies it according to
// its specific nature.
func (n *Node) Accept(visitor Visitor) (interface{}, error) {
	if n.lChild == nil && n.rChild == nil {
		// Expression node, it's a leaf of the tree
		return (visitor).VisitExpression(n)
	}
	return (visitor).VisitLogicalOperator(n)
}

// NewExpression creates a new expression and returns a
// pointer to it.
func NewExpression() *Expression {
	return &Expression{}
}

// String stringifies an expression, composing its
// operator with left and right operands using the
// infix notation.
func (e *Expression) String() string {
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
