package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := isPalindrome(1000021)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
	fmt.Println("https://leetcode.com/problems/palindrome-number/description/")
	str := strconv.Itoa(x)
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
