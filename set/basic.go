package set

type Set interface {
	Add(interface{})
	Remove(interface{})
	Contains(interface{}) bool
	Len() int
	IsEmpty() bool
}
