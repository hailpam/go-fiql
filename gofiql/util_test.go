package gofiql

import "testing"

func TestStack(t *testing.T) {
	s := newStack()
	a := "one"
	s.push(&a)
	b := "two"
	s.push(&b)
	c := "three"
	s.push(&c)

	if *s.top() != "three" {
		t.Logf("Expected 'three', got: %s", *s.top())
		t.Fail()
	}
	if s.len() != 3 {
		t.Logf("Expected 3 elements, got: %d", s.len())
		t.Fail()
	}
	s.pop()

	if *s.top() != "two" {
		t.Logf("Expected 'two', got: %s", *s.top())
		t.Fail()
	}
	if s.len() != 2 {
		t.Logf("Expected 2 elements, got: %d", s.len())
		t.Fail()
	}
	s.pop()

	if *s.top() != "one" {
		t.Logf("Expected 'one', got: %s", *s.top())
		t.Fail()
	}
	if s.len() != 1 {
		t.Logf("Expected 1 element, got: %d", s.len())
		t.Fail()
	}
	s.pop()

	if s.top() != nil {
		t.Logf("Expected 'nil', got: %s", *s.top())
		t.Fail()
	}
	if s.len() != 0 {
		t.Logf("Expected 0 elements, got: %d", s.len())
		t.Fail()
	}
}

func TestTabs(t *testing.T) {
	tt := tabs(0)
	if tt != "" {
		t.Logf("Expected 'nil', got: %s", tt)
		t.Fail()
	}
	tt = tabs(1)
	if tt != "\t" {
		t.Logf("Expected 1 tab, got: %s", tt)
		t.Fail()
	}
	tt = tabs(2)
	if tt != "\t\t" {
		t.Logf("Expected 2 tabs, got: %s", tt)
		t.Fail()
	}
	tt = tabs(3)
	if tt != "\t\t\t" {
		t.Logf("Expected 3 tabs, got: %s", tt)
		t.Fail()
	}
}

func TestCheckParenthesis(t *testing.T) {
	p := "()()()"
	if !checkParenthesis(&p) {
		t.Logf("Expected %s to be a correct combination", p)
		t.Fail()
	}

	pp := "((()))"
	if !checkParenthesis(&pp) {
		t.Logf("Expected %s to be a correct combination", pp)
		t.Fail()
	}

	ppp := "((()))()()()()()"
	if !checkParenthesis(&ppp) {
		t.Logf("Expected %s to be a correct combination", ppp)
		t.Fail()
	}

	pppp := ")()(()"
	if checkParenthesis(&pppp) {
		t.Logf("Expected %s to be an incorrect combination", pppp)
		t.Fail()
	}

	ppppp := "(((((product==\"Apple\",product==\"Google\");(name==\"Joe\",name==\"Alan\")));label=!~=\"text\";(qty=gte=1,qty=lte=10)))"
	if !checkParenthesis(&ppppp) {
		t.Logf("Expected %s to be an incorrect combination", ppppp)
		t.Fail()
	}
}
