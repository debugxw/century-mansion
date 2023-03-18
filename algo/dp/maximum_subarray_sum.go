package dp

// 力扣53.最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 典型的动态规划
func maxSubArray(nums []int) int {
	pre, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 这里是经过空间优化的代码
		pre = maxInt(pre+nums[i], nums[i])
		max = maxInt(max, pre)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}