package rpn

import "testing"

func TestEvalExpressionString(t *testing.T) {
	cases := []struct {
		expression string
		result     int
	}{
		{
			"2 * 3",
			6,
		},
		{
			"3 - 9",
			-6,
		},
		{
			"9/3",
			3,
		},
		{
			"9/3-1",
			2,
		},
		{
			"2 * 3! + (4 * 4 - 33)",
			-5,
		},
		{
			"(2 * 3! + (4 * 4 - 33))^2",
			25,
		},
	}

	for _, c := range cases {
		if v := EvalExpressionString(c.expression); v != c.result {
			t.Errorf("evalRPN error, excepted %d, but get %d", 6, v)
		}
	}
}

func TestReadCurNumFromExpression(t *testing.T) {
	expression := "2 * 3! + (4 * 4 - 33)"

	curIndex := 4
	if v := readCurNum(expression, &curIndex); v != 3 {
		t.Errorf("readCurNum error, excepted %d, but get %d", 3, v)
	}
	if curIndex != 5 {
		t.Errorf("readCurNum error, curIndex excepted %d, but get %d", 5, curIndex)
	}

	curIndex = 18
	if v := readCurNum(expression, &curIndex); v != 33 {
		t.Errorf("readCurNum error, excepted %d, but get %d", 33, v)
	}
	if curIndex != 20 {
		t.Errorf("readCurNum error, curIndex excepted %d, but get %d", 5, curIndex)
	}
}

func TestCalcRpnValue(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	if EvalRPN(tokens) != 22 {
		t.Errorf("evalRPN error, excepted 22, but get %d", EvalRPN(tokens))
	}

	tokens = []string{"4", "3", "-"}

	if EvalRPN(tokens) != 1 {
		t.Errorf("evalRPN error, excepted %d, but get %d", 1, EvalRPN(tokens))
	}
}

func TestFac(t *testing.T) {
	if v := fac(0); v != 1 {
		t.Errorf("fac error, excepted %d, but get %d", 1, v)
	}
	if v := fac(2); v != 2 {
		t.Errorf("fac error, excepted %d, but get %d", 2, v)
	}
	if v := fac(3); v != 6 {
		t.Errorf("fac error, excepted %d, but get %d", 6, v)
	}
	if v := fac(5); v != 120 {
		t.Errorf("fac error, excepted %d, but get %d", 2*3*4*5, v)
	}
}
