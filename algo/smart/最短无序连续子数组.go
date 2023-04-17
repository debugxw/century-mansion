package smart

import "math"

// 力扣 581、最短无序连续子数组
// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序
// 请你找出符合题意的 最短 子数组，并输出它的长度

// 方法一：单调栈
func findUnsortedSubarrayWithStack(nums []int) int {
	length := 0
	var stack []int
	firstIndex := -2
	for i := range nums {
		if len(stack) == 0 || nums[i] >= nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			var j int
			if firstIndex == -2 {
				j = stack[len(stack)-1]
			} else {
				j = firstIndex
			}
			for j >= 0 && nums[j] > nums[i] {
				j--
			}
			firstIndex = j
			length = i - j
		}
	}
	return length
}

// 方法二：聪明人才能想出来的方法
func findUnsortedSubarray(nums []int) int {
	start, end := 0, -1
	max, min := math.MinInt, math.MaxInt
	length := len(nums)
	for i := range nums {
		if max > nums[i] {
			end = i
		} else {
			max = nums[i]
		}
		if min < nums[length-i-1] {
			start = length - i - 1
		} else {
			min = nums[length-i-1]
		}
	}
	return end - start + 1
}
