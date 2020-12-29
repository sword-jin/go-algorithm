package stack

type arraystack struct {
	list []interface{}
}

func (a *arraystack) Size() int {
	return len(a.list)
}

func (a *arraystack) Empty() bool {
	return a.Size() == 0
}

func NewArrayStack() Stack {
	return &arraystack{list: []interface{}{}}
}

func (a *arraystack) Push(value interface{}) {
	a.list = append(a.list, value)
}

func (a *arraystack) Pop() (value interface{}, ok bool) {
	value, ok = a.Peek()
	if ok {
		a.list = a.list[:len(a.list)-1]
	}
	return
}

func (a *arraystack) Peek() (value interface{}, ok bool) {
	if len(a.list) == 0 {
		ok = false
		return
	}
	ok = true
	value = a.list[len(a.list)-1]
	return
}
