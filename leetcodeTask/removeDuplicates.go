package main

import "fmt"

func main() {
	result := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	count := 1
	prevElement := nums[0]

	for i := 1; i < len(nums); i++ {
		fmt.Println("https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/1619996420/")
		if nums[i] != prevElement {
			nums[count] = nums[i]
			count++
			prevElement = nums[i]
		}
	}

	return count
}
