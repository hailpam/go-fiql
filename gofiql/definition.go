package gofiql

import "fmt"

const (
	// expressionRx defines the regular expression to extract
	// group matches for the basic expression.
	expressionRx = `(?P<lOperand>[\w-]+)(?P<operator>=.{0,}=)(?P<rOperand>[\*\"\w-]+)`
)

var (
	// errNotImplemented defines a standard error to be used for
	// functions which require and implementation and that are only
	// a placeholder from the interface.
	errNotImplemented = fmt.Errorf("Should be implemented")
	// errMalformedOperandsStack
	errMalformedOperandsStack = fmt.Errorf("Operands stack should be sized to exactly 2 elements")
	// errMalformedOperator
	errMalformedOperator = fmt.Errorf("Operator cannot be null")
	// errMalformedOperand
	errMalformedOperand = fmt.Errorf("Operand cannot be null")
)
