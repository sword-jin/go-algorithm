package stack

import "github.com/rrylee/go-algorithm/container"

type Stack interface {
	container.Container
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
}
