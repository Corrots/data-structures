package stack

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	IsEmpty() bool
}

