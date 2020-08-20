package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

// https://leetcode-cn.com/problems/unique-morse-code-words/
func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, word := range words {
		var res strings.Builder
		for _, v := range word {
			res.WriteString(morse[v-'a'])
		}
		m[res.String()]++
	}
	return len(m)
}
