package container

type Container interface {
	Size() int
	Empty() bool
}
