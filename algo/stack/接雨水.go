package stack

import (
	"algo/utils"
)

// 力扣 42、接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水
// 动态规划实现
func trap(height []int) int {
	n := len(height)
	maxLeft, maxRight := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		maxLeft[i] = utils.MaxInt(maxLeft[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = utils.MaxInt(maxRight[i+1], height[i+1])
	}
	sum := 0
	for i := 1; i < n-1; i++ {
		min := utils.MinInt(maxLeft[i], maxRight[i])
		if height[i] < min {
			sum += min - height[i]
		}
	}
	return sum
}

// 双指针实现 其实就是在动态规划的基础上减少了leftMax rightMax数组的空间消耗
func trapWithPointers(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	sum := 0
	for left < right {
		leftMax = utils.MaxInt(leftMax, height[left])
		rightMax = utils.MaxInt(rightMax, height[right])
		if leftMax < rightMax {
			sum += leftMax - height[left]
			left++
		} else {
			sum += rightMax - height[right]
			right--
		}
	}
	return sum
}

// 利用栈实现，关键思路是一行一行的算
func trapWithStack(height []int) int {
	current := 0
	sum := 0
	var stack []int
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			min := minInt(height[stack[len(stack)-1]], height[current])
			sum += (min - height[index]) * (current - stack[len(stack)-1] - 1)
		}
		stack = append(stack, current)
		current++
	}
	return sum
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
