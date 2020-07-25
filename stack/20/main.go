package main

import (
	"fmt"

	"github.com/corrots/data-structures/stack/arraystack"
)

/**
#20
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/

func main() {
	s := "{[()]}"
	//for i := 0; i < len(s); i++ {
	//	fmt.Println("s[i]: ", s[i])
	//}
	fmt.Println(isValid(s))
	s1 := "(]"
	fmt.Println(isValid(s1))
	s2 := "()[]{}"
	fmt.Println(isValid(s2))
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := arraystack.NewStack(10)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack.Push(string(s[i]))
		} else {
			if stack.IsEmpty() {
				return false
			}
			topChar := stack.Pop()
			if string(s[i]) != ")" && topChar == "(" {
				return false
			}
			if string(s[i]) != "}" && topChar == "{" {
				return false
			}
			if string(s[i]) != "]" && topChar == "[" {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
