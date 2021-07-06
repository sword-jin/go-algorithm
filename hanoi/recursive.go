package hanoi

import "fmt"

type MoveFunc func(from, to string)

var PrintMove = func(from, to string) {
	fmt.Printf("%s->%s\n", from, to)
}

func HanoiRecursive(n int, from string, to string, via string, move MoveFunc) {
	if n == 1 {
		move(from, to)
	} else {
		HanoiRecursive(n-1, from, via, to, move)
		HanoiRecursive(1, from, to, via, move)
		HanoiRecursive(n-1, via, to, from, move)
	}
}
