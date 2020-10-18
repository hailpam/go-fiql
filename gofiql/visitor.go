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
	VisitLogicalOperator(node *LogicalOperatorNode) (interface{}, error)
	VisitExpression(n *ExpressionNode) (interface{}, error)
}

// SQLVisitor provides serialization services for SQL backends.
type SQLVisitor struct {
	// stack provides the means by which operands can be combined with
	// operators.
	stack *stack
}

// NewSQLVisitor creates a new instance of SQL visitor.
func NewSQLVisitor() *SQLVisitor {
	return &SQLVisitor{}
}

// VisitLogicalOperator is a specific implementation of the visit
// method to generate SQL for logical operator nodes.
func (s *SQLVisitor) VisitLogicalOperator(v *LogicalOperatorNode) (interface{}, error) {
	if s.stack.len() != 2 {
		return nil, errMalformedOperandsStack
	}
	r := s.stack.pop()
	l := s.stack.pop()

	o := v.Node.expression.operator
	if o == nil {
		return nil, errMalformedOperator
	}

	f := fmt.Sprintf("(%s %s %s)", *l, *o, *r)
	s.stack.push(&f)

	return f, nil
}

// VisitExpression is a specific implementation of the visit method
// to generate SQL for expression  nodes.
func (s *SQLVisitor) VisitExpression(v *ExpressionNode) (interface{}, error) {
	l := v.Node.expression.lOperand
	r := v.Node.expression.rOperand
	if l == nil || r == nil {
		return nil, errMalformedOperand
	}
	o := v.Node.expression.operator // TBD - to lookup operator
	if o == nil {
		return nil, errMalformedOperator
	}

	f := fmt.Sprintf("%s %s %s", *l, *o, *r)
	s.stack.push(&f)

	return f, nil
}

// Traverse implements a generic traversal which is going to be
// implemented according to the specific instance of Visitor.
func Traverse(root *Node, visitor *Visitor) (interface{}, error) {
	if root.lChild != nil {
		Traverse(root.lChild, visitor)
	}
	if root.rChild != nil {
		Traverse(root.rChild, visitor)
	}

	return root.Accept(visitor)
}
