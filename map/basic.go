package _map

type Map interface {
	Add(int, interface{})
	Remove(int) interface{}
	Contains(int) bool
	Get(int) interface{}
	Set(int, interface{})
	Len() int
	IsEmpty() bool
}
