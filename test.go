package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	a1 := Person{}
	a2 := map[string]int{
		"hi": 5,
		"he": 7,
	}
	a3 := true

	fmt.Printf("%T", a1) //main.Person
	fmt.Printf("%T", a2) //nmap[string]int
	fmt.Printf("%T", a3) //bool

}
