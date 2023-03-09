// 曼哈顿距离：坐标系中两个点的直线距离
// 给定n个坐标点，求两两坐标系之间的曼哈顿距离之和

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var array = [][]int{{1, 1}, {3, 3}, {6, 6}, {7, 7}, {9, 9}, {12, 12}}
	fmt.Println(enumerate(&array))
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
		res += (*array)[i]*i - sum
		sum += (*array)[i]
	}
	return res
}

// 枚举 时间复杂度O(M*N)
func enumerate(array *[][]int) int {
	var total = 0
	for i := 1; i < len(*array); i++ {
		for j := 0; j < i; j++ {
			total += abs((*array)[i][0]-(*array)[j][0]) + abs((*array)[i][1]-(*array)[j][1])
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

// 曼哈顿距离：坐标系中两个点的直线距离
// 给定n个坐标点，求任意两点间的曼哈顿距离的最大值

func MaxSmart(array *[][]int) int {
	aMax, aMin := math.MinInt64, math.MaxInt64
	bMax, bMin := math.MinInt64, math.MaxInt64
	cMax, cMin := math.MinInt64, math.MaxInt64
	dMax, dMin := math.MinInt64, math.MaxInt64
	for i := 0; i < len(*array); i++ {
		aMax = maxInt(+(*array)[i][0]+(*array)[i][1], aMax)
		bMax = maxInt(+(*array)[i][0]-(*array)[i][1], bMax)
		cMax = maxInt(-(*array)[i][0]+(*array)[i][1], cMax)
		dMax = maxInt(-(*array)[i][0]-(*array)[i][1], dMax)

		aMin = minInt(+(*array)[i][0]+(*array)[i][1], aMin)
		bMin = minInt(+(*array)[i][0]-(*array)[i][1], bMin)
		cMin = minInt(-(*array)[i][0]+(*array)[i][1], cMin)
		dMin = minInt(-(*array)[i][0]-(*array)[i][1], dMin)
	}
	max := maxInt(aMax-aMin, math.MinInt64)
	max = maxInt(bMax-bMin, max)
	max = maxInt(cMax-cMin, max)
	max = maxInt(dMax-dMin, max)
	return max
}

func MaxEnumerate(array *[][]int) int {
	max := 0
	for i := 0; i < len(*array); i++ {
		for j := i + 1; j < len(*array); j++ {
			temp := abs((*array)[i][0]-(*array)[j][0]) + abs((*array)[i][1]-(*array)[j][1])
			max = maxInt(temp, max)
		}
	}
	return max
}

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}