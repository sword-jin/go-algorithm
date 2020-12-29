package queue

import "github.com/rrylee/go-algorithm/container"

type Queue interface {
	container.Container
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
}
