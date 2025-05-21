package main

import "fmt"

func main() {
	fmt.Println(FindUnique([]int{1, 1, 1, 2, 1, 1}))
}

func FindUnique(arr []int) interface{} {
	obj := make(map[int]int)

	for _, value := range arr {
		if _, ok := obj[value]; ok {
			obj[value] += 1
		} else {
			obj[value] = 1
		}
	}

	for key, value := range obj {
		if value == 1 {
			return key
		}
	}

	return "Не найден"
}
