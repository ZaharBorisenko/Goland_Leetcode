package main

import (
	"fmt"
)

func main() {
	fmt.Println(Feast("заяц", "зефир"))
}

func Feast(beast string, dish string) bool {
	beastRunes := []rune(beast)
	dishRunes := []rune(dish)

	if beastRunes[0]+beastRunes[len(beastRunes)-1] == dishRunes[0]+dishRunes[len(dishRunes)-1] {
		return true
	}
	return false
}
