package calculation

import (
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		expression []string
		expected   []string
		err        error
	}{
		{[]string{"3", "+", "4"}, []string{"3", "4", "+"}, nil},
		{[]string{"3", "+", "4", "*", "2"}, []string{"3", "4", "2", "*", "+"}, nil},
		{[]string{"(", "3", "+", "4", ")", "*", "2"}, []string{"3", "4", "+", "2", "*"}, nil},
		{[]string{"(", "3", "+", "4"}, nil, ErrParse},
		{[]string{"3", "*", "(", "4", "+", "2", ")"}, []string{"3", "4", "2", "+", "*"}, nil},
		{[]string{"3", "*", "(", "4", "+", "2"}, nil, ErrParse},
		{[]string{"(", "3", "+", "4", ")", "*", "(", "2", "+", "1", ")"}, []string{"3", "4", "+", "2", "1", "+", "*"}, nil},
	}

	for _, test := range tests {
		result, err := infixToPostfix(test.expression)
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
		if !equal(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		expression []string
		expected   float64
		err        error
	}{
		{[]string{"3", "4", "+"}, 7, nil},
		{[]string{"3", "4", "2", "*", "+"}, 11, nil},
		{[]string{"3", "4", "+", "2", "*"}, 14, nil},
		{[]string{"3", "4", "/"}, 0.75, nil},
		{[]string{"3", "0", "/"}, 0, ErrDivisionByZero},
		{[]string{"3", "+"}, 0, ErrParse},
		{[]string{"3", "4", "2", "*", "/"}, 0.375, nil},
		{[]string{"3", "4", "2", "*", "+"}, 11, nil},
	}

	for _, test := range tests {
		result, err := evaluatePostfix(test.expression)
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		err        error
	}{
		{"3+4", 7, nil},
		{"3+4*2", 11, nil},
		{"3-2*2", -1, nil},
		{"(3+4)*2", 14, nil},
		{"3/4", 0.75, nil},
		{"3/0", 0, ErrDivisionByZero},
		{"3+", 0, ErrParse},
		{"3*(4+2)", 18, nil},
		{"3*(4+2", 0, ErrParse},
		{"10+(2*3)", 16, nil},
		{"10+(2*3", 0, ErrParse},
	}

	for _, test := range tests {
		result, err := Calc(test.expression)
		if err != test.err {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
