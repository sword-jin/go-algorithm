package container

type Container interface {
	Size() int
	Empty() bool
}

type CompareRet uint8

const (
	CompareEqual CompareRet = iota
	CompareGt
	CompareLt
)

type Compare interface {
	Compare(interface{}) CompareRet
}
