package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidPalindrome("A man, a plan, a canal: Panama"))
}

func isValidPalindrome(s string) bool {
	s = strip(s)

	left := 0
	right := len(s) - 1

	for left <= right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func strip(s string) string {
	s = strings.ToLower(s)
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if 'a' <= b && b <= 'z' || '0' <= b && b <= '9' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
