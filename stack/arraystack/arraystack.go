package arraystack

import (
	"fmt"
	"strings"

	"github.com/corrots/data-structures/stack"
)

type ArrayStack struct {
	array *Array
}

func NewStack(capacity int) stack.Stack {
	return &ArrayStack{
		array: NewArray(capacity),
	}
}

// 入栈 - 从栈顶压入元素
func (s *ArrayStack) Push(i interface{}) {
	s.array.AddLast(i)
}

// 出栈 - 从栈顶推出元素
func (s *ArrayStack) Pop() interface{} {
	return s.array.RemoveLast()
}

func (s *ArrayStack) Peek() interface{} {
	return s.array.GetLast()
}

func (s *ArrayStack) Len() int {
	return s.array.Len()
}

func (s *ArrayStack) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s *ArrayStack) String() string {
	var res strings.Builder
	res.WriteString("Stack: ")
	res.WriteString("[")
	for i := 0; i < s.array.Len(); i++ {
		res.WriteString(fmt.Sprintf("%v", s.array.Get(i)))
		if i != s.array.Len()-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("] top")
	return res.String()
}

// 动态数组特有的method
func (s *ArrayStack) Cap() int {
	return s.array.Cap()
}
