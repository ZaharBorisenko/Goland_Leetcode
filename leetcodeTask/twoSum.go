package main

import "fmt"

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	fmt.Println("https://leetcode.com/problems/two-sum/description/")
	var arr []int
	hashTable := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		el := nums[i]
		number := target - el

		val, ok := hashTable[number]

		if ok {
			arr = append(arr, val, i)
		}

		hashTable[el] = i
	}

	return arr
}
