package main

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(13))
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Чётное"
	}
	return "Нечётное"
}
