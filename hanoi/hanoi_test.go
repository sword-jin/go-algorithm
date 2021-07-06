package hanoi_test

import (
	"testing"

	"github.com/rrylee/go-algorithm/hanoi"
)

func TestHanoiRecursive(t *testing.T) {
	hanoi.HanoiRecursive(3, "A", "B", "C", hanoi.PrintMove)
}
