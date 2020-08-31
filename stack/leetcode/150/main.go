package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s1 := []string{"2", "1", "+", "3", "*"}
	//s1 := []string{"4", "13", "5", "/", "+"}
	s1 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(s1))
}

/**
逆波兰表达式主要有以下两个优点：
1.去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
2.适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
//https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	var res int
	if len(tokens) < 2 {
		res, _ = strconv.Atoi(tokens[0])
		return res
	}
	var stack []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			size := len(stack)
			first, second := stack[size-2], stack[size-1]
			res = calculate(first, second, tokens[i])
			stack = stack[:size-2]
			stack = append(stack, res)
		} else {
			n, _ := strconv.Atoi(tokens[i])
			stack = append(stack, n)
		}
	}
	return res
}

func calculate(i, j int, op string) int {
	var res int
	switch op {
	case "+":
		res = i + j
	case "-":
		res = i - j
	case "*":
		res = i * j
	case "/":
		res = i / j
	}
	return res
}
