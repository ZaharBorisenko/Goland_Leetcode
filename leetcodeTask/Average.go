package main

import (
	"math"
)

func main() {
}

func Average(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		el := array[i]
		sum += el
	}
	result := float64(sum) / float64(len(array))

	return int(math.Round(result))

}
