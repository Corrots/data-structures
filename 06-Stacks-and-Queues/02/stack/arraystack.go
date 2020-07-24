package stack

import "fmt"

type ArrayStack struct {
	array []interface{}
}

func New(n int) Stack {
	return &ArrayStack{
		array: make([]interface{}, n),
	}
}

func (s *ArrayStack) Push(i interface{}) {
	s.array = append(s.array, i)
	fmt.Println("s: ", s)
}

func (s *ArrayStack) Pop() interface{} {
	s.array = s.array[1:]
	return s.array[0]
}

func (s *ArrayStack) Peek() interface{} {
	fmt.Println("len: ", len(s.array))
	if len(s.array) > 0 {
		return s.array[len(s.array)-1]
	}
	return nil
}

func (s *ArrayStack) Len() int {
	return len(s.array)
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.array) == 0
}



