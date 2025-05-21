package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToWeirdCase("Привет мир!"))
}

func ToWeirdCase(str string) string {
	str = strings.Join(strings.Fields(str), "")
	result := ""
	runes := []rune(strings.ToLower(str))

	for i := 0; i < len(runes); i++ {
		el := string(runes[i])
		if i%2 == 0 {
			result += strings.ToUpper(el)
		} else {
			result += el
		}
	}

	return result
}
