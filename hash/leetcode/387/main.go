package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

//https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	record := make([]int, 26)
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if record[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
