package gofiql

import (
	"testing"
)

func TestNodeString(t *testing.T) {
	l := "name"
	r := "Joe"
	o := "=="

	n := NewNode()
	n.expression = NewExpression()
	n.expression.lOperand = &l
	n.expression.rOperand = &r
	n.expression.operator = &o

	e := "{  [ name == Joe ]  }"
	if n.String() != e {
		t.Logf("Expected %s, got %s:", e, n.String())
		t.Fail()
	}

	oo := ","
	n.expression = NewExpression()
	n.expression.operator = &oo

	ln := NewNode()
	ln.expression = NewExpression()
	rn := NewNode()
	rn.expression = NewExpression()
	o = "=="
	ln.expression.lOperand = &l
	ln.expression.rOperand = &r
	ln.expression.operator = &o
	r = "Tom"
	rn.expression.lOperand = &l
	rn.expression.rOperand = &r
	rn.expression.operator = &o

	n.lChild = ln
	n.rChild = rn

	e = "{ {  [ name == Tom ]  } [  ,  ] {  [ name == Tom ]  } }"
	if n.String() != e {
		t.Logf("Expected %s, got %s:", e, n.String())
		t.Fail()
	}
}

func TestNodeAccept(t *testing.T) {
	l := "name"
	r := "Joe"
	o := "=="

	n := NewNode()
	n.expression = NewExpression()
	n.expression.lOperand = &l
	n.expression.rOperand = &r
	n.expression.operator = &o

	v := NewSQLVisitor()
	i, err := n.Accept(v)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	e := "name = Joe"
	if i.(string) != e {
		t.Logf("Expected %s, got %s:", e, i.(string))
		t.Fail()
	}

	oo := ","
	n.expression = NewExpression()
	n.expression.operator = &oo

	ln := NewNode()
	ln.expression = NewExpression()
	rn := NewNode()
	rn.expression = NewExpression()
	o = "=="
	ln.expression.lOperand = &l
	ln.expression.rOperand = &r
	ln.expression.operator = &o
	r = "Tom"
	rn.expression.lOperand = &l
	rn.expression.rOperand = &r
	rn.expression.operator = &o

	n.lChild = ln
	n.rChild = rn

	nn1 := "name = Joe"
	nn2 := "name = Tom"
	v.stack.push(&nn1)
	v.stack.push(&nn2)
	i, err = n.Accept(v)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	e = "(name = Joe OR name = Tom)"
	if i.(string) != e {
		t.Logf("Expected %s, got %s:", e, i.(string))
		t.Fail()
	}
}
