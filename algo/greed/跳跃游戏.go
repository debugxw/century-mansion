package greed

// CanJump 力扣 55、跳跃游戏
// 给定一个非负整数数组 nums ，你最初位于数组的第一个下标
// 数组中的每个元素代表你在该位置可以跳跃的最大长度
// 判断你是否能够到达最后一个下标
func CanJump(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums) && i <= rightMost; i++ {
		if rightMost < i+nums[i] {
			rightMost = i + nums[i]
		}
		if rightMost >= len(nums)-1 {
			return true
		}
	}
	return false
}