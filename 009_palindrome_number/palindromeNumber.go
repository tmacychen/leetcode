package main

import (
	"fmt"
)

func main() {
	i := -2147447412
	fmt.Printf("The number %v is Palindrome ? :%v\n", i, isPalindrome(i))
}

func isPalindrome(i int) bool {
	return i == func(x int) int {
		var ret int
		for x/10 != 0 {
			ret = ret*10 + x%10
			x = (x - x%10) / 10
		}
		ret = ret*10 + x%10
		return ret
	}(i)
}
