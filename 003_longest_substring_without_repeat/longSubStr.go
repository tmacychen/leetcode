package main

import (
	"fmt"
	"log"
)

type HashSet struct {
	m map[rune]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[rune]bool)}
}

func (set *HashSet) Add(c rune) bool {
	if !set.m[c] {
		set.m[c] = true
		return true
	}
	return false
}
func (set *HashSet) Remove(c rune) {
	delete(set.m, c)
}
func (set *HashSet) Contains(c rune) bool {
	return set.m[c]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func main() {
	i := lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("i:%v\n", i)
	i = lengthOfLongestSubstring("bbb")
	fmt.Printf("i:%v\n", i)
	i = a("bbb")
	fmt.Printf("i:%v\n", i)
}

func lengthOfLongestSubstring(s string) int {
	lenSub := 0
	strBytes := []byte(s)
	strLen := len(s)
	set := NewHashSet()
	for i, j := 0, 0; i < strLen && j < strLen; {
		if !set.Contains(rune(strBytes[j])) {
			set.Add(rune(strBytes[j]))
			log.Printf("set add %c\n", rune(strBytes[j]))
			lenSub = max(lenSub, j-i)
			j++
		} else {
			set.Remove(rune(strBytes[i]))
			i++
		}
	}
	return lenSub + 1
}

// The best solution:
// The character arrary contains the last index of all characters.
// At every loop ,Max is updateed to ensure that the Max is the highest index
// of all characters
// the variable count  is the length of substring
// the variable last is the last index of substring that contains different characters

// func lengthOfLongestSubstring(s string) int {
// 	character := [128]int{}
// 	max, count, last := 0, 0, 0
// 	for index, ch := range s {
// 		if character[ch] < last {
// 			count += 1
// 		} else {
// 			last = character[ch]
// 			if count > max {
// 				max = count
// 			}
// 			count = index + 1 - last
// 		}
// 		character[ch] = index + 1
// 	}
// 	if count > max {
// 		max = count
// 	}
// 	return max
// }
//func lengthOfLongestSubstring(s string) int {
func a(s string) int {
	char := make(map[rune]int)
	max, last, count := 0, 0, 0
	for index, ch := range s {
		if char[ch] < last {
			count += 1
		} else {
			last = char[ch]
			if count > max {
				max = count
			}
			count = index + 1 - last
		}
		char[ch] = index + 1
	}
	if max < count {
		max = count
	}
	return max
}
