package calculation

import (
	"errors"
	"fmt"
)

var (
	ErrParse            = errors.New("parsing error")
	ErrExpression       = fmt.Errorf("wrong expression : %w", ErrParse)
	ErrInvalidOperation = errors.New("invalid operation")
	ErrDivisionByZero   = errors.New("division by zero")
)
