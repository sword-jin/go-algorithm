package hanoi

import "fmt"

var PrintMove = func(from, to string) {
	fmt.Printf("%s->%s\n", from, to)
}

func HanoiRecursive(n int, from string, to string, via string, move func(from string, to string)) {
	if n == 1 {
		move(from, to)
	} else {
		HanoiRecursive(n-1, from, via, to, move)
		HanoiRecursive(1, from, to, via, move)
		HanoiRecursive(n-1, via, to, from, move)
	}
}
