package main

import (
	"fmt"
)

func main() {
	i := lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("i:%v\n", i)
}

func lengthOfLongestSubstring(s string) int {
	lenSub := 0
	strBytes := []byte(s)
	strLen := len(s)
	var a []byte
	for i := 0; i < strLen; i++ {
		fmt.Printf("%c ", strBytes[i])
	}
	println()

	return lenSub
}
