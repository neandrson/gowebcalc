package calculation

import (
	"strconv"
)

type operation string

const (
	plus  operation = "+"
	minus operation = "-"
	multi operation = "*"
	div   operation = "/"
)

const (
	openingRoundBracket = "("
	closingRoundBracket = ")"
)

type expression struct {
	numbers    []float64
	operations []operation
}

func infixToPostfix(expression []string) ([]string, error) {
	precedence := map[operation]int{plus: 1, minus: 1, multi: 2, div: 2, openingRoundBracket: 0}
	var output []string
	var operators []string

	for _, token := range expression {
		if _, err := strconv.Atoi(token); err == nil {
			output = append(output, token)
		} else if token == openingRoundBracket {
			operators = append(operators, token)
		} else if token == closingRoundBracket {
			for len(operators) > 0 && operators[len(operators)-1] != openingRoundBracket {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, ErrParse
			}
			operators = operators[:len(operators)-1]
		} else {
			for len(operators) > 0 && precedence[operation(operators[len(operators)-1])] >= precedence[operation(token)] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == openingRoundBracket {
			return nil, ErrParse
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func evaluatePostfix(postfixExpression []string) (float64, error) {
	var stack []float64

	for _, token := range postfixExpression {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, ErrParse
			}
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch operation(token) {
			case plus:
				stack = append(stack, operand1+operand2)
			case minus:
				stack = append(stack, operand1-operand2)
			case multi:
				stack = append(stack, operand1*operand2)
			case div:
				if operand2 == 0 {
					return 0, ErrDivisionByZero
				}
				stack = append(stack, operand1/operand2)
			default:
				return 0, ErrParse
			}
		}
	}

	if len(stack) != 1 {
		return 0, ErrParse
	}

	return stack[0], nil
}

func calculate(expression string) (float64, error) {
	var symbols []string
	i := 0

	for i < len(expression) {
		if expression[i] >= '0' && expression[i] <= '9' {
			num := ""
			for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
				num += string(expression[i])
				i++
			}
			symbols = append(symbols, num)
		} else {
			symbols = append(symbols, string(expression[i]))
			i++
		}
	}
	postfixExpression, err := infixToPostfix(symbols)
	if err != nil {
		return 0, err
	}

	return evaluatePostfix(postfixExpression)
}

// Привет, меня зовут Иван, я разработал сие чудо.
// Я попытался показать почти все что изучил в данном курсе.
// С некоторыми аспектами языка был знаком ранее.
// Все же переписал на польскую нотацию после просмотора видео.
func Calc(expression string) (float64, error) {
	result, err := calculate(expression)
	if err != nil {
		return 0, err
	}
	return result, nil
}
