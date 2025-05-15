package main

import "fmt"

func main() {
	fmt.Println(ReverseNumbers([]int{100, 200, 300, 400}))
}
func ReverseNumbers(numbers []int) []int {
	result := []int{}
	if len(numbers) == 0 {
		return numbers
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		el := numbers[i]
		result = append(result, el)
	}

	return result
}
