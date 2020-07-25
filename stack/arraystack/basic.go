package arraystack

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	IsEmpty() bool
}
