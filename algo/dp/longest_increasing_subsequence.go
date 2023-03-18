package dp

// 力扣 300. 最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度，可以不连续
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
