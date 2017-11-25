package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("x :%v\n", reverse(123))
	fmt.Printf("x :%v\n", reverse(1534236469))
}

func reverse(x int) int {
	var ret int
	for x/10 != 0 {
		ret = ret*10 + x%10
		x = (x - x%10) / 10
	}
	ret = ret*10 + x%10
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret

}

// func reverse(x int) int {
// 	a := x / 10
// 	b := x % 10
// 	integer := []int{}
// 	for a != 0 {
// 		integer = append(integer, b)
// 		b = a % 10
// 		a = a / 10
// 	}
// 	integer = append(integer, b)
// 	if integer[0] == 0 {
// 		integer = integer[1:]
// 	}
// 	ret := 0
// 	for _, v := range integer {
// 		ret = ret*10 + v
// 	}
// 	if ret > 2147483647 || ret < -2147483647 {
// 		return 0
// 	}
// 	return ret
// }
