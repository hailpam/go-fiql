package gofiql

import (
	"bytes"
)

// findToken searches for logical operators and, in case of success,
// it returns the first index; otherwise, it returns a -1 which is
// the typical indicator of unsuccessful search.
func findToken(expression *string) (int, []byte) {
	tokens := []byte(",;()")
	comma := tokens[0]
	semicolon := tokens[1]
	lParenthesis := tokens[2]
	rParenthesis := tokens[3]
	chars := []byte(*expression)

	var cntr uint8
	for {
		if !bytes.ContainsAny(chars, ",;") {
			break
		}
		cntr = 0
		for i := 0; i < len(chars); i++ {
			if chars[i] == lParenthesis {
				cntr++
			}
			if chars[i] == rParenthesis {
				cntr--
			}
			if cntr == 0 && (chars[i] == comma || chars[i] == semicolon) {
				return i, chars
			}
		}
		if chars[0] == lParenthesis && chars[len(chars)-1] == rParenthesis {
			chars = chars[1 : len(chars)-1]
		}
	}

	return -1, chars
}

// Parse recursively parses the input string and builds an AST.
// The so built AST can be then traversed for interpretation or
// re-serialization purposes.
func Parse(expression string) (*Node, error) {
	idx, chars := findToken(&expression)
	expression = string(chars)
	node := NewNode()
	node.expression = NewExpression()
	var err error
	if idx == -1 {
		// It is a basic expression. Leaf of the tree.
		l, o, r, err := splitExpression(&expression)
		if err != nil {
			return nil, err
		}
		node.expression.lOperand = l
		node.expression.rOperand = r
		node.expression.operator = o
	} else {
		// It is a complex statement (expressions and lofical operators).
		// Intermediary node of the tree.
		node.lChild, err = Parse(expression[:idx])
		if err != nil {
			return nil, err
		}
		node.rChild, err = Parse(expression[idx+1:])
		if err != nil {
			return nil, err
		}
		o := string(expression[idx])
		node.expression.operator = &o
	}

	return node, nil
}
