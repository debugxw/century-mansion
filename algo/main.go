package main

import (
	"math"
)

func main() {
}

func maxInt(arr ...int) int {
	if len(arr) == 0 {
		panic("maxInt: the length of parameter arr should be greater than zero")
	}
	max := math.MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func minInt(arr ...int) int {
	if len(arr) == 0 {
		panic("minInt: the length of parameter arr should be greater than zero")
	}
	min := math.MaxInt
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}
