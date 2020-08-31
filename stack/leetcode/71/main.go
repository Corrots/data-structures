package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}

//https://leetcode-cn.com/problems/simplify-path/
func simplifyPath(path string) string {
	var res []string
	for _, part := range strings.Split(path, "/") {
		switch part {
		case "":
			break
		case ".":
			break
		case "..":
			if size := len(res); size > 0 {
				res = res[:size-1]
			}
		default:
			res = append(res, part)
		}
	}
	return "/" + strings.Join(res, "/")
}
