package gofiql

const (
	// expressionRx defines the regular expression to extract
	// group matches for the basic expression.
	expressionRx = `(?P<lOperand>[\w-]+)(?P<operator>=.{0,}=)(?P<rOperand>[\*\"\w-]+)`
)
