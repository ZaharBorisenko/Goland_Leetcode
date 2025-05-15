package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ReverseNumber(-123))
}

func ReverseNumber(n int) int {
	str := strconv.Itoa(n)
	isNegative := false
	reverseStr := ""
	if string(str[0]) == "-" {
		isNegative = true
		str = str[1:]
	}

	for i := len(str) - 1; i >= 0; i-- {
		el := string(str[i])
		reverseStr += el
	}

	if isNegative {
		reverseStr = "-" + reverseStr
	}
	result, _ := strconv.Atoi(reverseStr)

	return result
}
