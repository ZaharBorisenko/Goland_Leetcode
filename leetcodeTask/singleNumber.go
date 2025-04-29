package main

import "fmt"

func main() {
	result := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	fmt.Println("https://leetcode.com/problems/single-number/")
	m := make(map[int]int)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num] += 1
		} else {
			m[num] = 1
		}

	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}

	return 0
}
