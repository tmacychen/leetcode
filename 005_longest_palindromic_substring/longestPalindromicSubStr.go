package main

import "fmt"

func main() {
	//	s := "cabadda"
	//s := "caababa"
	s := "abcbaabcba"
	fmt.Printf("the %s longest palindromic substring :%v\n", s, longestPalindrome(s))
}

func isPalindromic(s []byte, head, tail int) bool {
	if s[head] == s[tail] {
		if head < tail {
			return isPalindromic(s, head+1, tail-1)
		}
		if head >= tail {
			return true
		}
	}
	return false
}
func maxString(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func longestPalindrome(s string) (maxStr string) {
	for i, j := len(s)-1, 0; i >= 0; i-- {
		for ; j < i; j++ {
			if isPalindromic([]byte(s), j, i) {
				maxStr = maxString(maxStr, s[j:i+1])
				break
			}

		}
		i = j
	}
	return
}

// func longestPalindrome(s string) string {
//     maxS := ""
//     for i := 0; i < len(s); {
//         j := i+1
//         for j < len(s) && s[i] == s[j] {
//             j++
//         }
//         maxS = maxString(maxS, getLongestPalindromeAt(s, i, j-1))
//         i = j
//     }
//     return maxS
// }

// func getLongestPalindromeAt(s string, i, j int) string {
//     for i >= 0 && j < len(s) && s[i] == s[j] {
//         i--
//         j++
//     }
//     return s[i+1:j]
// }

// func maxString(a, b string) string {
//     if len(a) > len(b) {
//         return a
//     }
//     return b
// }
