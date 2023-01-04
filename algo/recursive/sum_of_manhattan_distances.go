
package main

import (
	"fmt"
	_ "gf-test/internal/packed"
	"sort"
)

func main() {
	//cmd.Main.Run(gctx.New())
	var array = [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {4, -2}}
	//var array = [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}
	//var array = [2]int{1}
	fmt.Println(enumerate(&array))

	//array = [][]int{{2, 2}, {1, 1}, {3, 3}, {4, 4}, {5, 5}}
	fmt.Println(smart(&array))
}

func smart(array *[][]int) int {
	var x = make([]int, len(*array))
	var y = make([]int, len(*array))
	for i := 0; i < len(*array); i++ {
		x[i] = (*array)[i][0]
		y[i] = (*array)[i][1]
	}
	return cal(&x) + cal(&y)
}

func cal(array *[]int) int {
	sort.Ints(*array)
	var res, sum = 0, 0
	for i := 0; i < len(*array); i++ {
		res += (*array)[i] * i - sum
		sum += (*array)[i]
	}
	return res
}

func enumerate(array *[][]int) int {
	var total = 0
	for i := 1; i < len(*array); i++ {
		for j := 0; j < i; j++ {
			total += abs((*array)[i][0] - (*array)[j][0]) + abs((*array)[i][1] - (*array)[j][1])
		}
	}
	return total
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

