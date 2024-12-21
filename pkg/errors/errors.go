package errors

import "errors"

var (
	ErrMismatchedParentheses = errors.New("несоответствующие скобки")
	ErrInvalidNumber         = errors.New("недопустимая цифра в выражении")
	ErrInvalidExpression     = errors.New("недопустимое выражение")
	ErrNotEnoughOperands     = errors.New("недостаточно операндов")
	ErrDivisionByZero        = errors.New("деление на ноль")
	ErrUnknownOperator       = errors.New("неизвестный оператор")
)