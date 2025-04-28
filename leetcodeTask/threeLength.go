package main

import "fmt"

func main() {
	result := countSubarrays([]int{-1, -4, -1, 4})
	fmt.Println(result)
}

func countSubarrays(nums []int) int {
	fmt.Println("https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/?envType=daily-question&envId=2025-04-27")
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		if (nums[i]+nums[i+2])*2 == nums[i+1] {
			count++
		}
	}

	return count
}
