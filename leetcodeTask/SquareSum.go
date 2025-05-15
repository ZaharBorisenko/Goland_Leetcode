package main

import "fmt"

func main() {
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
}

func SquareSum(nums []int) (result int) {
	for i := 0; i <= len(nums)-1; i++ {
		el := nums[i]
		result += el * el
	}
	return
}
