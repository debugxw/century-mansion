package stack

import "algo/utils"

// 力扣 84、柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1
// 求在该柱状图中，能够勾勒出来的矩形的最大面积
func largestRectangleArea(heights []int) int {
	var stack []int
	largest := 0
	for cur, h := range heights {
		for len(stack) != 0 && h < heights[stack[len(stack)-1]] {
			calIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := cur
			if len(stack) != 0 {
				width = cur - stack[len(stack)-1] - 1
			}
			largest = utils.MaxInt(largest, heights[calIndex]*width)
		}
		stack = append(stack, cur)
	}
	cur := len(heights)
	for len(stack) != 0 {
		calIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := cur
		if len(stack) != 0 {
			width = cur - stack[len(stack)-1] - 1
		}
		largest = utils.MaxInt(largest, heights[calIndex]*width)
	}
	return largest
}

// 力扣 85、最大矩形
// 给定一个仅包含 0 和 1 、大小为 rows * cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积
// 本质上和84题是一样的，把每一层看做是柱状图即可
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 && heights[j] != 0 {
				heights[j] += 1
			} else {
				heights[j] = int(matrix[i][j])
			}
		}
		maxArea = utils.MaxInt(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}
