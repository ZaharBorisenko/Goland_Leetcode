package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}
