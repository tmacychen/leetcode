package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("string: " + convert("PAYPALISHIRING", 3))
	fmt.Println("string: " + convert("A", 1))
	fmt.Println("string: " + convert("ABCD", 2))
	fmt.Println("string: " + convert("ABCDEF", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	convertedStr := make([][]byte, numRows)
	n := len(s)
	ret := new(bytes.Buffer)

	for i := 0; i < n; {
		for j := 0; j < numRows && i < n; j++ {
			convertedStr[j] = append(convertedStr[j], s[i])
			i++
		}
		for j := numRows - 2; j > 0 && i < n; j-- {
			convertedStr[j] = append(convertedStr[j], s[i])
			i++
		}
	}
	for _, v := range convertedStr {
		ret.WriteString(string(v))

	}
	return ret.String()
}
