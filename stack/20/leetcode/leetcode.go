package main

import (
	"fmt"
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
	var stack []uint8
	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			topChar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i] != 41 && topChar == 40 {
				return false
			}
			if s[i] != 93 && topChar == 91 {
				return false
			}
			if s[i] != 125 && topChar == 123 {
				return false
			}
		}
	}
	return len(stack) == 0
}
