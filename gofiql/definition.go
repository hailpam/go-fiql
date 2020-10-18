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
	// errMalformedOperandsStack defines a standard error to be
	// used in case an operand stack containes any number of
	// elements different from 2 upon a traversal.
	errMalformedOperandsStack = fmt.Errorf("Operands stack should be sized to at least 2 elements")
	// errMalformedOperator defines a standard error to be used
	// in case an operator is null, for some reason or other
	// condition.
	errMalformedOperator = fmt.Errorf("Operator cannot be null")
	// errMalformedOperand defines a standard error to be used
	// in case an operand is null, for some reason or other
	// condition.
	errMalformedOperand = fmt.Errorf("Operand cannot be null")
)

var (
	// sqlLogicalOperators is a mapping of SQL logical operators
	// to the specific FIQL ones.
	sqlLogicalOperators = map[string]string{
		",": "OR",
		";": "AND",
	}
	// sqlOperators is a mapping of SQL operators/functions to the
	// specific FIQL ones.
	sqlOperators = map[string]string{
		"==":    "=",
		"=!=":   "!=",
		"=lt=":  "<",
		"=gt=":  ">",
		"=lte=": "<=",
		"=gte=": ">=",
		"=~=":   "~",
	}
)
