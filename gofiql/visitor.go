package gofiql

import "fmt"

// Visitor defines two primary visit functions: visit of a
// logical operator (i.e. node which only has children and
// a logical operator) an expression (i.e. node which only
// has an expression without children).
//
// Implementation of this Visitor interface will guarantee
// interpretation or specific serialization upon the traversal
// of the AST.
type Visitor interface {
	VisitLogicalOperator(node *Node) (interface{}, error)
	VisitExpression(n *Node) (interface{}, error)
}

// SQLVisitor provides serialization services for SQL backends.
type SQLVisitor struct {
	// stack provides the means by which operands can be combined with
	// operators.
	stack *stack
}

// NewSQLVisitor creates a new instance of SQL visitor.
func NewSQLVisitor() *SQLVisitor {
	return &SQLVisitor{
		stack: newStack(),
	}
}

// VisitLogicalOperator is a specific implementation of the visit
// method to generate SQL for logical operator nodes.
func (s *SQLVisitor) VisitLogicalOperator(v *Node) (interface{}, error) {
	if s.stack.len() < 2 {
		return nil, errMalformedOperandsStack
	}
	r := s.stack.pop()
	l := s.stack.pop()

	if v.expression.operator == nil {
		return nil, errMalformedOperator
	}
	o := sqlLogicalOperators[*v.expression.operator]

	f := fmt.Sprintf("(%s %s %s)", *l, o, *r)
	s.stack.push(&f)

	return f, nil
}

// VisitExpression is a specific implementation of the visit method
// to generate SQL for expression  nodes.
func (s *SQLVisitor) VisitExpression(v *Node) (interface{}, error) {
	l := v.expression.lOperand
	r := v.expression.rOperand
	if l == nil || r == nil {
		return nil, errMalformedOperand
	}

	if v.expression.operator == nil {
		return nil, errMalformedOperator
	}
	o := sqlOperators[*v.expression.operator]

	f := fmt.Sprintf("%s %s %s", *l, o, *r)
	s.stack.push(&f)

	return f, nil
}

// Traverse implements a generic traversal which is going to be
// implemented according to the specific instance of Visitor.
func Traverse(root *Node, visitor Visitor) (interface{}, error) {
	if root.lChild != nil {
		Traverse(root.lChild, visitor)
	}
	if root.rChild != nil {
		Traverse(root.rChild, visitor)
	}

	return root.Accept(visitor)
}
