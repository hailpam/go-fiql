package gofiql

import "fmt"

const (
	// expressionRx defines the regular expression to extract
	// group matches for the basic expression.
	expressionRx = `(?P<lOperand>[\w-]+)(?P<operator>=.{0,}=)(?P<rOperand>[\*\"\w-]+)`
)

var (
	// fiqlOr defines the OR operator of the FIQL language.
	fiqlOr = ","
	// fiqlAnd defines the AND operator of the FIQL language.
	fiqlAnd = ";"
	// fiqlOrByte defines the OR operator of the FIQL language and its
	// representative byte.
	fiqlOrByte = []byte(",")[0]
	// fiqlAndByte defines the ANd operator of the FIQL language and its
	// representative byte.
	fiqlAndByte = []byte(";")[0]
	// lParenthesis defines the left parenthesis expected in logical
	// compoound statements.
	lParenthesis = "("
	// rParenthesis defines the left parenthesis expected in logical
	// compoound statements.
	rParenthesis = ")"
	// lParenthesisByte defines the left parenthesis expected in logical
	// compoound statements and its representative byte.
	lParenthesisByte = []byte("(")[0]
	// rParenthesisByte defines the left parenthesis expected in logical
	// compoound statements and its representative byte.
	rParenthesisByte = []byte(")")[0]
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
	// errMalformedParenthesis defines a standard error to be
	// used in case of unbalancede parenthesis which make the
	// overall statement lofically meaningless.
	errMalformedParenthesis = fmt.Errorf("Parenthesis cannot be matched, they seems unbalanced")
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
		"=!~=":  "!~",
	}
)
