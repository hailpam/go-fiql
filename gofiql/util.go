package gofiql

import (
	"bytes"
	"fmt"
	"regexp"
)

// spitExpression takes in input a constraint expression and splits it
// into its component parts, i.e. left operand, right operand and
// operator.
// It makese uses of a regular expression.
func splitExpression(expression *string) (*string, *string, *string, error) {
	re := regexp.MustCompile(expressionRx)
	if re == nil {
		return nil, nil, nil, fmt.Errorf("Problem with the regexp: unable to compile")
	}
	match := re.FindAllStringSubmatch(*expression, -1)
	if len(match) == 0 {
		return nil, nil, nil, fmt.Errorf("No match found for %s", *expression)
	}
	names := re.SubexpNames()
	if len(names) != 4 {
		return nil, nil, nil, fmt.Errorf("Problem with the regexp: not enough names")
	}
	var lOperand, rOperand, operator string
	for i, v := range match[0] {
		if names[i] == "lOperand" {
			lOperand = v
		} else if names[i] == "rOperand" {
			rOperand = v
		} else if names[i] == "operator" {
			operator = v
		}
	}

	return &lOperand, &operator, &rOperand, nil
}

// tabs concatenas tabs according to the depth and so the
// indication of the tree level provided in input.
func tabs(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("\t")
	}

	return buffer.String()
}

// PrettyPrinting performs an AST traversal and prints out
// the tree in a pretty hierarchical format.
func PrettyPrinting(root *Node, depth int) {
	fmt.Printf("%s%s\n", tabs(depth), root.expression.String())
	if root.lChild != nil {
		PrettyPrinting(root.lChild, depth+1)
	}
	if root.rChild != nil {
		PrettyPrinting(root.rChild, depth+1)
	}
}

// stack provides a basic implementation of the stack logic
// based off a string slice.
type stack struct {
	stack []*string
}

// Push pushes an element on the stack
func (s *stack) push(v *string) {
	s.stack = append(s.stack, v)
}

// Top returns the topmost element without removing it from
// the datastructure.
func (s *stack) top() *string {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}
	return nil
}

// Pop returns the topmost elemet and removes it from the
// datastructure.
func (s *stack) pop() *string {
	if len(s.stack) > 0 {
		v := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return v
	}
	return nil
}

// Len returns the length in terms of element of the backing
// datastructure.
func (s *stack) len() int {
	return len(s.stack)
}
