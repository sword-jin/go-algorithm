package rpn

import (
	"math"
	"strconv"
)

// EvalExpressionString 计算算式表达式的值
// a+b-c*d/e，其中 a,b,c,d,e 都不能为负数
// 支持的符号 $()!+-*/^!
func EvalExpressionString(expression string) int {
	// 需要两个栈，一个存放操作符，一个存放数字
	var opStack RuneStack = []rune{'$'}
	var numStack IntStack = []int{}
	// 在表达式尾部添加一个站位符号
	expression += "$"

	for i := 0; i < len(expression); {
		if expression[i] == ' ' { // 空格处理
			i++
			continue
		}
		d := rune(expression[i])
		if runeIsDigit(d) {
			numStack.Push(readCurNum(expression, &i)) //把数字读出来
		} else {
			for {
				isBreak := true

				// 对比当前操作符和栈顶操作符的优先级
				switch compareOp(opStack.Peek(), d) {
				case '>':
					topOp := opStack.Pop()
					if topOp == '!' {
						numStack.Push(fac(numStack.Pop()))
					} else {
						right, left := numStack.Pop(), numStack.Pop()
						numStack.Push(calc(topOp, left, right))
					}
					isBreak = false
				case '<':
					opStack.Push(d)
				case '=':
					opStack.Pop()
				}

				if isBreak {
					break
				}
			}
			i++
		}
	}

	return numStack.Pop()
}

var compareResult = map[rune]map[rune]rune{
	'$': {
		'$': '=',
		'(': '<',
		')': ' ',
		'+': '<',
		'-': '<',
		'*': '<',
		'/': '<',
		'^': '<',
		'!': '<',
	},
	'(': {
		'$': ' ',
		'(': '<',
		')': '=',
		'+': '<',
		'-': '<',
		'*': '<',
		'/': '<',
		'^': '<',
		'!': '<',
	},
	')': {
		'$': ' ',
		'(': ' ',
		')': ' ',
		'+': ' ',
		'-': ' ',
		'*': ' ',
		'/': ' ',
		'^': ' ',
		'!': ' ',
	},
	'+': {
		'$': '>',
		'(': '<',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '<',
		'/': '<',
		'^': '<',
		'!': '<',
	},
	'-': {
		'$': '>',
		'(': '<',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '<',
		'/': '<',
		'^': '<',
		'!': '<',
	},
	'*': {
		'$': '>',
		'(': '<',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '>',
		'/': '>',
		'^': '<',
		'!': '<',
	},
	'/': {
		'$': '>',
		'(': '<',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '>',
		'/': '>',
		'^': '<',
		'!': '<',
	},
	'^': {
		'$': '>',
		'(': '<',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '>',
		'/': '>',
		'^': '>',
		'!': '<',
	},
	'!': {
		'$': '>',
		'(': ' ',
		')': '>',
		'+': '>',
		'-': '>',
		'*': '>',
		'/': '>',
		'^': '>',
		'!': '>',
	},
}

var digitRunes = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

var opstr = map[rune]string{
	'+': "+",
	'-': "-",
	'*': "*",
	'/': "/",
	'!': "!",
	'^': "^",
}

var ops = map[rune]func(a, b int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
	'^': func(a, b int) int { return pow(a, b) },
}

func compareOp(stackTopOp, curOp rune) rune {
	return compareResult[stackTopOp][curOp]
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func readCurNum(expression string, curIndex *int) int {
	num := 0
	for {
		d := rune(expression[*curIndex])
		if runeIsDigit(d) {
			num = num*10 + runeToNum(d)
			*curIndex += 1
		} else {
			break
		}
	}
	return num
}

func calc(op rune, left, right int) int {
	return ops[op](left, right)
}

func runeIsDigit(i rune) (ok bool) {
	_, ok = digitRunes[i]
	return
}

func runeToNum(i rune) int {
	return digitRunes[i]
}

func fac(a int) int {
	sum := 1
	for i := 1; i <= a; i++ {
		sum *= i
	}
	return sum
}

// EvalRPN 计算逆波兰式的值
func EvalRPN(tokens []string) int {
	ops := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	var s IntStack = []int{}

	for _, token := range tokens {
		i, err := strconv.Atoi(token)
		if err == nil {
			s.Push(i)
		} else {
			right, left := s.Pop(), s.Pop()
			s.Push(ops[token](left, right))
		}
	}

	return s.Pop()
}

type IntStack []int

func (s IntStack) Size() int {
	return len(s)
}

func (s *IntStack) Push(i int) {
	*s = append(*s, i)
}

func (s *IntStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *IntStack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

type RuneStack []rune

func (s RuneStack) Size() int {
	return len(s)
}

func (s *RuneStack) Push(str rune) {
	*s = append(*s, str)
}

func (s *RuneStack) Peek() rune {
	return (*s)[len(*s)-1]
}

func (s *RuneStack) Pop() rune {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
