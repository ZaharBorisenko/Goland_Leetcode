package main

import (
	"fmt"
	"strings"
)

func main() {
	result := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(result)
}

func lengthOfLastWord(s string) int {
	fmt.Println("https://leetcode.com/problems/length-of-last-word/")
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])
}
